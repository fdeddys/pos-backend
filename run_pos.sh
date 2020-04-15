#main.go
export MAXPROCS="1"
export PORT="8001"


#routers/route.go
export RESTO_BE_SERVICE_READ_TIMEOUT="120"
export RESTO_BE_SERVICE_WRITE_TIMEOUT="120"

#database/postgres.go
export RESTO_BE_POSTGRES_USER=pos_user
export RESTO_BE_POSTGRES_PASS=pos_db123
export RESTO_BE_POSTGRES_NAME=dbpos
#export RESTO_BE_POSTGRES_HOST=localhost
export RESTO_BE_POSTGRES_HOST=54.251.137.12
export RESTO_BE_POSTGRES_PORT=5432
export RESTO_BE_POSTGRES_DEBUG=true
export RESTO_BE__TYPE=POSTGRES
export RESTO_BE__POSTGRES_SSL_MODE=disable

#hosts/menustorage/http.go
export MENU_STORAGE_HOST=http://156.67.214.228:9001
export STORAGE_MINIO_URLACCESS=http://156.67.214.228:9000


export PATH_REPORT="/opt/report/"

go run main.go