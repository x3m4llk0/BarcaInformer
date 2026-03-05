package main

import (
	"BarcaInformer/internal/provider"
	"fmt"
)

func main() {
	response, err := provider.GetInfo()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}
