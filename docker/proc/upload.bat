SETLOCAL

SET S1=app-view
SET S2=auth
SET s3=comic
SET s4=news

cd D:\final\Projects\server\%S1%\cmd\grpc\ && go build -o main.out . && docker build -t chengche123/service-%S1%-grpc:1.0 .
cd D:\final\Projects\server\%S2%\cmd\grpc\ && go build -o main.out . && docker build -t chengche123/service-%S2%-grpc:1.0 .
cd D:\final\Projects\server\%S3%\cmd\grpc\ && go build -o main.out . && docker build -t chengche123/service-%S3%-grpc:1.0 .
cd D:\final\Projects\server\%S4%\cmd\grpc\ && go build -o main.out . && docker build -t chengche123/service-%S4%-grpc:1.0 .

ENDLOCAL