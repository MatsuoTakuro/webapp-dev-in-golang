artifact_name := "webapp-dev-in-golang"

run: && exec
  @go build -o {{artifact_name}} main.go

run-mac-m1: && exec
  @CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o {{artifact_name}} -trimpath -ldflags='-s -w -X main.version=1.0.0' main.go

exec:
  @./{{artifact_name}}

create-db:
  docker run -d --name my-postgres -e POSTGRES_USER=testuser -e POSTGRES_PASSWORD=pass -e POSTGRES_DB=testdb -p 5432:5432 postgres

restart-db:
  docker start my-postgres

stop-db:
  docker stop my-postgres

db-in:
  docker exec -it my-postgres bash -c "psql testdb -U testuser"

create-mysql:
  docker run -d --name my-mysql --platform=linux/x86_64 -e MYSQL_USER=testuser -e MYSQL_PASSWORD=pass -e MYSQL_ROOT_PASSWORD=pass -e MYSQL_DATABASE=testdb -h localhost -p 3306:3306 mysql

mysql-in:
  docker exec -it my-mysql bash -c "mysql -h localhost -u testuser -p"
