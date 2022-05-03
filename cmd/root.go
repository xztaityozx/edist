package cmd

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
)

const (
	OptionNameNormalize = "normalize"
)

var rootCmd = &cobra.Command{
	Use: "edist",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

func Execute() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	rootCmd.PersistentFlags().BoolP(OptionNameNormalize, "n", false, "apply normalization")

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal().Msgf("%v\n", err)
	}
}
