// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"github.com/telekom/pubsub-horizon-probe/internal/messaging"
	"net/http"
)

type sseHandler struct{}

func (h *sseHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case http.MethodHead:
		writer.WriteHeader(http.StatusOK)
		writer.(http.Flusher).Flush()

	case http.MethodGet:
		writer.Header().Add("Content-Type", "application/stream+json")
		for {
			select {

			case <-req.Context().Done():
				return

			default:
				var message = messaging.NewMessage()
				_, err := writer.Write(message.Bytes())
				if err != nil {
					panic(err)
				}
				writer.(http.Flusher).Flush()

			}
		}

	default:
		writer.WriteHeader(http.StatusNotFound)
		writer.(http.Flusher).Flush()

	}
}
