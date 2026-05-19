# Caddy

Caddy image built for rootless Podman deployments.

## Features

- rootless execution with Podman
- non-root execution inside the container
- custom Caddy binary with the OVH DNS module
- ACME DNS challenge support through OVH

## Image

The Caddy binary is built with [`xcaddy`](https://github.com/caddyserver/xcaddy) and includes:

- `github.com/caddy-dns/ovh`

## Runtime

The image is intended to run with Podman and a non-privileged user inside the container.

The Quadlet example in this repository:

- uses `UserNS=keep-id`
- adds `CAP_NET_BIND_SERVICE`
- mounts Caddy configuration and data from the host

## ACME DNS challenge

The OVH module allows Caddy to solve ACME DNS challenges without relying on an HTTP challenge exposed on port `80`.

Required OVH credentials must be provided through environment variables, typically from an environment file used by Quadlet.

## Healthcheck

The image includes a reusable HTTP probe binary and defines a `HEALTHCHECK` against Caddy's local admin API:

- `PROBE_URL=http://127.0.0.1:2019/config/`
