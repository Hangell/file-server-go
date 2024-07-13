package main

import auth "github.com/abbot/go-http-auth"

// Secret function to validate user credentials
func Secret(user, realm string) string {
	if user == "admin" || user == "dev" || user == "dev" {
		return "$1$/q9tTBog$KtdXn0eQTpiUP9g//xLaS1" //https://unix4lyfe.org/crypt/ MD5 Crypt: md5 salt
	}
	return ""
}

// NewAuthenticator creates a new BasicAuthenticator
func NewAuthenticator() *auth.BasicAuth {
	return auth.NewBasicAuthenticator("meuserver.com", Secret)
}
