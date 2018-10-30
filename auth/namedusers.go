// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "auth": namedusers Resource Client
//
// Command:
// $ goagen
// --design=github.com/fabric8-services/fabric8-auth/design
// --out=$(GOPATH)/src/github.com/fabric8-services/fabric8-auth-client
// --pkg=auth
// --version=v1.3.0

package auth

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// DeprovisionNamedusersPath computes a request path to the deprovision action of namedusers.
func DeprovisionNamedusersPath(username string) string {
	param0 := username

	return fmt.Sprintf("/api/namedusers/%s/deprovision", param0)
}

// deprovision the user
func (c *Client) DeprovisionNamedusers(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeprovisionNamedusersRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeprovisionNamedusersRequest create the request corresponding to the deprovision action endpoint of the namedusers resource.
func (c *Client) NewDeprovisionNamedusersRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PATCH", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}
