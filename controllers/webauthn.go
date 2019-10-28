package controllers

import (
	"net/http"

	"github.com/duo-labs/webauthn/webauthn"
	"github.com/equinor/no-factor-auth/services"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type WebAuthN struct {
	UserStore services.UserStore
	Authn     *webauthn.WebAuthn
}

func (w *WebAuthN) BeginLogin(c echo.Context) error {
	c.Param("name")
	user, err := w.UserStore.GetUser(c.Param("name")) // Find the user
	if err != nil {

		return err
	}
	options, sessionData, err := w.Authn.BeginLogin(user)
	if err != nil {

		return err
	}
	// handle errors if present
	// store the sessionData values
	s, err := session.Get("login-session", c)
	if err != nil {

		return err
	}
	s.Values["creds"] = sessionData
	c.JSON(http.StatusOK, options)
	// JSONResponse(w, options, http.StatusOK) // return the options generated
	// options.publicKey contain our registration options
	return nil
}

func (w *WebAuthN) FinishLogin(c echo.Context) error {
	user, err := w.UserStore.GetUser(c.Param("name")) // Get the user
	if err != nil {

		return err
	}
	// Get the session data stored from the function above
	// using gorilla/sessions it could look like this
	s, err := session.Get("login-session", c)
	if err != nil {

		return err
	}
	sd := s.Values["creds"].(webauthn.SessionData)
	credential, err := w.Authn.FinishLogin(user, sd, c.Request())
	if err != nil {

		return err
	}
	err = w.UserStore.StoreUserCreds(user, credential)
	if err != nil {

		return err
	}
	// Handle validation or input errors
	// If login was successful, handle next steps
	c.JSON(http.StatusOK, "Login Success")
	return nil
}

func (w *WebAuthN) BeginRegistration(c echo.Context) error {
	user, err := w.UserStore.GetUser(c.Param("name")) // Find or create the new user
	if err != nil {

		return err
	}
	options, sessionData, err := w.Authn.BeginRegistration(user)
	if err != nil {

		return err
	}
	// handle errors if present
	// store the sessionData values
	s, err := session.Get("registration-session", c)
	if err != nil {

		return err
	}
	s.Values["creds"] = sessionData
	c.JSON(http.StatusOK, options) // return the options generated
	// options.publicKey contain our registration options
	return nil
}

func (w *WebAuthN) FinishRegistration(c echo.Context) error {
	user, err := w.UserStore.GetUser(c.Param("name")) // Get the user
	if err != nil {

		return err
	}
	// Get the session data stored from the function above
	// using gorilla/sessions it could look like this
	// sessionData := store.Get(r, "registration-session")
	s, err := session.Get("registration-session", c)
	if err != nil {

		return err
	}
	sd := s.Values["creds"].(webauthn.SessionData)
	credential, err := w.Authn.FinishRegistration(user, sd, c.Request())

	if err != nil {

		return err
	}

	err = w.UserStore.StoreUserCreds(user, credential)
	if err != nil {

		return err
	}
	// Handle validation or input errors
	// If creation was successful, store the credential object
	// Handle next steps
	c.JSON(http.StatusOK, "Registration Success")
	return nil
}
