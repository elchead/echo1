FROM golang:alpine AS build
WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod .
COPY main.go .


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./echo ./main.go

FROM scratch AS export-stage
COPY --from=build /usr/src/app/echo .
CMD ["./echo"]
