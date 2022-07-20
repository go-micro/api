# API Gateway

An api gateway for Go Micro services

## Overview

The API gateway dynamically serves requests via HTTP json to go-micro services using path based resolution.

## Example

The request http://localhost:8080/helloworld/call will route to the service `helloworld` and endpoint `Helloworld.Call`.

### For differing handlers

The request http://localhost:8080/helloworld/Greeter/Call will route to the service `helloworld` and endpoint `Greeter.Call`

## Usage

Install

```
go install github.com/go-micro/api/cmd/api
```

Run (listens on :8080)

```
api
```

## TODO

- Enable changing registry, client/server, etc
