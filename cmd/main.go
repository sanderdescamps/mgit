package cmd

func Run() {
	main()
}

// func printMapRecursive(m map[string]interface{}, level int) {
// 	for k, v := range m {
// 		for i := 0; i < level; i++ {
// 			fmt.Print("\t")
// 		}
// 		fmt.Printf("%v: ", k)
// 		switch v := v.(type) {
// 		case map[string]interface{}:
// 			fmt.Println()
// 			printMapRecursive(v, level+1)
// 		default:
// 			fmt.Println(v)
// 		}
// 	}
// }

// func Run() {
// 	Init()
// 	paths := []string{"/home/sander/git/mgit/config.yaml", "/home/sander/git/mgit/repos.yaml"}
// 	config, err := config.ReadConfig(paths...)
// 	if err != nil {
// 		display.Errorf("%v", err)
// 		panic(err)
// 	}

// 	for _, r := range config.RepoList() {
// 		display.Print("Repo %s", r.GitUrl)
// 		if _, err := os.Stat(r.PrettyLocalPath()); errors.Is(err, os.ErrNotExist) {
// 			display.Debugf("create directory: %s", r.PrettyLocalPath())
// 			err := os.MkdirAll(r.PrettyLocalPath(), os.ModePerm)
// 			if err != nil {
// 				display.Warningf("failed to create directory: %s", r.PrettyLocalPath())
// 				continue
// 			}
// 		} else if err != nil {
// 			display.Warningf("unable to create directory: %s", r.PrettyLocalPath())
// 			continue
// 		} else {
// 			display.Debugf("directory already exists: %s", r.PrettyLocalPath())
// 		}
// 		if r.Clone {
// 			repo := repo.Repo{
// 				Url:      r.GitUrl,
// 				RepoPath: r.PrettyLocalPath(),
// 				Display:  display,
// 			}
// 			repo.Clone()
// 		} else {
// 			display.Skipf("skip repo: %s", r.GitUrl)
// 		}
// 	}
// }

// func Run() {
// 	o := Display{}

// 	o.Error("error test message")
// 	o.Warning("warning test message")
// 	o.Info("info test message")
// 	o.Change("change test message")
// 	o.Ok("ok test message")
// 	o.Skip("skip test message")

// }
