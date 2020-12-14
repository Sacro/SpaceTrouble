package spacex

import (
	"encoding/json"
	"net/http"
)

func GetLaunches(c *http.Client) (LaunchPads, error) {
	res, err := c.Get("https://api.spacexdata.com/v4/launches")

	if err != nil {
		return nil, err
	}

	var launches LaunchPads
	err = json.NewDecoder(res.Body).Decode(&launches)

	if err != nil {
		return nil, err
	}

	return launches, nil
}
