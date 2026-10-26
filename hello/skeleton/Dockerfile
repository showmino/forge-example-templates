# Docker Hub's official golang image through Google's mirror: CI runners share
# IPs, and Docker Hub rate-limits anonymous pulls.
FROM mirror.gcr.io/library/golang:1.27 AS build
WORKDIR /src
COPY go.mod ./
COPY *.go ./
RUN CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o /out/{% .Service.Name %} .

FROM gcr.io/distroless/static-debian12:nonroot
# Links the image on ghcr.io to its repository, so the repository's workflow
# may push new versions of it.
LABEL org.opencontainers.image.source={% .Repo.URL %}
COPY --from=build /out/{% .Service.Name %} /{% .Service.Name %}
USER 65532:65532
EXPOSE {% .Params.port %}
ENTRYPOINT ["/{% .Service.Name %}"]
