package files

import (
	"encoding/json"
	"fmt"
	"log"
	 
)

func jsonRead(file string) Experience {
	str := fileRead(file)
	res := Experience{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	return res
}

func jsonWrite(file string, experience Experience) {
	
	json, err := json.Marshal(experience)
	if err != nil {
		log.Fatalf("Error reading json %s", file)
		return
	}
	str := fmt.Sprintf("%d", json)
	fileSave(str, file)

}