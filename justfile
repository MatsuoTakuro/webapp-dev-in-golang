create-db:
  docker run -d --name my-postgres -e POSTGRES_USER=testuser -e POSTGRES_PASSWORD=pass -e POSTGRES_DB=testdb -p 5432:5432 postgres

restart-db:
  docker start my-postgres

stop-db:
  docker stop my-postgres

db-in:
  docker exec -it my-postgres bash -c "psql testdb -U testuser"
