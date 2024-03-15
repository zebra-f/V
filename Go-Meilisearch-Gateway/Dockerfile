FROM golang:1.22.1-alpine3.19 as builder 

WORKDIR /go_app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o go_meilisearch_gateway

FROM alpine:3.19.1

RUN mkdir /go_app

COPY --from=builder /go_app/go_meilisearch_gateway /go_app/go_meilisearch_gateway

CMD ["/go_app/go_meilisearch_gateway", "container", "8080"]
# CMD ["tail", "-f", "/dev/null"]
