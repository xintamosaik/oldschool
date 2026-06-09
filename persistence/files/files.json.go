package files

import (
	"encoding/json"
	"fmt"
	"log"
)

// Experience represents the top-level JSON wrapper structure
type Experience struct {
	Jobs []Job `json:"jobs"`
}

// jsonRead now correctly extracts and returns the []Job slice
func jsonRead(file string) []Job {
	str := fileRead(file)
	var res []Job // Read directly into a slice of Jobs instead of the struct
	
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		log.Printf("Error unmarshaling json from %s: %s", file, err)
		return []Job{}
	}
	
	fmt.Println(res)
	return res
}
// jsonWrite now wraps the []Job back into an Experience struct before marshaling
func jsonWrite(file string, experience []Job) {
	// Marshal the slice directly to keep it a flat array in the JSON file
	jsonData, err := json.MarshalIndent(experience, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling json %s: %s", file, err)
		return
	}
	
	str := string(jsonData) 
	fileSave(str, file)
}