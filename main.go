package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type TitleResponse struct {
	Hits []Items	`json:"hits"`
}

type Items struct {
	Title string `json:"title"`
	Url string `json:"url"`
}
func main() {
	response, err := http.Get("https://hn.algolia.com/api/v1/search?tags=front_page&hitsPerPage=30&page=0")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}
	var returnData TitleResponse

	json.Unmarshal(responseData, &returnData)
	for i := 0; i < len(returnData.Hits); i++ {
		fmt.Printf("items %s %s\n", returnData.Hits[i].Title, returnData.Hits[i].Url)
	}
}