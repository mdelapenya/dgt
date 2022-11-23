FROM mysql:5

COPY sql/000_create.sql /docker-entrypoint-initdb.d/000_create.sql
COPY conf.d/* /etc/mysql/conf.d/

CMD ["mysqld", "--character-set-server=utf8mb4", "--collation-server=utf8mb4_unicode_ci"]