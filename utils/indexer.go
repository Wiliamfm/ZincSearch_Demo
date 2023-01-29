package indexer

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Wiliamfm/ZincSearch_Demo/models"
)

func SetEmails(path string) {
	email := models.Email{}
	fmt.Println(email)
	err := filepath.WalkDir(path, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
			return err
		}
		if info.IsDir() && info.Name() == "maildir" {
			fmt.Println(info.Name())
			return nil
		}
		//fmt.Println(info.Name())
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	/*
		files, err := os.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}
		for _, elem := range files {
			if elem.IsDir() && elem.Name() == "maildir" {
				files2, err := os.ReadDir(elem.Name())
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(files2)
			}
		}
	*/

}
