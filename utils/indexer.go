package indexer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/Wiliamfm/ZincSearch_Demo/models"
)

func SetEmails(path string) models.Emails {
	emails := models.Emails{}
	err := filepath.WalkDir(path, func(path string, info os.DirEntry, err error) error {
		email := models.Email{MailFolders: map[string][]models.File{}}
		path = filepath.ToSlash(path)
		basePath := filepath.Base(path)
		if err != nil {
			return err
		}
		if m, err := regexp.MatchString(".+maildir$", filepath.Dir(path)); m {
			if err != nil {
				return err
			}
			email.Username = basePath
			emails.Emails = append(emails.Emails, email)
		}
		for i, email := range emails.Emails {
			if m, err := regexp.MatchString(email.Username, path); m && !info.IsDir() && basePath != email.Username {
				folder := filepath.Dir(path)
				if err != nil {
					return err
				}
				if _, ok := email.MailFolders[folder]; !ok {
					email.MailFolders[folder] = make([]models.File, 0)
				}
				file := setFile(path)
				email.MailFolders[folder] = append(email.MailFolders[folder], file)
				emails.Emails[i] = email
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return emails
}

func SetToJson(emails models.Emails) []byte {
	emailsJson, err := json.Marshal(emails)
	if err != nil {
		log.Fatal(err)
	}
	return emailsJson
}

func setFile(path string) models.File {
	file, err := os.Stat(path)
	data, err2 := os.ReadFile(path)
	if err != nil || file.IsDir() || err2 != nil {
		log.Fatal(err)
	}
	return models.File{FileName: file.Name(), Content: string(data)}
}

func LoadData(body []byte) bool {
	responseBody := bytes.NewBuffer(body)
	//resp, err := http.Post("http://localhost:4080/api/_bulkV2", "application/json", responseBody)
	resp, err := http.Post("http://localhost:4080/api/test/_doc", "application/json", responseBody)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	return true
}
