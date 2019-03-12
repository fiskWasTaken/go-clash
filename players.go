package clash

import (
	"errors"
	"fmt"
	"time"
)

type Card struct {
	Name      string   `json:"name"`
	Level     int      `json:"level"`
	MaxLevel  int      `json:"maxLevel"`
	Count     int      `json:"count"`
	IconUrls  IconUrls `json:"iconUrls"`
	StarLevel int      `json:"starLevel"`
}

// Return the internal client level for the card, as these are zero-indexed
func (c *Card) ClientLevel() int {
	return c.Level - 1
}

type FavouriteCard struct {
	Name     string   `json:"name"`
	ID       int      `json:"id"`
	MaxLevel int      `json:"maxLevel"`
	IconUrls IconUrls `json:"iconUrls"`
}

type IconUrls struct {
	Medium string `json:"medium"`
}

type Achievement struct {
	Name   string `json:"name"`
	Stars  int    `json:"stars"`
	Value  int    `json:"value"`
	Target int    `json:"target"`
	Info   string `json:"info"`
}

type Badge struct {
	Name     string `json:"name"`
	Level    int    `json:"level,omitempty"`
	MaxLevel int    `json:"maxLevel,omitempty"`
	Progress int    `json:"progress,omitempty"`
	Target   int    `json:"target,omitempty"`
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
	Rank         int    `json:"rank"`
	Trophies     int    `json:"trophies"`
	BestTrophies int    `json:"bestTrophies"`
	ID           string `json:"id"`
}

type LeagueStats struct {
	BestSeason     Season `json:"bestSeason"`
	PreviousSeason Season `json:"previousSeason"`
	CurrentSeason  Season `json:"currentSeason"`
}

type Player struct {
	Tag                   string        `json:"tag"`
	Name                  string        `json:"name"`
	ExpLevel              int           `json:"expLevel"`
	Trophies              int           `json:"trophies"`
	BestTrophies          int           `json:"bestTrophies"`
	Wins                  int           `json:"wins"`
	Losses                int           `json:"losses"`
	BattleCount           int           `json:"battleCount"`
	ThreeCrownWins        int           `json:"threeCrownWins"`
	ChallengeCardsWon     int           `json:"challengeCardsWon"`
	ChallengeMaxWins      int           `json:"challengeMaxWins"`
	TournamentCardsWon    int           `json:"tournamentCardsWon"`
	TournamentBattleCount int           `json:"tournamentBattleCount"`
	Role                  string        `json:"role"`
	Donations             int           `json:"donations"`
	DonationsReceived     int           `json:"donationsReceived"`
	TotalDonations        int           `json:"totalDonations"`
	WarDayWins            int           `json:"warDayWins"`
	ClanCardsCollected    int           `json:"clanCardsCollected"`
	Clan                  PlayerClan    `json:"clan"`
	Arena                 Arena         `json:"arena"`
	Achievements          []Achievement `json:"achievements"`
	Badges                []Badge       `json:"badges"`
	Cards                 []Card        `json:"cards"`
	CurrentDeck           []Card        `json:"currentDeck"`
	CurrentFavouriteCard  FavouriteCard `json:"currentFavouriteCard"`
	LeagueStatistics      LeagueStats   `json:"leagueStatistics"`
	StarPoints            int           `json:"starPoints"`
}

type VerificationResult struct {
	Tag    string `json:"tag"`
	Token  string `json:"token"`
	Status string `json:"status"`
}

func (v *VerificationResult) IsValid() bool {
	return v.Status == "ok"
}

type GameMode struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type BattlePlayer struct {
	Tag              string `json:"tag"`
	Name             string `json:"name"`
	StartingTrophies int    `json:"startingTrophies"`
	TrophyChange     int    `json:"trophyChange"`
	Crowns           int    `json:"crowns"`
	// king tower hit points is missing from the response if it is zero.
	KingTowerHitPoints int `json:"kingTowerHitPoints,omitempty"`
	// ditto. princess tower hit points are an array of values.
	PrincessTowersHitPoints []int `json:"princessTowersHitPoints,omitempty"`
	Clan struct {
		Tag     string `json:"tag"`
		Name    string `json:"name"`
		BadgeId int    `json:"badgeId"`
	} `json:"clan"`
	Cards []Card `json:"cards"`
}

type Battle struct {
	Type                    string         `json:"type"`
	RawBattleTime           string         `json:"battleTime"`
	Arena                   Arena          `json:"arena"`
	GameMode                GameMode       `json:"gameMode"`
	DeckSelection           string         `json:"deckSelection"`
	Team                    []BattlePlayer `json:"team"`
	Opponent                []BattlePlayer `json:"opponent"`
	TournamentTag           string         `json:"tournamentTag"`
	ChallengeId             int            `json:"challengeId"`
	ChallengeWinCountBefore int            `json:"challengeWinCountBefore"`
	ReplayTag               string         `json:"replayTag,omitempty"`
}

type BattleOutcome struct {
	IsDraw  bool
	Winners []BattlePlayer
	Losers  []BattlePlayer
}

// Find a player in the battle by tag. Return an error if the tag could not be found.
func (b *Battle) PlayerByTag(tag string) (BattlePlayer, error) {
	tag = NormaliseTag(tag)

	for _, player := range append(b.Team, b.Opponent...) {
		if player.Tag == tag {
			return player, nil
		}
	}

	return BattlePlayer{}, errors.New("player does not exist in battle")
}

// Get a struct describing the outcome of the battle.
func (b *Battle) Outcome() BattleOutcome {
	if b.Opponent[0].Crowns > b.Team[0].Crowns {
		return BattleOutcome{Winners: b.Opponent, Losers: b.Team}
	} else if b.Team[0].Crowns > b.Opponent[0].Crowns {
		return BattleOutcome{Losers: b.Opponent, Winners: b.Team}
	}

	return BattleOutcome{IsDraw: true}
}

// Get the time of the battle.
func (b *Battle) BattleTime() time.Time {
	parsed, _ := time.Parse(TimeLayout, b.RawBattleTime)
	return parsed
}

type Battles []Battle

type UpcomingChest struct {
	Index int    `json:"index"`
	Name  string `json:"name"`
}

type UpcomingChests struct {
	Items []UpcomingChest `json:"items"`
}

type PlayerService struct {
	c   *Client
	tag string
}

func (c *Client) Player(tag string) *PlayerService {
	return &PlayerService{c, tag}
}

// Get list of reward chests that the player will receive next in the game.
func (i *PlayerService) UpcomingChests() (UpcomingChests, error) {
	url := fmt.Sprintf("/v1/players/%s/upcomingchests", NormaliseTag(i.tag))
	req, err := i.c.NewRequest("GET", url, nil)
	var chests UpcomingChests

	if err == nil {
		_, err = i.c.Do(req, &chests)
	}

	return chests, err
}

// Get list of recent battle results for a player.
func (i *PlayerService) BattleLog() (Battles, error) {
	url := fmt.Sprintf("/v1/players/%s/battlelog", NormaliseTag(i.tag))
	req, err := i.c.NewRequest("GET", url, nil)
	var list Battles

	if err == nil {
		_, err = i.c.Do(req, &list)
	}

	return list, err
}

// Get information about a single player by player tag. Player tags
// can be found either in game or by from clan member lists.
func (i *PlayerService) Get() (Player, error) {
	url := fmt.Sprintf("/v1/players/%s", NormaliseTag(i.tag))
	req, err := i.c.NewRequest("GET", url, nil)
	var player Player

	if err == nil {
		_, err = i.c.Do(req, &player)
	}

	return player, err
}

// Verifies a player token and returns whether or not the token was associated with the given player.
//
// This API call can be used by a player to prove that they own a particular game account as the token
// can only be retrieved inside the game from settings view.
func (i *PlayerService) VerifyToken(token string) (VerificationResult, error) {
	url := fmt.Sprintf("/v1/players/%s/verifytoken", NormaliseTag(i.tag))
	req, err := i.c.NewRequest("POST", url, map[string]string{"token": token})
	var result VerificationResult

	if err == nil {
		_, err = i.c.Do(req, &result)
	}

	return result, err
}
