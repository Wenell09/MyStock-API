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
	mutex    sync.RWMutex
)

func GetJWKSet() (jwk.Set, error) {
	jwksURL := os.Getenv("SUPABASE_JWKS_URL")
	if jwksURL == "" {
		log.Println("[JWKS] ENV SUPABASE_JWKS_URL not set")
		return nil, errors.New("SUPABASE_JWKS_URL not set")
	}
	mutex.RLock()
	if jwkSet != nil && time.Since(lastLoad) < time.Hour {
		defer mutex.RUnlock()
		return jwkSet, nil
	}
	mutex.RUnlock()
	mutex.Lock()
	defer mutex.Unlock()
	if jwkSet != nil && time.Since(lastLoad) < time.Hour {
		return jwkSet, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	set, err := jwk.Fetch(ctx, jwksURL)
	if err != nil {
		if jwkSet != nil {
			log.Println("[JWKS] fetch failed, using cache:", err)
			return jwkSet, nil
		}
		log.Println("[JWKS] fetch failed and no cache:", err)
		return nil, err
	}
	jwkSet = set
	lastLoad = time.Now()
	return jwkSet, nil
}
