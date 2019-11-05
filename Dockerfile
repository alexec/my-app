FROM golang:1.12 as builder

ADD . /src
RUN go build -o /dist/my-app /src/my-app.go

FROM scratch

COPY --from=builder /dist/my-app /my-app
EXPOSE 8080
ENV GREETING "Hello world!"
ENTRYPOINT ["/my-app"]
