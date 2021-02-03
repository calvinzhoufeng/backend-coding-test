package pkg

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

var (
	MarketClient *client
)

type Funny interface {
	UserGenerator() (string, string, error)
	ContentGenerator(firstName string, lastName string) (string, error)
}

type client struct {
	userUrl    string
	icndbUrl   string
	httpClient *http.Client
}

func NewClient(userUrl string, icndbUrl string) *client {
	if MarketClient == nil {
		httpClient := &http.Client{Timeout: 10 * time.Second}
		MarketClient = &client{
			userUrl:    userUrl,
			icndbUrl:   icndbUrl,
			httpClient: httpClient,
		}
	}

	return MarketClient
}

type Result struct {
	Results []PersonName `json:"results"`
}

type PersonName struct {
	Name struct {
		Last  string `json:"last"`
		First string `json:"first"`
	} `json:"name"`
}

func (c *client) UserGenerator() (string, string, error) {
	result := &Result{}

	err := c.getJson(c.userUrl, result)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get content")
		return "", "", err
	}

	log.Debug().Msgf("Person %v\n", result)
	// TODO: assume the result is always returned with 1 set of person names
	person := result.Results[0]

	return person.Name.First, person.Name.Last, nil
}

type UserResponse struct {
	Value struct {
		Content string `json:"joke"`
	} `json:"value"`
}

func (c *client) ContentGenerator(firstName string, lastName string) (string, error) {
	user := &UserResponse{}

	err := c.getJson(c.icndbUrl+"firstName="+firstName+"&lastName="+lastName, user)
	if err != nil {
		log.Error().Msg("Failed to get content")
		return "", err
	}

	return user.Value.Content, err
}

func (c *client) getJson(url string, target interface{}) error {

	r, err := c.httpClient.Get(url)
	if err != nil {
		return err
	}
	// dump, _ := httputil.DumpResponse(r, true)
	// log.Debug().Msgf("Response body %s %q\n", url, dump)
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
