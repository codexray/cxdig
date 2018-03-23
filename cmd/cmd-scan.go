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
	Short: "Collect data from a given repository",
	Long:  "Scan a given repository for its commits and source files",
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
