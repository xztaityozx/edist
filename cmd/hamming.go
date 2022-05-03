package cmd

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"xztaityozx/edist/lib"
)

var hammingCmd = &cobra.Command{
	Use:     "hamming",
	Aliases: []string{"H"},
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		result, err := lib.NewHammingDistance().Walk(args[0], args[1])
		if err != nil {
			log.Fatal().Msgf("%v", err)
		}

		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(hammingCmd)
}
