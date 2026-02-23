package main

import (
	"flag"
	"fmt"
	"leagueClient/client"
	"log"
)

func main() {
	inGameName := flag.String("inGameName", "", "In game name")
	tagline := flag.String("tagline", "NA1", "User id tag. Typically found after # in your in game name")
	flag.Parse()
	log.Println(*inGameName)
	account, err := client.GetAccountByGameName(*inGameName, *tagline)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(account.AccountId)
	accountRegion, err := client.GetAccountRegionByGame(account.AccountId, "lol")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(accountRegion.Region)
}
