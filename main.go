package main

import (
	"fmt"
	"os"

	linkparser "github.com/sarathkumar17/linkparser/cmd"
)

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	nodes, err := linkparser.Parse(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, node := range nodes {
		fmt.Println(node.Href, node.Text)
	}
}
