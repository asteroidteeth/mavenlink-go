package mavenlink

import (
	"flag"
	"fmt"
	"testing"
	"time"

	"github.com/rakyll/globalconf"
)

var my_token = flag.String("my_token", "", "")

func setup(t *testing.T) {
	configPath := "../config.ini"
	config, configErr := globalconf.NewWithOptions(&globalconf.Options{
		Filename: configPath,
	})
	if configErr != nil {
		t.Fatalf("Error loading config! \"%s\"\n", configErr.Error())
	}
	config.Parse()
}

// func TestTokenRequest(t *testing.T) {
// 	setup(t)
// 	const userCode string = "asdf"
// 	request, _ := createTokenRequest(userCode)
// 	if !(request.ContentLength > 0) {
// 		t.Errorf("Content length is %d", request.ContentLength)
// 	}
// 	body := make([]byte, request.ContentLength)
// 	request.Body.Read(body)
// }
// func TestGetSelf(t *testing.T) {
// 	myName := GetMavenlinkUserName(*my_token)
// 	if myName != "Andrew Taggart" {
// 		t.Fatalf("Expected Andrew Taggart but got %s", myName)
// 	}
// }

// func TestGetUsers(t *testing.T) {
// 	setup(t)
// 	client := NewClient(*my_token)
// 	users := client.GetUsers()
// 	fmt.Print(users.GetUserNames())
// }

// func TestGetTimeEntries(t *testing.T) {
// 	setup(t)
// 	client := NewClient(*my_token)
// 	timeEntries := client.GetTimeEntries(time.Now())
// 	fmt.Println(timeEntries)
// 	if timeEntries.Count == 0 {
// 		t.Error("Got 0 time entries")
// 	}
// }

// func TestGetTimePerformed(t *testing.T) {
// 	setup(t)
// 	client := NewClient(*my_token)
// 	timeEntries := client.GetTimeEntries(time.Now())
// 	if timeEntries.Count == 0 {
// 		t.Fatal("Got 0 time entries")
// 	}
// 	for _, v := range timeEntries.Time_entries {
// 		fmt.Println(v.Date_performed)
// 	}
// }

// func TestGetMyTimeEntries(t *testing.T) {
// 	setup(t)
// 	client := NewClient(*my_token)
// 	me := client.GetSelf()
// 	timeEntries := client.GetTimeEntries(time.Now())
// 	if timeEntries.Count == 0 {
// 		t.Fatal("Got 0 time entries")
// 	}

// }
