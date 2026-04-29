package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// agentInstructionsCmd prints the llms.txt content at runtime so agents can
// include it in their system prompt without needing a separate file:
//
//	INSTRUCTIONS=$(discovery-api agent-instructions)
var agentInstructionsCmd = &cobra.Command{
	Use:   "agent-instructions",
	Short: "Print machine-readable instructions for AI agents (llms.txt format)",
	Long: `Prints a complete description of this CLI's commands, flags, exit codes,
and usage patterns optimised for inclusion in an AI agent's system prompt.

Example:
  # Include in Claude Code context:
  discovery-api agent-instructions > CLAUDE.md

  # Capture inline:
  INSTRUCTIONS=$(discovery-api agent-instructions)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print(agentInstructionsContent)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(agentInstructionsCmd)
}

// agentInstructionsContent is the full llms.txt baked into the binary at build time.
// Regenerate by re-running the CLI generator against the same OpenAPI spec.
const agentInstructionsContent = `# Discovery API CLI
# Version: v2
# This file is machine-generated. Include it in your system prompt for zero-doc CLI usage.

## OVERVIEW
The Ticketmaster Discovery API allows you to search for events, attractions, or venues.
Binary: discovery-api
Base URL: https://app.ticketmaster.com

## AUTHENTICATION
API key — set env var DISCOVERY_API_API_KEY or pass --api-key <key>

## OUTPUT CONTROL
--output-format / -o   json (default) | compact | yaml | raw | table | pretty
--jq <query>           Extract fields using GJSON syntax (see examples below)
--agent-mode           Force compact output + structured errors (auto-detected in agent envs)

GJSON query examples:
  --jq id                        → scalar field
  --jq user.email                → nested field
  --jq items.#.id                → all ids from array
  --jq items.0.name              → first element
  --jq "items.#(active==true)#"  → filter array by condition
  --jq items.#                   → count items in array

## OPERATIONAL FLAGS
--dry-run        Print the exact HTTP request without sending it (safe preview)
--schema         Print full input/output schema for a command as JSON (same as --help in agent mode)
--debug          Log request + response details to stderr
--no-retries     Disable automatic retries on 429 and 5xx responses
--base-url       Override the API base URL at runtime

## EXIT CODES
Branch on $? without parsing stderr:
  0  success
  1  unknown error
  2  auth failed       (401 / 403)
  3  not found         (404)
  4  validation error  (400 / 422)
  5  rate limited      (429)
  6  server error      (5xx)
  7  network error

## ERROR FORMAT
All errors go to stderr as a single JSON line. stdout is always clean for piping.
{"error":true,"code":"not_found","status":404,"message":"Resource not found","suggestion":"Verify the ID or path is correct","exit_code":3}

Error codes: auth_failed | forbidden | not_found | validation_error | rate_limited | server_error | network_error | request_failed

## HELP SYSTEM
--help in human mode  → readable prose
--help in agent mode  → JSON schema (same as --schema)
--help on group cmd   → JSON list of subcommands

## SKILLS
For deeper, per-group guidance (workflows, gotchas, exact flag examples), load a skill on demand instead of ingesting this whole file:
  discovery-api skills list                # JSON index — name + description + when_to_use per group
  discovery-api skills get <group>         # full markdown body for one group
  discovery-api skills get _global         # cross-cutting reference (auth, output, exit codes)
Prefer skills over this file once you know which group(s) you need — they're scoped, smaller, and editable in the studio.

## COMMANDS

### GROUP: v2

  discovery-api v2 event-search
  Description : Event Search
  HTTP        : GET /discovery/v2/events
  Safe        : true | Idempotent: true | Reversible: true | Impact: low
  Auth        : true
  Flags:
    --sort (string) Sorting order of the search result. Allowable values : 'name,asc', 'name,desc', 'date,asc', 'date,desc', 'relevance,asc', 'relevance,desc', 'distance,asc', 'name,date,asc', 'name,date,desc', 'date,name,asc', 'date,name,desc','onsaleStartDate,asc', 'id,asc'
    --start-date-time (string) Filter events with a start date after this date
    --end-date-time (string) Filter events with a start date before this date
    --onsale-start-date-time (string) Filter events with onsale start date after this date
    --onsale-on-start-date (string) Filter events with onsale start date on this date
    --onsale-on-after-start-date (string) Filter events with onsale range within this date
    --onsale-end-date-time (string) Filter events with onsale end date before this date
    --city (string) Filter events by city
    --country-code (string) Filter events by country code
    --state-code (string) Filter events by state code
    --postal-code (string) Filter events by postal code / zipcode
    --venue-id (string) Filter events by venue id
    --attraction-id (string) Filter events by attraction id
    --segment-id (string) Filter events by segment id
    --segment-name (string) Filter events by segment name
    --classification-name (array) Filter events by classification name: name of any segment, genre, sub-genre, type, sub-type
    --classification-id (array) Filter events by classification id: id of any segment, genre, sub-genre, type, sub-type
    --market-id (string) Filter events by market id
    --promoter-id (string) Filter events by promoter id
    --dma-id (string) Filter events by dma id
    --include-tba (string) True, to include events with date to be announce (TBA) [yes, no, only]
    --include-tbd (string) True, to include event with a date to be defined (TBD) [yes, no, only]
    --client-visibility (string) Filter events by clientName
    --latlong (string) Filter events by latitude and longitude, this filter is deprecated and maybe removed in a future release, please use geoPoint instead
    --radius (string) Radius of the area in which we want to search for events.
    --unit (string) Unit of the radius [miles,km]
    --geo-point (string) filter events by geoHash
    --keyword (string) Keyword to search on
    --id (string) Filter entities by its id
    --source (string) Filter entities by its source name [ticketmaster, universe, frontgate, tmr]
    --include-test (string) True if you want to have entities flag as test in the response. Only, if you only wanted test entities [yes, no, only]
    --page (string) Page number
    --size (string) Page size of the response
    --locale (string) The locale in ISO code format. Multiple comma-separated values can be provided. When omitting the country part of the code (e.g. only 'en' or 'fr') then the first matching locale is used. When using a '*' it matches all locales. '*' can only be used at the end (e.g. 'en-us,en,*') 
    --include-licensed-content (string) Yes if you want to display licensed content [yes, no]
    --include-spellcheck (string) yes, to include spell check suggestions in the response. [yes, no]

  discovery-api v2 find-suggest
  Description : Find Suggest
  HTTP        : GET /discovery/v2/suggest
  Safe        : true | Idempotent: true | Reversible: true | Impact: low
  Auth        : true
  Flags:
    --keyword (string) Keyword to search on
    --source (string) Filter entities by its source name [ticketmaster, universe, frontgate, tmr]
    --latlong (string) Filter events by latitude and longitude, this filter is deprecated and maybe removed in a future release, please use geoPoint instead
    --radius (string) Radius of the area in which we want to search for events.
    --unit (string) Unit of the radius [miles,km]
    --size (string) Size of every entity returned in the response
    --include-fuzzy (string) yes, to include fuzzy matches in the search. This has performance impact. [yes, no]
    --client-visibility (string) Filter events to clientName
    --country-code (string) Filter suggestions by country code
    --include-tba (string) True, to include events with date to be announce (TBA) [yes, no, only]
    --include-tbd (string) True, to include event with a date to be defined (TBD) [yes, no, only]
    --segment-id (string) Filter suggestions by segment id
    --geo-point (string) filter events by geoHash
    --locale (string) The locale in ISO code format. Multiple comma-separated values can be provided. When omitting the country part of the code (e.g. only 'en' or 'fr') then the first matching locale is used. When using a '*' it matches all locales. '*' can only be used at the end (e.g. 'en-us,en,*') 
    --include-licensed-content (string) Yes if you want to display licensed content [yes, no]
    --include-spellcheck (string) yes, to include spell check suggestions in the response. [yes, no]

  discovery-api v2 find-venues
  Description : Venue Search
  HTTP        : GET /discovery/v2/venues
  Safe        : true | Idempotent: true | Reversible: true | Impact: low
  Auth        : true
  Flags:
    --sort (string) Sorting order of the search result. Allowable Values: 'name,asc', 'name,desc', 'relevance,asc', 'relevance,desc', 'distance,asc', 'distance,desc'
    --state-code (string) Filter venues by state / province code
    --country-code (string) Filter venues by country code
    --latlong (string) Filter events by latitude and longitude, this filter is deprecated and maybe removed in a future release, please use geoPoint instead
    --radius (string) Radius of the area in which we want to search for events.
    --unit (string) Unit of the radius [miles,km]
    --geo-point (string) filter events by geoHash
    --keyword (string) Keyword to search on
    --id (string) Filter entities by its id
    --source (string) Filter entities by its source name [ticketmaster, universe, frontgate, tmr]
    --include-test (string) True if you want to have entities flag as test in the response. Only, if you only wanted test entities [yes, no, only]
    --page (string) Page number
    --size (string) Page size of the response
    --locale (string) The locale in ISO code format. Multiple comma-separated values can be provided. When omitting the country part of the code (e.g. only 'en' or 'fr') then the first matching locale is used. When using a '*' it matches all locales. '*' can only be used at the end (e.g. 'en-us,en,*') 
    --include-licensed-content (string) Yes if you want to display licensed content [yes, no]
    --include-spellcheck (string) yes, to include spell check suggestions in the response. [yes, no]

  discovery-api v2 get-attraction-details
  Description : Get Attraction Details
  HTTP        : GET /discovery/v2/attractions/{id}
  Safe        : true | Idempotent: true | Reversible: true | Impact: low
  Auth        : true
  Flags:
    --id* (string) ID of the attraction
    --locale (string) The locale in ISO code format. Multiple comma-separated values can be provided. When omitting the country part of the code (e.g. only 'en' or 'fr') then the first matching locale is used. When using a '*' it matches all locales. '*' can only be used at the end (e.g. 'en-us,en,*') 
    --include-licensed-content (string) True if you want to display licensed content [yes, no]

  discovery-api v2 get-classification-details
  Description : Get Classification Details
  HTTP        : GET /discovery/v2/classifications/{id}
  Safe        : true | Idempotent: true | Reversible: true | Impact: low
  Auth        : true
  Flags:
    --id* (string) ID of the segment, genre, or sub-genre
    --locale (string) The locale in ISO code format. Multiple comma-separated values can be provided. When omitting the country part of the code (e.g. only 'en' or 'fr') then the first matching locale is used. When using a '*' it matches all locales. '*' can only be used at the end (e.g. 'en-us,en,*') 
    --include-licensed-content (string) True if you want to display licensed content [yes, no]

  discovery-api v2 get-event-details
  Description : Get Event Details
  HTTP        : GET /discovery/v2/events/{id}
  Safe        : true | Idempotent: true | Reversible: true | Impact: low
  Auth        : true
  Flags:
    --id* (string) ID of the event
    --locale (string) The locale in ISO code format. Multiple comma-separated values can be provided. When omitting the country part of the code (e.g. only 'en' or 'fr') then the first matching locale is used. When using a '*' it matches all locales. '*' can only be used at the end (e.g. 'en-us,en,*') 
    --include-licensed-content (string) True if you want to display licensed content [yes, no]

  discovery-api v2 get-event-images
  Description : Get Event Images
  HTTP        : GET /discovery/v2/events/{id}/images
  Safe        : true | Idempotent: true | Reversible: true | Impact: low
  Auth        : true
  Flags:
    --id* (string) ID of the event
    --locale (string) The locale in ISO code format. Multiple comma-separated values can be provided. When omitting the country part of the code (e.g. only 'en' or 'fr') then the first matching locale is used. When using a '*' it matches all locales. '*' can only be used at the end (e.g. 'en-us,en,*') 
    --include-licensed-content (string) True if you want to display licensed content [yes, no]

  discovery-api v2 get-genre-details
  Description : Get Genre Details
  HTTP        : GET /discovery/v2/classifications/genres/{id}
  Safe        : true | Idempotent: true | Reversible: true | Impact: low
  Auth        : true
  Flags:
    --id* (string) ID of the genre
    --locale (string) The locale in ISO code format. Multiple comma-separated values can be provided. When omitting the country part of the code (e.g. only 'en' or 'fr') then the first matching locale is used. When using a '*' it matches all locales. '*' can only be used at the end (e.g. 'en-us,en,*') 
    --include-licensed-content (string) True if you want to display licensed content [yes, no]

  discovery-api v2 get-segment-details
  Description : Get Segment Details
  HTTP        : GET /discovery/v2/classifications/segments/{id}
  Safe        : true | Idempotent: true | Reversible: true | Impact: low
  Auth        : true
  Flags:
    --id* (string) ID of the segment
    --locale (string) The locale in ISO code format. Multiple comma-separated values can be provided. When omitting the country part of the code (e.g. only 'en' or 'fr') then the first matching locale is used. When using a '*' it matches all locales. '*' can only be used at the end (e.g. 'en-us,en,*') 
    --include-licensed-content (string) True if you want to display licensed content [yes, no]

  discovery-api v2 get-subgenre-details
  Description : Get Sub-Genre Details
  HTTP        : GET /discovery/v2/classifications/subgenres/{id}
  Safe        : true | Idempotent: true | Reversible: true | Impact: low
  Auth        : true
  Flags:
    --id* (string) ID of the subgenre
    --locale (string) The locale in ISO code format. Multiple comma-separated values can be provided. When omitting the country part of the code (e.g. only 'en' or 'fr') then the first matching locale is used. When using a '*' it matches all locales. '*' can only be used at the end (e.g. 'en-us,en,*') 
    --include-licensed-content (string) True if you want to display licensed content [yes, no]

  discovery-api v2 get-venue-details
  Description : Get Venue Details
  HTTP        : GET /discovery/v2/venues/{id}
  Safe        : true | Idempotent: true | Reversible: true | Impact: low
  Auth        : true
  Flags:
    --id* (string) ID of the venue
    --locale (string) The locale in ISO code format. Multiple comma-separated values can be provided. When omitting the country part of the code (e.g. only 'en' or 'fr') then the first matching locale is used. When using a '*' it matches all locales. '*' can only be used at the end (e.g. 'en-us,en,*') 
    --include-licensed-content (string) True if you want to display licensed content [yes, no]

  discovery-api v2 search-attractions
  Description : Attraction Search
  HTTP        : GET /discovery/v2/attractions
  Safe        : true | Idempotent: true | Reversible: true | Impact: low
  Auth        : true
  Flags:
    --sort (string) Sorting order of the search result. Allowable Values : 'name,asc', 'name,desc', 'relevance,asc', 'relevance,desc'
    --classification-name (array) Filter attractions by classification name: name of any segment, genre, sub-genre, type, sub-type
    --classification-id (array) Filter attractions by classification id: id of any segment, genre, sub-genre, type, sub-type
    --keyword (string) Keyword to search on
    --id (string) Filter entities by its id
    --source (string) Filter entities by its source name [ticketmaster, universe, frontgate, tmr]
    --include-test (string) True if you want to have entities flag as test in the response. Only, if you only wanted test entities [yes, no, only]
    --page (string) Page number
    --size (string) Page size of the response
    --locale (string) The locale in ISO code format. Multiple comma-separated values can be provided. When omitting the country part of the code (e.g. only 'en' or 'fr') then the first matching locale is used. When using a '*' it matches all locales. '*' can only be used at the end (e.g. 'en-us,en,*') 
    --include-licensed-content (string) Yes if you want to display licensed content [yes, no]
    --include-spellcheck (string) yes, to include spell check suggestions in the response. [yes, no]

  discovery-api v2 search-classifications
  Description : Classification Search
  HTTP        : GET /discovery/v2/classifications
  Safe        : true | Idempotent: true | Reversible: true | Impact: low
  Auth        : true
  Flags:
    --sort (string) Sorting order of the search result
    --keyword (string) Keyword to search on
    --id (string) Filter entities by its id
    --source (string) Filter entities by its source name [ticketmaster, universe, frontgate, tmr]
    --include-test (string) True if you want to have entities flag as test in the response. Only, if you only wanted test entities [yes, no, only]
    --page (string) Page number
    --size (string) Page size of the response
    --locale (string) The locale in ISO code format. Multiple comma-separated values can be provided. When omitting the country part of the code (e.g. only 'en' or 'fr') then the first matching locale is used. When using a '*' it matches all locales. '*' can only be used at the end (e.g. 'en-us,en,*') 
    --include-licensed-content (string) Yes if you want to display licensed content [yes, no]
    --include-spellcheck (string) yes, to include spell check suggestions in the response. [yes, no]

## AGENT WORKFLOWS

### Safely delete a resource
  discovery-api <group> delete --id X --dry-run   # verify request before sending
  discovery-api <group> delete --id X             # execute; exit 0 = deleted, exit 3 = not found

### Extract a field without jq installed
  ID=$(discovery-api <group> create [flags] --jq id)
  discovery-api <group> get --id "$ID"

### Get all emails from a list response
  discovery-api <group> list --jq "items.#.email"

### Handle errors by exit code
  discovery-api <group> get --id X
  case $? in
    0) : success ;;
    2) echo "fix credentials" ;;
    3) echo "not found, skip" ;;
    5) sleep 10 ; retry ;;
    6) echo "server down, abort" ;;
  esac

### Pipe between commands (auto compact JSONL when stdout is a pipe)
  discovery-api <group> list | while IFS= read -r line; do
    ID=$(echo "$line" | discovery-api <group> process --body "$line" --jq id)
  done

### Discover a command's interface (agent mode)
  CLAUDE_CODE=1 discovery-api <group> <command> --help   # returns JSON schema
  discovery-api <group> <command> --schema               # explicit, same result

### Reload these instructions at runtime
  discovery-api agent-instructions
`
