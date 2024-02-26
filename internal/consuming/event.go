// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package consuming

type ConsumedEvent struct {
	Id        string `json:"id"`
	Type      string `json:"type"`
	Timestamp string `json:"time"`
}
