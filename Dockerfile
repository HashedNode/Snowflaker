FROM golang:1.24.4-bookworm AS build

LABEL authors="HashedNode"

RUN useradd -u 1001 nonroot
WORKDIR /app
COPY go.mod ./

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download


COPY . .

RUN go build \
    -ldflags="-linkmode external -extldflags -static -s -w" \
    -tags netgo \
    -o crystal-snowflake


FROM scratch

COPY --from=build /etc/passwd /etc/passwd

COPY --from=build /app/crystal-snowflake /crystal-snowflake

USER nonroot

EXPOSE 8080

CMD ["./crystal-snowflake"]