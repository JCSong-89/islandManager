package filereader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func FileReader(filename string) map[string][][]int {
	in, err := os.Open(filename)

	if err != nil {
		fmt.Println("파일을 읽지 못해 종료합니다.")
		panic(err)
	}

	byteValue, err := ioutil.ReadAll(in)
	
	if err != nil {
		fmt.Println("파일을 읽지 못해 종료합니다.")
		panic(err)
	}
	
	var fileData map[string][][]int
	json.Unmarshal(byteValue, &fileData)

	in.Close()
	return fileData
}