package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (cfg *Config) Parse(filename string) error {
	if filename == "" {
		return fmt.Errorf("config filename is empty")
	}

	byteValue, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(byteValue, cfg); err != nil {
		return err
	}
	return nil
}

type Config struct {
	MongoDB Mongo `json:"mongo"`
	Bot     Bot   `json:"bot"`
}

type Bot struct {
	Token     Token `json:"token,required"`
	CreatorID int64 `json:"creator_id,required"`
}

type Mongo struct {
	Addrs    []string `json:"addrs"`
	UserName string   `json:"username"`
	Password string   `json:"password"`
	DBName   string   `json:"db_name"`
}

func (b Bot) GetToken() Token {
	return b.Token
}

type Token string

func (t Token) String() string {
	return string(t)
}
