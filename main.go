package main

import (
	_ "fmt"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discordgo.New("Bot " + "authentication token")
}
