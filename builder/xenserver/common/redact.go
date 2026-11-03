// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"net/url"
)

const redactedValue = "***"

// session_id is a credential: XAPI's HTTP handlers accept it as a bearer token for the whole pool.
var sensitiveQueryParams = []string{"session_id", "password", "pwd"}

// redactURL also covers the query string: net/url's own redaction handles only the userinfo password.
func redactURL(urlString string) string {
	u, err := url.Parse(urlString)
	if err != nil {
		// Echoing an unparsable string back could print the credential it carries.
		return "<unparsable URL>"
	}

	if u.User != nil {
		u.User = url.UserPassword(redactedValue, redactedValue)
	}

	query := u.Query()
	for _, key := range sensitiveQueryParams {
		if query.Has(key) {
			query.Set(key, redactedValue)
		}
	}
	u.RawQuery = query.Encode()

	return u.String()
}

// net/http puts the request URL, session_id included, in the *url.Error it returns for a failed request.
func redactURLError(err error) error {
	urlErr, ok := err.(*url.Error)
	if !ok {
		return err
	}

	return &url.Error{Op: urlErr.Op, URL: redactURL(urlErr.URL), Err: urlErr.Err}
}

func redactArgs(args []string) []string {
	redacted := make([]string, len(args))
	copy(redacted, args)

	for i, arg := range redacted {
		if arg == "-pw" && i+1 < len(redacted) {
			redacted[i+1] = redactedValue
		}
	}

	return redacted
}
