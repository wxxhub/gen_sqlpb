syntax = "proto3";

option go_package ="./pb";

package pb;

// ------------------------------------
// Messages
// ------------------------------------

{{range $table := .Tables}}
//--------------------------------{{$table}}--------------------------------
message {{$table}} {
  int64 id = 1; //id
  string name = 2; //name
}

message Add{{$table}}Req {
  string name = 1; //name
}

message Add{{$table}}Resp {
}

message Update{{$table}}Req {
  int64 id = 1; //id
  string name = 2; //name
}

message Update{{$table}}Resp {
}

message Del{{$table}}Req {
  int64 id = 1; //id
}

message Del{{$table}}Resp {
}

message Get{{$table}}ByIdReq {
  int64 id = 1; //id
}

message Get{{$table}}ByIdResp {
  {{$table}} {{$table}} = 1; //{{$table}}
}

message Search{{$table}}Req {
  int64 page = 1;       //page
  int64 pageSize = 2;   //pageSize
  int64 id = 3;         //id
  string name = 4;      //name
}

message Search{{$table}}Resp {
  repeated {{$table}} {{$table}} = 1; //{{$table}}
}
{{end}}


service {{.Srv}}{

    {{range $table := .Tables}}
    //-----------------------{{$table}}-----------------------
    rpc Add{{$table}}(Add{{$table}}Req) returns (Add{{$table}}Resp);
    rpc Update{{$table}}(Update{{$table}}Req) returns (Update{{$table}}Resp);
    rpc Del{{$table}}(Del{{$table}}Req) returns (Del{{$table}}Resp);
    rpc Get{{$table}}ById(Get{{$table}}ByIdReq) returns (Get{{$table}}ByIdResp);
    rpc Search{{$table}}(Search{{$table}}Req) returns (Search{{$table}}Resp);
    {{end}}
}