// Package natsrt is a NATS transport for togo realtime. It implements togo.Broker:
// Publish fans events out over a NATS subject (so every app instance and any other
// service connected to NATS receives them), and Handler() is an SSE endpoint that
// subscribes to NATS and streams events to browser clients. Overrides the default
// in-process SSE broker when installed.
//
//	togo install togo-framework/realtime-nats
//
// Env: NATS_URL (default nats://127.0.0.1:4222), NATS_SUBJECT (default togo.realtime).
package natsrt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/nats-io/nats.go"

	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("realtime-nats", togo.PriorityService+10, func(k *togo.Kernel) error {
		url := os.Getenv("NATS_URL")
		if url == "" {
			url = nats.DefaultURL
		}
		nc, err := nats.Connect(url, nats.Name("togo-realtime"), nats.MaxReconnects(-1))
		if err != nil {
			if k.Log != nil {
				k.Log.Warn("realtime-nats: connect failed (%v); keeping default broker", err)
			}
			return nil
		}
		subject := os.Getenv("NATS_SUBJECT")
		if subject == "" {
			subject = "togo.realtime"
		}
		k.Realtime = &broker{nc: nc, subject: subject}
		return nil
	})
}

type event struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

type broker struct {
	nc      *nats.Conn
	subject string
}

// Publish marshals the event and publishes it to the NATS subject.
func (b *broker) Publish(ev, data string) {
	buf, _ := json.Marshal(event{Event: ev, Data: data})
	_ = b.nc.Publish(b.subject, buf)
}

// Handler streams events to an SSE client by subscribing to the NATS subject.
func (b *broker) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "streaming unsupported", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		ch := make(chan *nats.Msg, 64)
		sub, err := b.nc.ChanSubscribe(b.subject, ch)
		if err != nil {
			http.Error(w, "subscribe failed", http.StatusInternalServerError)
			return
		}
		defer sub.Unsubscribe()

		fmt.Fprint(w, ": connected\n\n")
		flusher.Flush()

		for {
			select {
			case <-r.Context().Done():
				return
			case m := <-ch:
				var ev event
				if json.Unmarshal(m.Data, &ev) != nil {
					continue
				}
				fmt.Fprintf(w, "event: %s\ndata: %s\n\n", ev.Event, ev.Data)
				flusher.Flush()
			}
		}
	}
}
