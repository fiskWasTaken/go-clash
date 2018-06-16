package clash

import "fmt"

type Tournament struct {
}

// todo
func (c *Client) GetTournament(id int) (Tournament, error) {
	url := fmt.Sprintf("/v1/tournaments/%s", id)
	req, err := c.newRequest("GET", url, nil)
	var tournament Tournament

	if err == nil {
		_, err = c.do(req, &tournament)
	}

	return tournament, err
}

// todo
func (c *Client) SearchTournaments(name string) ([]Tournament, error) {
	req, err := c.newRequest("GET", "/v1/tournaments", map[string]string{"name": name})
	var tournaments []Tournament

	if err == nil {
		_, err = c.do(req, &tournaments)
	}

	return tournaments, err
}
