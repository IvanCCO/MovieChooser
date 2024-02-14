package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func main() {

	discord, err := discordgo.New("Bot " + "authentication token")

	fmt.Println(discord.Client)
	fmt.Println(err)

	fmt.Println("Hello, World!")
}
