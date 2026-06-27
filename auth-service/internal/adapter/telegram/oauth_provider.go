package telegram

import (
	"context"
	"fmt"
	"task-tracker/auth-service/internal/domain"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type OAuthProvider struct {
	clientID string
}

func NewOAuthProvider(clientID string) *OAuthProvider {
	return &OAuthProvider{
		clientID: clientID,
	}
}

func (o OAuthProvider) GetUserInfo(ctx context.Context, provider, code, redirectURI string) (domain.OAuthUserInfo, error) {
	token := code

	keySet, err := jwk.Fetch(ctx, "https://oauth.telegram.org/.well-known/jwks.json")
	if err != nil {
		return domain.OAuthUserInfo{}, err
	}

	parsed, err := jwt.Parse([]byte(token),
		jwt.WithKeySet(keySet),
		jwt.WithValidate(true),
		jwt.WithIssuer("https://oauth.telegram.org"),
		jwt.WithAudience(o.clientID),
	)
	if err != nil {
		return domain.OAuthUserInfo{}, err
	}

	id, _ := parsed.Get("id")
	name, _ := parsed.Get("name")
	username, _ := parsed.Get("preferred_username")
	info := domain.OAuthUserInfo{
		ID:         fmt.Sprintf("%v", id),
		Name:       name.(string),
		TGUsername: new(username.(string)),
	}

	return info, nil
}
