package jira

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Client struct {
	auth    string
	baseUrl string
	headers map[string]string
}

func CreateJiraClient(username, password, baseUrl string) *Client {
	return &Client{base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password))), baseUrl, map[string]string{}}
}

func CreateJiraClientWithToken(authToken, baseUrl string) *Client {
	return &Client{authToken, baseUrl, map[string]string{}}
}

// Perform an action on the API against this path
func (c *Client) doRequest(method string, path string, body io.Reader) (*http.Response, error) {
	c.headers["Accept"] = "application/json"
	c.headers["Authorization"] = "Basic " + c.auth
	url := c.baseUrl + path
	client := &http.Client{}
	log.Println("jira-sdk Request:", method, url)
	req, _ := http.NewRequest(method, url, body)
	for k, v := range c.headers {
		req.Header.Add(k, v)
	}
	return client.Do(req)
}
