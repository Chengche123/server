docker run --link etcd \ 
-e MICRO_REGISTRY=etcd -e MICRO_REGISTRY_ADDRESS=etcd:2379 \
-p 8082:8082 \ 
-d --rm micro/micro:v2.9.3 web