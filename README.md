# Simple API Server

This is a Simple API Server built from the ground up with Go.

## Usage

```bash
go run main.go
```

## Endpoints

### /time

Returns the current server time and presents a list of 600+ timezones to choose from.

### /time/{region}/{timezone}

Returns the current time in the specified location.

Supported locations can be found [here](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List).

If the region or timezone is not found, it will list all the regions and timezones available.
