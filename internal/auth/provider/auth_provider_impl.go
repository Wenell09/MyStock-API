package provider

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

type AuthProviderImpl struct{}

func NewAuthProvider() AuthProvider {
	return &AuthProviderImpl{}
}

// Login implements [AuthProvider].
func (a *AuthProviderImpl) Login(email string, password string) (string, error) {
	url := os.Getenv("SUPABASE_URL")
	if url == "" {
		return "", errors.New("SUPABASE_URL empty")
	}
	payload := map[string]string{
		"email":    email,
		"password": password,
	}
	body, _ := json.Marshal(payload)
	req, err := http.NewRequest(
		"POST",
		url+"/auth/v1/token?grant_type=password",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return "", err
	}
	req.Header.Set("apikey", os.Getenv("SUPABASE_ANON_KEY"))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", errors.New("invalid credentials")
	}
	var res struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}
	return res.AccessToken, nil
}
