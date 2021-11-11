package types

import (
	"io"
	"net/url"
	"strings"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/internal/util"
)

type SpacesLookupIDParams struct {
	accessToken string

	// Path parameter
	ID string

	// Query parameters
	Expansions  fields.ExpansionList
	SpaceFields fields.SpaceFieldList
	UserFields  fields.UserFieldList
}

var SpacesLookupIDQueryParams = map[string]struct{}{
	"expansions":   {},
	"space.fields": {},
	"user.fields":  {},
}

func (p *SpacesLookupIDParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *SpacesLookupIDParams) AccessToken() string {
	return p.accessToken
}

func (p *SpacesLookupIDParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	qs := util.QueryString(pm, SpacesLookupIDQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *SpacesLookupIDParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *SpacesLookupIDParams) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Expansions, p.SpaceFields, p.UserFields)
	return m
}