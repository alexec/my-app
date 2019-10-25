FROM scratch

ADD my-app /

EXPOSE 8080

ENTRYPOINT ["/my-app"]
