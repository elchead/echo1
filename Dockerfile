FROM golang:1.19 AS build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod main.go

COPY . .
RUN CGO_ENABLED=0 go build -v -o ./echo ./...


FROM scratch AS export-stage
COPY --from=build ./echo .
#CMD ["echo"]