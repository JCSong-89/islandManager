package generater

import (
	"math/rand"
	"strconv"
)

type IslandGeneraterInterface interface{
	generate(n int, m *[][]int) int
}

type IslandGenerater struct {
}

func (i *IslandGenerater) Generate(n int, m *[][]int, l *int, ip *[]string) int {
	if 	*l == 0 {
		return 0
	}
	
	count := 0
	for {
		x := len(*m)
		y := len((*m)[0])
		x = rand.Intn(x)
		y = rand.Intn(y)

		if 	*l == 0 || count >= n {
			break
		}

		v := (*m)[x][y]

		if v == 1 {
			continue
		}

		count +=1

		(*m)[x][y] = 1
		strX := strconv.Itoa(x)
		strY := strconv.Itoa(y)

		point := strX + "," + strY
		*ip = append(*ip, point)
		
		*l -= 1
	} 

	return count
}
