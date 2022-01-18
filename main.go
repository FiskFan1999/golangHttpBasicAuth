package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

const (
	username = "abc"
	password = "123"
)

var hash []byte

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

	if !ok || u != username || bcrypt.CompareHashAndPassword(hash, []byte(p)) != nil {

		// Username or password is not correct, or the request
		// did not provide any authentication. Reply with a
		// 401 error with header including "WWW-Authenticate=Basic"

		w.Header().Set(`WWW-Authenticate`, `Basic realm="admin"`)
		/*
			For information about "realm", refer to
			https://stackoverflow.com/questions/12701085/what-is-the-realm-in-basic-authentication
			and https://datatracker.ietf.org/doc/html/rfc7235#section-2.2
		*/
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}

	// authentication was successful.

	return true
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if !PasswordAuthSucceeds(w, r) {
		return
	}
	/*
		At this point, you can safely handle the
		request as a successful authentication
	*/
	u, p, _ := r.BasicAuth()
	fmt.Fprintln(w, "Authentication successful.")
	fmt.Fprintln(w, "username: ", u)
	fmt.Fprintln(w, "password: ", p)

}

func main() {
	var err error
	hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}
