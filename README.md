# My App

Build:

```
GOOS=linux GOARCH=amd64  go build .
docker build -t my-app:v1 .
```

Run:

```
docker run -P -p 8080:8080 my-app:v1
```