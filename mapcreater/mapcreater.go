package mapcreater

type Island struct {
	IslandsMap   [][]int
	IslanddPoint []string
	IslandCount  int
	MapSize int
}

type MapCreaterInferface interface {
	make(int, int) Island
}

type MapCreater struct {
}

func (mc *MapCreater) Make(n int, m int) *Island {
	islandMap := make([][]int, 0)

	for i := 0; i < n; i++ {
		islandMap = append(islandMap, make([]int, n))
	}

	i := Island{IslandsMap: islandMap, IslandCount: m * n, IslanddPoint: []string{}, MapSize: m * n}
	return &i
}