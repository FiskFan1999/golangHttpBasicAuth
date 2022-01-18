package main

/*
NOTE: for testing purposes (temporary server)
username=abc, password=123 (see extras.go)
*/

import (
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func PasswordAuthSucceeds(w http.ResponseWriter, r *http.Request) bool {
	/*
		Returns false if authentication FAILS
		true is authentication SUCEEDS
	*/

	// See https://pkg.go.dev/net/http#Request.BasicAuth
	u, p, ok := r.BasicAuth()

	/*
		ok is boolean, false if no basic authentication is
		provided in the request; true if basic authentication
		is present (wether or not the username and password is
		correct)
	*/

	if !ok || // ok = false if no authentication provided
		u != username || // check username matches
		bcrypt.CompareHashAndPassword(hash, []byte(p)) != nil /* check password matches */ {

		/*
			Authentication error. Reply with WWW-Authenticate
			header and 401 error.
		*/
		w.Header().Set(`WWW-Authenticate`, `Basic realm="admin"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false

		/*
			For information about "realm", refer to
			https://stackoverflow.com/questions/12701085/what-is-the-realm-in-basic-authentication
			and https://datatracker.ietf.org/doc/html/rfc7235#section-2.2
		*/
	}

	// authentication was successful.

	return true
}
