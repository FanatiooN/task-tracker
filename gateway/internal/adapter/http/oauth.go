package http

import (
	"fmt"
	"net/http"
	"strings"
	authpb "task-tracker/gen/proto/auth"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OAuthHandler struct {
	client            authpb.AuthServiceClient
	googleClientID    string
	googleRedirectURI string
	frontendURL       string
}

func NewOAuthHandler(client authpb.AuthServiceClient, googleClientID, googleRedirectURI, frontendURL string) *OAuthHandler {
	return &OAuthHandler{client: client, googleClientID: googleClientID, googleRedirectURI: googleRedirectURI, frontendURL: frontendURL}
}

func (h *OAuthHandler) LoginWithGoogle(w http.ResponseWriter, r *http.Request) {
	oauthConfig := oauth2.Config{
		ClientID:    h.googleClientID,
		RedirectURL: h.googleRedirectURI,
		Endpoint:    google.Endpoint,
		Scopes:      []string{"openid", "email", "profile"},
	}

	url := oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *OAuthHandler) LoginCallbackWithGoogle(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimSpace(r.URL.Query().Get("code"))

	response, err := h.client.LoginByOAuth(r.Context(), &authpb.LoginByOAuthRequest{
		Provider:    "google",
		RedirectUri: h.googleRedirectURI,
		Code:        code,
	})
	if err != nil {
		writeGRPCError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf(
		"%s?access_token=%s&refresh_token=%s",
		h.frontendURL,
		response.AccessToken,
		response.RefreshToken,
	), http.StatusTemporaryRedirect)
}
