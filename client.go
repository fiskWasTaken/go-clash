package clash

import (
	"net/url"
	"net/http"
	"io"
	"bytes"
	"encoding/json"
	"fmt"
)

type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	Bearer     string
	httpClient http.Client
}

type ErrorBody struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

type APIError struct {
	Response *http.Response
	Body     *ErrorBody
}

func (e *APIError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Response.StatusCode, e.Body.Message)
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Bearer))
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		errorResponse := &ErrorBody{}
		err = json.NewDecoder(resp.Body).Decode(errorResponse)

		if err == nil {
			err = &APIError{resp, errorResponse}
		}
	} else {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}

func normaliseHashtag(hashtag string) string {
	if hashtag[0] == '#' {
		return hashtag
	}

	return "#" + hashtag
}
