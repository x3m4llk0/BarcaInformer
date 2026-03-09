package main

import (
	"BarcaInformer/internal/notifier"
	"BarcaInformer/internal/provider"
	"fmt"
)

func main() {
	response, err := provider.GetInfo()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)

	err = notifier.SendMessage(response)
}
