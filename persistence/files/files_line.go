package files

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func lineSave(text string, filename string) {
	path := filepath.Join(data_path + filename)
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("error writing file: %s", err)
	}

	defer f.Close()

	n, err := f.WriteString(text)
	if err != nil {
		log.Fatalf("error writing file: %s", err)
	} else {
		log.Printf("We did this: %d", n)
	}

	f.Sync()
}

func lineRead(filename string) string {
	// Open the file
	file, err := os.Open(data_path + filename)
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

		if len(line) > 5 {
			return line
		}

	}

	// Check for errors during the scan
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	return ""
}

