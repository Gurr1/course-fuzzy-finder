package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("enter a website to scrape")
	reader := bufio.NewReader(os.Stdin)
	url, err := reader.ReadString('\n')
	url = strings.Replace(url, " ", "", -1)
	url = strings.Replace(url, "\n", "", -1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(url)
	fmt.Println("started scraping " + url)
	//ScrapeWebsite(url)
	split := strings.SplitAfter(url, "://")
	root := directory + "/" + split[1]
	fmt.Println(root)
	ReadSubfiles(strings.TrimSuffix(root, "/"))

}
