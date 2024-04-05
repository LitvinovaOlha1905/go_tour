package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go запущено на ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Перемикач без умови

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Доброго ранку!")
	case t.Hour() < 17:
		fmt.Println("Добрий день.")
	default:
		fmt.Println("Добрий вечір.")
	}

}
