syntax = "proto3";

option go_package ="{{.GoPackage}}";

package {{.Package}};

service {{.Srv}} {
    //-----------------------{.TableInfo.CamelName}}-----------------------
    rpc Add{{.TableInfo.CamelName}}(Add{{.TableInfo.CamelName}}Req) returns (Add{{.TableInfo.CamelName}}Resp);
    rpc Update{{.TableInfo.CamelName}}(Update{{.TableInfo.CamelName}}Req) returns (Update{{.TableInfo.CamelName}}Resp);
    rpc Del{{.TableInfo.CamelName}}(Del{{.TableInfo.CamelName}}Req) returns (Del{{.TableInfo.CamelName}}Resp);
    rpc Get{{.TableInfo.CamelName}}By{{.ProtoContent.PrimaryIndexItem.GenItem.CamelName}}(Get{{.TableInfo.CamelName}}By{{.ProtoContent.PrimaryIndexItem.GenItem.CamelName}}Req) returns (Get{{.TableInfo.CamelName}}By{{.ProtoContent.PrimaryIndexItem.GenItem.CamelName}}Resp);
    rpc Mget{{.TableInfo.CamelName}}By{{.ProtoContent.PrimaryIndexItem.GenItem.CamelName}}(Mget{{.TableInfo.CamelName}}By{{.ProtoContent.PrimaryIndexItem.GenItem.CamelName}}Req) returns (Mget{{.TableInfo.CamelName}}By{{.ProtoContent.PrimaryIndexItem.GenItem.CamelName}}Res);
    rpc Search{{.TableInfo.CamelName}}(Search{{.TableInfo.CamelName}}Req) returns (Search{{.TableInfo.CamelName}}Resp);
}

message Extra {

}

//--------------------------------{{.TableInfo.CamelName}}--------------------------------
message {{.TableInfo.CamelName}} {
{{- range $item := .ProtoContent.ProtoItems}}
  {{$item.Type}} {{$item.Name}} = {{$item.Index}};
{{- end}}
}

message Add{{.TableInfo.CamelName}}Req {
  {{.TableInfo.CamelName}} {{.TableInfo.Name}} = 1;
}

message Add{{.TableInfo.CamelName}}Resp {
  {{.TableInfo.CamelName}} {{.TableInfo.Name}} = 1;
}

message Update{{.TableInfo.CamelName}}Req {
  {{.TableInfo.CamelName}} {{.TableInfo.Name}} = 1;
}

message Update{{.TableInfo.CamelName}}Resp {
  {{.TableInfo.CamelName}} {{.TableInfo.Name}} = 1;
}

message Del{{.TableInfo.CamelName}}Req {
  {{.ProtoContent.PrimaryIndexItem.GenItem.Type}} {{.ProtoContent.PrimaryIndexItem.GenItem.Name}} = 1; // {{.ProtoContent.PrimaryIndexItem.GenItem.Name}}
}

message Del{{.TableInfo.CamelName}}Resp {
}

message Get{{.TableInfo.CamelName}}By{{.ProtoContent.PrimaryIndexItem.GenItem.CamelName}}Req {
  {{.ProtoContent.PrimaryIndexItem.GenItem.Type}} {{.ProtoContent.PrimaryIndexItem.GenItem.Name}} = 1; // {{.ProtoContent.PrimaryIndexItem.GenItem.Name}}
}

message Get{{.TableInfo.CamelName}}By{{.ProtoContent.PrimaryIndexItem.GenItem.CamelName}}Resp {
  {{.TableInfo.CamelName}} {{.TableInfo.Name}} = 1; //{{.TableInfo.CamelName}}
}

message Mget{{.TableInfo.CamelName}}By{{.ProtoContent.PrimaryIndexItem.GenItem.CamelName}}Req {
  repeated {{.ProtoContent.PrimaryIndexItem.GenItem.Type}} {{.ProtoContent.PrimaryIndexItem.GenItem.Name}} = 1;
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

message Search{{.TableInfo.CamelName}}Resp {
  repeated {{.TableInfo.CamelName}} {{.TableInfo.Name}} = 1; //{{.TableInfo.CamelName}}
}
