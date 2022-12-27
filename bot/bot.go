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

	err = goBot.Open()

	if err != nil {
		fmt.Println("Erro ao iniciar a conexão: ", err.Error())
		return
	}
	fmt.Println("Bot iniciado com sucesso!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "@${m.Author.Username} Pong!")
	}
}
