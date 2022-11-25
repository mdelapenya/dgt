#/bin/bash

docker exec -it dgt-db-1 mysqldump -u root -ppassw0rd dgt > ./database/sql/001_dgt.sql 
