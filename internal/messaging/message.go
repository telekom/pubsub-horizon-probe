// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package messaging

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Message struct {
	Id        string `json:"id"`
	Timestamp string `json:"time"`
}

func NewMessage() Message {
	var id, _ = uuid.NewRandom()
	return Message{
		Id:        id.String(),
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
}

func (m *Message) Bytes() []byte {
	bytes, _ := json.Marshal(m)
	return append(bytes, byte('\n'))
}
