package services

import "github.com/duo-labs/webauthn/webauthn"

type UserStore interface {
	GetUser(string) (webauthn.User, error)
	StoreUserCreds(webauthn.User, *webauthn.Credential) error
}
