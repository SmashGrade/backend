package auth

import (
	"context"
	"log"

	c "github.com/SmashGrade/backend/app/config"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/lestrrat-go/jwx/jwk"
)

// AuthProvider is used to authenticate users
type AuthProvider struct {
	config *c.APIConfig // Contains the API configuration
	keySet jwk.Set      // Contains the valid keys for authentication
}

// Creates a new AuthProvider
// Downloads and caches the keys from the OAuthKeyDiscoveryURL
func NewAuthProvider(config *c.APIConfig) *AuthProvider {
	keySet, err := jwk.Fetch(context.Background(), config.AuthConfig.OAuthKeyDiscoveryURL)
	if err != nil {
		log.Fatalf("Failed to fetch the oauth keys from URL %s: %s", config.AuthConfig.OAuthKeyDiscoveryURL, err)
	}
	// Return the new AuthProvider
	return &AuthProvider{
		config: config,
		keySet: keySet,
	}
}

// Validates the token and returns the public key
// This is used within the middleware to validate the token is from the correct issuer
func (a *AuthProvider) ValidateToken(token *jwt.Token) (interface{}, error) {
	// Get the key ID from the token
	kid, ok := token.Header["kid"].(string)
	if !ok {
		return nil, jwt.ErrInvalidKey
	}
	// Check if the key is part of the cached keyset
	key, ok := a.keySet.LookupKeyID(kid)
	if !ok {
		return nil, jwt.ErrInvalidKey
	}
	// Create a new interface to store the public key
	var publicKey interface{}
	// Mashal the raw key into the public key interface
	if err := key.Raw(&publicKey); err != nil {
		return nil, err
	}
	// Return the public key interface
	return publicKey, nil
}

// Returns the JWT configuration for the middleware to use
func (a *AuthProvider) GetJWTConfig() echojwt.Config {
	return echojwt.Config{
		KeyFunc: a.ValidateToken,
	}
}
