syntax = "proto3";

option go_package ="./pb";

package pb;

// ------------------------------------
// Messages
// ------------------------------------

{{range $table := .Tables}}
//--------------------------------{{$table.Name}}--------------------------------
message {{$table.Name}} {
  int64 id = 1; //id
  string name = 2; //name
}

message Add{{$table.Name}}Req {
  string name = 1; //name
}

message Add{{$table.Name}}Resp {
}

message Update{{$table.Name}}Req {
  int64 id = 1; //id
  string name = 2; //name
}

message Update{{$table.Name}}Resp {
}

message Del{{$table.Name}}Req {
  int64 id = 1; //id
}

message Del{{$table.Name}}Resp {
}

message Get{{$table.Name}}ByIdReq {
  int64 id = 1; //id
}

message Get{{$table.Name}}ByIdResp {
  {{$table.Name}} {{$table.Name}} = 1; //{{$table.Name}}
}

message Search{{$table.Name}}Req {
  int64 page = 1;       //page
  int64 pageSize = 2;   //pageSize
  int64 id = 3;         //id
  string name = 4;      //name
}

message Search{{$table.Name}}Resp {
  repeated {{$table.Name}} {{$table.Name}} = 1; //{{$table.Name}}
}
{{end}}


service {{.Srv}}{

    {{range $table := .Tables}}
    //-----------------------{{$table.Name}}-----------------------
    rpc Add{{$table.Name}}(Add{{$table.Name}}Req) returns (Add{{$table.Name}}Resp);
    rpc Update{{$table.Name}}(Update{{$table.Name}}Req) returns (Update{{$table.Name}}Resp);
    rpc Del{{$table.Name}}(Del{{$table.Name}}Req) returns (Del{{$table.Name}}Resp);
    rpc Get{{$table.Name}}ById(Get{{$table.Name}}ByIdReq) returns (Get{{$table.Name}}ByIdResp);
    rpc Search{{$table.Name}}(Search{{$table.Name}}Req) returns (Search{{$table.Name}}Resp);
    {{end}}
}
