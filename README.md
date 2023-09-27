# devproxy

A proxy for simple local development that can route paths to different remote targets

## Installation

```sh
go install github.com/seanpfeifer/devproxy
```

## Usage

```sh
# Proxy /api/* to the local server on port 9191
# Proxy all other paths to the local server on port 9595
devproxy  -proxy /api/->http://0.0.0.0:9191 -proxy /->http://0.0.0.0:9595
```
