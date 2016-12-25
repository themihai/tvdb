package tvdb

import (
	"github.com/dghubble/sling"
	"net/http"
)

// Token struct
type Token struct {
	Token string `json:"token"`
}

// TokenService tv series service
type TokenService struct {
	sling *sling.Sling
}

// newTokenService initialize a new TokenService
func newTokenService(sling *sling.Sling) *TokenService {
	return &TokenService{
		sling: sling,
	}
}

// Login requests and applies a new Token to the base client
func (s *TokenService) Login(auth *Auth) (*Token, *http.Response, error) {
	token := new(Token)
	jsonError := new(JSONError)
	res, err := s.sling.New().Post("/login").BodyJSON(auth).Receive(token, jsonError)
	if err == nil {
		err = jsonError
	}

	s.sling.Set("Authorization", "Bearer "+token.Token)
	return token, res, err
}
