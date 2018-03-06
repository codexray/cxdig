package referential

import (
	"time"

	"path/filepath"

	"codexray/cxdig/filetypes"
	"codexray/cxdig/types"

	log "github.com/sirupsen/logrus"
)

// ReferentialBuilder is responsible of tracking the changes done on all the source files of a code base
type ReferentialBuilder struct {
	currentFileID types.LocalFileID
	files         map[string]*types.LocalFile
}

func NewReferentialBuilder() *ReferentialBuilder {
	return &ReferentialBuilder{
		files: map[string]*types.LocalFile{},
	}
}

func (r *ReferentialBuilder) finalize() types.ProjectReferential {
	result := make(types.ProjectReferential, 0, len(r.files))
	for _, f := range r.files {
		if f != nil {
			result = append(result, *f)
		}
	}
	return result
}

func (r *ReferentialBuilder) addFile(ch types.FileChange, commitID types.CommitID, authorID string, dateTime time.Time) {
	f := r.files[ch.FilePath]
	if f != nil {
		if f.DeletionDate == nil {
			log.WithField("file", ch.FilePath).Debug("Added file already exists, treating as modified")
			r.modifyFile(ch, commitID, authorID, dateTime)
			return
		}
		// when the exact same file name was removed then added again, consider it to be a file restore
		f.DeletionDate = nil
		act := types.ActivityInfo{
			CommitID: commitID,
			Date:     dateTime,
		}
		f.Activity.UndeletionDates = append(f.Activity.UndeletionDates, act)
		f.AuthorCommits[authorID]++
	} else {
		ftype, lang := filetypes.IdentifyFileTypeAndLanguage(ch.FilePath)

		r.currentFileID = r.currentFileID.Next()
		r.files[ch.FilePath] = &types.LocalFile{
			ID:            r.currentFileID,
			LatestPath:    ch.FilePath,
			CreationDate:  dateTime,
			FileType:      ftype,
			Language:      lang,
			AuthorCommits: make(map[string]int),
		}
		r.files[ch.FilePath].AuthorCommits[authorID]++
	}
}

func (r *ReferentialBuilder) modifyFile(ch types.FileChange, commitID types.CommitID, authorID string, dateTime time.Time) {
	f := r.files[ch.FilePath]
	if f == nil {
		log.WithField("file", ch.FilePath).Debug("Modified file does not exist, treating as added")
		r.addFile(ch, commitID, authorID, dateTime)
		return
	}
	act := types.ActivityInfo{
		CommitID: commitID,
		Date:     dateTime,
	}
	f.Activity.ModificationDates = append(f.Activity.ModificationDates, act)
	f.AuthorCommits[authorID]++
}

func (r *ReferentialBuilder) renameFile(ch types.FileChange, commitID types.CommitID, authorID string, dateTime time.Time) {
	oldName := ch.FilePath
	newName := ch.RenamedFile
	if newName == "" {
		log.WithField("file", ch.FilePath).Panic("renamed name value is missing")
	}

	f := r.files[oldName]
	if f == nil {
		// check if the file was already renamed
		if r.files[newName] != nil {
			names := r.files[newName].PreviousNames
			if len(names) > 0 {
				latest := names[len(names)-1]
				if latest.FullPath == oldName {
					log.WithField("file", oldName).Debug("File renaming was already applied, ignoring")
					return
				}
			}
		}
		log.WithField("file", oldName).Warn("Renamed file does not exist, treating as added")
		r.addFile(ch, commitID, authorID, dateTime)
		return
	}

	// compare old/new file types
	oldType, oldLang := filetypes.IdentifyFileTypeAndLanguage(oldName)
	newType, newLang := filetypes.IdentifyFileTypeAndLanguage(newName)
	if oldType != filetypes.FileTypeUnknown &&
		newType != oldType &&
		newType != filetypes.FileTypeGenerator { // it's quite common to move a source file to a generated source file
		log.WithFields(log.Fields{
			"old-name": oldName,
			"new-name": newName}).Warn("File type was changed by renaming")
	} else if oldLang != filetypes.LanguageUnknown &&
		newLang != oldLang {
		log.WithFields(log.Fields{
			"old-name": oldName,
			"new-name": newName}).Debug("File language was changed by renaming")
	}

	// Process the renaming
	// Note: with Git, when a file was tagged as renamed, we are sure it was only renamed and not modified
	f.PreviousNames = append(f.PreviousNames, types.FileNameInfo{
		FullPath:      oldName,
		EndOfValidity: dateTime,
	})
	f.LatestPath = newName
	f.FileType = newType
	f.Language = newLang
	f.AuthorCommits[authorID]++

	// is it a file rename and/or a file move?
	if filepath.Base(oldName) != filepath.Base(newName) {
		ren := types.RenamingInfo{
			ActivityInfo: types.ActivityInfo{
				CommitID: commitID,
				Date:     dateTime,
			},
			PreviousName: oldName,
		}
		f.Activity.RenamingDates = append(f.Activity.RenamingDates, ren)
	}
	if filepath.Dir(oldName) != filepath.Dir(newName) {
		act := types.ActivityInfo{
			CommitID: commitID,
			Date:     dateTime,
		}
		f.Activity.RelocationDates = append(f.Activity.RelocationDates, act)
	}

	// apply
	r.files[oldName] = nil
	r.files[newName] = f
}

func (r *ReferentialBuilder) deleteFile(ch types.FileChange, authorID string, dateTime time.Time) {
	f := r.files[ch.FilePath]
	if f == nil {
		log.WithField("file", ch.FilePath).Debug("Deleted file was not found, ignoring")
		return
	}
	f.DeletionDate = &dateTime
	f.AuthorCommits[authorID]++
}
