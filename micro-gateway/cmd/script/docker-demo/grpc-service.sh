docker run --link etcd --link mysql  \ 
-e COMIC_MYSQL_ADDR=mysql:3306 -e COMIC_REGISTRY_ADDR=etcd:2379 \ 
-d chengche123/comic-auth-micro-grpc-service:1.0