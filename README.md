# hexagonal-todo-gin

https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749

https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3
```
├── cmd
├── pkg
└── internal
    ├── core
    │   ├── domain
    │   │   ├── game.go
    │   │   └── board.go
    │   ├── ports
    │   │   ├── repositories.go
    │   │   └── services.go
    │   └── services
    │       └── gamesrv
    │           └── service.go
    ├── handlers
    └── repositories
```

https://www.youtube.com/watch?v=th4AgBcrEHA&list=PLGl1Jc8ErU1w27y8-7Gdcloy1tHO7NriL

https://github.com/pallat/todohexagonal

```shell
go mod init github.com/panutat-p/hexagonal-core-gin

go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/joho/godotenv
```

[hexagonal image](https://miro.medium.com/max/700/1*ERYx0IB1pN-5ZX98cKAoUw.png)
