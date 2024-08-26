## Docker command used

- `docker run -it --rm  --network host --volume "$(pwd)/db:/db"  migrate/migrate:latest create -ext sql -dir /db/migrations init_schema`

- `docker exec -i project1-mysql mysql -u<USERNAME> -p<PASSWORD><<< "CREATE DATABASE ecomm;"`

- `docker run -it --rm --network host --volume "$(pwd)/db:/db" migrate/migrate:latest -source file://db/migrations -database "mysql://root:password@tcp(host.docker.internal:3307)/ecomm" up`
