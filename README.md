# Helper Functions for Go

## Installation
`go get github.com/loyalid/helper-go`

## Functions
```
// ToDateTimeString converts DateTime into string with Y-m-d H:i:s format
func ToDateTimeString(dateTime time.Time) string {}

// LogError logs the error with message
func LogError(err error, message string, vendor map[string]interface{}) {}

// RandomLetter generates random letter with custom length
func RandomLetter(n int) string {}

// RandomInteger returns random integer between parameters
func RandomInteger(min int, max int) int {}

// JSONEncode converts data into JSON string
func JSONEncode(data interface{}) (string, error) {}

// InArray checks if a value exists in an array
func InArray(needle string, haystack []interface{}) bool {}
```