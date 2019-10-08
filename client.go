package goblin

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/mchrobak/goblin/models"
)

const (
	baseURL        string = `https://api.pubg.com`
	shardsEndpoint string = "/shards"
)

var (
	errNoShardSpecified = errors.New("no shard specified")
)

// APIClient holds a REST API Client
type APIClient struct {
	apiURL *url.URL
	apiKey string

	httpClient *http.Client
}

// NewAPIClient returns a new APIClient with APIKey attached
func NewAPIClient(apiKey string) (*APIClient, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	// defaultClient := http.DefaultClient
	return &APIClient{
		apiURL:     parsedURL,
		apiKey:     apiKey,
		httpClient: http.DefaultClient,
	}, nil
}

// Get is the basic HTTP GET wrapper
func (c *APIClient) Get(route *url.URL, auth bool, responseObject interface{}) error {
	req, err := http.NewRequest("GET", route.String(), nil)
	if err != nil {
		return err
	}

	if auth {
		req.Header.Set("Accept", "application/vnd.api+json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		switch res.StatusCode {
		case http.StatusUnauthorized:
			return errors.New("API key invalid or missing")
		case http.StatusNotFound:
			return errors.New("The specified resource was not found")
		case http.StatusUnsupportedMediaType:
			return errors.New("Content type incorrect or not specified")
		case http.StatusTooManyRequests:
			return errors.New("Too many requests")
		default:
			return errors.New(res.Status)
		}
	}

	if responseObject != nil {
		if err := json.NewDecoder(res.Body).Decode(responseObject); err != nil {
			return fmt.Errorf("Failed to parse API response from '%s': %s",
				req.RequestURI, err.Error())
		}
	}

	return nil
}

// GetStatus returns the status of the API
func (c *APIClient) GetStatus() (*models.Status, error) {
	path, err := url.Parse(models.StatusEndpoint)
	if err != nil {
		return nil, err
	}

	u := c.apiURL.ResolveReference(path)

	status := new(models.Status)
	if err := c.Get(u, false, status); err != nil {
		return nil, err
	}

	return status, nil
}

// GetPlayer returns a single Player from the PlayerID
func (c *APIClient) GetPlayer(playerID, shard string) (*models.Player, error) {
	if shard == "" {
		return nil, errNoShardSpecified
	}

	path := fmt.Sprintf("%s/%s", shardsEndpoint, shard)
	path += fmt.Sprintf("%s/%s", models.PlayersEndpoint, playerID)

	p, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	u := c.apiURL.ResolveReference(p)

	player := new(models.Player)
	if err := c.Get(u, true, player); err != nil {
		return nil, err
	}

	return player, nil
}

// GetPlayers finds a player based on Player name or PlayerID
func (c *APIClient) GetPlayers(playerName, playerID, shard string) (*models.PlayerList, error) {
	if shard == "" {
		return nil, errors.New("shard must be specified")
	}

	path := fmt.Sprintf("%s/%s", shardsEndpoint, shard)
	path += fmt.Sprintf("%s", models.PlayersEndpoint)

	if playerName != "" { // PlayerName will take precedence if both are defined
		path += fmt.Sprintf("?filter[playerName]=%s", playerName)
	} else if playerID != "" {
		path += fmt.Sprintf("?filter[playerIds]=%s", playerID)
	} else {
		return nil, errors.New("A filter is required for /players endpoint")
	}

	p, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	u := c.apiURL.ResolveReference(p)

	players := new(models.PlayerList)
	if err := c.Get(u, true, players); err != nil {
		return nil, err
	}

	return players, nil
}

// GetMatch returns single match from the MatchID
func (c *APIClient) GetMatch(matchID, shard string) (*models.Match, error) {
	if shard == "" {
		return nil, errors.New("shard must be specified")
	}

	path := fmt.Sprintf("%s/%s", shardsEndpoint, shard)
	path += fmt.Sprintf("%s/%s", models.MatchesEndpoint, matchID)

	p, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	u := c.apiURL.ResolveReference(p)

	match := new(models.Match)
	if err := c.Get(u, true, match); err != nil {
		return nil, err
	}

	return match, nil
}

func (c *APIClient) GetSeasons() (*models.SeasonsList, error) {
	path, err := url.Parse(models.SeasonsEndpoint)
	if err != nil {
		return nil, err
	}

	u := c.apiURL.ResolveReference(path)

	seasons := new(models.SeasonsList)
	if err := c.Get(u, false, seasons); err != nil {
		return nil, err
	}

	return seasons, nil
}
