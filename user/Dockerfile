FROM golang:1.24.2-alpine AS builder

WORKDIR /build

COPY ./user/go.mod go.mod
COPY ./user/go.sum go.sum

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /user ./user/cmd/user/main.go

FROM scratch

COPY --from=builder user /bin/user/cmd
COPY ./user/envs/values_docker.yaml /bin/user/envs/values_docker.yaml

ENV CONFIG_FILE=/bin/user/envs/values_docker.yaml

ENTRYPOINT ["/bin/user/cmd"]
