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
	path := os.Args[1]
	emails := indexer.SetEmails(path)
	fmt.Println(len(emails))
	//printFiles(emails)
	if indexer.LoadDataBulkV2(emails, "http://localhost:4080/api/_bulkv2", "admin", "Complexpass#123") {
		fmt.Println("Data loaded")
	}
}

func printFiles(files []models.File) {
	for _, file := range files {
		fmt.Println(file.Folder)
		//fmt.Println("Content:\t", file.Content)
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
