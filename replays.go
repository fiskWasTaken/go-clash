package clash

import "fmt"

type Replay struct {
	BattleTime string `json:"battleTime"`
	// Replay data is hideously unstructured, so let's save some time.
	ReplayData map[string]interface{} `json:"replayData"`
	ShareCount int `json:"shareCount"`
	Tag string `json:"tag"`
	ViewCount int `json:"viewCount"`
}

type ReplayService struct {
	c   *Client
	tag string
}

func (c *Client) Replay(tag string) *ReplayService {
	return &ReplayService{c, tag}
}

// Get information about a single replay by a replay tag.
func (i *ReplayService) Get() (Replay, error) {
	url := fmt.Sprintf("/v1/replays/%s", NormaliseTag(i.tag))
	req, err := i.c.NewRequest("GET", url, nil)
	var replay Replay

	if err == nil {
		_, err = i.c.Do(req, &replay)
	}

	return replay, err
}