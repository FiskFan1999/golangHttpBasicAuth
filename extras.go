/*

MIT License

Copyright (c) 2022 William Rehwinkel

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

GOLANG NET/HTTP BASIC AUTH EXAMPLE
by William Rehwinkel, 2022.

Simple working example for how to require
simple authentication from the user to
access a protected page. Includes an example
of crypto hashing.

Note: compiling this program and running
will spin up a temporary HTTP server on
port 8080, which will require the login
username:password (see const)

Note: almost all browsers have disabled the
username:password@example.com functionality (although
it still works with some tools like curl). Instead,
the browser will initially make a request to
the server and recieve a 401 error, which will
cause the browser to show a prompt for the user to
enter the username and password.

Please feel free to leave an issue or pull request
on the GitHub page if I made a mistake such as
incorrectly interpreting an RFC or sending the wrong
response code.

*/

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
