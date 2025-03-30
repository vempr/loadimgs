package main

import (
	"fmt"
	"strings"

	"github.com/vempr/loadimgs/helpers"
)

func main() {
	var link string
	var singleInput string

	fmt.Println("\n====== File Downloader ======")
	fmt.Println()

	for {
		fmt.Print("Enter the file source link: ")
		fmt.Scanln(&link)

		if strings.TrimSpace(link) != "" {
			break
		}

		fmt.Printf("\033[1A\033[K")
	}
	if (!strings.Contains(link, "/") || !strings.Contains(link, ".")) {
		helpers.Seperate()
		fmt.Printf("ERROR: Invalid link [%v]\n", link)
		return
	}

	for strings.ToUpper(singleInput) != "Y" && strings.ToUpper(singleInput) != "N" {
		fmt.Print("Download a single file? [Y/N]: ")
		fmt.Scan(&singleInput)
	}
	single := strings.ToUpper(singleInput) == "Y"

	helpers.Seperate()

	if single {
		DownloadSingle(link)
	} else {
		DownloadMultiple(link)
	}
}
