package gitlog

// TODO: change to get -u command
/*
func cloneGitRepo(args []string) {
	rootClone := ""
	if rootClone != "" {
		if len(args) != 1 {
			logrus.Panic("--clone: path to project is missing")
		}
		rootPath := args[0]
		_, err := os.Stat("/" + filepath.Join(strings.Split(args[0], "/")[:len(strings.Split(rootPath, "/"))-1]...))
		if err != nil {
			logrus.Panic("--clone: path to project is invalid, " + filepath.Join(strings.Split(args[0], "/")[:len(strings.Split(args[0], "/"))-1]...) + " not existing")
		}
		_, err = os.Stat(rootPath)
		if err != nil {
			gitlog.RunGitCommandOnDir(strings.TrimSuffix(args[0], strings.Split(args[0], "/")[len(strings.Split(args[0], "/"))-1]), []string{"clone", rootClone, rootPath}, false)
		} else {
			vcsType, err := vcslog.IdentifyVCSType(rootPath)
			if vcsType == vcslog.GitType && err == nil {
				if !gitlog.CheckGitStatus(rootPath) {
					logrus.Panic("--clone: the existing working tree contains local changes, please commit your changes and retry")
				} else {
					origin := gitlog.RunGitCommandOnDir(rootPath, []string{"config", "--get", "remote.origin.url"}, false)[0]
					if origin != rootClone {
						logrus.Panic("--clone: repository already exists and its origin is different of the clone url given")
					}
					gitlog.RunGitCommandOnDir(rootPath, []string{"pull"}, true)
				}
			} else {
				logrus.Panic("--clone: " + rootPath + " already exists and is not a git repository")
			}
		}
	}
}
*/
