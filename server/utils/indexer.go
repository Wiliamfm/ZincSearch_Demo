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
	emailsChannel := make(chan []models.File)
	dirs := loadUserDir(path)
	//var wg sync.WaitGroup
	//wg.Add(len(dirs))
	for _, dir := range dirs {
		fmt.Println("Launch go: ", dir)
		go func(dir string) {
			//defer wg.Done()
			emailsChannel <- setFiles(dir)
		}(dir)
		//emails = setFiles(emails, dir)
	}
	//wg.Wait()
	for range dirs {
		emails = append(emails, <-emailsChannel...)
	}
	return emails
}

func setFiles(path string) []models.File {
	//fmt.Println(path)
	emails := make([]models.File, 0)
	folders, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Error loading dir: %s: %+v", path, err)
	}
	for _, folder := range folders {
		if folder.IsDir() {
			//emails = setFiles(path + "/" + folder.Name())
			emails = append(emails, setFiles(path+"/"+folder.Name())...)
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
