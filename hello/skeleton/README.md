# {% .Service.Name %}

A small Go web service, made by Forge from the `hello` template
({% .Template.Version %}). It answers `GET /` with a greeting and
`GET /healthz` with 200.

```sh
go test ./...
go run .            # listens on :{% .Params.port %}, or :$PORT when PORT is set
```

## How it ships

Every push to `main` runs `.github/workflows/ci.yml` on GitHub-hosted runners:

1. **test**: `go vet` and `go test -race`.
2. **image**: builds the Dockerfile and pushes
   `ghcr.io/{% .Repo.Owner | lower %}/{% .Service.Name %}:<first 7 characters of the commit>`.
3. **deploy**: downloads the `forge` CLI from your Forge and runs
   `forge deploy start --env staging`. The job signs in with its GitHub OIDC
   token, which only allows deploying this service to staging at this commit.

Promote to production yourself:

```sh
forge deploy promote --yes
```

The deploy job finds Forge through the organization variable
`FORGE_API_URL`, or the address written into the workflow when the service
was created. Change the `env` at the top of `ci.yml` if Forge moves.
