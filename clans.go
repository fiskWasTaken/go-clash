package clash

import (
	"fmt"
	"time"
)

type ClanQuery struct {
	PagedQuery
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
	Items  []Clan `json:"items"`
	Paging Paging `json:"paging"`
}

type MemberPaging struct {
	Items  []ClanMember `json:"items"`
	Paging Paging       `json:"paging"`
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
	SeasonId       int                  `json:"seasonId"`
	RawCreatedDate string               `json:"createdDate"`
	Participants   []ClanWarParticipant `json:"participants"`
	Standings      []WarLogStanding     `json:"standings"`
}

func (w *WarLogEntry) CreatedDate() time.Time {
	parsed, _ := time.Parse(TimeLayout, w.RawCreatedDate)
	return parsed
}

type WarLogPaging struct {
	Items  []WarLogEntry `json:"items"`
	Paging Paging        `json:"paging"`
}

type ClanWar struct {
	State                string               `json:"state"`
	RawCollectionEndTime string               `json:"collectionEndTime"`
	Clan                 ClanWarStanding      `json:"clan"`
	Participants         []ClanWarParticipant `json:"participants"`
}

func (w *ClanWar) CollectionEndTime() time.Time {
	parsed, _ := time.Parse(TimeLayout, w.RawCollectionEndTime)
	return parsed
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

type ClansInterface struct {
	c *Client
}

type ClanInterface struct {
	c   *Client
	tag string
}

func (c *Client) Clans() *ClansInterface {
	return &ClansInterface{c}
}

func (c *Client) Clan(tag string) *ClanInterface {
	return &ClanInterface{c, tag}
}

// Get information about a single clan by clan tag.
// Clan tags can be found using clan search operation.
func (i *ClanInterface) Get() (Clan, error) {
	url := fmt.Sprintf("/v1/clans/%s", normaliseTag(i.tag))
	req, err := i.c.newRequest("GET", url, nil)
	var clan Clan

	if err == nil {
		_, err = i.c.do(req, &clan)
	}

	return clan, err
}

// Retrieve information about clan's current clan war
func (i *ClanInterface) CurrentWar() (ClanWar, error) {
	url := fmt.Sprintf("/v1/clans/%s/currentwar", normaliseTag(i.tag))
	req, err := i.c.newRequest("GET", url, nil)
	var war ClanWar

	if err == nil {
		_, err = i.c.do(req, &war)
	}

	return war, err
}

// Retrieve clan's clan war log
func (i *ClanInterface) WarLog() (WarLogPaging, error) {
	url := fmt.Sprintf("/v1/clans/%s/warlog", normaliseTag(i.tag))
	req, err := i.c.newRequest("GET", url, nil)
	var warLog WarLogPaging

	if err == nil {
		_, err = i.c.do(req, &warLog)
	}

	return warLog, err
}

// List clan members
func (i *ClanInterface) Members() (MemberPaging, error) {
	url := fmt.Sprintf("/v1/clans/%s/members", normaliseTag(i.tag))
	req, err := i.c.newRequest("GET", url, nil)
	var members MemberPaging

	if err == nil {
		_, err = i.c.do(req, &members)
	}

	return members, err
}

// Search all clans by name and/or filtering the results using various criteria.
// At least one filtering criteria must be defined and if name is used
// as part of search, it is required to be at least three characters long.
func (i *ClansInterface) Search(query *ClanQuery) (ClanPaging, error) {
	req, err := i.c.newRequest("GET", "/v1/clans", nil)
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

	if len(query.Name) >= 3 {
		q.Add("name", query.Name)
	}

	if query.Limit > 0 {
		q.Add("limit", fmt.Sprintf("%d", query.Limit))
	}

	if query.After > 0 {
		q.Add("after", fmt.Sprintf("%d", query.After))
	}

	if query.Before > 0 {
		q.Add("before", fmt.Sprintf("%d", query.Before))
	}

	req.URL.RawQuery = q.Encode()

	var clans ClanPaging

	if err == nil {
		_, err = i.c.do(req, &clans)
	}

	return clans, err
}
