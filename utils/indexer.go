package indexer

import (
	"bytes"
	"encoding/json"
	"io"
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

func LoadDataBulkV2(emails models.Emails, url, username, password string) bool {
	data := models.RequestData{Index: "emails", Records: emails.Emails}
	jsonData, err := json.Marshal(data)
	//fmt.Println(string(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	body := bytes.NewBuffer(jsonData)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(username, password)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	switch {
	case resp.StatusCode == 400:
		log.Fatal("Bad Request: ", string(bodyResponse))
	case resp.StatusCode != 200:
		log.Fatal("Wrong request: ", resp.Status)
	}
	return true
}
