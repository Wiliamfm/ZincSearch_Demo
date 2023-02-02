package indexer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Wiliamfm/ZincSearch_Demo/models"
)

func loadUserDir(path string) []string {
	dirs := make([]string, 0)
	path = path + "/maildir"
	users, err := os.ReadDir(path)
	if err != nil {
		log.Fatal("Error loading directory of users: ", err)
	}
	for _, dir := range users {
		if dir.IsDir() {
			folderPath := path + "/" + dir.Name()
			dirs = append(dirs, folderPath)
		} else {
			fmt.Printf("File %s found on users directory, file not handled", dir.Name())
		}
	}
	return dirs
}

func SetEmails(path string) []models.File {
	emails := make([]models.File, 0)
	for _, dir := range loadUserDir(path) {
		emails = append(emails, setFiles(emails, dir)...)
	}
	return emails
}

func setFiles(emails []models.File, path string) []models.File {
	folders, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Error loading dir: %s: %+v", path, err)
	}
	for _, folder := range folders {
		if folder.IsDir() {
			emails = setFiles(emails, path+"/"+folder.Name())
		} else {
			emails = append(emails, addFile(path+"/"+folder.Name()))
		}
	}
	return emails
}

func addFile(path string) models.File {
	file := models.File{Folder: path}
	fileContent, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}
	file.Content = string(fileContent[:])
	return file
}

func SetEmailsV2(path string) []models.File {
	//emails := make([]models.Email, 0)
	emails := make([]models.File, 0) //V2
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal("Error reading dirs", err)
	}
	for _, file := range files {
		if !file.IsDir() {
			log.Fatalf("File found in mailfolder: %s; file not handled", file.Name())
		}
		folderPath := path + "/" + file.Name()
		/*
			email := models.Email{Username: file.Name(), MailFolders: make(map[string][]models.File)}
			emails = append(emails, loadEmail(email, folderPath))
		*/
		//emails = loadEmailV2(emails, folderPath) //V2
		emailsChanel := loadEmailGoRoutine(&emails, folderPath)
		emails = append(emails, <-emailsChanel...)
	}
	return emails
}

func loadEmailGoRoutine(emails *[]models.File, path string) <-chan []models.File {
	fmt.Println(path)
	emailsChanel := make(chan []models.File)
	items, err := os.ReadDir(path)
	go func() {
		if err != nil {
			log.Fatalf("Error reading folders of %s:\n%+v", path, err)
		}
		for _, item := range items {
			if item.IsDir() {
				emailsChanel <- loadEmailV2(*emails, path+"/"+item.Name())
			} else {
				fileContent, err := os.ReadFile(path + "/" + item.Name())
				if err != nil {
					log.Fatal("Error reading file ", err)
				}
				content := string(fileContent[:])
				file := models.File{Folder: path + "/" + item.Name(), Content: content}
				*emails = append(*emails, file)
			}
		}
	}()
	return emailsChanel
}

func loadEmailV2(emails []models.File, path string) []models.File {
	//fmt.Println(path)
	items, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Error reading folders of %s:\n%+v", path, err)
	}
	for _, item := range items {
		if item.IsDir() {
			emails = loadEmailV2(emails, path+"/"+item.Name())
		} else {
			fileContent, err := os.ReadFile(path + "/" + item.Name())
			if err != nil {
				log.Fatal("Error reading file ", err)
			}
			content := string(fileContent[:])
			file := models.File{Folder: path + "/" + item.Name(), Content: content}
			emails = append(emails, file)
		}
	}
	return emails
}

func LoadDataBulkV2(emails []models.File, url, username, password string) bool {
	data := models.RequestData{Index: "emails", Records: emails}
	jsonData, err := json.Marshal(data)
	createJsonFile(jsonData)
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

func createJsonFile(data []byte) {
	fmt.Println("creating Json file")
	file, err := os.Create("json_data.json")
	if err != nil {
		log.Fatal("Could not create json File: ", err)
	}
	_, err = file.Write(data)
	if err != nil {
		log.Fatal("Could not write json file: ", err)
	}
}
