package shikimori

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cli/browser"
	"github.com/zalando/go-keyring"
	"golang.org/x/oauth2"
)

const (
	// keyring related constants
	keyringServiceName = "shikimori-authorization-code"
	keyringUsername    = "naruw"
	// oauth2 related constants
	shikimoriOauthURL          = "https://shikimori.one/oauth"
	shikimoriOAuthAppName      = "narutoep"
	shikimoriOAuthClientID     = "uDvhUSh3iibbH6IhJqrGtxLtLixWfIkx8bJ36C4Hcvg"
	shikimoriOAuthClientSecret = "Q6XInqsVBDSWMYrWKr6ciHAJrbVZt3Fl5zsT9MEWFJA"
)

var (
	_conf   *oauth2.Config
	_client *http.Client
)

func initConfig() {
	_conf = &oauth2.Config{
		ClientID:     shikimoriOAuthClientID,
		ClientSecret: shikimoriOAuthClientSecret,
		Scopes:       []string{"user_rates"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  fmt.Sprintf("%s/authorize", shikimoriOauthURL),
			TokenURL: fmt.Sprintf("%s/token", shikimoriOauthURL),
		},
		RedirectURL: "urn:ietf:wg:oauth:2.0:oob",
	}
}

func getConfig() *oauth2.Config {
	if _conf == nil {
		initConfig()
	}

	return _conf
}

func initClient() error {
	conf := getConfig()
	ctx := context.Background()
	token, err := getToken(conf, ctx)
	if err != nil {
		return err
	}

	// The HTTP Client returned by conf.Client will refresh the token as necessary.
	_client = conf.Client(ctx, token)

	return nil
}

func getClient() (*http.Client, error) {
	if _client == nil {
		if err := initClient(); err != nil {
			return nil, err
		}
	}

	return _client, nil
}

func getToken(conf *oauth2.Config, ctx context.Context) (*oauth2.Token, error) {
	tokenStr, err := keyring.Get(keyringServiceName, keyringUsername)
	if err != nil {
		return nil, fmt.Errorf("you are not authenticated: use 'naruw auth'")
	}
	token := &oauth2.Token{}
	err = json.Unmarshal([]byte(tokenStr), &token)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal token: %w", err)
	}

	// TODO: check token validity

	return token, nil
}

func saveToken(token *oauth2.Token) error {
	tokenBytes, err := json.Marshal(token)
	if err != nil {
		return fmt.Errorf("unable to marshal token: %w", err)
	}

	// Save token to keyring
	if err := keyring.Set(keyringServiceName, keyringUsername, string(tokenBytes)); err != nil {
		return fmt.Errorf("unable to save token: %w", err)
	}

	return nil
}

func Authenticate() error {
	conf := getConfig()
	ctx := context.Background()

	// Redirect user to consent page to ask for permission for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	if err := browser.OpenURL(url); err != nil {
		fmt.Printf("Visit the URL and get Authorization code:\n%v\n", url)
	}

	// Get the authorization code
	var code string
	fmt.Printf("Enter Authorization code: ")
	if _, err := fmt.Scan(&code); err != nil {
		return fmt.Errorf("unable to read authorization code: %w", err)
	}

	// Exchange will do the handshake to retrieve the initial access token.
	token, err := conf.Exchange(ctx, code)
	if err != nil {
		return fmt.Errorf("unable to exchange authorization code: %w", err)
	}

	if err = saveToken(token); err != nil {
		return err
	}

	return nil
}

func doAPIRequest(method string, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("https://shikimori.one/api%s", path), body)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Add("User-Agent", shikimoriOAuthAppName)

	client, err := getClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to get user rate: %w", err)
	}

	return resp, nil
}

type ShikimoriUserRate struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	TargetID   int       `json:"target_id"`
	TargetType string    `json:"target_type"`
	Score      int       `json:"score"`
	Status     string    `json:"status"`
	Rewatches  int       `json:"rewatches"`
	Episodes   int       `json:"episodes"`
	Volumes    int       `json:"volumes"`
	Chapters   int       `json:"chapters"`
	Text       string    `json:"text"`
	TextHTML   string    `json:"text_html"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func GetUserRate(animeID int) (ShikimoriUserRate, error) {
	resp, err := doAPIRequest("GET", fmt.Sprintf("/v2/user_rates/%d", animeID), nil)
	if err != nil {
		return ShikimoriUserRate{}, err
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return ShikimoriUserRate{}, fmt.Errorf("unable to get user rate: %s", resp.Status)
	}

	var userRate ShikimoriUserRate
	if err := json.NewDecoder(resp.Body).Decode(&userRate); err != nil {
		return ShikimoriUserRate{}, fmt.Errorf("unable to decode JSON data: %w", err)
	}

	return userRate, nil
}

func IncrementEpisodes(animeID int) error {
	resp, err := doAPIRequest("POST", fmt.Sprintf("/v2/user_rates/%d/increment", animeID), nil)
	if err != nil {
		return fmt.Errorf("unable to increment episodes: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unable to increment episode: %s", resp.Status)
	}

	return nil
}
