FROM mysql:8.0.24
LABEL maintainer="chengche123"
ENV MYSQL_ROOT_PASSWORD root

COPY user_account.sql /docker-entrypoint-initdb.d/
