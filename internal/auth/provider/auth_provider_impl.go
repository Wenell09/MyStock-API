package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type AuthProviderImpl struct{}

func NewAuthProvider() AuthProvider {
	return &AuthProviderImpl{}
}

// Login implements [AuthProvider].
func (a *AuthProviderImpl) Login(email, password string) (string, error) {
	supabaseURL := os.Getenv("SUPABASE_URL")
	anonKey := os.Getenv("SUPABASE_ANON_KEY")
	if supabaseURL == "" || anonKey == "" {
		return "", errors.New("SUPABASE_URL or SUPABASE_ANON_KEY not set")
	}
	payload := map[string]string{
		"email":    email,
		"password": password,
	}
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", supabaseURL+"/auth/v1/token?grant_type=password", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", anonKey)
	client := &http.Client{Timeout: 10 * time.Second}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var res struct {
		AccessToken string `json:"access_token"`
		Error       string `json:"error"`
		ErrorDesc   string `json:"error_description"`
	}
	if err := json.Unmarshal(bodyData, &res); err != nil {
		log.Println("[LOGIN] Failed to decode Supabase response:", string(bodyData))
		return "", errors.New("invalid response from Supabase")
	}
	if res.AccessToken == "" {
		msg := "login failed"
		if res.ErrorDesc != "" {
			msg = res.ErrorDesc
		} else if res.Error != "" {
			msg = res.Error
		}
		return "", errors.New(msg)
	}
	return res.AccessToken, nil
}
