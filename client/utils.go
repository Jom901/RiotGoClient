package client

import (
	"os"
)

type EnvKey string

func (key EnvKey) GetValue() string {
	return os.Getenv(string(key))
}

const (
	RiotDevKey EnvKey = "RIOT_DEV_KEY"
)

var apiKey = RiotDevKey.GetValue()
