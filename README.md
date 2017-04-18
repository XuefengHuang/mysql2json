# mysql2json
Dumping MySQL tables to JSON.

## Installation
- Install mysql2json dependency (if not already installed):
- `go get github.com/go-sql-driver/mysql`

## Usage
```
import "github.com/XuefengHuang/mysql2json"

mysql2json.Getjson(username, password, address, database, sql_statement)
```

## Example
`mysql2json.Getjson("root", "123456", "127.0.0.1", "test", "select * from test")`
