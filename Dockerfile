# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM golang:1.17.8-alpine as build

WORKDIR /app

ENV GO111MODULE=on
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY handler/ ./handler/ 
COPY store/ ./store/ 

ARG TARGETOS TARGETARCH

RUN GOOS=$TARGETOS GOARCH=$TARGETARCH CGO_ENABLED=0 go build -o /drinkpipe-ui

##
## Deploy
##
FROM gcr.io/distroless/base-debian11

EXPOSE 3000

USER nonroot:nonroot
WORKDIR /

COPY views/ /views/ 
COPY static/ /static/ 

COPY --from=build /drinkpipe-ui /drinkpipe-ui

ENTRYPOINT ["/drinkpipe-ui"]
