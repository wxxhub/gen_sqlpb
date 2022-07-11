syntax = "proto3";

option go_package ="{{.GoPackage}}";

package {{.Package}};

service {{.Srv}} {
    //-----------------------{.TableInfo.CamelName}}-----------------------
    rpc Add{{.TableInfo.CamelName}}(Add{{.TableInfo.CamelName}}Req) returns (Add{{.TableInfo.CamelName}}Res);
    rpc Update{{.TableInfo.CamelName}}(Update{{.TableInfo.CamelName}}Req) returns (Update{{.TableInfo.CamelName}}Res);
    rpc Del{{.TableInfo.CamelName}}(Del{{.TableInfo.CamelName}}Req) returns (Del{{.TableInfo.CamelName}}Res);
    rpc Get{{.TableInfo.CamelName}}By{{.ProtoContent.PrimaryIndexItem.GenItem.CamelName}}(Get{{.TableInfo.CamelName}}By{{.ProtoContent.PrimaryIndexItem.GenItem.CamelName}}Req) returns (Get{{.TableInfo.CamelName}}By{{.ProtoContent.PrimaryIndexItem.GenItem.CamelName}}Res);
    rpc Mget{{.TableInfo.CamelName}}By{{.ProtoContent.PrimaryIndexItem.GenItem.CamelName}}(Mget{{.TableInfo.CamelName}}By{{.ProtoContent.PrimaryIndexItem.GenItem.CamelName}}Req) returns (Mget{{.TableInfo.CamelName}}By{{.ProtoContent.PrimaryIndexItem.GenItem.CamelName}}Res);
    rpc Search{{.TableInfo.CamelName}}(Search{{.TableInfo.CamelName}}Req) returns (Search{{.TableInfo.CamelName}}Res);
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

//--------------------------------{{.TableInfo.CamelName}}--------------------------------
message {{.TableInfo.CamelName}} {
{{- range $item := .ProtoContent.ProtoItems}}
  {{$item.Type}} {{$item.Name}} = {{$item.Index}};
{{- end}}
}

message Add{{.TableInfo.CamelName}}Req {
  {{.TableInfo.CamelName}} {{.TableInfo.Name}} = 1;
  ExtraOpt extra_opt = 2;
}

message Add{{.TableInfo.CamelName}}Res {
  {{.TableInfo.CamelName}} {{.TableInfo.Name}} = 1;
}

message Update{{.TableInfo.CamelName}}Req {
  {{.TableInfo.CamelName}} {{.TableInfo.Name}} = 1;
  ExtraOpt extra_opt = 2;
}

message Update{{.TableInfo.CamelName}}Res {
  {{.TableInfo.CamelName}} {{.TableInfo.Name}} = 1;
}

message Del{{.TableInfo.CamelName}}Req {
  {{.ProtoContent.PrimaryIndexItem.GenItem.Type}} {{.ProtoContent.PrimaryIndexItem.GenItem.Name}} = 1; // {{.ProtoContent.PrimaryIndexItem.GenItem.Name}}
  ExtraOpt extra_opt = 2;
}

message Del{{.TableInfo.CamelName}}Res {
}

message Get{{.TableInfo.CamelName}}By{{.ProtoContent.PrimaryIndexItem.GenItem.CamelName}}Req {
  {{.ProtoContent.PrimaryIndexItem.GenItem.Type}} {{.ProtoContent.PrimaryIndexItem.GenItem.Name}} = 1; // {{.ProtoContent.PrimaryIndexItem.GenItem.Name}}
  ExtraOpt extra_opt = 2;
}

message Get{{.TableInfo.CamelName}}By{{.ProtoContent.PrimaryIndexItem.GenItem.CamelName}}Res {
  {{.TableInfo.CamelName}} {{.TableInfo.Name}} = 1; //{{.TableInfo.CamelName}}
  ExtraOpt extra_opt = 2;
}

message Mget{{.TableInfo.CamelName}}By{{.ProtoContent.PrimaryIndexItem.GenItem.CamelName}}Req {
  repeated {{.ProtoContent.PrimaryIndexItem.GenItem.Type}} {{.ProtoContent.PrimaryIndexItem.GenItem.Name}} = 1;
  ExtraOpt extra_opt = 2;
}

message Mget{{.TableInfo.CamelName}}By{{.ProtoContent.PrimaryIndexItem.GenItem.CamelName}}Res {
  repeated {{.TableInfo.CamelName}} {{.TableInfo.Name}} = 1;
}


message Search{{.TableInfo.CamelName}}Req {
  int64 page = 1;       //page
  int64 pageSize = 2;   //pageSize
  int64 id = 3;         //id
  string name = 4;      //name
}

message Search{{.TableInfo.CamelName}}Res {
  repeated {{.TableInfo.CamelName}} {{.TableInfo.Name}} = 1; //{{.TableInfo.CamelName}}
}
