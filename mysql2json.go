package mysql2json

import (
  "fmt"
  "encoding/json"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func Getjson(username string, password string, address string, database string, sqlstatement string) (string, error) {
  dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", username, password, address, database)
  db, err := sql.Open("mysql", dsn)
  rows, err := db.Query(sqlstatement)
  if err != nil {
      return "", err
  }
  defer rows.Close()
  columns, err := rows.Columns()
  if err != nil {
      return "", err
  }
  count := len(columns)
  tableData := make([]map[string]interface{}, 0)
  values := make([]interface{}, count)
  valuePtrs := make([]interface{}, count)
  for rows.Next() {
      for i := 0; i < count; i++ {
          valuePtrs[i] = &values[i]
      }
      rows.Scan(valuePtrs...)
      entry := make(map[string]interface{})
      for i, col := range columns {
          var v interface{}
          val := values[i]
          b, ok := val.([]byte)
          if ok {
              v = string(b)
          } else {
              v = val
          }
          entry[col] = v
      }
      tableData = append(tableData, entry)
  }

  jsonData, err := json.Marshal(tableData)
  if err != nil {
      return "", err
  }

  return string(jsonData), nil 
}
