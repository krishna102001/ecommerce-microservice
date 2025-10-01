FROM postgres:10.3

COPY up.sql /docker-entrypoiny-initdb.d/1.sql

CMD [ "postgres" ]