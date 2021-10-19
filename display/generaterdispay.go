package display

import (
	"fmt"

	"island.com/generater"
	"island.com/keyboard"
	"island.com/mapcreater"
)

type GenerateDisplay struct{}

func (g *GenerateDisplay) Show(islands *mapcreater.Island) {
	for {
		fmt.Println("생성할 섬의 n의 갯수를 입력하세요")
		kb:= keyboard.Keyboard{}
		gn := kb.Input()

		if gn > 20000 {
			fmt.Println("최대치 2만을 초과하였으므로 자동으로 최대치로 조정합니다.")
			gn = 20000
		}

		g := generater.IslandGenerater{}
		fmt.Printf("생성된 갯수는 %v \n", g.Generate(gn, &islands.IslandsMap, &islands.IslandCount, &islands.IslanddPoint))

		if islands.IslandCount == 0 {
			fmt.Println("모든 격자에 섬이 생성되어 더 이상 생성할 수 없습니다.")
			break
		}

		fmt.Println("계속하시겠습니까? 그만하시려면 0, 아무키나 입력해주세요.")
		gControl := kb.Input()

		if gControl == 0 {
			break
		}
		
		continue
	}
}