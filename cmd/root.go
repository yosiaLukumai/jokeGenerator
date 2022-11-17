package cmd

import (
	"github.com/spf13/cobra"
)

var JokeApiUrl = "https://v2.jokeapi.dev/joke/"
var JokesCategories = []string{"Programming", "Miscellaneous", "Dark", "Christmas", "Pun"}

var (
	rootCmd = &cobra.Command{
		Use:   "jokeGenerator",
		Short: "A simple joke generator that provide with category of the joke you want",
		Long: `Get any sort of the Joke one can eve provide the flag for category type such as programming
		or Misc or Christmas or a random category or even specifying the amount of jokes one wants`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
