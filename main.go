package main

import (
	"fmt"
	"os"

	"github.com/Wiliamfm/ZincSearch_Demo/models"
	indexer "github.com/Wiliamfm/ZincSearch_Demo/utils"
)

func main() {
	path := os.Args[1] + "/maildir"
	loadData(path)
	/*
		emails := indexer.SetEmails(path)
		if indexer.LoadDataBulkV2(emails, "http://localhost:4080/api/_bulkv2", "admin", "Complexpass#123") {
			fmt.Println("Data loaded")
		}
	*/
}

func loadData(path string) bool {
	emails := indexer.SetEmailsV2(path)
	printEmails(emails)
	return true
}

func printEmails(emails []models.Email) {
	for _, email := range emails {
		fmt.Println(email.Username)
		for k, v := range email.MailFolders {
			fmt.Println("Folder: ", k)
			for _, file := range v {
				fmt.Println("\tFiles:\t", file.FileName)
				//fmt.Println("Content:\n", file.Content)
			}
		}
	}
}
