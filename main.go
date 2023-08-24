package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type Config struct {
	ApplicationID string `json:"applicationID"`
	GuildID       string `json:"guildID"`
	Token         string `json:"token"`
}

type Commands struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Options     []Option `json:"options"`
}

type Option struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Type        int      `json:"type"`
	Required    bool     `json:"required"`
	Choices     []Choice `json:"choices"`
}

type Choice struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func main() {
	log.Println("start")
	if err := run(); err != nil {
		log.Fatal(err)
	}
	log.Println("finish")
}

func run() error {
	c := Config{}
	if err := json.Unmarshal(configBytes, &c); err != nil {
		return err
	}

	postRequest(discordgo.EndpointApplicationGuildCommands(c.ApplicationID, c.GuildID), c.Token, getCommands())

	return nil
}

func getCommands() Commands {
	c1 := Choice{
		Name:  "start",
		Value: "start",
	}
	c2 := Choice{
		Name:  "stop",
		Value: "stop",
	}
	c3 := Choice{
		Name:  "test",
		Value: "test",
	}
	choices := []Choice{c1, c2, c3}

	o1 := Option{
		Name:        "action",
		Description: "start/stop",
		Type:        3,
		Required:    true,
		Choices:     choices,
	}
	options := []Option{o1}

	return Commands{
		Name:        "sss",
		Description: "command to start/stop server",
		Options:     options,
	}
}

func postRequest(endpoint string, token string, commands Commands) error {
	jsonString, err := json.Marshal(commands)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonString))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bot "+token)
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	byteArray, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("%#v", string(byteArray))
	return nil
}
