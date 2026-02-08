package middleware

import (
	"context"
	"errors"
	"log"
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
	jwksURL := os.Getenv("SUPABASE_JWKS_URL")
	if jwksURL == "" {
		return nil, errors.New("SUPABASE_JWKS_URL is not set")
	}
	mutex.Lock()
	defer mutex.Unlock()
	if jwkSet != nil && time.Since(lastLoad) < time.Hour {
		return jwkSet, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	set, err := jwk.Fetch(ctx, jwksURL)
	if err != nil {
		if jwkSet != nil {
			log.Println("[JWKS] fetch failed, using cached keys:", err)
			return jwkSet, nil
		}
		return nil, err
	}
	jwkSet = set
	lastLoad = time.Now()
	return jwkSet, nil
}
