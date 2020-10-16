package main

import (
	"strings"

	requests "github.com/asmcos/requests"
)

func main() {
	apiID := "changeMe!"
	message := "Hello world!"
	phoneNumber := "89211232323"
	smsURL := "https://sms.ru/sms/send?api_id=key&to=number&msg=message&json=1"
	smsURL = strings.Replace(smsURL, "key", apiID, 1)
	smsURL = strings.Replace(smsURL, "number", phoneNumber, 1)
	smsURL = strings.Replace(smsURL, "message", message, 1)
	resp, err := requests.Get(smsURL)
	if err != nil {
		return
	}
	println(resp.Text())
}
