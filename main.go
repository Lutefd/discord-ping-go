package main

import (
	"fmt"
	"github.com/Lutefd/discord-ping-go/bot"
	"github.com/Lutefd/discord-ping-go/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bot.Start()

	<-make(chan struct{})
	return
}
