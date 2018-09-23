package main

import (
	"fmt"

	"github.com/sticreations/spotigopher/spotigopher"
)

func main() {

	client := spotigopher.NewClient()
	information, err := client.GetInfo()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Printf("Your Spotify Informations: %v", information)

}
