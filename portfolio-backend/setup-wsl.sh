package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	_, _ = fmt.Println, log.Println
	_ = filepath.Dir(os.Args[0])
	log.Println("WSL setup placeholder")
}
