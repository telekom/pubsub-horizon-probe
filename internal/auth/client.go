// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/telekom/pubsub-horizon-probe/internal/utils"
	"io"
	"net/http"
	"time"
)

var Client utils.HttpClient = &http.Client{Timeout: 10 * time.Second}

func RetrieveToken(url string, clientId string, clientSecret string) (string, error) {
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte("grant_type=client_credentials")))
	if err != nil {
		return "", err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.SetBasicAuth(clientId, clientSecret)

	response, err := Client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var body struct {
		AccessToken string `json:"access_token"`
	}

	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		return "", err
	}

	return body.AccessToken, nil
}
