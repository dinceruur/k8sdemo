FROM golang:1.11-alpine as build

WORKDIR /src/
COPY main.go go.* /src/
RUN apk add --no-cache git
RUN CGO_ENABLED=0 go get gopkg.in/yaml.v2
RUN CGO_ENABLED=0 go build -o /bin/demo

FROM scratch
COPY --from=build /bin/demo /bin/demo
ENTRYPOINT ["/bin/demo"]