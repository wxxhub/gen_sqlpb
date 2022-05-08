# gen_sqlpb
根据数据库生成protobuf文件

#### install
```shell
go install github.com/wxxhub/gen_sqlpb@latest
```

###
```shell
gen_sqlpb --srvName Test --dsn "root:123456@tcp(www.wxxhome.com:3306)/test?tableName=test" --dsn "root:123456@tcp(www.wxxhome.com:3306)/test?tableName=user" 
```