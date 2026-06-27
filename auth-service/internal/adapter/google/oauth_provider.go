package google

import (
	"context"
	"encoding/json"
	"task-tracker/auth-service/internal/domain"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OAuthProvider struct {
	clientID     string
	clientSecret string
}

func NewOAuthProvider(clientID, clientSecret string) *OAuthProvider {
	return &OAuthProvider{
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

func (o OAuthProvider) GetUserInfo(ctx context.Context, provider, code, redirectURI string) (domain.OAuthUserInfo, error) {
	oAuthConf := oauth2.Config{
		ClientID:     o.clientID,
		ClientSecret: o.clientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  redirectURI,
		Scopes:       []string{"openid", "email", "profile"},
	}

	token, err := oAuthConf.Exchange(ctx, code)
	if err != nil {
		return domain.OAuthUserInfo{}, err
	}

	client := oAuthConf.Client(ctx, token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return domain.OAuthUserInfo{}, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var req struct {
		ID    string  `json:"id"`
		Name  string  `json:"name"`
		Email *string `json:"email"`
	}

	if decodeErr := json.NewDecoder(resp.Body).Decode(&req); decodeErr != nil {
		return domain.OAuthUserInfo{}, decodeErr
	}

	info := domain.OAuthUserInfo{
		ID:    req.ID,
		Name:  req.Name,
		Email: req.Email,
	}

	return info, nil
}
