package helper

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"
	"unsafe"

	"github.com/loyalid/slack-incoming-webhook-go"
)

// ToDateTimeString converts DateTime into string with Y-m-d H:i:s format
func ToDateTimeString(dateTime time.Time) string {
	return dateTime.Format("2006-01-02 15:04:05")
}

// LogError logs the error with message
func LogError(err error, message string, vendor map[string]interface{}) {
	log.Printf("%s: %s", message, err)

	if vendor["slack_enabled"].(bool) {
		slackMessage := "[" + vendor["app_name"].(string) + "]\r\n" + message + ": " + err.Error()
		err = slack.PostMessage(vendor["slack_webhook_url"].(string), slackMessage)
		if err != nil {
			log.Printf("Failed to post message to Slack: %s", err.Error())
		}
	}
}

// RandomLetter generates random letter with custom length
func RandomLetter(n int) string {
	const (
		letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		letterIdxBits = 6
		letterIdxMask = 1<<letterIdxBits - 1
		letterIdxMax  = 63 / letterIdxBits
	)

	var (
		src = rand.NewSource(time.Now().UnixNano())
		b   = make([]byte, n)
	)

	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}

		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}

		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// RandomInteger returns random integer between parameters
func RandomInteger(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

// JSONEncode converts data into JSON string
func JSONEncode(data interface{}) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// InArray checks if a value exists in an array
func InArray(needle string, haystack []interface{}) bool {
	for _, value := range haystack {
		if needle == value {
			return true
		}
	}

	return false
}
