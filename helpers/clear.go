package helpers

import "fmt"

func ClearInput() {
	fmt.Print("\033[F")
	fmt.Print("\r\033[K")
	fmt.Print("\033[E")
	fmt.Print("\r\033[K")
	fmt.Print("\033[F")
}