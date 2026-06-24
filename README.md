<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/realtime-nats</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/realtime-nats"><img src="https://pkg.go.dev/badge/github.com/togo-framework/realtime-nats.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/realtime-nats
```

<!-- /togo-header -->

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

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
