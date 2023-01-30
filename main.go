package main

import (
	"fmt"
	"os"

	indexer "github.com/Wiliamfm/ZincSearch_Demo/utils"
)

func main() {
	path := os.Args[1]
	emails := indexer.SetEmails(path)
	//emailsJson := indexer.SetToJson(emails)
	for _, email := range emails.Emails {
		fmt.Println(email.Username)
		for k, v := range email.MailFolders {
			fmt.Println("Folder: ", k)
			for _, file := range v {
				fmt.Println("Files: \t", file.FileName)
			}
		}
	}
	//s := indexer.Indexer(emailsJson)
	//if s {
	//fmt.Println("Success")
	//}
}
