package jira_sdk

import (
	"io"
	"log"
	"net/http"
)

type Client struct {
	username string
	password string
	baseUrl  string
	headers  map[string]string
}

func CreateJiraClient(username, password, baseUrl string) *Client {
	return &Client{username, password, baseUrl, map[string]string{}}
}

// Perform an action on the API against this path
func (c *Client) doRequest(method string, path string, body io.Reader) (*http.Response, error) {
	c.headers["Accept"] = "application/json"
	url := c.baseUrl + path
	client := &http.Client{}
	log.Println("jira-sdk Request:", method, url)
	req, _ := http.NewRequest(method, url, body)
	for k, v := range c.headers {
		req.Header.Add(k, v)
	}
	return client.Do(req)
}
