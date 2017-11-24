FROM postgres:latest

ENV POSTGRES_DB opus47
COPY opus47.sql /docker-entrypoint-initdb.d/opus47.sql
COPY seed.sql /docker-entrypoint-initdb.d/seed.sql

