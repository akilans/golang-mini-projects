package main

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"time"
)

// websites needs to be monitored
var websites = map[string]int{
	"https://facebook.com": 200,
	"https://google.com":   200,
	"https://twitter.com":  200,
	"http://localhost":     200,
}

// 5 seconds interval
const checkInterval int = 5

// 5 mins interval for alerting user for same web down
const reminderInterval int = 5

// to track web status
type webStatus struct {
	web         string
	status      string
	lastFailure time.Time
}

// email config

// Sender data.
var from string = os.Getenv("GMAIL_ID")
var password string = os.Getenv("GMAIL_PASSWORD")

// Receiver email address.
var to = []string{
	os.Getenv("GMAIL_ID"),
}

// smtp server configuration.
var smtpHost string = "smtp.gmail.com"
var smtpPort string = "587"

func main() {

	// initial web status slice - empty.
	// It start stores web status info in case of down
	webStatusSlice := []webStatus{}

	// infinite loop to keep on looping
	for {

		if len(webStatusSlice) == 0 {
			fmt.Println("All websites are up!!!")
		}

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
		fmt.Printf("Sleep for %v seconds \n", checkInterval)
		time.Sleep(time.Duration(checkInterval) * time.Second)
	}
}

func alertUser(web string, err error, webStatusSlice *[]webStatus) {

	// this info has to be appended in status slice
	downWebInfo := webStatus{web, "down", time.Now()}

	if len(*webStatusSlice) > 0 {
		// check for this web down is tracked in webstatus slice
		// if no, then it is down for first time then trigger email
		// if yes, then check for which time it triggered email
		// If it is above the reminderInterval then trigger email
		previousAlert := checkForPreviousAlert(webStatusSlice, web)

		if !previousAlert {
			fmt.Printf("%v added to alert list\n", web)
			*webStatusSlice = append(*webStatusSlice, downWebInfo)
			triggerEmail(web)
		} else {
			fmt.Printf("%v already in alert list\n", web)
			// check when it triggered email last time
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
}

// alert user by sending an email
func triggerEmail(web string) {

	//message
	//message := []byte("This is a test email message.")
	message := []byte("Subject: Web Monitor Alert \r\n\r\n" + web + " - Website is down\r\n")
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}

// check for the web is in webstatus slice
func checkForPreviousAlert(webStatusSlice *[]webStatus, web string) bool {

	alreadyDown := false

	for _, webStatusInfo := range *webStatusSlice {
		if webStatusInfo.web == web {
			alreadyDown = true
		}
	}

	if !alreadyDown {
		//fmt.Printf("%v not in the alert list\n", web)
		return false
	} else {
		//fmt.Printf("%v already in list\n", web)
		return true
	}

}

// to when it trigger email last
// last time triggered email + reminder interval > current timem then return true
// then update time
func checkForReminderInterval(webStatusSlice *[]webStatus, web string) bool {
	triggerAnother := false

	for i, webStatusInfo := range *webStatusSlice {
		if webStatusInfo.web == web {
			lastFailurePlusReminderMins := webStatusInfo.lastFailure.Add(time.Duration(reminderInterval) * time.Minute)
			if lastFailurePlusReminderMins.Before(time.Now()) {
				triggerAnother = true
				// update with current time
				(*webStatusSlice)[i] = webStatus{web, "down", time.Now()}
				fmt.Printf("%v - Time for new alert\n", web)
			} else {
				fmt.Printf("%v - Next alert will be send after reminder interval!!!\n", web)
			}
		}
	}

	return triggerAnother
}
