FROM golang:1.14.4-alpine3.12 as builder

WORKDIR /opt/math-custom

COPY . .

RUN go build

RUN rm -Rf *.go

FROM hello-world:latest

COPY --from=builder /opt/math-custom .

ENTRYPOINT [ "./math-custom" ]