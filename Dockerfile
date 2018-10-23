FROM golang:1.10 AS BUILD
COPY . /work
ENV CGO_ENABLED=0
RUN go get github.com/go-stomp/stomp && \
  cd /work/ && \
  go build main.go && \
  go build reader.go
FROM alpine
COPY --from=BUILD /work/main /usr/bin/publisher
COPY --from=BUILD /work/reader /usr/bin/subscriber
