package config

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

const (
	ShikimoriOAuthAppNameKey      = "shikimori.oauth.app_name"
	ShikimoriOAuthClientIDKey     = "shikimori.oauth.client_id"
	ShikimoriOAuthClientSecretKey = "shikimori.oauth.client_secret"
)

type OAuthCredentials struct {
	AppName      string
	ClientID     string
	ClientSecret string
}

func InitConfig() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting user home directory:", err)
		return
	}

	configDir := path.Join(homeDir, ".config", "naruw")
	// if config folder doesn't exist, create it
	if err := os.MkdirAll(configDir, 0755); err != nil {
		fmt.Fprintln(os.Stderr, "Error creating config directory:", err)
	}

	configName := "config.yaml"
	// if config file doesn't exist, create it
	if _, err := os.Stat(path.Join(configDir, configName)); os.IsNotExist(err) {
		if _, err := os.Create(path.Join(configDir, configName)); err != nil {
			fmt.Fprintln(os.Stderr, "Error creating config file:", err)
		}
	}

	viper.AddConfigPath(configDir)
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.ReadInConfig()
}

func GetShikimoriOAuthCredentials() OAuthCredentials {
	creds := OAuthCredentials{
		AppName:      viper.GetString(ShikimoriOAuthAppNameKey),
		ClientID:     viper.GetString(ShikimoriOAuthClientIDKey),
		ClientSecret: viper.GetString(ShikimoriOAuthClientSecretKey),
	}

	if creds.AppName == "" || creds.ClientID == "" || creds.ClientSecret == "" {
		fmt.Println("Shikimori OAuth credentials are not set")
		AskShikimoriOAuthCredentials()
		return GetShikimoriOAuthCredentials()
	}

	return creds
}

func AskShikimoriOAuthCredentials() {
	// ask app name
	var appName string
	fmt.Print("Enter Shikimori OAuth app name: ")
	if _, err := fmt.Scan(&appName); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		return
	}
	viper.Set(ShikimoriOAuthAppNameKey, appName)

	// ask client id
	var clientID string
	fmt.Print("Enter Shikimori OAuth app client ID: ")
	if _, err := fmt.Scan(&clientID); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		return
	}
	viper.Set(ShikimoriOAuthClientIDKey, clientID)

	// ask client secret
	var clientSecret string
	fmt.Print("Enter Shikimori OAuth app client secret: ")
	if _, err := fmt.Scan(&clientSecret); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		return
	}
	viper.Set(ShikimoriOAuthClientSecretKey, clientSecret)

	// save config
	if err := viper.WriteConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Error writing config file:", err)
	}
}
