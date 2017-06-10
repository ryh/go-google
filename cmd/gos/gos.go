package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/ryh/gosearch"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		q := strings.Join(os.Args[1:], " ")
		fmt.Println("Go Searching %s", q)

		client := gosearch.NewClient()
		fmt.Print(client.FormatURL(q, 20, 0))
		if os.Getenv("https_proxy") != "" {
			fmt.Print("https proxy used")
			client.SetProxy(os.Getenv("https_proxy"))
		}

		googleResponse, err := client.Search(q)
		if err != nil {
			fmt.Println(err)
		}
		for _, result := range googleResponse {
			fmt.Println()
			color.Cyan(result.Name)
			fmt.Println(result.Desc)
			color.Green(result.Link)
		}

	} else {

		fmt.Println(`
        https://github.com/ryh/gosearch

        Searching Something? 
        `)
	}
}
