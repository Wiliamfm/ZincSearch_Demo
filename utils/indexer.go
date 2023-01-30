package indexer

import (
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/Wiliamfm/ZincSearch_Demo/models"
)

func SetEmails(path string) models.Emails {
	emails := models.Emails{}
	err := filepath.WalkDir(path, func(path string, info os.DirEntry, err error) error {
		email := models.Email{MailFolders: map[string][]string{}}
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
					email.MailFolders[folder] = make([]string, 0)
				}
				email.MailFolders[folder] = append(email.MailFolders[folder], info.Name())
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
