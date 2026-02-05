package middleware

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

var (
	jwkSet   jwk.Set
	lastLoad time.Time
	mutex    sync.Mutex
)

func GetJWKSet() (jwk.Set, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if time.Since(lastLoad) < time.Hour && jwkSet != nil {
		return jwkSet, nil
	}
	set, err := jwk.Fetch(context.TODO(), os.Getenv("SUPABASE_JWKS_URL"))
	if err != nil {
		return nil, err
	}
	jwkSet = set
	lastLoad = time.Now()
	return jwkSet, nil
}
