# gen_sqlpb
根据数据库生成protobuf文件。  
generate protobuf from mysql.  

## install
```shell
go install github.com/wxxhub/gen_sqlpb@latest
```

## params
| param      | description            |
|:-----------|:-----------------------|
| --srvName  | Service name           |
| --dsn      | Database source name   |
| --debug    | Print debug info       |
| --SavePath | The path to save proto |


## use
```shell
gen_sqlpb --srvName ServiceName --dsn "[username]:[password]@tcp([ip:port])/[database]?tableName=[tablename]" 
```