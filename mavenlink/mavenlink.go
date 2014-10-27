package mavenlink

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

var (
	AppId       = flag.String("app_id", "", "Mavenlink app id")
	SecretToken = flag.String("secret", "", "Mavenlink app secret ")
	RedirectUri = flag.String("redirect_uri", "", "Mavenlink app callback uri")
)

const (
	ymdFormat = "2006-01-02"
	// ISO8601_dayStart = "T000000"
	// ISO8601_dayEnd   = "T115959"
)

const AuthUrl = "https://app.mavenlink.com/oauth/authorize?response_type=code&client_id=%s&redirect_uri=%s"

type AccessData struct {
	Access_token string
}

type Client struct {
	Token string
	Root  url.URL
}

func NewClient(token string) Client {
	u, _ := url.Parse("https://app.mavenlink.com/api/v1/")
	return Client{token, *u}
}

func (c *Client) get(resourcePath string, query *map[string]string) []byte {
	resource, _ := c.Root.Parse(resourcePath)
	request, _ := http.NewRequest("GET", resource.String(), nil)
	urlQuery := request.URL.Query()
	if query != nil {
		for k, v := range *query {
			urlQuery.Add(k, v)
		}
	}
	request.URL.RawQuery = urlQuery.Encode()
	log.Printf("FILTER=%s", request.URL.Query().Encode())

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
	}
	if response.ContentLength < 1 {
		log.Println(request.Header.Get("Authorization"))
		log.Printf("Response code: %d, Response message: %s\n\n\n", response.StatusCode, response.Status)
	}
	return readBody(response)

}

func (c *Client) GetUsers() *Users {
	jsonData := c.get("users.json", nil)
	users := &Users{}
	json.Unmarshal(jsonData, users)
	return users
}

func (c *Client) GetSelf() *User {
	jsonData := c.get("users/me.json", nil)
	users := &Users{}
	json.Unmarshal(jsonData, users)
	user := users.Users[users.GetUserIds()[0]]
	return &user
}

func (c *Client) GetTimeEntries(day time.Time) *TimeEntries {
	dayDuration := time.Hour * 24 * 2
	startTime := day.Add(-dayDuration).Format(ymdFormat)
	endTime := day.Add(dayDuration).Format(ymdFormat)
	timeRange := fmt.Sprintf("%s:%s", startTime, endTime)

	log.Printf("DATE PERFORMED RANGE: %s", timeRange)
	jsonData := c.get("time_entries.json", &map[string]string{"date_performed_between": timeRange})
	fmt.Println(string(jsonData))
	entries := &TimeEntries{}
	err := json.Unmarshal(jsonData, entries)
	if err != nil {
		log.Println(err)
	}
	return entries
}

func readBody(r *http.Response) []byte {
	var buffer bytes.Buffer
	buffer.ReadFrom(r.Body)
	return buffer.Bytes()
}

func ExchangeForToken(code string) string {
	log.Println("Getting token!")
	tokenRequest, err := createTokenRequest(code)
	client := http.Client{}
	response, err := client.Do(tokenRequest)

	if err != nil {
		log.Println("Error in request:")
		log.Println(err.Error())
		buf := new(bytes.Buffer)
		buf.ReadFrom(tokenRequest.Body)
		s := buf.String()
		log.Println(s)
	}
	if response.ContentLength < 0 {
		responsedump, _ := httputil.DumpResponse(response, true)
		log.Println(string(responsedump))
		log.Printf("Response code: %d, Response message: %s\n\n\n", response.StatusCode, response.Status)
	}
	log.Println(response.ContentLength)

	var buffer bytes.Buffer
	buffer.ReadFrom(response.Body)
	data := buffer.Bytes()

	accessData := AccessData{}

	err = json.Unmarshal(data, &accessData)
	if err != nil {
		log.Printf("Error reading JSON: %s", err.Error())
	}
	return accessData.Access_token
}

func createTokenRequest(code string) (*http.Request, error) {
	params := fmt.Sprintf("client_id=%s&client_secret=%s&grant_type=authorization_code&code=%s&redirect_uri=%s",
		*AppId, *SecretToken, code, url.QueryEscape(*RedirectUri))
	request, err := http.NewRequest("POST", "https://app.mavenlink.com/oauth/token", bytes.NewBufferString(params))
	return request, err
}

func GetMavenlinkUserName(code string) string {

	log.Println("Getting user name!")
	request, _ := http.NewRequest("GET", "https://api.mavenlink.com/api/v1/users/me.json", nil)

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", code))
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
	}
	if response.ContentLength < 1 {
		log.Println(request.Header.Get("Authorization"))
		log.Printf("Response code: %d, Response message: %s\n\n\n", response.StatusCode, response.Status)
	}
	data := readBody(response)
	users := Users{}
	json.Unmarshal(data, &users)
	jsonrep, _ := json.Marshal(&users)
	log.Println(string(jsonrep))

	return users.Users[users.GetUserIds()[0]].Full_name
}
