package display

import (
	"fmt"

	"island.com/remover"
	"island.com/keyboard"
	"island.com/mapcreater"
)

type RemoveDisplay struct{}

func (rd *RemoveDisplay) Show(islands *mapcreater.Island) {
	for {
		fmt.Println("지울 섬의 n의 갯수를 입력하세요")
		kb := keyboard.Keyboard{}
		rn := kb.Input()

		if rn > 20000 {
			fmt.Println("최대치 2만을 초과하였으므로 자동으로 최대치로 조정합니다.")
			rn = 20000
		}

	if rn <= islands.MapSize &&  islands.MapSize - islands.IslandCount != 0{
		r := remover.IslandsRemover{}
		fmt.Printf("지워진 갯수는 %v \n", r.Remove(rn, &islands.IslandsMap, &islands.IslandCount,  &islands.IslanddPoint, islands.MapSize))
		fmt.Println("계속하시겠습니까? 그만하시려면 0, 계속하시려면 아무키나 입력해주세요.")
		rControl := kb.Input()

		if rControl == 0 {
			break
		}
		continue
	 } else if  islands.MapSize - islands.IslandCount == 0 {
		fmt.Println("지울 수 있는 갯수가 없습니다.")
		fmt.Println("지울 수 있는 갯수: ",  islands.MapSize - islands.IslandCount)
		break
	 }	else {
		fmt.Println("지울 수 있는 갯수를 초과 하였습니다.")
		fmt.Println("지울 수 있는 갯수: ",  islands.MapSize - islands.IslandCount)
		continue
	}
}
}