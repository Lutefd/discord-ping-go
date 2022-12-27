package bot

import (
	"fmt"
	"github.com/Lutefd/discord-ping-go/config"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	fmt.Println("Iniciando o bot ...")
	discord, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println("Erro ao iniciar o bot: ", err)
		return
	}

}
