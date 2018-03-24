package cmd

import (
	"codexray/cxdig/core"
	"codexray/cxdig/core/progress"
	"codexray/cxdig/output"
	"codexray/cxdig/repos"
	"codexray/cxdig/repos/vcs"
	"codexray/cxdig/types"
	"errors"

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

	path, err := getRepositoryPathFromCmdArgs(args)
	core.DieOnError(err)

	err = execOpts.checkFlagCombination()
	core.DieOnError(err)

	repo, err := vcs.OpenRepository(path)
	core.DieOnError(err)

	rate, err := repos.DecodeSamplingRate(execOpts.rate)
	core.DieOnError(err)

	if !execOpts.force {
		err = repo.CheckIgnoredFilesExistence()
		core.DieOnError(err)
	}

	existCommitsFile, err := output.CheckFileExistence(repo, "commits.json")
	core.DieOnError(err)

	if !existCommitsFile {
		core.Info("Scanning repository...")
		cmdScanProject(cmd, args)
	}

	var commits []types.CommitInfo
	err = output.ReadJSONFile(repo, "commits.json", &commits)
	core.DieOnError(err)
	commits = repos.SortCommitByDateDecr(commits)

	if execOpts.input == "" {
		exist, err := output.CheckFileExistence(repo, "samples."+rate.String()+".json")
		core.DieOnError(err)
		if !exist {
			repo.ConstructSampleList(rate, commits, execOpts.limit, execOpts.output)
			core.DieOnError(err)
		}
	} else {
		if exist, _ := output.CheckFileExistence(repo, execOpts.input); !exist {
			core.DieOnError(errors.New("the file given in input doesn't exists"))
		}
	}

	if commits != nil && execOpts.cmd != "" {
		pb := &progress.ProgressBar{}
		execOpts.input = execOpts.output
		tool := repos.NewExternalTool(execOpts.cmd)
		err = repo.SampleWithCmd(tool, rate, commits, execOpts.input, pb)
		core.DieOnError(err)
	}
}

func init() {
	sampleCmd.Flags().IntVarP(&execOpts.limit, "limit", "l", 0, "Maximum number of samples to process")
	sampleCmd.Flags().StringVarP(&execOpts.rate, "rate", "r", defaultSamplingRate, "Time difference between two samples (10c, 2d, 1m, 3y, etc.)")
	sampleCmd.Flags().StringVarP(&execOpts.cmd, "cmd", "c", "", "External command to be executed for each sample")
	sampleCmd.Flags().StringVarP(&execOpts.input, "input", "i", "", "Existing sample file to be reused rather than generating a new sampling list")
	sampleCmd.Flags().StringVarP(&execOpts.output, "output", "o", "", "Save the generated sampling list with the given name")
	sampleCmd.Flags().BoolVarP(&execOpts.force, "force", "f", false, "Delete local changes in the repository if required")
}
