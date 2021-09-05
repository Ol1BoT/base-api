package auth

import (
	"github.com/Ol1BoT/base-api/helper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GoogleAuthConfig() (*oauth2.Config, error) {

	creds, err := helper.NewConfig("./config.yml")
	if err != nil {
		return nil, err
	}

	conf := &oauth2.Config{
		ClientID:     creds.Google.ClientID,
		ClientSecret: creds.Google.Secret,
		RedirectURL:  creds.Google.RedirectURI,
		Scopes: []string{
			"openid",
			"email",
			"profile",
		},
		Endpoint: google.Endpoint,
	}

	return conf, nil
}