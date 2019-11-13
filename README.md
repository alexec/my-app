# My App

Build:

```
GOOS=linux GOARCH=amd64 go build .
docker build -t my-app:v2 .
```

Run:

```
docker run -P -p 8080:8080 -e GREETING=Howdy my-app:v2
```

### Release:

```
docker tag my-app:v2 alexcollinsintuit/my-app:v2
docker push alexcollinsintuit/my-app:v2
```

### Clean Up

```
docker rmi my-app:v2
```
