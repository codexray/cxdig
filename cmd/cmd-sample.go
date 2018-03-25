package cmd

import (
	"codexray/cxdig/core"
	"codexray/cxdig/core/progress"
	"codexray/cxdig/output"
	"codexray/cxdig/repos"
	"codexray/cxdig/repos/vcs"
	"codexray/cxdig/types"
	"fmt"

	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

var sampleCmd = &cobra.Command{
	Use:   "sample",
	Short: "Run a sampling opoeration at a given rate on the repository",
	Long:  "Run a sampling tool on a repository at regular time points in its history (sampling rate)",
	Run:   cmdSample,
}

const defaultSamplingRate = "1w"

type execOptions struct {
	limit  int
	rate   string
	cmd    string
	input  string
	output string
	force  bool
}

func (opts *execOptions) checkFlagCombination() error {
	if opts.input != "" {
		if opts.cmd == "" {
			return errors.New("--input must be used in combination with --cmd")
		}
		if opts.rate != defaultSamplingRate {
			return errors.New("--input cannot be used in combination with --rate")
		}
		if opts.output != "" {
			return errors.New("--input and --output cannot be mixed together")
		}
	}
	return nil
}

var execOpts execOptions

func cmdSample(cmd *cobra.Command, args []string) {
	// TODO: group in a dedicated function
	err := execOpts.checkFlagCombination()
	core.DieOnError(err)
	rate, err := repos.DecodeSamplingRate(execOpts.rate)
	core.DieOnError(err)
	path, err := getRepositoryPathFromCmdArgs(args)
	core.DieOnError(err)

	repo, err := vcs.OpenRepository(path)
	core.DieOnError(err)

	hasLocalModif, err := repo.HasLocalModifications()
	core.DieOnError(err)
	if hasLocalModif {
		if execOpts.force {
			core.Info("Repository has local changes that will be deleted (force mode enabled)")
		} else {
			err = errors.New("repository has local changes that would be destroyed by sampling operation (use --force to ignore)")
			core.DieOnError(err)
		}
	}

	// if the repo scan was not already done, do it now
	if !output.FileExists(repo, "commits.json") {
		err = extractRepoCommitsAndSaveResult(repo)
		core.DieOnError(err)
	}
	commits, err := loadCommitsFromFile(repo, "commits.json")
	core.DieOnError(err)

	// if we are given a sampling list file, use it
	var samples []types.SampleInfo
	if execOpts.input != "" {
		inputFile := execOpts.input
		//if !output.FileExists(repo, inputFile) {
		//	core.DieOnError(errors.New("the file given in input doesn't exists"))
		//}
		core.Infof("Loading sampling list from file '%s'", inputFile)
		err = output.ReadJSONFile(repo, inputFile, &samples)
		core.DieOnError(errors.Wrap(err, "failed to load sample file"))
	} else {
		// generate default file name if nothing specified
		outputName := execOpts.output
		if outputName == "" {
			outputName = fmt.Sprintf("samples.%s.json", rate.String())
		}

		core.Infof("Saving sampling list to file '%s'", outputName)
		samples := repos.FilterCommitsByStep(commits, rate, execOpts.limit)
		if len(samples) == 0 {
			core.Warn("the generated sample list is empty")
		}

		err := output.WriteJSONFile(repo, outputName, samples)
		core.DieOnError(errors.Wrap(err, "failed to save sampling list"))
	}

	// now that we have the sampling list, run the sampling tool
	if execOpts.cmd != "" {
		pb := &progress.ProgressBar{}
		tool := repos.NewExternalTool(execOpts.cmd)
		core.Info("Sampling repository...")
		err = repo.SampleWithCmd(tool, rate, commits, samples, pb)
		core.DieOnError(err)
	}
}

func loadCommitsFromFile(repo repos.Repository, fileName string) ([]types.CommitInfo, error) {
	var commits []types.CommitInfo
	if err := output.ReadJSONFile(repo, fileName, &commits); err != nil {
		return nil, err
	}
	return repos.SortCommitByDateDecr(commits), nil
}

func init() {
	sampleCmd.Flags().IntVarP(&execOpts.limit, "limit", "l", 0, "Maximum number of samples to generate")
	sampleCmd.Flags().StringVarP(&execOpts.rate, "rate", "r", defaultSamplingRate, "Time difference between two samples (10c, 2d, 1m, 3y, etc.)")
	sampleCmd.Flags().StringVarP(&execOpts.cmd, "cmd", "c", "", "External command to be executed for each sample")
	sampleCmd.Flags().StringVarP(&execOpts.input, "input", "i", "", "Existing sample file to be reused rather than generating a new sampling list")
	sampleCmd.Flags().StringVarP(&execOpts.output, "output", "o", "", "Save the generated sampling list with the given name")
	sampleCmd.Flags().BoolVarP(&execOpts.force, "force", "f", false, "Delete local changes in the repository if required")
}
