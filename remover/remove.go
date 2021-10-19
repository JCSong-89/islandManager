package remover

import (
	"fmt"
	"strconv"
	"strings"
)

type IslandsRemoverInferface interface {
	remove(n int, m *[][]int, l *int) int
}

type IslandsRemover struct {
}

func (i *IslandsRemover) Remove(n int, m *[][]int, l *int, ip *[]string, size int) int {
	if *l == size {
		return 0
	}


	count := 0
	for {
		if *l == size || count >= n {
			break
		}

		pop := (*ip)[len(*ip) - 1:]
		*ip = (*ip)[:len(*ip)-1]
		
		points := strings.Split(pop[0], ",")
		pointX, err := strconv.Atoi(points[0])

		if err != nil {
			fmt.Println(err)
			return 0
		}

		pointY, err := strconv.Atoi(points[1])

		if err != nil {
			fmt.Println(err)
			return 0
		}

		(*m)[pointX][pointY] = 0

		count += 1
		*l += 1
		}
	
	return count
}
