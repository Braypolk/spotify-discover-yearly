package main

import (
	"log"
	"os"
)

func WriteSliceToFile(stuff []string, file string)  {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range stuff{
		if _, err := f.WriteString(line + "\n"); err != nil {
			log.Fatal(err)
		}
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}