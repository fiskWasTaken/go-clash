package clash

import (
	"fmt"
	"time"
)

type TournamentQuery struct {
	PagedQuery
	Name string
}

type TournamentMember struct {
	Tag   string     `json:"tag"`
	Name  string     `json:"name"`
	Score int        `json:"score"`
	Rank  int        `json:"rank"`
	Clan  PlayerClan `json:"clan"`
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
	RawCreatedTime      string             `json:"createdTime"`
	RawStartedTime      string             `json:"startedTime"`
	MembersList         []TournamentMember `json:"membersList"`
}

func (t *Tournament) CreatedTime() time.Time {
	parsed, _ := time.Parse(TimeLayout, t.RawCreatedTime)
	return parsed
}

func (t *Tournament) StartedTime() time.Time {
	parsed, _ := time.Parse(TimeLayout, t.RawStartedTime)
	return parsed
}

type TournamentPager struct {
	Items  []Tournament `json:"items"`
	Paging Paging       `json:"paging"`
}

type TournamentService struct {
	c   *Client
	tag string
}

type TournamentsService struct {
	c *Client
}

func (c *Client) Tournaments() *TournamentsService {
	return &TournamentsService{c}
}

func (c *Client) Tournament(tag string) *TournamentService {
	return &TournamentService{c, tag}
}

// Get information about a single tournament by a tournament tag.
func (i *TournamentService) Get() (Tournament, error) {
	url := fmt.Sprintf("/v1/tournaments/%s", normaliseTag(i.tag))
	req, err := i.c.newRequest("GET", url, nil)
	var tournament Tournament

	if err == nil {
		_, err = i.c.do(req, &tournament)
	}

	return tournament, err
}

// Search all tournaments by name.
//
// It is not possible to specify ordering for results so clients should not
// rely on any specific ordering as that may change in the future releases of the API.
func (i *TournamentsService) Search(query *TournamentQuery) (TournamentPager, error) {
	req, err := i.c.newRequest("GET", "/v1/tournaments", nil)
	q := req.URL.Query()

	q.Add("name", query.Name)

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

	var tournaments TournamentPager

	if err == nil {
		_, err = i.c.do(req, &tournaments)
	}

	return tournaments, err
}
