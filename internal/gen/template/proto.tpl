syntax = "proto3";

option go_package ="../{{.GoPackage}}";

package {{.Package}};

service {{.Srv}} {
    //-----------------------{{.TableInfo.Name|StringCamel}}-----------------------
    rpc Add{{.TableInfo.Name|StringCamel}}(Add{{.TableInfo.Name|StringCamel}}Req) returns (Add{{.TableInfo.Name|StringCamel}}Res);
    rpc Update{{.TableInfo.Name|StringCamel}}(Update{{.TableInfo.Name|StringCamel}}Req) returns (Update{{.TableInfo.Name|StringCamel}}Res);

    {{StringJoin "Get" (.TableInfo.Name|StringCamel) "By" (.ProtoContent.PrimaryIndexItem.Fields|ToCamelJoinAnd)| RpcLine}};
    {{StringJoin "Update" (.TableInfo.Name|StringCamel) "By" (.ProtoContent.PrimaryIndexItem.Fields|ToCamelJoinAnd)| RpcLine}};
    {{StringJoin "Del" (.TableInfo.Name|StringCamel) "By" (.ProtoContent.PrimaryIndexItem.Fields|ToCamelJoinAnd)| RpcLine}};
    {{StringJoin "Mget" (.TableInfo.Name|StringCamel) "By" (.ProtoContent.PrimaryIndexItem.Fields|ToCamelJoinAnd)| RpcLine}};

    rpc Search{{.TableInfo.Name|StringCamel}}(Search{{.TableInfo.Name|StringCamel}}Req) returns (Search{{.TableInfo.Name|StringCamel}}Res);
}

message ExtraOpt {
  bool use_cache = 1;
  bool use_local_cache = 2;
  bool del_cache = 3;
  bool refresh_cache = 4;
  bool only_cache = 5;
  uint32 expire = 6;
}

message Where {

}

message Condition {
  uint32 offset = 1;
  uint32 limit = 2;
  string order_by = 3;
  string group_by = 4;
  Where having = 5;
}

//--------------------------------{{.TableInfo.Name|StringCamel}}--------------------------------
message {{.TableInfo.Name|StringCamel}} {
{{- range $index, $item := .ProtoContent.ProtoItems}}
  {{$item.Type}} {{$item.Name}} = {{$index|ItemIndex}}; {{$item.Comment|AddNote}}
{{- end}}
}

message Add{{.TableInfo.Name|StringCamel}}Req {
  {{.TableInfo.Name|StringCamel}} {{.TableInfo.Name}} = 1;
  ExtraOpt extra_opt = 2;
}

message Add{{.TableInfo.Name|StringCamel}}Res {
  {{.TableInfo.Name|StringCamel}} {{.TableInfo.Name}} = 1;
}

message Update{{.TableInfo.Name|StringCamel}}Req {
  {{.TableInfo.Name|StringCamel}} {{.TableInfo.Name}} = 1;
  ExtraOpt extra_opt = 2;
}

message Update{{.TableInfo.Name|StringCamel}}Res {
  {{.TableInfo.Name|StringCamel}} {{.TableInfo.Name}} = 1;
}

{{/*primary index*/}}
message Key {
{{- range $index, $item := .ProtoContent.PrimaryIndexItem.IndexItems}}
  {{$item.Type}} {{$item.Name}} = {{$index|ItemIndex}}; {{$item.Comment|AddNote}}
{{- end}}
}

message {{StringJoin "Get" (.TableInfo.Name|StringCamel) "By" (.ProtoContent.PrimaryIndexItem.Fields|ToCamelJoinAnd) "Req"}} {
  Key key = 1;
  ExtraOpt extra_opt = 2;
}

message {{StringJoin "Get" (.TableInfo.Name|StringCamel) "By" (.ProtoContent.PrimaryIndexItem.Fields|ToCamelJoinAnd) "Res"}} {
  {{.TableInfo.Name|StringCamel}} {{.TableInfo.Name}} = 1; //{{.TableInfo.Name|StringCamel}}
  ExtraOpt extra_opt = 2;
}

message {{StringJoin "Update" (.TableInfo.Name|StringCamel) "By" (.ProtoContent.PrimaryIndexItem.Fields|ToCamelJoinAnd) "Req"}} {
  Key key = 1;
  {{.TableInfo.Name|StringCamel}} {{.TableInfo.Name}} = 2;
  ExtraOpt extra_opt = 3;
}

message {{StringJoin "Update" (.TableInfo.Name|StringCamel) "By" (.ProtoContent.PrimaryIndexItem.Fields|ToCamelJoinAnd) "Res"}} {
  {{.TableInfo.Name|StringCamel}} {{.TableInfo.Name}} = 1; //{{.TableInfo.Name|StringCamel}}
  ExtraOpt extra_opt = 2;
}

message {{StringJoin "Del" (.TableInfo.Name|StringCamel) "By" (.ProtoContent.PrimaryIndexItem.Fields|ToCamelJoinAnd) "Req"}} {
  Key key = 1;
  ExtraOpt extra_opt = 2;
}

message {{StringJoin "Del" (.TableInfo.Name|StringCamel) "By" (.ProtoContent.PrimaryIndexItem.Fields|ToCamelJoinAnd) "Res"}} {
  ExtraOpt extra_opt = 2;
}

message {{StringJoin "Mget" (.TableInfo.Name|StringCamel) "By" (.ProtoContent.PrimaryIndexItem.Fields|ToCamelJoinAnd) "Req"}} {
  repeated Key key = 1;
  ExtraOpt extra_opt = 2;
}

message {{StringJoin "Mget" (.TableInfo.Name|StringCamel) "By" (.ProtoContent.PrimaryIndexItem.Fields|ToCamelJoinAnd) "Res"}} {
repeated {{.TableInfo.Name|StringCamel}} {{.TableInfo.Name}} = 1;
}

message Search{{.TableInfo.Name|StringCamel}}Req {
  int64 page = 1;       //page
  int64 pageSize = 2;   //pageSize
  int64 id = 3;         //id
  string name = 4;      //name
}

message Search{{.TableInfo.Name|StringCamel}}Res {
  repeated {{.TableInfo.Name|StringCamel}} {{.TableInfo.Name}} = 1; //{{.TableInfo.Name|StringCamel}}
}

