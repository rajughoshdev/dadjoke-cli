/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a random dad joke",
	Long:  `The get command fetches a random dad joke from the icanhazdadjoke api`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

type Joke struct {
	Id     string `json:"id"`
	Joke   string `json: "joke"`
	Status int    `json: "status"`
}

func getRandomJoke() {
	viper.AddConfigPath("config/")
	viper.SetConfigName("api")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	apiUrl := viper.GetString("apiUrl")
	resBytes := getJokeData(apiUrl)
	dadJoke := Joke{}
	if err := json.Unmarshal(resBytes, &dadJoke); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dadJoke.Joke))
}

func getJokeData(apiUrl string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		apiUrl,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "DadJoke Cli for learning purpose")

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	resBytes, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	return resBytes
}
