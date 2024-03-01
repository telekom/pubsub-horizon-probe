// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"context"
	"errors"
	"net/http"
	"time"
)

type MockServer struct {
	server *http.Server
}

func NewMockServer() MockServer {
	var mux = http.NewServeMux()
	mux.Handle("/sse", new(sseHandler))

	var server = &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return MockServer{server}
}

func (m *MockServer) Start() {
	go func() {
		if err := m.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
	m.awaitReadiness()
}

func (m *MockServer) awaitReadiness() {
	for i := 0; i < 10; i++ {
		res, _ := http.Head("http://localhost:8080/sse")
		if res != nil && res.StatusCode == http.StatusOK {
			return
		}
		time.Sleep(1 * time.Second)
	}
	panic(errors.New("mock server didn't come up in time"))
}

func (m *MockServer) Stop() {
	_ = m.server.Shutdown(context.Background())
}
