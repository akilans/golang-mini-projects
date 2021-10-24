package main

import (
	"fmt"
	"net/http"
	"time"
)

// websites needs to be monitored
var websites = map[string]int{
	"https://facebook.com":   202,
	"https://google.com":     200,
	"https://twitter.com":    200,
	"https://localhost:8000": 200,
}

// 5 seconds interval
const checkInterval int = 5

type webStatus struct {
	web         string
	status      string
	lastFailure time.Time
}

// 15 mins interval for alerting user for same web
const reminderInterval int = 1

func main() {

	webStatusSlice := []webStatus{}

	// infinite loop to keep on looping
	for {

		for web, expectedStatusCode := range websites {

			res, err := http.Get(web)

			if err != nil {
				// in case of website down/connection refused
				alertUser(web, err, &webStatusSlice)
				continue
			} else {
				// check the response code
				if res.StatusCode != expectedStatusCode {
					errMsg := fmt.Errorf("%v is down", web)
					alertUser(web, errMsg, &webStatusSlice)
				}
			}
		}
		// sleep for $checkInterval seconds
		time.Sleep(time.Duration(checkInterval) * time.Second)
	}
}

func alertUser(web string, err error, webStatusSlice *[]webStatus) {

	downWebInfo := webStatus{web, "down", time.Now()}

	if len(*webStatusSlice) > 0 {
		previousAlert := checkForPreviousAlert(webStatusSlice, web)

		if !previousAlert {
			fmt.Printf("%v added to alert list\n", web)
			*webStatusSlice = append(*webStatusSlice, downWebInfo)
			triggerEmail(web)
		} else {
			fmt.Printf("%v already in alert list\n", web)
			triggerAnother := checkForReminderInterval(webStatusSlice, web)

			if triggerAnother {
				triggerEmail(web)
			}
		}
	} else {
		fmt.Printf("%v added to alert list\n", web)
		*webStatusSlice = append(*webStatusSlice, downWebInfo)
		triggerEmail(web)
	}

	//fmt.Printf("%v is down with error %v \n", web, err)
	//fmt.Println(webStatusSlice)

}

func triggerEmail(web string) {
	fmt.Printf("%v - Triggered email\n", web)
}

func checkForPreviousAlert(webStatusSlice *[]webStatus, web string) bool {

	alreadyDown := false

	for _, webStatusInfo := range *webStatusSlice {
		if webStatusInfo.web == web {
			alreadyDown = true
		}
	}

	if !alreadyDown {
		fmt.Printf("%v not in the alert list\n", web)
		return false
	} else {
		fmt.Printf("%v already in list\n", web)
		return true
	}

}

func checkForReminderInterval(webStatusSlice *[]webStatus, web string) bool {
	triggerAnother := false

	for i, webStatusInfo := range *webStatusSlice {
		if webStatusInfo.web == web {
			lastFailurePlus15Mins := webStatusInfo.lastFailure.Add(time.Duration(reminderInterval) * time.Minute)
			if lastFailurePlus15Mins.Before(time.Now()) {
				triggerAnother = true
				(*webStatusSlice)[i] = webStatus{web, "down", time.Now()}
				fmt.Printf("%v - Time for new alert\n", web)
			}
		}
	}

	return triggerAnother
}
