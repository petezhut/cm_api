package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/petezhut/cm_api/go/api/auth"
	"github.com/petezhut/cm_api/go/logging"
)

// Get - Wrapper for doing a GET call.
func Get(auth *auth.Auth, url *url.URL) []byte {
	return do("GET", auth, url)
}

// Post - Wrapper for doing a POST call.
func Post(auth *auth.Auth, url *url.URL) []byte {
	return do("POST", auth, url)
}

func do(method string, auth *auth.Auth, url *url.URL) []byte {
	/* Do - This is the unified requesting method.  It is unexported and
	only reachable via the Get and Post methods.  This is to make sure that
	people are reaching their URL correctly.
	Requires a method (POST/GET)
	Requires an Auth Object
	Requires a URL
	Returns a slice of bytes.
	*/

	client := &http.Client{}
	logging.DEBUG(fmt.Sprintf("querying: %s", url))
	req, err := http.NewRequest(method, fmt.Sprintf("%s", url), nil)
	req.SetBasicAuth(auth.Username, auth.Password)
	resp, err := client.Do(req)
	if err != nil {
		logging.FATAL(fmt.Sprintf("%s", err))
	}
	bodytext, err := ioutil.ReadAll(resp.Body)
	return bodytext
}
