FROM mysql:5.7
LABEL maintainer="chengche123"
# RUN

ENV MYSQL_ROOT_PASSWORD root
ENV MYSQL_DATABASE comic
# EXPOSE
# VOLUME

COPY bak-comic-20210723.sql /docker-entrypoint-initdb.d/

# WORKDIR
# CMD
# ENTRYPOINT