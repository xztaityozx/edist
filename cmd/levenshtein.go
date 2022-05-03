package cmd

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"xztaityozx/edist/lib"
)

const (
	OptionNameReplaceCost = "replace-cost"

	OptionNameSwapCost = "swap-cost"
	OptionNameLimited  = "limited"
)

var levenshteinCmd = &cobra.Command{
	Use:     "levenshtein",
	Aliases: []string{"L"},
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		a, b := args[0], args[1]
		cost, err := cmd.Flags().GetInt(OptionNameReplaceCost)
		if err != nil {
			log.Fatal().Msgf("%v\n", err)
		}

		result, _ := lib.NewLevenshtein(cost).Walk(a, b)
		if v, _ := cmd.Flags().GetBool(OptionNameNormalize); v {
			fmt.Println(lib.Normalize(a, b, result))
		} else {
			fmt.Println(result)
		}
	},
}

var damerauLevenshteinCmd = &cobra.Command{
	Use:     "damerau-levenshtein",
	Aliases: []string{"DL"},
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		a, b := args[0], args[1]
		replaceCost, err := cmd.Flags().GetInt(OptionNameReplaceCost)
		if err != nil {
			log.Fatal().Msgf("%v\n", err)
		}
		swapCost, err := cmd.Flags().GetInt(OptionNameSwapCost)
		if err != nil {
			log.Fatal().Msgf("%v\n", err)
		}

		limited, err := cmd.Flags().GetBool(OptionNameLimited)
		if err != nil {
			log.Fatal().Msgf("%v\n", err)
		}

		result, _ := lib.NewDamerauLevenshtein(replaceCost, swapCost, limited).Walk(a, b)
		if v, _ := cmd.Flags().GetBool(OptionNameNormalize); v {
			fmt.Println(lib.Normalize(a, b, result))
		} else {
			fmt.Println(result)
		}
	},
}

func init() {
	levenshteinCmd.Flags().IntP(OptionNameReplaceCost, "c", 1, "cost for replace")
	damerauLevenshteinCmd.Flags().IntP(OptionNameReplaceCost, "c", 1, "cost for replace")
	damerauLevenshteinCmd.Flags().IntP(OptionNameSwapCost, "s", 1, "cost for swap")
	damerauLevenshteinCmd.Flags().BoolP(OptionNameLimited, "l", true, "apply limited damerau-levenshtein distance algorithm")

	rootCmd.AddCommand(levenshteinCmd)
	rootCmd.AddCommand(damerauLevenshteinCmd)
}
