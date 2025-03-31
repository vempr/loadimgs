# loadimgs

Is a simple command-line application written in Go for downloading files from a specified URL. It supports downloading single files and multiple files in a numerical sequence. loadimgs additionally provides an option to convert downloaded image files into a PDF.

Modern command-line shells are recommended (e.g. Bash) for displaying messages containing [ANSI escape sequences](https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797).

## Installation
### Prerequisites
Ensure you have Go installed on your system. If not, download and install it from [Go's official website](https://go.dev/dl/).

### Clone the Repository
```sh
git clone https://github.com/vempr/loadimgs.git
cd loadimgs
```

### Install Dependencies
```sh
go mod tidy
```

## Usage
Run the application with:
```sh
go run .
```

## Examples
### Download a single file
```sh
Enter the file source link: https://example.com/image.jpg
Download a single file? [Y/N]: Y
```

### Download multiple files
```sh
Enter the file source link: https://example.com/image1.jpg
Download a single file? [Y/N]: N
Enter starting index: 1
Enter ending index: 5
```
This will download `image1.jpg` to `image5.jpg` from the specified base URL (https://example.com/).

## Why?

I wanted to download a chemistry book to learn in the holidays (skipped an entire year of chemistry) and luckily all the pages were online. However, the publisher's website used the most doodoo interactive JavaScript book which was too annoying to navigate. This was the perfect chance to learn Go, so I made <b>loadimgs</b> and spared my sanity by not having to click `Save image as...` 200 times!