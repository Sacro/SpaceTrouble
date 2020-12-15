package spacex

import (
	"context"
	"encoding/json"
	"net/http"
)

func GetLaunches(ctx context.Context, c *http.Client) (LaunchPads, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://api.spacexdata.com/v4/launches", nil)
	if err != nil {
		return nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var launches LaunchPads
	err = json.NewDecoder(res.Body).Decode(&launches)

	if err != nil {
		return nil, err
	}

	return launches, nil
}
