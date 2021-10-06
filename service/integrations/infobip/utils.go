package infobip

import (
	"io/ioutil"
	"net/http"
)

// wrapper function to execute external requests (used for mocking in unit testing)
var executeRequest = func(client *http.Client, req *http.Request) ([]byte, int, error) {
	res, err := client.Do(req)
	if err != nil {
		return nil, -1, err
	}

	var body []byte = nil
	if req.Method != "GET" {
		defer res.Body.Close()
		body, err = ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
	}

	return body, res.StatusCode, err
}