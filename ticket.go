package jira_sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type TicketSearchResults struct {
	Total  int          `json:"total"`
	Issues []JiraTicket `json:"issues"`
}

type JiraTicket struct {
	Id          string `json:"id"`
	Self        string `json:"self"`
	Key         string `json:"key"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Fields      Fields `json:"fields"`
}

type Fields struct {
	IssueType IssueType `json:"issuetype"`
}

type IssueType struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	IconUrl     string `json:"iconUrl"`
	Name        string `json:"name"`
	SubTask     string `json:"subtask"`
}

func (c Client) GetTicketByKey(id string) (*JiraTicket, error) {
	path := fmt.Sprintf("/rest/api/2/search?jql=id=%s", id)
	resp, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	bo, be := ioutil.ReadAll(resp.Body)
	if be != nil {
		return nil, be
	}
	list := &TicketSearchResults{}
	je := json.Unmarshal(bo, list)
	if je != nil {
		return nil, je
	}
	return &list.Issues[0], nil
}
