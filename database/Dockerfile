FROM mysql:5

COPY sql/000_create.sql /docker-entrypoint-initdb.d/000_create.sql
COPY conf.d/* /etc/mysql/conf.d/