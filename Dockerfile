FROM postgres:10.3

COPY database/up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["postgres"]