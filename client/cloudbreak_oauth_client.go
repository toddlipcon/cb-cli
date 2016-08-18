package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	httptransport "github.com/go-swagger/go-swagger/httpkit/client"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// Default cloudbreak HTTP client.
var DefaultOAuth2 = NewOAuth2HTTPClient(nil)

// NewHTTPClient creates a new cloudbreak HTTP client.
func NewOAuth2HTTPClient(formats strfmt.Registry) *Cloudbreak {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New("localhost:9091", "/cb/api/v1", []string{"http"})
	return New(transport, formats)
}
