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
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println("Erro ao iniciar o bot: ", err.Error())
		return
	}
	user, err := goBot.User("@me")

	if err != nil {
		fmt.Println("Erro ao obter o usuário: ", err.Error())
	}

	BotID = user.ID

	goBot.AddHandler(messageHandler)
}
