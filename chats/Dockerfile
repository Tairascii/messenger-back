FROM golang:1.24.2-alpine AS builder

WORKDIR /build

COPY ./chats/go.mod go.mod
COPY ./chats/go.sum go.sum

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /chats ./chats/cmd/chats/main.go

FROM scratch

COPY --from=builder chats /bin/chats/cmd
COPY ./chats/envs/values_docker.yaml /bin/chats/envs/values_docker.yaml

ENV CONFIG_FILE=/bin/chats/envs/values_docker.yaml

ENTRYPOINT ["/bin/chats/cmd"]
