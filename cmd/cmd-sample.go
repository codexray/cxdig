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

const defaultSamplingFreq = "1w"

type execOptions struct {
	limit  int
	freq   string
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
		if opts.freq != defaultSamplingFreq {
			return errors.New("--input cannot be used in combination with --freq")
		}
		if opts.output != "" {
			return errors.New("--input and --output cannot be mixed together")
		}
	}
	return nil
}

var execOpts execOptions

func cmdSample(cmd *cobra.Command, args []string) error {

	path, err := getRepositoryPathFromCmdArgs(args)
	if err != nil {
		return err
	}

	if err = execOpts.checkFlagCombination(); err != nil {
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

	if !execOpts.force {
		err = repo.CheckIgnoredFilesExistence()
		if err != nil {
			core.Error(err)
			return nil
		}
	}

	tool := repos.NewExternalTool(execOpts.cmd)

	existCommitsFile, err := output.CheckFileExistence(repo, "commits.json")
	if err != nil {
		core.Error(err)
		return nil
	}
	if !existCommitsFile {
		core.Info("Scanning repository...")
		cmdScanProject(cmd, args)
	}
	var commits []types.CommitInfo
	if err = output.ReadJSONFile(repo, "commits.json", &commits); err != nil {
		core.Error(err)
		return nil
	}
	commits = repos.SortCommitByDateDecr(commits)

	if execOpts.input == "" {
		exist, err := output.CheckFileExistence(repo, "samples."+freq.String()+".json")
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
		err = repo.SampleWithCmd(tool, freq, commits, execOpts.input, pb)
		if err != nil {
			core.Error(err)
		}
	}
	return nil
}

func init() {
	sampleCmd.Flags().IntVarP(&execOpts.limit, "limit", "l", 0, "Set the number of commits used")
	sampleCmd.Flags().StringVarP(&execOpts.freq, "freq", "f", defaultSamplingFreq, "Set the frequence separating the commits treated (must be of the form : 10c, 2d, 1m, 3y, etc.")
	sampleCmd.Flags().StringVarP(&execOpts.cmd, "cmd", "c", "", "Command to be executed for each sample (default give just the list of the commits'sha for the freq given")
	sampleCmd.Flags().StringVarP(&execOpts.input, "input", "i", "", "Existing sample file to be used rather than generating a new sampling list")
	sampleCmd.Flags().StringVarP(&execOpts.output, "output", "o", "", "Save the generated sampling list with the given name")
	sampleCmd.Flags().BoolVarP(&execOpts.force, "force", "f", false, "Force the deletion of git ignored files")
}
