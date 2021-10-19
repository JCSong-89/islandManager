package main

import (
	"fmt"

	"island.com/filereader"
	"island.com/filewriter"
	"island.com/keyboard"
	"island.com/mapcreater"
	"island.com/display"
)

type key struct {
	IslandsMap [][]int
}

func main () {
	fmt.Println("생성할 격자의 m의 갯수를 입력하세요")
	kb := keyboard.Keyboard{}
	m := kb.Input()

	if m > 20000 {
		fmt.Println("최대치 2만을 초과하였으므로 자동으로 최대치로 조정합니다.")
		m = 20000
	}

	fmt.Println("생성할 격자의 n의 갯수를 입력하세요")
	n := kb.Input()

	if n > 20000 {
		fmt.Println("최대치 2만을 초과하였으므로 자동으로 최대치로 조정합니다.")
		n = 20000
	}

	mc := mapcreater.MapCreater{}
	islands := mc.Make(n, m)

	gd := display.GenerateDisplay{}
	gd.Show(islands)

	rd := display.RemoveDisplay{}
	rd.Show(islands)

	k := key{IslandsMap: islands.IslandsMap}

	fmt.Println("파일을 생성합니다.")
	filewriter.FileWriter(k)
	fmt.Println("파일을 읽습니다.")

	readFile := filereader.FileReader("test.json")

	for k, v := range readFile {
		for _, el := range v {
			fmt.Println("Key: ", k, "value: ", el)
		}
	}
}