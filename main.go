package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/Wiliamfm/ZincSearch_Demo/models"
	indexer "github.com/Wiliamfm/ZincSearch_Demo/utils"
)

var cpuprofile = flag.String("cpuprofile", "cpu.prof", "Write cpu profile to file")
var memprofile = flag.String("memprofile", "mem.prof", "write memory profile to `file`")
var heapprofile = flag.String("heapprofile", "heap.prof", "write memory profile to `file`")

func main() {
	flag.Parse()
	//profiles()
	path := os.Args[1] + "/maildir"
	//emails := indexer.SetEmails(path)
	//listEmails := indexer.SetEmailsV2("/home/william/Downloads/enron_mail_test/maildir")
	listEmails := indexer.SetEmailsV2(path)
	//emails := models.Emails{Emails: listEmails}
	if indexer.LoadDataBulkV2V2(listEmails, "http://localhost:4080/api/_bulkv2", "admin", "Complexpass#123") {
		fmt.Println("Data loaded")
	}
	/*
		if indexer.LoadDataBulkV2(emails, "http://localhost:4080/api/_bulkv2", "admin", "Complexpass#123") {
			fmt.Println("Data loaded")
		}
	*/
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

func profiles() {
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}

	if *heapprofile != "" {
		f, err := os.Create(*heapprofile)
		if err != nil {
			log.Fatal("could not create heap profile: ", err)
		}
		if err = pprof.Lookup("heap").WriteTo(f, 0); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
	}
}
