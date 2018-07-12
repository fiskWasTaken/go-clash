package clash

import (
	"fmt"
	"time"
)

type ClanQuery struct {
	LocationId int
	MinScore   int
	MinMembers int
	MaxMembers int
	Name       string
}

type Clan struct {
	Tag               string       `json:"tag"`
	Name              string       `json:"name"`
	Type              string       `json:"type"`
	Description       string       `json:"description"`
	BadgeId           int          `json:"badgeId"`
	ClanScore         int          `json:"clanScore"`
	Location          Location     `json:"location"`
	RequiredTrophies  int          `json:"requiredTrophies"`
	DonationsPerWeek  int          `json:"donationsPerWeek"`
	ClanChestStatus   string       `json:"clanChestStatus"`
	ClanChestPoints   int          `json:"clanChestPoints"`
	ClanChestLevel    int          `json:"clanChestLevel"`
	ClanChestMaxLevel int          `json:"clanChestMaxLevel"`
	Members           int          `json:"members"`
	MemberList        []ClanMember `json:"memberList"`
}

type ClanPaging struct {
	Items []Clan `json:"items"`
	Paging struct {
		Cursors struct{} `json:"cursors"`
	} `json:"paging"`
}

type MemberPaging struct {
	Items []ClanMember `json:"items"`
	Paging struct {
		Cursors struct{} `json:"cursors"`
	} `json:"paging"`
}

type ClanWarParticipant struct {
	Tag           string `json:"tag"`
	Name          string `json:"name"`
	CardsEarned   int    `json:"cardsEarned"`
	BattlesPlayed int    `json:"battlesPlayed"`
	Wins          int    `json:"wins"`
}

type ClanWarStanding struct {
	Tag           string `json:"tag"`
	Name          string `json:"name"`
	BadgeId       int    `json:"badgeId"`
	ClanScore     int    `json:"clanScore"`
	Participants  int    `json:"participants"`
	BattlesPlayed int    `json:"battlesPlayed"`
	Wins          int    `json:"wins"`
	Crowns        int    `json:"crowns"`
}

type WarLogStanding struct {
	Clan         ClanWarStanding `json:"clan"`
	TrophyChange int             `json:"trophyChange"`
}

type WarLogEntry struct {
	SeasonId     int                  `json:"seasonId"`
	CreatedDate  string               `json:"createdDate"`
	Participants []ClanWarParticipant `json:"participants"`
	Standings    []WarLogStanding     `json:"standings"`
}

func (w *WarLogEntry) ParseCreatedDate() (time.Time, error) {
	return time.Parse(TimeLayout, w.CreatedDate)
}

type WarLogPaging struct {
	Items []WarLogEntry `json:"items"`
	Paging struct {
		Cursors struct{} `json:"cursors"`
	} `json:"paging"`
}

type ClanWar struct {
	State             string               `json:"state"`
	CollectionEndTime string               `json:"collectionEndTime"`
	Clan              ClanWarStanding      `json:"clan"`
	Participants      []ClanWarParticipant `json:"participants"`
}

func (w *ClanWar) ParseCollectionEndTime() (time.Time, error) {
	return time.Parse(TimeLayout, w.CollectionEndTime)
}

type ClanMember struct {
	Tag               string `json:"tag"`
	Name              string `json:"name"`
	Role              string `json:"role"`
	ExpLevel          int    `json:"expLevel"`
	Trophies          int    `json:"trophies"`
	Arena             Arena  `json:"arena"`
	ClanRank          int    `json:"clanRank"`
	PreviousClanRank  int    `json:"previousClanRank"`
	Donations         int    `json:"donations"`
	DonationsReceived int    `json:"donationsReceived"`
	ClanChestPoints   int    `json:"clanChestPoints"`
}

func (c *Client) GetClan(hashtag string) (Clan, error) {
	url := fmt.Sprintf("/v1/clans/%s", normaliseHashtag(hashtag))
	req, err := c.newRequest("GET", url, nil)
	var clan Clan

	if err == nil {
		_, err = c.do(req, &clan)
	}

	return clan, err
}

func (c *Client) GetClanCurrentWar(hashtag string) (ClanWar, error) {
	url := fmt.Sprintf("/v1/clans/%s/currentwar", normaliseHashtag(hashtag))
	req, err := c.newRequest("GET", url, nil)
	var war ClanWar

	if err == nil {
		_, err = c.do(req, &war)
	}

	return war, err
}

func (c *Client) GetClanWarLog(hashtag string) (WarLogPaging, error) {
	url := fmt.Sprintf("/v1/clans/%s/warlog", normaliseHashtag(hashtag))
	req, err := c.newRequest("GET", url, nil)
	var warLog WarLogPaging

	if err == nil {
		_, err = c.do(req, &warLog)
	}

	return warLog, err
}

func (c *Client) GetClanMembers(hashtag string) (MemberPaging, error) {
	url := fmt.Sprintf("/v1/clans/%s/members", normaliseHashtag(hashtag))
	req, err := c.newRequest("GET", url, nil)
	var members MemberPaging

	if err == nil {
		_, err = c.do(req, &members)
	}

	return members, err
}

func (c *Client) SearchClans(query ClanQuery) (ClanPaging, error) {
	req, err := c.newRequest("GET", "/v1/clans", nil)
	q := req.URL.Query()

	if query.LocationId > 0 {
		q.Add("locationId", fmt.Sprintf("%d", query.LocationId))
	}

	if query.MinScore > 0 {
		q.Add("minScore", fmt.Sprintf("%d", query.MinScore))
	}

	if query.MinMembers > 1 {
		q.Add("minMembers", fmt.Sprintf("%d", query.MinMembers))
	}

	if query.MaxMembers <= 50 {
		q.Add("maxMembers", fmt.Sprintf("%d", query.MaxMembers))
	}

	if len(query.Name) > 3 {
		q.Add("name", fmt.Sprintf("%d", query.Name))
	}

	req.URL.RawQuery = q.Encode()

	var clans ClanPaging

	if err == nil {
		_, err = c.do(req, &clans)
	}

	return clans, err
}
