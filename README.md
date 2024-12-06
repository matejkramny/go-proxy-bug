# README

This repo attempts to replicate a bug with httputil.ReverseProxy which causes client disconnects when using websockets in specific conditions.

Reproduction steps:
1. build: `make build`
1. start websocket server `make server`
  - the server sends 2 messages with 5 second timeouts
1. (another tab) start proxy to server `./bin/proxy-bug proxy`
1. (another tab) start client to websocket `./bin/proxy-bug client raw`

## Correct/desired output

```
client called
Received: �server closed request body...
Received: �this isn't received?.
2024/12/06 10:05:19 EOF
```

## Bug with proxy

When we add the `--disconnect` flag: `./bin/proxy-bug client raw --disconnect`, this makes the client disconnect the outbound connection after receiving some data:

> // CloseWrite shuts down the writing side of the TCP connection.

You will get the following from the client:
```
client called
Received: �server closed request body...
Closing write
<nil>
2024/12/06 10:05:06 EOF
```

## Without the proxy

The proxy is closing the inbound connection in the above example, as it works correctly when you connect directly to the server.

`./bin/proxy-bug client raw --disconnect --direct`

```
client called
Received: �server closed request body...
Closing write
<nil>
Received: �this isn't received?.
Closing write
<nil>
2024/12/06 10:19:15 EOF
```
