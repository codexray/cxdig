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
	Short: "Repeated source code analysis over time",
	Long:  "Run a sampling tool on the source code at different points in time (sampling frequency)",
	RunE:  cmdSample,
}

type execOptions struct {
	limit  int
	freq   string
	cmd    string
	input  string
	output string
}

var execOpts execOptions

func cmdSample(cmd *cobra.Command, args []string) error {

	path, err := getRepositoryPathFromCmdArgs(args)
	if err != nil {
		return err
	}
	err = checkFlagInteracting()
	if err != nil {
		return err
	}

	repo, err := vcs.OpenRepository(path)
	if err != nil {
		core.Error(err)
		return nil
	}

	freq, err := repos.DecodeSamplingFreq(execOpts.freq)
	if err != nil {
		core.Error(err)
		return nil
	}

	tool := repos.NewExternalTool(execOpts.cmd)

	core.Info("Scanning repository")
	existCommitsFile, err := output.CheckFileExistence(repo, "commits.json")
	if err != nil {
		core.Error(err)
		return nil
	}
	if !existCommitsFile {
		cmdScanProject(cmd, args)
	}
	var commits []types.CommitInfo
	if err = output.ReadJSONFile(repo, "commits.json", &commits); err != nil {
		core.Error(err)
		return nil
	}
	commits = repos.SortCommitByDateDecr(commits)

	if execOpts.input == "" {
		exist, err := output.CheckFileExistence(repo, "sample.json")
		if err != nil {
			core.Error(err)
			return nil
		}
		if !exist {
			err = repo.ConstructSampleList(freq, commits, execOpts.limit, execOpts.output)
			if err != nil {
				core.Error(err)
				return nil
			}
		}
	} else {
		exist, _ := output.CheckFileExistence(repo, execOpts.input)
		if !exist {
			core.Error(errors.New("the file given in input doesn't exists"))
			return nil
		}
	}

	if commits != nil && execOpts.cmd != "" {
		pb := &progress.ProgressBar{}
		execOpts.input = execOpts.output
		err = repo.SampleWithCmd(tool, commits, execOpts.input, pb)
		if err != nil {
			core.Error(err)
		}
	}
	return nil
}

func checkFlagInteracting() error {
	if execOpts.input != "" && execOpts.cmd == "" {
		return errors.New("-i/--input flag cannot be used without -c/--cmd flag")
	}
	if execOpts.output != "" && execOpts.input != "" {
		return errors.New("-i/--input flag cannot be used with -o/--output flag")
	}
	return nil
}

func init() {
	sampleCmd.Flags().IntVarP(&execOpts.limit, "limit", "l", 0, "Set the number of commits used")
	sampleCmd.Flags().StringVarP(&execOpts.freq, "freq", "f", "1w", "Set the frequence separating the commits treated (must be of the form : 10c, 2d, 1m, 3y, etc.")
	sampleCmd.Flags().StringVarP(&execOpts.cmd, "cmd", "c", "", "Command to be executed for each sample (default give just the list of the commits'sha for the freq given")
	sampleCmd.Flags().StringVarP(&execOpts.input, "input", "i", "", "Specify an sample file to be load in place of generate it, must be combined with -c")
	sampleCmd.Flags().StringVarP(&execOpts.output, "output", "o", "", "Specify the name for the generated sample file, cannot be combined with -i")
}
