package main

import (
	"fmt"
	"os"

	indexer "github.com/Wiliamfm/ZincSearch_Demo/utils"
)

func main() {
	path := os.Args[1]
	emails := indexer.SetEmails(path)
	for _, email := range emails.Emails {
		fmt.Println(email.Username, len(email.MailFolders))
		for folder, files := range email.MailFolders {
			fmt.Println("\t", folder)
			for _, file := range files {
				fmt.Println(file.FileName)
			}
		}
	}
}
