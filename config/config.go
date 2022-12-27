package config

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token  string `json:"token"`
	Prefix string `json:"prefix"`
}
