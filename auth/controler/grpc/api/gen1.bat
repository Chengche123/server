SETLOCAL

SET DOMAIN=%1%

protoc -I=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative %DOMAIN%.proto
protoc -I=. --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative,grpc_api_configuration=%DOMAIN%.yaml %DOMAIN%.proto

set PBTS_BIN_DIR=E:\final\Web\client\miniprogram\node_modules\.bin

%PBTS_BIN_DIR%\pbjs -t static -w es6 %DOMAIN%.proto --no-create --no-encode --no-decode --no-verify --no-delimited -o %DOMAIN%_pb.js

%PBTS_BIN_DIR%\pbts -o %DOMAIN%_pb.d.ts %DOMAIN%_pb.js

ENDLOCAL

