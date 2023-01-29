package main

import (
	"os"

	indexer "github.com/Wiliamfm/ZincSearch_Demo/utils"
)

func main() {
	path := os.Args[1]
	indexer.SetEmails(path)
}
