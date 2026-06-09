package files

import (
	"log"
	"os"
	"path/filepath"
)


// fileSave writes the entire text string to a file at once
func fileSave(text string, filename string) {
	path := filepath.Join(data_path, filename)
	
	// os.WriteFile creates or truncates the file, writes the data, and closes it automatically
	err := os.WriteFile(path, []byte(text), 0644)
	if err != nil {
		log.Fatalf("error writing file: %s", err)
	}
	log.Printf("Successfully wrote to %s", filename)
}

// reads the entire file into a single string block
func fileRead(filename string) string {
	path := filepath.Join(data_path, filename)
	
	// os.ReadFile reads the whole file into memory at once
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}
	
	return string(content)
}