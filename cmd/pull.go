package cmd

// var cmdPull = &cobra.Command{
// 	Use:   "pull [string to print]",
// 	Short: "Pull all repositories",
// 	Long:  `iterate over all the repositories and pull the latest changes`,
// 	Args:  cobra.MinimumNArgs(1),
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("Print: " + strings.Join(args, " "))
// 	},
// }

// var cmdClone = &cobra.Command{
// 	Use:   "clone",
// 	Short: "Clone all repositories",
// 	Long:  `Create the local directory structure and clone all repositories`,
// 	Args:  cobra.MinimumNArgs(1),
// 	Run: func(cmd *cobra.Command, args []string) {
// 		paths := []string{"/home/sander/git/mgit/config.yaml", "/home/sander/git/mgit/repos.yaml"}
// 		config, err := config.ReadConfig(paths...)
// 		if err != nil {
// 			display.Errorf("%v", err)
// 			panic(err)
// 		}

// 		for _, r := range config.RepoList() {
// 			display.Print("Repo %s", r.GitUrl)
// 			if _, err := os.Stat(r.PrettyLocalPath()); errors.Is(err, os.ErrNotExist) {
// 				display.Debugf("create directory: %s", r.PrettyLocalPath())
// 				err := os.MkdirAll(r.PrettyLocalPath(), os.ModePerm)
// 				if err != nil {
// 					display.Warningf("failed to create directory: %s", r.PrettyLocalPath())
// 					continue
// 				}
// 			} else if err != nil {
// 				display.Warningf("unable to create directory: %s", r.PrettyLocalPath())
// 				continue
// 			} else {
// 				display.Debugf("directory already exists: %s", r.PrettyLocalPath())
// 			}
// 			if r.Clone {
// 				repo := repo.Repo{
// 					Url:      r.GitUrl,
// 					RepoPath: r.PrettyLocalPath(),
// 					Display:  display,
// 				}
// 				repo.Clone()
// 			} else {
// 				display.Skipf("skip repo: %s", r.GitUrl)
// 			}
// 		}
// 	},
// }
