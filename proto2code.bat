set cur_path=%cd%

set client_in_path=%cur_path%\clientproto
set server_in_path=%cur_path%\serverproto


set go_out_path=%cur_path%\cache\go\pb
set cplusplus_out_path=%cur_path%\cache\cplusplus
set csharp_out_paht=%cur_path%\cache\csharp

::go has import protobuf error must manual write
if not %go_out_path% == "" (
    ::protoc.exe  --plugin=protoc-gen-go=.\protoc-gen-go.exe  --proto_path %client_in_path%  --go_out %go_out_path%   %client_in_path%\*.proto
    ::protoc.exe  --plugin=protoc-gen-go=.\protoc-gen-go.exe  --proto_path %server_in_path%  --go_out %go_out_path%   %server_in_path%\*.proto
    protoc.exe  --plugin=protoc-gen-go=.\protoc-gen-go.exe  --proto_path %client_in_path% --proto_path %server_in_path%  --go_out %go_out_path%  %client_in_path%\*.proto  %server_in_path%\*.proto
)

if not %go_out_path% == "" (
    protoc.exe  --proto_path %client_in_path% --proto_path %server_in_path%  --cpp_out  %cplusplus_out_path%  %client_in_path%\*.proto  %server_in_path%\*.proto
)

if not %go_out_path% == "" (
    protoc.exe  --proto_path %client_in_path% --proto_path %server_in_path% --csharp_out  %csharp_out_paht%  %client_in_path%\*.proto  %server_in_path%\*.proto
)
pause