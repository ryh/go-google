package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jessevdk/go-flags"
	"github.com/ryh/gosearch"
	"os"
	"strings"
)

func main() {

	var opts struct {
		// Slice of bool will append 'true' each time the option
		// is encountered (can be set multiple times, like -vvv)
		Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`

		// Example of automatic marshalling to desired type (uint)
		Offset uint `long:"offset" description:"Search from Offset"`
	}

	args, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		os.Exit(1)

	}

	if len(args) > 1 {
		q := strings.Join(args[1:], " ")
		fmt.Println("Go Searching :", q)
		client := gosearch.NewClient()
		if len(opts.Verbose) > 0 {
			fmt.Print(client.FormatURL(q, 10, 0))
		}
		if os.Getenv("https_proxy") != "" {
			if len(opts.Verbose) > 0 {
				fmt.Println("https proxy used")
			}
			client.SetProxy(os.Getenv("https_proxy"))
		}

		googleResponse, err := client.Search(q)
		if err != nil {
			fmt.Println("error: ", err)
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

        gos -h for help
        `)
	}
}
