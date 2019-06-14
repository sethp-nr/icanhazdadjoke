package icanhazdadjoke

import (
	"io/ioutil"
	"net/http"
)

const UserAgent = "ICanHazDadJoke go client (https://github.com/sethp-nr/icanhazdadjoke)"
const DefaultUrl = "https://icanhazdadjoke.com/"

var DefaultClient = &Client{
	Client: http.DefaultClient,

	url: DefaultUrl,
}

type Client struct {
	*http.Client

	url string
}

type FuncOpt func(*Client)

func NewClient(opts ...FuncOpt) *Client {
	client := &Client{
		Client: http.DefaultClient,

		url: DefaultUrl,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func GetRandomJokeText() (string, error) {
	return DefaultClient.GetRandomJokeText()
}

func (c *Client) GetRandomJokeText() (string, error) {
	req, err := http.NewRequest("GET", c.url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Accept", "text/plain")

	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
