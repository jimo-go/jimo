# JIMO (Skeleton Application)

This repository is the **starter/skeleton application** for the JIMO ecosystem.

- The **framework (core engine)** lives in: `github.com/jimo-go/framework`
- This repo (`github.com/jimo-go/jimo`) is a **clean application template**.

## What you get

- `cmd/server/main.go` as the application entrypoint
- `routes/` for route definitions
- `app/http/controllers/` for HTTP controllers
- `app/models/` for domain models
- `resources/views/` for HTML templates
- `configs/` for application configuration

## Requirements

- Go 1.22+

## Running the app

From the project root:

```bash
jimo serve
```

Then open:

- `http://localhost:8000`

You can change the port using:

```bash
jimo serve --port 9000
```

## Local framework development

This skeleton depends on `github.com/jimo-go/framework`.

For local development against a local checkout of the framework, you have two options:

```bash
go work init
go work use ../framework
go work use .
```

Or, add a temporary `replace` in your app's `go.mod`:

```go
replace github.com/jimo-go/framework => ../framework
```

## Next steps

- Create your own controllers in `app/http/controllers/`
- Register routes in `routes/web.go`
- Add templates in `resources/views/`
