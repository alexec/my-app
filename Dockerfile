FROM scratch

ADD my-app /

EXPOSE 8080

ENV GREETING "Hello world!"

ENTRYPOINT ["/my-app"]
