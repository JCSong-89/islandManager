package filewriter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func FileWriter(key interface{}){
	file, _ := json.MarshalIndent(key, "", " ")
	err := ioutil.WriteFile("test.json", file, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
}