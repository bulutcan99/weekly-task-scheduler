# build stage
FROM golang:1.22.0-alpine3.18 AS build

# set working directory
WORKDIR /app

# copy source code
COPY . .

# install dependencies
RUN go mod download

# build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main main.go

# final stage
FROM alpine:3.18 AS final
LABEL maintainer="golang"
# set working directory
WORKDIR /app

# copy binary
COPY --from=build /app/bin/main ./

EXPOSE 8080

ENTRYPOINT [ "./main" ]
