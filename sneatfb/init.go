package sneatfb

import "github.com/sneat-co/sneat-go-core/apicore"

func init() { // TODO: Consider providing value for api-core initialization
	apicore.GetAuthTokenFromHttpRequest = getAuthTokenFromHttpRequest
	apicore.NewAuthContext = newAuthContext
}
