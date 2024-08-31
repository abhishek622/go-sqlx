## Docker command used

create mysql container in docker -
`docker run --name project1-mysql -p 3307:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql:9.0.1`

create database ecomm -
`docker exec -i project1-mysql mysql -uroot -ppassword<<< "CREATE DATABASE ecomm;"`

create migrations -
`docker run -it --rm  --network host --volume "$(pwd)/db:/db"  migrate/migrate:latest create -ext sql -dir /db/migrations init_schema`

run migration -
`docker run -it --rm --network host --volume ./db:/db migrate/migrate:latest -path=/db/migrations -database "mysql://root:password@tcp(localhost:3307)/ecomm" up`

proto cmd to compile new proto definition

```
protoc --proto_path=ecomm-grpc/pb --go_out=ecomm-grpc/pb --go_opt=paths=source_relative \
    --go-grpc_out=ecomm-grpc/pb --go-grpc_opt=paths=source_relative \
    ecomm-grpc/pb/api.proto
```
