# vps-health

A CLI tool that checks the health of multiple web services in parallel — status code, latency, and SSL certificate expiry.

## Requirements

- Go 1.21+

## Install & run

```bash
git clone https://github.com/ibrahima-gh/vps-health
cd vps-health
go run .
```

### Output

```
checking 3 targets (timeout: 10s)...

✓  Google                200  143ms  SSL ok (53d)
✗  mysite.dev            request failed: context deadline exceeded
```

- Green `✓` — service is up
- Red `✗` — request failed (timeout, DNS error, etc.)
- SSL warning in yellow when certificate expires in less than 14 days

## Configuration

Edit `config.yaml` to add or remove targets:

```yaml
timeout_seconds: 10

targets:
  - name: Google
    url: https://google.com

  - name: My site
    url: https://example.com
```

| Field             | Description                        |
|-------------------|------------------------------------|
| `timeout_seconds` | Max wait per request (default: 10) |
| `targets[].name`  | Display name shown in output       |
| `targets[].url`   | Full URL to check (http or https)  |

## Project structure

```
.
├── main.go                    # Entry point, output rendering
├── config.yaml                # Target list
├── internal/
│   ├── config/config.go       # YAML loading and validation
│   └── checker/checker.go     # Parallel HTTP checks
```
