# gen_sqlpb
根据数据库生成protobuf文件。  
generate protobuf from mysql.  

## install
```shell
go install github.com/wxxhub/gen_sqlpb@latest
```

## params
| param       | description           | must |
|:------------|:----------------------|------|
| --dsn       | Database source name  | *    |
| --debug     | Print debug info      |      |
| --savePath  | The path to save      |      |
| --package   | protobuf package name |      |
| --goPackage | golang package name   |      |


## use
```shell
gen_sqlpb --dsn "username:password@tcp(ip:port)/database?tableName=tablename" 
```
