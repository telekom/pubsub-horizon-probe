// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"sync"
	"time"
)

type ResultSet struct {
	results map[string]*Result
	mutex   sync.Mutex
}

func NewResultSet() ResultSet {
	return ResultSet{
		results: make(map[string]*Result),
		mutex:   sync.Mutex{},
	}
}

func (s *ResultSet) getOrCreate(key string) *Result {
	result, ok := s.results[key]
	if !ok {
		var result = new(Result)
		s.results[key] = result
		return result
	}
	return result
}

func (s *ResultSet) RecordAsSent(key string) {
	var timeSent = time.Now()
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.getOrCreate(key).Sent = timeSent
}

func (s *ResultSet) RecordAsReceived(key string) {
	var timeReceived = time.Now()
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.getOrCreate(key).Received = timeReceived
}

func (s *ResultSet) HasRecorded(id string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, ok := s.results[id]
	return ok
}

func (s *ResultSet) GetResult(id string) *Result {
	s.mutex.Lock()
	var result = s.results[id]
	s.mutex.Unlock()
	return result
}

func (s *ResultSet) IsComplete() bool {
	for _, result := range s.results {
		if !result.IsComplete() {
			return false
		}
	}
	return true
}

func (s *ResultSet) Surpasses(duration time.Duration) bool {
	for _, result := range s.results {
		if !result.IsComplete() {
			continue
		}

		if result.GetLatency() > duration {
			return true
		}
	}
	return false
}
