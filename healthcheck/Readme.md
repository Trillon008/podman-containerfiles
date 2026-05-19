# Healthcheck

A small Go HTTP probe binary for container `HEALTHCHECK`.

## Build

```bash
podman build -t ghcr.io/trillon008/healthcheck:1.0.0 .
```

## Usage

```bash
podman run --rm \
  -e PROBE_URL=http://127.0.0.1:8080/healthz \
  ghcr.io/trillon008/healthcheck:1.0.0
```

## Reuse in another Containerfile

```Dockerfile
COPY --from=ghcr.io/trillon008/healthcheck:1.0.0 /healthcheck /usr/local/bin/healthcheck

ENV PROBE_URL=http://127.0.0.1:8080/healthz
HEALTHCHECK CMD ["/usr/local/bin/healthcheck"]
```
