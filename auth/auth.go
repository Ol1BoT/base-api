package auth

import (
	"github.com/ol1bot/base-api/helper"
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

func ExistsInDatabase(email string) error {

	//TODO: Check D to see if Email exists

	// svc, err :=controllers.CreateServices()
	// if err != nil {
	// 	return err
	// }

	// if err := svc.User

	// if us.Email == "" {
	// // 	return fmt.Errorf("User does not exist in the Database.")
	// // }

	return nil

}


type GoogleProfileBody struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
}

type TokenPayload struct {
	Email     string
	YearLevel string
	Staff     bool
}
