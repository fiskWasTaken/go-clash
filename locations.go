package clash

import "fmt"

type LocationPager struct {
	Items []Location `json:"items"`
	Paging struct {
		Cursors struct{} `json:"cursors"`
	} `json:"paging"`
}

type Location struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	IsCountry   bool   `json:"isCountry"`
	CountryCode string `json:"countryCode,omitempty"`
}

type LocationRankingPager struct {
	Items []LocationRanking `json:"items"`
	Paging struct {
		Cursors struct{} `json:"cursors"`
	} `json:"paging"`
}

type LocationRanking struct {
	Tag          string   `json:"tag"`
	Name         string   `json:"name"`
	Rank         int      `json:"rank"`
	PreviousRank int      `json:"previousRank"`
	Location     Location `json:"location"`
	BadgeId      int      `json:"badgeId"`
	ClanScore    int      `json:"clanScore"`
	Members      int      `json:"members"`
}

func (c *Client) GetLocations() (LocationPager, error) {
	req, err := c.newRequest("GET", "/v1/locations", nil)

	var locations LocationPager

	if err == nil {
		_, err = c.do(req, &locations)
	}

	return locations, err
}

func (c *Client) GetLocation(id int) (Location, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/v1/locations/%d", id), nil)

	var location Location

	if err == nil {
		_, err = c.do(req, &location)
	}

	return location, err
}

func (c *Client) GetLocationRankings(id int, rankingType string) (LocationRankingPager, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/v1/locations/%d/rankings/%s", id, rankingType), nil)

	var rankings LocationRankingPager

	if err == nil {
		_, err = c.do(req, &rankings)
	}

	return rankings, err
}
