package main

import (
	"fmt"
	"github.com/MarcosSmeets/GoOportunities.git/config"
	"github.com/MarcosSmeets/GoOportunities.git/router"
)

func main() {
	//Initialize config
	err := config.InitDb()

	if err == nil {
		fmt.Println(err)
		return
	}

	//Initialize Router
	router.Initialize()
}
