package clash

import (
	"fmt"
)

type ClanLocation struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsCountry bool   `json:"isCountry"`
}

type MemberList struct {
	Items []ClanMember `json:"items"`
	Paging struct {
		Cursors struct{} `json:"cursors"`
	} `json:"paging"`
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

type Clan struct {
	Tag               string       `json:"tag"`
	Name              string       `json:"name"`
	Type              string       `json:"type"`
	Description       string       `json:"description"`
	BadgeId           int          `json:"badgeId"`
	ClanScore         int          `json:"clanScore"`
	Location          ClanLocation `json:"location"`
	RequiredTrophies  int          `json:"requiredTrophies"`
	DonationsPerWeek  int          `json:"donationsPerWeek"`
	ClanChestStatus   string       `json:"clanChestStatus"`
	ClanChestPoints   int          `json:"clanChestPoints"`
	ClanChestLevel    int          `json:"clanChestLevel"`
	ClanChestMaxLevel int          `json:"clanChestMaxLevel"`
	Members           int          `json:"members"`
	MemberList        []ClanMember `json:"memberList"`
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

// todo
func (c *Client) GetClanCurrentWar(hashtag string) (Clan, error) {
	url := fmt.Sprintf("/v1/clans/%s/currentwar", normaliseHashtag(hashtag))
	req, err := c.newRequest("GET", url, nil)
	var clan Clan

	if err == nil {
		_, err = c.do(req, &clan)
	}

	return clan, err
}

// todo
func (c *Client) GetClanWarLog(hashtag string) (Clan, error) {
	url := fmt.Sprintf("/v1/clans/%s/warlog", normaliseHashtag(hashtag))
	req, err := c.newRequest("GET", url, nil)
	var clan Clan

	if err == nil {
		_, err = c.do(req, &clan)
	}

	return clan, err
}

func (c *Client) GetClanMembers(hashtag string) (MemberList, error) {
	url := fmt.Sprintf("/v1/clans/%s/members", normaliseHashtag(hashtag))
	req, err := c.newRequest("GET", url, nil)
	var members MemberList

	if err == nil {
		_, err = c.do(req, &members)
	}

	return members, err
}

// todo
func (c *Client) SearchClans(name string) ([]Clan, error) {
	req, err := c.newRequest("GET", "/v1/clans", map[string]string{"name": name})
	var clans []Clan

	if err == nil {
		_, err = c.do(req, &clans)
	}

	return clans, err
}
