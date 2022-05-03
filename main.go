package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/sheepla/qiitaz/client"
	"github.com/sheepla/qiitaz/ui"
	"github.com/toqueteos/webbrowser"
)

const (
	appName    = "qiitaz"
	appVersion = "0.1.0"
	appUsage   = "[OPTIONS] QUERY..."
)

type exitCode int

type options struct {
	Version bool   `short:"V" long:"version" description:"Show version"`
	Sort    string `short:"s" long:"sort" description:"Sort key to search e.g. \"created\", \"like\", \"stock\", \"rel\",  (default: \"rel\")" `
	Open    bool   `short:"o" long:"open" description:"Open URL in your web browser"`
	Preview bool   `short:"p" long:"preview" description:"Preview page on your terminal"`
	PageNo  int    `short:"n" long:"pageno" description:"Max page number of search page" default:"1"`
	Json    bool   `short:"j" long:"json" description:"Output result in JSON format"`
}

const (
	exitCodeOK exitCode = iota
	exitCodeErrArgs
	exitCodeErrRequest
	exitCodeErrFuzzyFinder
	exitCodeErrWebbrowser
	exitCodeErrPreview
)

func main() {
	os.Exit(int(Main(os.Args[1:])))
}

func Main(cliArgs []string) exitCode {
	var opts options
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = appName
	parser.Usage = appUsage

	args, err := parser.ParseArgs(cliArgs)
	if err != nil {
		if flags.WroteHelp(err) {
			return exitCodeOK
		} else {
			fmt.Fprintf(os.Stderr, "Argument parsing failed: %s", err)
			return exitCodeErrArgs
		}
	}

	if opts.Version {
		fmt.Printf("%s: v%s\n", appName, appVersion)
		return exitCodeOK
	}

	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Must require argument (s)")
		return exitCodeErrArgs
	}

	if opts.PageNo <= 0 {
		fmt.Fprintln(os.Stderr, "The page number must be a positive value.")
		return exitCodeErrArgs
	}

	var urls []string
	for i := 1; i <= opts.PageNo; i++ {
		u, err := client.NewSearchURL(strings.Join(args, " "), client.SortBy(opts.Sort), i)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return exitCodeErrArgs
		}
		urls = append(urls, u)
	}

	var results []client.Result
	for _, u := range urls {
		r, err := client.Search(u)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return exitCodeErrRequest
		}
		results = append(results, r...)
	}

	if len(results) == 0 {
		fmt.Fprintln(os.Stderr, "No results found.")
		return exitCodeOK
	}

	if opts.Json {
		bytes, err := json.Marshal(&results)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		stdout := bufio.NewWriter(os.Stdout)
		fmt.Fprintln(stdout, string(bytes))
		stdout.Flush()
		return exitCodeOK
	}

	choices, err := ui.Find(results)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitCodeErrFuzzyFinder
	}

	if opts.Open {
		for _, idx := range choices {
			url := client.NewPageURL(results[idx].Link)
			if err := webbrowser.Open(url); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return exitCodeErrWebbrowser
			}
		}
	}

	if opts.Preview {
		for _, idx := range choices {
			url := client.NewPageURL((results[idx].Link + ".md"))
			title := results[idx].Title
			if err := ui.Preview(url, title); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return exitCodeErrPreview
			}
		}
	}

	for _, idx := range choices {
		fmt.Println(client.NewPageURL(results[idx].Link))
	}

	return exitCodeOK
}
