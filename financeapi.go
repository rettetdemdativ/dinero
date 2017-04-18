// Author(s): Michael Koeppl

package dinero

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	url     = "http://rate-exchange-1.appspot.com/currency?from=%s&to=%s"
	timeout = time.Second * 5
)

// Result represents the JSON response from the Yahoo finance API.
type result struct {
	// The ",string" part tells json's Unmarshal function
	// that Rate is a string encoded float.
	Rate float64 `json:"rate"`
	From string  `json:"from"`
	To   string  `json:"to"`
}

// URL generates a URL for querying the Yahoo finance API.
func URL(from, to string) string {
	return fmt.Sprintf(url, from, to)
}

// Returns a custom client with timeout set to 5 seconds. This ensures
// that the library doesn't make an app freeze.
func customClient() *http.Client {
	cc := &http.Client{
		Timeout: timeout,
	}
	return cc
}

// Sends a GET request to the Yahoo finance API and returns
// the response containing the relevant fields.
func queryAPI(from, to string) (*result, error) {
	apiResult := &result{}
	response, err := customClient().Get(URL(from, to))
	if err != nil {
		return nil, err
	}
	buf, _ := ioutil.ReadAll(response.Body)

	if err := json.Unmarshal(buf, &apiResult); err != nil {
		return nil, err
	}

	return apiResult, nil
}

// rate is sort of a convenience function to quickly fetch the most
// recent exchance rate from the Yahoo finance API.
// It returns only the exchance rate and any error that might occur.
func rate(from, to string) (float64, error) {
	res, err := queryAPI(from, to)
	if err != nil {
		return -1, err
	}
	return res.Rate, nil
}
