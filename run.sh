#main.go
export MAXPROCS="1"
export PORT="8001"


#routers/route.go
export RESTO_BE_SERVICE_READ_TIMEOUT="120"
export RESTO_BE_SERVICE_WRITE_TIMEOUT="120"

#database/postgres.go
export RESTO_BE_POSTGRES_USER=resto
export RESTO_BE_POSTGRES_PASS=Resto#123
export RESTO_BE_POSTGRES_NAME=restodb
#export RESTO_BE_POSTGRES_HOST=localhost
export RESTO_BE_POSTGRES_HOST=156.67.214.228
export RESTO_BE_POSTGRES_PORT=5432
export RESTO_BE_POSTGRES_DEBUG=true
export RESTO_BE__TYPE=POSTGRES
export RESTO_BE__POSTGRES_SSL_MODE=disable

#hosts/menustorage/http.go
export STORAGE_MINIO_URLACCESS=http://156.67.214.228:9001

go run main.go