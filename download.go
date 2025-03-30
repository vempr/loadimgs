package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

	fmt.Printf("File successfully saved to .\\%v\n", resp.Filename)
}

func DownloadMultiple(link string) {
	var start, end int

	fmt.Println("Extracting base link...")

	lastSlash := strings.LastIndex(link, "/")
	baseLink := link[:lastSlash+1]
	fmt.Printf("Extracted base link [%v]\n", baseLink)

	start = getValidInteger("Enter starting index: ")
	end = getValidInteger("Enter ending index: ")

	if (start > end) {
		helpers.Seperate()
		fmt.Println("ERROR: Starting index can not be larger than ending index")
		return
	}

	format := link[strings.LastIndex(link, ".")+1:]

	helpers.Seperate()

	fmt.Println("  Creating necessary directories...")
	newPath := createDirectory()

	fmt.Printf("  Downloading multiple files from [%v] in the range of [%v] to [%v]...\n", baseLink, start, end)

	for start <= end {
		downloadLink := fmt.Sprint(baseLink, start, ".", format)
		
		resp, err := grab.Get(newPath, downloadLink)
		fmt.Printf("  [%v]: Imported image [%v]\n", resp.HTTPResponse.Status, downloadLink)
		if err != nil {
			helpers.Seperate()
			fmt.Println("ERROR:")
			log.Fatal(err)
		}

		start += 1
	}

	helpers.Seperate()

	fmt.Printf("Files successfully saved to .\\%v\n", newPath[strings.LastIndex(newPath, "downloads"):])
}

func createDirectory() string {
	date := time.Now().Format("2006-01-02_15-04-05")
	newPath := filepath.Join(".", "downloads", date)
	os.MkdirAll(newPath, os.ModePerm)

	return newPath
}

func getValidInteger(prompt string) int {
	var input string

	for {
 		fmt.Printf("\r%s", prompt)
		fmt.Scanln(&input)

		value, err := strconv.Atoi(input)
		if err == nil {
			return value
		}

		fmt.Print("\033[F")
		fmt.Print("\r\033[K")
		fmt.Print("\033[E")
		fmt.Print("\r\033[K")
		fmt.Print("\033[F")
	}
}