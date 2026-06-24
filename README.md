# realtime-nats

A **NATS** transport for [togo](https://to-go.dev) realtime. Implements `togo.Broker`:
`Publish` fans events over a NATS subject (so every app instance — and any other
service on your NATS — receives them), and the SSE `Handler` subscribes to NATS and
streams events to browsers. Overrides the default in-process SSE broker.

Great for **multi-instance fan-out** and bridging your message queue to services.

## Install

```bash
togo install togo-framework/realtime-nats
```

## Configure (`.env`)

```ini
NATS_URL=nats://127.0.0.1:4222
NATS_SUBJECT=togo.realtime
```

If NATS is unreachable at boot, the plugin logs and keeps the default broker.

MIT © togo-framework
