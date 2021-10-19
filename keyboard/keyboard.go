package keyboard

import (
	"fmt"
	"strconv"
)

type KeyboardInterface interface {
	Input() int
}

type Keyboard struct{}

func (k *Keyboard) Input() int {
	var inputValue string
	_, err := fmt.Scanln(&inputValue)
	if err != nil {
		panic(err)
	}

	v, err := strconv.Atoi(inputValue)
	if err != nil {
		panic(err)
	}
	
	return v
}
