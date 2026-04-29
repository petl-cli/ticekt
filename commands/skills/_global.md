---
name: _global
description: Cross-cutting reference — auth, output flags, exit codes, error format. Always relevant.
when_to_use: Load alongside any group skill. Covers behavior shared by every command.
---

# discovery-api — global reference

The Ticketmaster Discovery API allows you to search for events, attractions, or venues.

- **Binary**: `discovery-api`
- **Base URL**: `https://app.ticketmaster.com`
- **Version**: `v2`

## Authentication

- **apiKey** — set `DISCOVERY_API_API_KEY` or run `discovery-api configure`


## Output flags

| Flag | Effect |
|------|--------|
| `--output-format`, `-o` | `json` (default), `compact`, `yaml`, `raw`, `table`, `pretty` |
| `--jq <query>` | Extract fields with GJSON syntax — see below |
| `--agent-mode` | Force compact output + structured errors (auto-detected in agent envs) |
| `--dry-run` | Print the exact HTTP request without sending it |
| `--schema` | Print JSON schema for the command's inputs and outputs |
| `--debug` | Log request + response details to stderr |
| `--no-retries` | Disable retries on 429 / 5xx |
| `--base-url` | Override the API base URL |

### GJSON query reference

```
--jq id                        scalar field
--jq user.email                nested field
--jq items.#.id                all ids from an array
--jq items.0.name              first element
--jq "items.#(active==true)#"  filter array by condition
--jq items.#                   count items in an array
```

## Exit codes

Branch on `$?` — never parse stderr.

| Code | Meaning |
|------|---------|
| 0 | success |
| 1 | unknown error |
| 2 | auth failed (401 / 403) |
| 3 | not found (404) |
| 4 | validation error (400 / 422) |
| 5 | rate limited (429) |
| 6 | server error (5xx) |
| 7 | network error |

## Error format

All errors go to stderr as a single JSON line. stdout stays clean.

```json
{"error":true,"code":"not_found","status":404,"message":"Resource not found","suggestion":"Verify the ID or path is correct","exit_code":3}
```

Codes: `auth_failed` | `forbidden` | `not_found` | `validation_error` | `rate_limited` | `server_error` | `network_error` | `request_failed`

## Discovery

```bash
discovery-api skills list                # index of all skills (one per command group)
discovery-api skills get <group>          # full skill for that group
discovery-api agent-instructions          # consolidated llms.txt
discovery-api <group> <command> --schema  # per-command JSON schema
```

## Available skills

- `v2` — operations on v2
