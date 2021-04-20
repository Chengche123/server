docker run --link etcd \
-e MICRO_REGISTRY=etcd -e MICRO_REGISTRY_ADDRESS=etcd:2379 \
-p 8080:8080 \
-d micro/micro:v2.9.3 api --handler=api 