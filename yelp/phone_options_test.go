package yelp

import (
	"testing"
)

// TestPhoneSearch will check PhoneSearch using PhoneOptions
func TestPhoneSearch(t *testing.T) {
	client := getClient(t)
	options := PhoneOptions{Phone: "4158632800"} // # of Mission Chinese
	result, err := client.PhoneSearch(options)
	check(t, err)
	assert(t, len(result.Businesses) > 0, containsResults)
}

// TestPhoneSearchWithoutNumber ensures validation for a phone number when doing a PhoneSearch
func TestPhoneSearchWithoutNumber(t *testing.T) {
	client := getClient(t)
	options := PhoneOptions{}
	_, err := client.PhoneSearch(options)
	assert(t, err == errUnspecifiedPhone, shouldRequirePhone)
}
