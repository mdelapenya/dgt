#/bin/bash

# stop processors
docker stop dgt-go-j-1 dgt-go-l-1 dgt-go-f-1 dgt-go-h-1 dgt-go-c-1 dgt-go-k-1 dgt-go-m-1 dgt-go-b-1 dgt-go-g-1 dgt-go-d-1

echo 'PLATE' > ./plates.txt
echo '=====' >> ./plates.txt
docker exec -it dgt-db-1 mysql -u root -ppassw0rd --database=dgt -e "select plate from plates where plate_id=(select plate_id from plates where plate like '____B__' order by plate_id desc limit 1);" >> ./plates.txt
docker exec -it dgt-db-1 mysql -u root -ppassw0rd --database=dgt -e "select plate from plates where plate_id=(select plate_id from plates where plate like '____C__' order by plate_id desc limit 1);" >> ./plates.txt
docker exec -it dgt-db-1 mysql -u root -ppassw0rd --database=dgt -e "select plate from plates where plate_id=(select plate_id from plates where plate like '____D__' order by plate_id desc limit 1);" >> ./plates.txt
docker exec -it dgt-db-1 mysql -u root -ppassw0rd --database=dgt -e "select plate from plates where plate_id=(select plate_id from plates where plate like '____F__' order by plate_id desc limit 1);" >> ./plates.txt
docker exec -it dgt-db-1 mysql -u root -ppassw0rd --database=dgt -e "select plate from plates where plate_id=(select plate_id from plates where plate like '____G__' order by plate_id desc limit 1);" >> ./plates.txt
docker exec -it dgt-db-1 mysql -u root -ppassw0rd --database=dgt -e "select plate from plates where plate_id=(select plate_id from plates where plate like '____H__' order by plate_id desc limit 1);" >> ./plates.txt
docker exec -it dgt-db-1 mysql -u root -ppassw0rd --database=dgt -e "select plate from plates where plate_id=(select plate_id from plates where plate like '____J__' order by plate_id desc limit 1);" >> ./plates.txt
docker exec -it dgt-db-1 mysql -u root -ppassw0rd --database=dgt -e "select plate from plates where plate_id=(select plate_id from plates where plate like '____K__' order by plate_id desc limit 1);" >> ./plates.txt
docker exec -it dgt-db-1 mysql -u root -ppassw0rd --database=dgt -e "select plate from plates where plate_id=(select plate_id from plates where plate like '____L__' order by plate_id desc limit 1);" >> ./plates.txt
docker exec -it dgt-db-1 mysql -u root -ppassw0rd --database=dgt -e "select plate from plates where plate_id=(select plate_id from plates where plate like '____M__' order by plate_id desc limit 1);" >> ./plates.txt

cat ./plates.txt | grep -v "mysql" | grep -v '-' | grep -v 'plate'