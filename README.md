# Post Manager API
Post manager API 

### Tech
Dependencies of this project:

* [GORM] - Fantastic ORM library for Golang
* [Gin] - Web framework written in Go

##### Build
```sh
$ go build *.go
```

##### Run
```sh
$ go run *.go
```

### Requests
Making requests in terminal

##### POST
```sh
$ curl -i -X POST http://localhost:8080/posts/ -d '{"title" :  "...", "description" : "...", "body": "..."}'
```

##### GET
```sh
$ curl -i -X GET http://localhost:8080/posts/
```

##### GET ONE
```sh
$ curl -i -X GET http://localhost:8080/posts/id
```

##### PUT
```sh
$ curl -i -X PUT http://localhost:8080/posts/ -d '{"title" :  ".^.", "description" : ".^.", "body": ".^."}'
```

[//]: # (Reference links)
   [matheus]: <https://github.com/matheusps>
   [git-repo-url]: <https://github.com/matheusps/personal-page-react>
   [Gin]: <https://github.com/gin-gonic/gin>
   [GORM]: <http://jinzhu.me/gorm/#>