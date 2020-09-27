# docker-golang-gin-hot_reload-mysql-nginx-starter

## Structure

### api

- golang:ver 1.15
  - gorm (ORM)
  - gin (REST API)
  - CompileDaemon (Hot reload)
- dependency:go mod

### MySQL

- mysql:ver 8.0.21
- sample database:[world](https://dev.mysql.com/doc/index-other.html)

### Nginx

- nginx:ver 1.19
  - port:8888

## Start Development

### First

```shell

# Build Container
docker-compose build

# Login Container Bash
docker exec -it docker_db_1 bash

# if vary yours db container name.
## verify container name.
docker-compose ps
docker exec -it {container_name} bash


# login mysql.
mysql -p
# please type password.
root

# Import sample database.
source /docker-entrypoint-initdb.d/world.sql

# exit mysql
exit

# exit db container
exit

# you can start development life!
docker-compose up -d
```

### after ~

```sh
# you can start development life!
docker-compose up -d
```

### Documemt

You can see, sample rest api response in this url.

```
http://localhost:8888
```

This is response.

```json
{
  data : {
    ID : 1,
    Name : "Kabul",
    CountryCode : "AFG",
    District : "Kabol",
    Population : 1780000,
  }
}
```