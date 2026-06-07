package files

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func EmailSave() {

}

func EmailRead() {
	// Open the file
	file, err := os.Open(data_path + "email.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Loop through the file and read each line
	for scanner.Scan() {
		line := scanner.Text() // Get the line as a string
		fmt.Println(line)
	}

	// Check for errors during the scan
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
}
