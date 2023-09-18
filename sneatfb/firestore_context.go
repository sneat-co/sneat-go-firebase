package sneatfb

//import (
//	"context"
//	"net/http"
//)
//
//// NewFirestoreContext creates a context with a Firebase ContactID token
//var NewFirestoreContext = func(r *http.Request, authRequired bool) (firestoreContext *FirestoreContext, err error) {
//	var ctx context.Context
//	if ctx, err = getSneatAuthTokenFromHttpRequest(r, authRequired); err != nil {
//		return nil, err
//	}
//	return NewContextWithFirestoreClient(ctx)
//}
