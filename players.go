package clash

import (
	"fmt"
)

type Card struct {
	Name     string       `json:"name"`
	Level    int          `json:"level"`
	MaxLevel int          `json:"maxLevel"`
	Count    int          `json:"count"`
	IconUrls CardIconUrls `json:"iconUrls"`
}

type FavouriteCard struct {
	Name     string       `json:"name"`
	ID       int          `json:"id"`
	MaxLevel int          `json:"maxLevel"`
	IconUrls CardIconUrls `json:"iconUrls"`
}

type CardIconUrls struct {
	Medium string `json:"medium"`
}

type Achievement struct {
	Name   string `json:"name"`
	Stars  int    `json:"stars"`
	Value  int    `json:"value"`
	Target int    `json:"target"`
	Info   string `json:"info"`
}

type PlayerClan struct {
	Tag     string `json:"tag"`
	Name    string `json:"name"`
	BadgeID int    `json:"badgeId"`
}

type Arena struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Season struct {
	Rank         int `json:"rank"`
	Trophies     int `json:"trophies"`
	BestTrophies int `json:"bestTrophies"`
	ID           int `json:"string"`
}

type LeagueStatistics struct {
	BestSeason     Season `json:"bestSeason"`
	PreviousSeason Season `json:"previousSeason"`
	CurrentSeason  Season `json:"currentSeason"`
}

type Player struct {
	Tag                   string           `json:"tag"`
	Name                  string           `json:"name"`
	ExpLevel              int              `json:"expLevel"`
	Trophies              int              `json:"trophies"`
	BestTrophies          int              `json:"bestTrophies"`
	Wins                  int              `json:"wins"`
	Losses                int              `json:"losses"`
	BattleCount           int              `json:"battleCount"`
	ThreeCrownWins        int              `json:"threeCrownWins"`
	ChallengeCardsWon     int              `json:"challengeCardsWon"`
	ChallengeMaxWins      int              `json:"challengeMaxWins"`
	TournamentCardsWon    int              `json:"tournamentCardsWon"`
	TournamentBattleCount int              `json:"tournamentBattleCount"`
	Role                  string           `json:"role"`
	Donations             int              `json:"donations"`
	DonationsReceived     int              `json:"donationsReceived"`
	TotalDonations        int              `json:"totalDonations"`
	WarDayWins            int              `json:"warDayWins"`
	ClanCardsCollected    int              `json:"clanCardsCollected"`
	Clan                  PlayerClan       `json:"clan"`
	Arena                 Arena            `json:"arena"`
	Achievements          []Achievement    `json:"achievements"`
	Cards                 []Card           `json:"cards"`
	CurrentDeck           []Card           `json:"currentDeck"`
	CurrentFavouriteCard  FavouriteCard    `json:"currentFavouriteCard"`
	LeagueStatistics      LeagueStatistics `json:"leagueStatistics"`
}

type VerificationResult struct {
	Tag    string `json:"tag"`
	Token  string `json:"token"`
	Status string `json:"status"`
}

func (v *VerificationResult) isValid() bool {
	return v.Status == "ok"
}

type GameMode struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type BattleLogPlayer struct {
	Tag              string `json:"tag"`
	Name             string `json:"name"`
	StartingTrophies int    `json:"startingTrophies"`
	TrophyChange     int    `json:"trophyChange"`
	Crowns           int    `json:"crowns"`
	Clan struct {
		Tag     string `json:"tag"`
		Name    string `json:"name"`
		BadgeId int    `json:"badgeId"`
	} `json:"clan"`
	Cards []Card `json:"cards"`
}

type BattleLogEntry struct {
	Type          string            `json:"type"`
	BattleTime    string            `json:"battleTime"`
	Arena         Arena             `json:"arena"`
	GameMode      GameMode          `json:"gameMode"`
	DeckSelection string            `json:"deckSelection"`
	Team          []BattleLogPlayer `json:"team"`
	Opponent      []BattleLogPlayer `json:"opponent"`
}

type BattleLogEntries []BattleLogEntry

type UpcomingChest struct {
	Index int    `json:"index"`
	Name  string `json:"name"`
}

type UpcomingChests struct {
	Items []UpcomingChest `json:"items"`
}

func (c *Client) GetPlayerUpcomingChests(hashtag string) (UpcomingChests, error) {
	url := fmt.Sprintf("/v1/players/%s/upcomingchests", normaliseHashtag(hashtag))
	req, err := c.newRequest("GET", url, nil)
	var chests UpcomingChests

	if err == nil {
		_, err = c.do(req, &chests)
	}

	return chests, err
}

func (c *Client) GetPlayerBattleLog(hashtag string) (BattleLogEntries, error) {
	url := fmt.Sprintf("/v1/players/%s/battlelog", normaliseHashtag(hashtag))
	req, err := c.newRequest("GET", url, nil)
	var list BattleLogEntries

	if err == nil {
		_, err = c.do(req, &list)
	}

	return list, err
}

func (c *Client) GetPlayer(hashtag string) (Player, error) {
	url := fmt.Sprintf("/v1/players/%s", normaliseHashtag(hashtag))
	req, err := c.newRequest("GET", url, nil)
	var player Player

	if err == nil {
		_, err = c.do(req, &player)
	}

	return player, err
}

func (c *Client) VerifyPlayerToken(hashtag string, token string) (VerificationResult, error) {
	url := fmt.Sprintf("/v1/players/%s/verifytoken", normaliseHashtag(hashtag))
	req, err := c.newRequest("POST", url, map[string]string{"token": token})
	var result VerificationResult

	if err == nil {
		_, err = c.do(req, &result)
	}

	return result, err
}
