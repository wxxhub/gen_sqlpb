syntax = "proto3";

option go_package ="{{.GoPackage}}";

package {{.Package}};

service {{.Srv}} {
    //-----------------------{.TableInfo.UpperName}}-----------------------
    rpc Add{{.TableInfo.UpperName}}(Add{{.TableInfo.UpperName}}Req) returns (Add{{.TableInfo.UpperName}}Resp);
    rpc Update{{.TableInfo.UpperName}}(Update{{.TableInfo.UpperName}}Req) returns (Update{{.TableInfo.UpperName}}Resp);
    rpc Del{{.TableInfo.UpperName}}(Del{{.TableInfo.UpperName}}Req) returns (Del{{.TableInfo.UpperName}}Resp);
    rpc Get{{.TableInfo.UpperName}}ById(Get{{.TableInfo.UpperName}}ByIdReq) returns (Get{{.TableInfo.UpperName}}ByIdResp);
    rpc Search{{.TableInfo.UpperName}}(Search{{.TableInfo.UpperName}}Req) returns (Search{{.TableInfo.UpperName}}Resp);
}

//--------------------------------{{.TableInfo.UpperName}}--------------------------------
message {{.TableInfo.UpperName}} {
{{range $item := .ProtoItems}}
  {{$item.Type}} {{$item.Name}} = {{$item.Index}};
{{- end}}
}

message Add{{.TableInfo.UpperName}}Req {
{{.TableInfo.UpperName}} {{.TableInfo.Name}} = 1;
}

message Add{{.TableInfo.UpperName}}Resp {
{{.TableInfo.UpperName}} {{.TableInfo.Name}} = 1;
}

message Update{{.TableInfo.UpperName}}Req {
  {{.TableInfo.UpperName}} {{.TableInfo.Name}} = 1;
}

message Update{{.TableInfo.UpperName}}Resp {
}

message Del{{.TableInfo.UpperName}}Req {
  int64 id = 1; //id
}

message Del{{.TableInfo.UpperName}}Resp {
}

message Get{{.TableInfo.UpperName}}ByIdReq {
  int64 id = 1; //id
}

message Get{{.TableInfo.UpperName}}ByIdResp {
  {{.TableInfo.UpperName}} {{.TableInfo.Name}} = 1; //{{.TableInfo.UpperName}}
}

message Search{{.TableInfo.UpperName}}Req {
  int64 page = 1;       //page
  int64 pageSize = 2;   //pageSize
  int64 id = 3;         //id
  string name = 4;      //name
}

message Search{{.TableInfo.UpperName}}Resp {
  repeated {{.TableInfo.UpperName}} {{.TableInfo.Name}} = 1; //{{.TableInfo.UpperName}}
}
