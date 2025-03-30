package main

import (
	"fmt"
	"strings"

	"github.com/vempr/loadimgs/helpers"
)

func main() {
	var link string
	var singleInput string
	var start, end int

	fmt.Println("\n====== Image Downloader ======")

	for {
		fmt.Print("Enter the image source link: ")
		fmt.Scanln(&link)

		if strings.TrimSpace(link) != "" {
			break
		}

		fmt.Printf("\033[1A\033[K")
	}

	for strings.ToUpper(singleInput) != "Y" && strings.ToUpper(singleInput) != "N" {
		fmt.Print("Download a single image? [Y/N]: ")
		fmt.Scan(&singleInput)
	}
	single := strings.ToUpper(singleInput) == "Y"

	helpers.Seperate()

	if single {
		DownloadSingle(link)
	} else {
		fmt.Println("Configuring multiple image download...")

		lastSlash := strings.LastIndex(link, "/")
		imagesLink := link[:lastSlash+1]
		fmt.Printf("Extracted base link [%v]\n", imagesLink)

		fmt.Print("Enter starting index: ")
		fmt.Scan(&start)
		fmt.Print("Enter ending index: ")
		fmt.Scan(&end)

		lastDot := strings.LastIndex(link, ".")
		format := link[lastDot+1:]

		helpers.Seperate()

		fmt.Printf("Downloading images from [%v]\n", imagesLink)
		fmt.Println(start, end, format)
	}
}
