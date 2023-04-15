# How To Run API

* Run api and database with docker compose ( makesure port 3306 available )
```
docker compose up -d
```

* Run migration with cmd
```
migrate -database 'mysql://root:root@tcp(localhost:3306)/registration?parseTime=true' -path ./db/migrations up
```


* Open Swagger with [URL](http://localhost:8880/swagger/index.html)
