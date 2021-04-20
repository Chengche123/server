SETLOCAL

SET DOMAIN=%1%

set PBTS_BIN_DIR=E:\final\Web\client\miniprogram\node_modules\.bin
%PBTS_BIN_DIR%\pbts -o %DOMAIN%_pb.d.ts %DOMAIN%_pb.js

ENDLOCAL