syntax = "proto3";

option go_package ="{{.GoPackage}}";

package {{.Package}};

service {{.Srv}} {
    {{range $table := .Tables}}
    //-----------------------{{$table.UpperName}}-----------------------
    rpc Add{{$table.UpperName}}(Add{{$table.UpperName}}Req) returns (Add{{$table.UpperName}}Resp);
    rpc Update{{$table.UpperName}}(Update{{$table.UpperName}}Req) returns (Update{{$table.UpperName}}Resp);
    rpc Del{{$table.UpperName}}(Del{{$table.UpperName}}Req) returns (Del{{$table.UpperName}}Resp);
    rpc Get{{$table.UpperName}}ById(Get{{$table.UpperName}}ByIdReq) returns (Get{{$table.UpperName}}ByIdResp);
    rpc Search{{$table.UpperName}}(Search{{$table.UpperName}}Req) returns (Search{{$table.UpperName}}Resp);
    {{- end}}
}

{{range $table := .Tables}}
//--------------------------------{{$table.UpperName}}--------------------------------
message {{$table.UpperName}} {
  int64 id = 1; //id
  string name = 2; //name
}

message Add{{$table.UpperName}}Req {
  string name = 1; //name
}

message Add{{$table.UpperName}}Resp {
}

message Update{{$table.UpperName}}Req {
  int64 id = 1; //id
  string name = 2; //name
}

message Update{{$table.UpperName}}Resp {
}

message Del{{$table.UpperName}}Req {
  int64 id = 1; //id
}

message Del{{$table.UpperName}}Resp {
}

message Get{{$table.UpperName}}ByIdReq {
  int64 id = 1; //id
}

message Get{{$table.UpperName}}ByIdResp {
  {{$table.UpperName}} {{$table.Name}} = 1; //{{$table.UpperName}}
}

message Search{{$table.UpperName}}Req {
  int64 page = 1;       //page
  int64 pageSize = 2;   //pageSize
  int64 id = 3;         //id
  string name = 4;      //name
}

message Search{{$table.UpperName}}Resp {
  repeated {{$table.UpperName}} {{$table.Name}} = 1; //{{$table.UpperName}}
}
{{end}}
