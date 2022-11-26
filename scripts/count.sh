#/bin/bash

docker exec -it dgt-db-1 mysql -u root -ppassw0rd --database=dgt -e "select s.description as Distintivo, s.emoji as Label, count(1) as total from plates p, stickers s where s.sticker_id=p.sticker_id group by s.sticker_id;"
