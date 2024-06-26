package main

import (
	"os"
	"path"

	"github.com/kefniark/mango/pkg/mango-cli/config"
	"github.com/kefniark/mango/pkg/mango-cli/generate"
	"github.com/kefniark/mango/pkg/mango-cli/prepare"
	"github.com/spf13/cobra"
)

var preparer = []config.Executer{}
var generater = []config.Executer{}

func initExec(cfg *config.Config) {
	// preparer
	preparer = append(preparer, prepare.AirPrepare{Config: cfg})
	preparer = append(preparer, prepare.DBPrepare{Config: cfg})
	preparer = append(preparer, prepare.NodeJSPrepare{})
	preparer = append(preparer, prepare.OpenAPIPrepare{Config: cfg})
	preparer = append(preparer, prepare.SQLCPrepare{Config: cfg})
	preparer = append(preparer, prepare.StaticBuildPrepare{Config: cfg})
	preparer = append(preparer, prepare.AssetsFilePrepare{})

	// generater
	generater = append(generater, generate.ProtoGenerator{})
	generater = append(generater, generate.SQLCGenerator{})
	generater = append(generater, generate.MarkdownGenerator{})
	generater = append(generater, generate.TemplGenerator{})
}

func main() {
	cfg := config.Parse()
	initExec(cfg)

	rootCmd := &cobra.Command{
		Use:   "mango",
		Short: "Mango is a Web Development Framework for Go.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
	}

	rootCmd.AddCommand(buildCmd(cfg))
	rootCmd.AddCommand(cleanCmd(cfg))
	rootCmd.AddCommand(devCmd(cfg))
	rootCmd.AddCommand(prepareCmd(cfg))
	rootCmd.AddCommand(generateCmd(cfg))
	rootCmd.AddCommand(lintCmd(cfg))
	rootCmd.AddCommand(formatCmd(cfg))

	if err := rootCmd.Execute(); err != nil {
		config.Logger.Err(err).Msg("Cannot execute command")
		os.Exit(1)
	}
}

func checkAppReady(name string) {
	if _, err := os.Stat(path.Join(name, ".mango")); err != nil {
		config.Logger.Warn().Msg("Need preparation first, wait a moment ...")
		for _, exec := range preparer {
			config.Logger.Debug().Str("app", name).Str("exec", exec.Name()).Msg("Start")
			if err := exec.Execute(name); err != nil {
				config.Logger.Err(err).Msg("Prepare Failed")
			}
			config.Logger.Debug().Str("app", name).Str("exec", exec.Name()).Msg("End")
		}

		for _, exec := range generater {
			config.Logger.Debug().Str("app", name).Str("exec", exec.Name()).Msg("Start")
			if err := exec.Execute(name); err != nil {
				config.Logger.Err(err).Msg("Generate Failed")
			}
			config.Logger.Debug().Str("app", name).Str("exec", exec.Name()).Msg("End")
		}
	}
}

func checkAppPrepared(name string) {
	if _, err := os.Stat(path.Join(name, ".mango")); err != nil {
		config.Logger.Warn().Msg("Need preparation first, wait a moment ...")
		for _, exec := range preparer {
			config.Logger.Debug().Str("app", name).Str("exec", exec.Name()).Msg("Start")
			if err := exec.Execute(name); err != nil {
				config.Logger.Err(err).Msg("Prepare Failed")
			}
			config.Logger.Debug().Str("app", name).Str("exec", exec.Name()).Msg("End")
		}
	}
}
