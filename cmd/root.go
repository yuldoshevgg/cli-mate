package cmd

import (
	"cli-mate/server/config"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	version string = "1.0.0"
	Debug   bool
	rootCmd = cobra.Command{
		Use:     "cli-mate",
		Short:   "cli-mate - Simple cli weather application",
		Version: version,
		PreRun:  toggleDebug,
		Example: "cli-mate",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 && args[0] == "version" {
				fmt.Printf("cli-mate %vs\n", version)
			}
		},
	}
)

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Flags().BoolVarP(&Debug, "debug", "d", false, "verbose logging")
}

func Execute() {
	var (
		cfg config.Config
	)

	if err := cfg.Load(); err != nil {
		log.Fatalf("failed to load cfg: %v", err)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
