# devproxy

A reverse proxy for simple local development that can route paths to different remote targets.

This can be useful in cases where you are developing a webapp that needs to make requests to an API that should look like it's on the same host.

> [!WARNING]
> This is intended for local development only. It should not be used in production.

## Installation

```sh
go install github.com/seanpfeifer/devproxy
```

## Usage

```sh
# Set up the reverse proxy to listen on port 9090 (8080 default)
# Proxy /api/* to the local server on port 9191
# Proxy all other paths to the local server on port 9595
devproxy -port 9090 -proxy /api/->http://0.0.0.0:9191 -proxy /->http://0.0.0.0:9595
```
