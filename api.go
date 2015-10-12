package circleci

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func NewClient(config *Config) *Client {
	return &Client{config.Client()}
}

type Client struct {
	hc *http.Client
}

func get(c *http.Client, url string, res interface{}) error {
	r, err := c.Get(url)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	r.Body.Close()

	return json.Unmarshal(body, res)
}

// MeResponce
type MeResponce struct {
	SelectedEmail       string    `json:"selected_email"`
	AvatarURL           string    `json:"avatar_url"`
	TrialEnd            string    `json:"trial_end"`
	Admin               bool      `json:"admin"`
	BasicEmailPrefs     string    `json:"basic_email_prefs"`
	SignInCount         int       `json:"sign_in_count"`
	GithubOauthScopes   []string  `json:"github_oauth_scopes"`
	Name                string    `json:"name"`
	GravatarID          string    `json:"gravatar_id"`
	DaysLeftInTrial     int       `json:"days_left_in_trial"`
	Parallelism         int       `json:"parallelism"`
	GitHubID            int       `json:"github_id"`
	LastViewedChangeLog time.Time `json:"last_viewed_changelog"`
	DevAdmin            bool      `json:"dev_admin"`
	AllEmails           []string  `json:"all_emails"`
	CreatedAt           time.Time `json:"created_at"`
	Plan                string    `json:"plan"`
	HerokuApiKey        string    `json:"heroku_api_key"`
	Projects            map[string]struct {
		OnDashboard bool   `json:"on_dashboard"`
		Emails      string `json:"emails"`
	} `json:"projects"`
	Login      string `json:"login"`
	Containers int    `json:"containers"`
}

func (c *Client) Me() (*MeResponce, error) {
	r := &MeResponce{}
	if err := get(c.hc, "https://circleci.com/api/v1/me", r); err != nil {
		return nil, err
	}
	return r, nil
}
