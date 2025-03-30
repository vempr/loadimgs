package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"github.com/vempr/loadimgs/helpers"
)

func DownloadSingle(link string) {
	fmt.Println("  Creating necessary directories...")

	newPath := createDirectory()

	fmt.Printf("  Downloading single image from [%v]...\n", link)

	resp, err := grab.Get(newPath, link)
	fmt.Printf("  Response Status: %v\n", resp.HTTPResponse.Status)
	if err != nil {
		log.Fatal(err)
	}

	helpers.Seperate()

	fmt.Printf("File successfully saved to .%v\n", resp.Filename)
}

func createDirectory() (string) {
	date := time.Now().Format("2006-01-02_15-04-05")
	newPath := filepath.Join(".", "downloads", date)
	os.MkdirAll(newPath, os.ModePerm)

	return newPath
}