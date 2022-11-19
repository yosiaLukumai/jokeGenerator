package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type OutputInterface interface {
	PrintOutput()
}

// struct for decoding the json
type JokeStruct struct {
	Category string `json:"category"`
	Error    bool   `json:"error"`
	Id       int    `json:"id"`
	Joke     string `json:"joke"`
	Safe     bool   `json:"safe"`
}

type JokeStructTwopart struct {
	Category string `json:"category"`
	Error    bool   `json:"error"`
	Id       int    `json:"id"`
	Setup    string `json:"setup"`
	Delivery string `json:"delivery"`
	Safe     bool   `json:"safe"`
}

type ArrayJokes struct {
	Error  bool         `json:"error"`
	Amount int          `json:"amount"`
	Jokes  []JokeStruct `json:"jokes"`
}

type ArrayJokesTwoParts struct {
	Error  bool                `json:"error"`
	Amount int                 `json:"amount"`
	Jokes  []JokeStructTwopart `json:"jokes"`
}

// handling output err
func (joke JokeStructTwopart) PrintOutput() {
	if !joke.Error {
		fmt.Printf("Id: %v,\nCategory:  %v, \nSetup: %v,\nDelivery: %v\n", joke.Id, joke.Category, joke.Setup, joke.Delivery)
	} else {
		fmt.Println("( ᵔ ͜ʖ ᵔ ) No such filter for the given joke")
	}
}

func (joke JokeStruct) PrintOutput() {
	if !joke.Error {
		fmt.Printf("Id: %v,\nCategory:  %v,\nJoke: %v \n", joke.Id, joke.Category, joke.Joke)
	} else {
		fmt.Println("( ❛ ͜ʖ ❛ )No such filter for the given joke")
	}
}

func printerOfOutput(i OutputInterface) {
	i.PrintOutput()
}

func init() {
	// joke.Flags().BoolP("jokeCount", "c", false, "How many joke counts can be pooled")
	joke.Flags().IntP("amount", "a", 1, "Specify the number of the jokes to be pooled..")
	joke.Flags().BoolP("type", "t", false, "A joke type such as twopart or single-part is false")
	joke.Flags().StringP("category", "c", "any", "A joke category can be like programming")
	rootCmd.AddCommand(joke)
}

var joke = &cobra.Command{
	Use:   "joke",
	Short: "Get a random Joke",
	Long:  `Getting you a random joke of any category`,
	Run: func(cmd *cobra.Command, args []string) {
		counts, _ := cmd.Flags().GetInt("amount")
		category, _ := cmd.Flags().GetString("category")
		ty, _ := cmd.Flags().GetBool("type")
		url := formUrl(category, counts, ty)
		fmt.Printf("--> Api: %v \n \n", url)
		b := GetJoke(url)
		var jokeSingle JokeStruct
		var jokeTwopart JokeStructTwopart
		if counts == 1 {
			if ty {
				err := json.Unmarshal(b, &jokeTwopart)
				HandleError(err, "Failed to decode the json object")
				printerOfOutput(jokeTwopart)
			}
			if !ty {
				err := json.Unmarshal(b, &jokeSingle)
				HandleError(err, "Failed to decode the json object")
				printerOfOutput(jokeSingle)
			}

		} else {
			if !ty {
				var arrayjokeSingle ArrayJokes
				err := json.Unmarshal(b, &arrayjokeSingle)
				HandleError(err, "Failed to decode the json object")
				for i, e := range arrayjokeSingle.Jokes {
					fmt.Printf("--------Joke: %v --------\n", i+1)
					printerOfOutput(e)
					fmt.Println()
				}
			}
			if ty {
				var arrayjokeTwopart ArrayJokesTwoParts
				err := json.Unmarshal(b, &arrayjokeTwopart)
				HandleError(err, "Failed to decode the json object")
				for i, e := range arrayjokeTwopart.Jokes {
					fmt.Printf("-------Joke: %v --------\n", i+1)
					printerOfOutput(e)
					fmt.Println()
				}
			}
		}
	},
}

func formUrl(category string, jokeCount int, jokeType bool) string {
	// using  the params provided we create the url
	// JokeApiUrl = JokeApiUrl + strconv.Itoa(jokeCount) + "/" + category
	if category == "any" {
		// no modificarion has been done on the category
		return JokeApiUrl + category + createJokeType(jokeType) + createAmount(jokeCount)
	}
	for _, t := range JokesCategories {
		if strings.EqualFold(t, category) {
			return JokeApiUrl + t + createJokeType(jokeType) + createAmount(jokeCount)
		}
	}
	fmt.Println("---> The flag for type joke is't correct")
	fmt.Println("---> Default any category will be used")
	return JokeApiUrl + "Any" + createJokeType(jokeType) + createAmount(jokeCount)
}

func createAmount(jokeCount int) string {
	if jokeCount == 1 {
		return ""
	} else {
		return "&amount=" + strconv.Itoa(jokeCount)
	}
}
func createJokeType(jokeType bool) string {
	if jokeType {
		return "?type=twopart"
	}
	return "?type=single"
}

func GetJoke(url string) []byte {
	resp, err := http.Get(url)
	HandleError(err, "Failed to fetch the data")
	byte, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "Failed to parse the output")
	return byte
}

func HandleError(err error, message string) {
	if err != nil {
		log.Fatal(message)
	}
}
