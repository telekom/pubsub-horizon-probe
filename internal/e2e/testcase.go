// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/telekom/pubsub-horizon-probe/internal/config"
	"github.com/telekom/pubsub-horizon-probe/internal/consuming"
	"github.com/telekom/pubsub-horizon-probe/internal/publishing"
	"io"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

type TestCase struct {
	results      ResultSet
	messageCount int
	testFile     string

	ctx    context.Context
	cancel context.CancelFunc
}

func NewTestCase(messageCount int, timeout time.Duration, testFile string) TestCase {
	var ctx, cancel = context.WithTimeout(context.Background(), timeout)
	return TestCase{
		results:      NewResultSet(),
		messageCount: messageCount,
		testFile:     testFile,

		ctx:    ctx,
		cancel: cancel,
	}
}

func (t *TestCase) Start() bool {
	var testWg = new(sync.WaitGroup)
	testWg.Add(2)

	var errs = make(chan error, 1)

	defer t.cancel()
	go t.publish(testWg, errs)
	go t.consume(testWg)

	var complete = make(chan bool, 1)
	go func() {
		defer close(complete)
		testWg.Wait()
	}()

	select {

	case err := <-errs:
		log.Error().Err(err).Msg("An error occurred while publishing")
		return false

	case <-complete:
		return true // Completed without complications

	case <-t.ctx.Done():
		return false // Reached timeout (so context has already been cancelled)

	}
}

func (t *TestCase) consume(wg *sync.WaitGroup) {
	defer wg.Done()

	var consumer = consuming.NewConsumer(&config.Current.Consuming)
	go t.openConnection(consumer)

	for {
		select {

		case <-t.ctx.Done():
			return

		default:
			msg := <-consumer.Events

			if t.results.HasRecorded(msg.Id) {
				t.results.RecordAsReceived(msg.Id)
				log.Info().Fields(map[string]any{
					"eventId": msg.Id,
					"latency": fmt.Sprintf("%dms", t.results.GetResult(msg.Id).GetLatencyMs()),
				}).Msgf("Received message")

				if t.results.IsComplete() {
					return
				}
			} else {
				log.Warn().Msgf("Received unexpected message with id %s. Skipping...", msg.Id)
			}

		}
	}
}

func (t *TestCase) openConnection(consumer *consuming.Consumer) {
	for {
		select {

		case <-t.ctx.Done():
			return

		default:
			if err := consumer.Start(); err != nil {
				if os.IsTimeout(err) {
					log.Debug().Msg("Connection timed out. Reconnecting...")
					continue
				}

				if errors.Is(err, io.EOF) {
					log.Debug().Msg("Received end of stream (EOF). Reconnecting...")
					continue
				}
				log.Fatal().Err(err).Msg("Error while consuming events")
			}

		}
	}
}

func (t *TestCase) publish(wg *sync.WaitGroup, errs chan<- error) {
	defer wg.Done()

	var publishingWorkers = new(sync.WaitGroup)
	publishingWorkers.Add(t.messageCount)

	var publishingCount = new(atomic.Int32)

	for i := 0; i < t.messageCount; i++ {
		go func() {
			defer publishingWorkers.Done()
			eventId, err := publishing.Publish(&config.Current.Publishing, t.testFile)
			if err != nil {
				errs <- err
				return
			}

			t.results.RecordAsSent(eventId)
			log.Info().Fields(map[string]any{
				"eventId": eventId,
			}).Msgf("Published message %d of %d", publishingCount.Add(1), t.messageCount)
		}()
	}

	publishingWorkers.Wait()
}
