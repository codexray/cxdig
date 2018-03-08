package cmd

import (
	"codexray/cxdig/config"
	"codexray/cxdig/core"
	"codexray/cxdig/output"
	"codexray/cxdig/repos"
	"codexray/cxdig/repos/referential"
	"codexray/cxdig/repos/vcs"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Extract information from a repository",
	Long:  "Scan a given repository for its commits and source files.",
	RunE:  cmdScanProject,
}

func cmdScanProject(cmd *cobra.Command, args []string) error {
	path, err := getRepositoryPathFromCmdArgs(args)
	if err != nil {
		return err
	}

	repo, err := vcs.OpenRepository(path)
	if err != nil {
		core.Error(err)
		return nil
	}

	err = extractRepoCommitsAndSaveResult(repo)
	if err != nil {
		core.Error(err)
	}
	return nil
}

func extractRepoCommitsAndSaveResult(repo repos.Repository) error {
	r := config.NewFileTypeRegistry()
	core.Infof("Processing project '%s'...", repo.Name())

	commits, err := repo.ExtractCommits()
	if err != nil {
		return errors.Wrap(err, "failed to extract commits from the repository")
	}

	if err := output.WriteJSONFile(repo, "commits.json", commits); err != nil {
		return errors.Wrap(err, "failed to save commits to JSON file")
	}

	ref := referential.BuildProjectReferential(commits, r)
	core.Infof("Saving results to JSON")
	if err := output.WriteJSONFile(repo, "referential.json", ref); err != nil {
		return errors.Wrap(err, "failed to save referential to JSON file")
	}

	return nil // ok
}

/*
func saveCodeChangesToJSON(name ProjectName, diff []core.CommitChanges) error {
	fileName := fmt.Sprintf("%s.[codechanges].json", name.String())
	logrus.WithField("file-name", fileName).Info("Saving code changes to JSON file")

	return core.WriteJSONFile(fileName, diff)
}

func extractCommitChanges(name ProjectName, repoPath string, commits []types.CommitInfo) {
	const nbCommitsToExtract = 50
	ids := make([]types.CommitID, 0, nbCommitsToExtract)
	for _, c := range commits {
		if len(c.Changes) > 0 {
			ids = append(ids, c.CommitID)
			if len(ids) == nbCommitsToExtract {
				break
			}
		}
	}

	printProgress := func(current int, total int) {
		fmt.Printf("%d / %d\n", current, total)
	}
	diffs := codechanges.ExtractCommitChanges(repoPath, ids, printProgress)

	err := saveCodeChangesToJSON(name, diffs)
	DieOnError(err, "Failed to save code changes to JSON file")
}

func init() {
	scanCmd.PersistentFlags().String("name", "", "Name of the project. Default is repository name put in lower case.")
}
*/
