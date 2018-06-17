package clash

import "fmt"

type TournamentMember struct {
	Tag string `json:"tag"`
	Name string `json:"name"`
	Score int `json:"score"`
	Rank int `json:"rank"`
	Clan PlayerClan`json:"clan"`
}

type Tournament struct {
	Tag                 string             `json:"tag"`
	Type                string             `json:"type"`
	Status              string             `json:"status"`
	CreatorTag          string             `json:"creatorTag"`
	Name                string             `json:"name"`
	Description         string             `json:"description"`
	Capacity            int                `json:"capacity"`
	MaxCapacity         int                `json:"maxCapacity"`
	PreparationDuration int                `json:"preparationDuration"`
	Duration            int                `json:"duration"`
	CreatedTime         string             `json:"createdTime"`
	MembersList         []TournamentMember `json:"membersList"`
}

type TournamentPaging struct {
	Items []Tournament `json:"items"`
	Paging struct {
		Cursors struct{} `json:"cursors"`
	} `json:"paging"`
}

func (c *Client) GetTournament(hashtag string) (Tournament, error) {
	url := fmt.Sprintf("/v1/tournaments/%s", normaliseHashtag(hashtag))
	req, err := c.newRequest("GET", url, nil)
	var tournament Tournament

	if err == nil {
		_, err = c.do(req, &tournament)
	}

	return tournament, err
}

func (c *Client) SearchTournaments(name string) (TournamentPaging, error) {
	req, err := c.newRequest("GET", "/v1/tournaments", nil)
	q := req.URL.Query()
	q.Add("name", name)
	req.URL.RawQuery = q.Encode()

	var tournaments TournamentPaging

	if err == nil {
		_, err = c.do(req, &tournaments)
	}

	return tournaments, err
}
