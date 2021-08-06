SETLOCAL

@REM SET FILE_PREFIX=%1%

SET S1=app-view
SET S2=auth
SET s3=comic
SET s4=news

docker image rm chengche123/service-%S1%-grpc:1.0 .
docker image rm chengche123/service-%S2%-grpc:1.0 .
docker image rm chengche123/service-%S3%-grpc:1.0 .
docker image rm chengche123/service-%S4%-grpc:1.0 .

docker image rm chengche123/dev-service-%S1%-grpc:1.0 .
docker image rm chengche123/dev-service-%S2%-grpc:1.0 .
docker image rm chengche123/dev-service-%S3%-grpc:1.0 .
docker image rm chengche123/dev-service-%S4%-grpc:1.0 .

ENDLOCAL