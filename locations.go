package clash

import "fmt"

type LocationPager struct {
	Items  []Location `json:"items"`
	Paging Paging     `json:"paging"`
}

type Location struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	IsCountry   bool   `json:"isCountry"`
	CountryCode string `json:"countryCode,omitempty"`
}

type LocationClanRankingPager struct {
	Items  []LocationClanRanking `json:"items"`
	Paging Paging                `json:"paging"`
}

type LocationPlayerRankingPager struct {
	Items  []LocationPlayerRanking `json:"items"`
	Paging Paging                  `json:"paging"`
}

type LocationClanRanking struct {
	Tag          string   `json:"tag"`
	Name         string   `json:"name"`
	Rank         int      `json:"rank"`
	PreviousRank int      `json:"previousRank"`
	Location     Location `json:"location"`
	BadgeId      int      `json:"badgeId"`
	ClanScore    int      `json:"clanScore"`
	Members      int      `json:"members"`
}

type LocationPlayerRanking struct {
	Tag          string     `json:"tag"`
	Name         string     `json:"name"`
	ExpLevel     int        `json:"expLevel"`
	Trophies     int        `json:"trophies"`
	Clan         PlayerClan `json:"clan"`
	Rank         int        `json:"Rank"`
	PreviousRank int        `json:"previousRank"`
	Arena        Arena      `json:"arena"`
}

type LocationsInterface struct {
	c *Client
}

type LocationInterface struct {
	c  *Client
	id int
}

func (c *Client) Locations() *LocationsInterface {
	return &LocationsInterface{c}
}

func (c *Client) Location(id int) *LocationInterface {
	return &LocationInterface{c, id}
}

// List all available locations
func (i *LocationsInterface) All() (LocationPager, error) {
	req, err := i.c.newRequest("GET", "/v1/locations", nil)

	var locations LocationPager

	if err == nil {
		_, err = i.c.do(req, &locations)
	}

	return locations, err
}

// Get information about specific location
func (i *LocationInterface) Get() (Location, error) {
	req, err := i.c.newRequest("GET", fmt.Sprintf("/v1/locations/%d", i.id), nil)

	var location Location

	if err == nil {
		_, err = i.c.do(req, &location)
	}

	return location, err
}

// Get clan rankings for a specific location
func (i *LocationInterface) ClanRankings(query *PagedQuery) (LocationClanRankingPager, error) {
	req, err := i.c.newRequest("GET", fmt.Sprintf("/v1/locations/%d/rankings/clans", i.id), nil)

	q := req.URL.Query()

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

	var rankings LocationClanRankingPager

	if err == nil {
		_, err = i.c.do(req, &rankings)
	}

	return rankings, err
}

// Get player rankings for a specific location
func (i *LocationInterface) PlayerRankings(query *PagedQuery) (LocationPlayerRankingPager, error) {
	req, err := i.c.newRequest("GET", fmt.Sprintf("/v1/locations/%d/rankings/players", i.id), nil)

	q := req.URL.Query()

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

	var rankings LocationPlayerRankingPager

	if err == nil {
		_, err = i.c.do(req, &rankings)
	}

	return rankings, err
}

// Get clan war rankings for a specific location
func (i *LocationInterface) ClanWarRankings(query *PagedQuery) (LocationClanRankingPager, error) {
	req, err := i.c.newRequest("GET", fmt.Sprintf("/v1/locations/%d/rankings/clanwars", i.id), nil)

	q := req.URL.Query()

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

	var rankings LocationClanRankingPager

	if err == nil {
		_, err = i.c.do(req, &rankings)
	}

	return rankings, err
}
