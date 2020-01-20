package jira

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_GetTicketByKey_TestParsing(t *testing.T) {
	srv := createTestHttpServer(response)
	c := CreateJiraClient("nil", "nil", srv.URL)
	ti, err := c.GetTicketByKey("HS-1015")
	defer srv.Close()
	assert.NoError(t, err)
	assert.Equal(t, "HS-1015", ti.Key)
}

func createTestHttpServer(response string) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Println("Values:", k, v)
		}
		fmt.Println("URL:", r.URL)
		w.Write([]byte(response))
	}))
	return ts
}

var response = `{
	"expand": "names,schema",
	"startAt": 0,
	"maxResults": 50,
	"total": 1,
	"issues": [{
		"expand": "operations,versionedRepresentations,editmeta,changelog,transitions,renderedFields",
		"id": "60829",
		"self": "http://localhost:8080/rest/api/2/issue/60829",
		"key": "HS-1015",
		"fields": {
			"issuetype": {
				"self": "http://localhost:8080/rest/api/2/issuetype/3",
				"id": "3",
				"description": "A task that needs to be done.",
				"iconUrl": "http://localhost:8080/secure/viewavatar?size=xsmall&avatarId=10618&avatarType=issuetype",
				"name": "Task",
				"subtask": false,
				"avatarId": 10618
			},
			"timespent": null,
			"project": {
				"self": "http://localhost:8080/rest/api/2/project/13400",
				"id": "13400",
				"key": "HS",
				"name": "HiHi Software",
				"avatarUrls": {
					"48x48": "http://localhost:8080/secure/projectavatar?pid=13400&avatarId=14300",
					"24x24": "http://localhost:8080/secure/projectavatar?size=small&pid=13400&avatarId=14300",
					"16x16": "http://localhost:8080/secure/projectavatar?size=xsmall&pid=13400&avatarId=14300",
					"32x32": "http://localhost:8080/secure/projectavatar?size=medium&pid=13400&avatarId=14300"
				}
			},
			"fixVersions": [],
			"customfield_11001": null,
			"aggregatetimespent": null,
			"resolution": null,
			"resolutiondate": null,
			"workratio": -1,
			"lastViewed": "2020-01-20T09:25:43.731+0000",
			"watches": {
				"self": "http://localhost:8080/rest/api/2/issue/HS-1015/watchers",
				"watchCount": 2,
				"isWatching": true
			},
			"created": "2020-01-06T11:41:31.000+0000",
			"priority": {
				"self": "http://localhost:8080/rest/api/2/priority/3",
				"iconUrl": "http://localhost:8080/images/icons/priorities/major.svg",
				"name": "Major",
				"id": "3"
			},
			"labels": [],
			"timeestimate": null,
			"aggregatetimeoriginalestimate": null,
			"versions": [],
			"issuelinks": [],
			"assignee": {
				"self": "http://localhost:8080/rest/api/2/user?username=name",
				"name": "name",
				"key": "name",
				"emailAddress": "person.name@email.com",
				"avatarUrls": {
					"48x48": "http://localhost:8080/secure/useravatar",
					"24x24": "http://localhost:8080/secure/useravatar",
					"16x16": "http://localhost:8080/secure/useravatar",
					"32x32": "http://localhost:8080/secure/useravatar"
				},
				"displayName": "Person Name",
				"active": true,
				"timeZone": "Europe/London"
			},
			"updated": "2020-01-20T08:16:32.000+0000",
			"status": {
				"self": "http://localhost:8080/rest/api/2/status/3",
				"description": "This issue is being actively worked on at the moment by the assignee.",
				"iconUrl": "http://localhost:8080/images/icons/statuses/inprogress.png",
				"name": "In Progress",
				"id": "3",
				"statusCategory": {
					"self": "http://localhost:8080/rest/api/2/statuscategory/4",
					"id": 4,
					"key": "indeterminate",
					"colorName": "yellow",
					"name": "In Progress"
				}
			},
			"components": [],
			"timeoriginalestimate": null,
			"description": "Upgrade the next 2 TeamCity build agents",
			"aggregatetimeestimate": null,
			"summary": "Upgrade the next 2 TeamCity build agents",
			"creator": {
				"self": "http://localhost:8080/rest/api/2/user?username=underwoodk",
				"name": "name",
				"key": "name",
				"emailAddress": "person.name@email.com",
				"avatarUrls": {
					"48x48": "http://localhost:8080/secure/useravatar?ownerId=underwoodk&avatarId=14100",
					"24x24": "http://localhost:8080/secure/useravatar?size=small&ownerId=underwoodk&avatarId=14100",
					"16x16": "http://localhost:8080/secure/useravatar?size=xsmall&ownerId=underwoodk&avatarId=14100",
					"32x32": "http://localhost:8080/secure/useravatar?size=medium&ownerId=underwoodk&avatarId=14100"
				},
				"displayName": "Person Name",
				"active": true,
				"timeZone": "Europe/London"
			},
			"subtasks": [],
			"reporter": {
				"self": "http://localhost:8080/rest/api/2/user?username=name",
				"name": "Name",
				"key": "name",
				"emailAddress": "person.name@email.com",
				"avatarUrls": {
					"48x48": "http://localhost:8080/secure/useravatar",
					"24x24": "http://localhost:8080/secure/useravatar",
					"16x16": "http://localhost:8080/secure/useravatar",
					"32x32": "http://localhost:8080/secure/useravatar"
				},
				"displayName": "Person Name",
				"active": true,
				"timeZone": "Europe/London"
			},
			"customfield_10000": null,
			"aggregateprogress": {
				"progress": 0,
				"total": 0
			},
			"environment": null,
			"duedate": null,
			"progress": {
				"progress": 0,
				"total": 0
			},
			"votes": {
				"self": "http://localhost:8080/rest/api/2/issue/HS-1015/votes",
				"votes": 0,
				"hasVoted": false
			}
		}
	}]
}`
