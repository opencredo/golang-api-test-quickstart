package utils

import (
	"errors"
	"gopkg.in/resty.v0"
	"os"
)

// for some reason resty doesn't by default return errors
func RegisterErrors() {
	resty.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {

		if (r.StatusCode() > 399) {
			return errors.New("Status code error: " + string(r.StatusCode()))
		}

		return nil
	})
}

func GetApiKey() (string) {
	apiKey := os.Getenv("TWFY_KEY")

	if (apiKey == "") {
		panic("No api key (TWFY_KEY) present in environment variables")
	}

	return apiKey
}