# IDDD_go_sample
sample of iddd for golang

# Usage

```
http://localhost:8000/article/1
http://localhost:8000/articles?user_id=1
```

## mock mode

### Requires
nothing

```
$ go build +build=mock
$ ./IDDD_go_sample
```

## implemented mode


### Requires

- Docker
- docker-compose

```
$ docker-compose -f Docker/docker-compose.yml up
```
