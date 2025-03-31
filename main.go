package main

import (
	"fmt"
	"strings"

	"github.com/br3w0r/goitopdf/itopdf"
	"github.com/vempr/loadimgs/helpers"
)

func main() {
	var link string
	var singleInput string
	var savedPath string
	var pdfInput string

	fmt.Println("\n====== File Downloader ======")
	fmt.Println()

	for {
		fmt.Print("Enter the file source link: ")
		fmt.Scanln(&link)

		if strings.TrimSpace(link) != "" {
			break
		}

		helpers.ClearInput()
	}
	if !strings.Contains(link, "/") || !strings.Contains(link, ".") {
		helpers.Seperate()
		fmt.Printf("ERROR: Invalid link [%v]\n", link)
		return
	}

	for {
		fmt.Print("Download a single file? [Y/N]: ")
		fmt.Scanln(&singleInput)

		if strings.ToUpper(singleInput) == "Y" || strings.ToUpper(singleInput) == "N" {
			break
		}

		helpers.ClearInput()
	}
	single := strings.ToUpper(singleInput) == "Y"

	helpers.Seperate()

	if single {
		savedPath = DownloadSingle(link)
	} else {
		savedPath = DownloadMultiple(link)
	}

	helpers.Seperate()

	for {
		fmt.Print("Convert image file(s) to a PDF file? [Y/N]: ")
		fmt.Scanln(&pdfInput)

		if strings.ToUpper(pdfInput) == "Y" || strings.ToUpper(pdfInput) == "N" {
			break
		}

		helpers.ClearInput()
	}
	pdf := strings.ToUpper(pdfInput) == "Y"

	helpers.Seperate()

	if pdf {
		fmt.Println("  Creating itopdf instance...")
		pdf := itopdf.NewInstance()

		fmt.Println("  Iterating through images...")
		err := pdf.WalkDir(savedPath, nil)
		if err != nil {
			helpers.Seperate()
			fmt.Printf("ERROR: %q", err)
			return
		}

		fmt.Println("  Saving PDF file...")
		outputPath := savedPath + "\\output.pdf"
		err = pdf.Save(outputPath)
		if err != nil {
			helpers.Seperate()
			fmt.Printf("ERROR: %q", err)
			return
		}

		helpers.Seperate()

		fmt.Printf("PDF file successfully generated and saved to '.%v'\n", outputPath)
	}
}
