package utils

import (
	"database/sql"
	"strconv"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

var (
	Username = "root"
	Password = "111111"
	Host     = "localhost"
	DbName   = "lu_ban"
)

func ExecuteSql(s string, args ...interface{}) int64 {
	db, err := sql.Open("mysql", Username+":"+Password+"@tcp("+Host+")/"+DbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtIns, err := db.Prepare(s)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	res, err := stmtIns.Exec(args...)
	if err != nil {
		panic(err.Error())
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	return affectedRows
}

func QuerySql(s string, args ...interface{}) []map[string]interface{} {
	// Open database connection
	db, err := sql.Open("mysql", Username+":"+Password+"@tcp("+Host+")/"+DbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(s, args...)
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	result := make([]map[string]interface{}, 0)

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		m := make(map[string]interface{})

		columnTypes, err := rows.ColumnTypes()
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {
			t := columnTypes[i]
			if col == nil {
				m[columns[i]] = nil
				continue
			} else {
				value = string(col)
			}
			switch t.DatabaseTypeName() {
			case "BIT":
				m[columns[i]], _ = strconv.Atoi(value)
			case "TINYINT":
				m[columns[i]], _ = strconv.Atoi(value)
			case "SMALLINT":
				m[columns[i]], _ = strconv.Atoi(value)
			case "MEDIUMINT":
				m[columns[i]], _ = strconv.Atoi(value)
			case "INT":
				m[columns[i]], _ = strconv.Atoi(value)
			case "BIGINT":
				m[columns[i]], _ = strconv.ParseInt(value, 10, 64)
			case "CHAR":
				m[columns[i]] = value
			case "VARCHAR":
				m[columns[i]] = value
			case "TINYTEXT":
				m[columns[i]] = value
			case "TEXT":
				m[columns[i]] = value
			case "MEDIUMTEXT":
				m[columns[i]] = value
			case "LONGTEXT":
				m[columns[i]] = value
			case "TINYBLOB":
				m[columns[i]] = value
			case "BLOB":
				m[columns[i]] = value
			case "MEDIUMBLOB":
				m[columns[i]] = value
			case "LONGBLOB":
				m[columns[i]] = value
			case "DATE":
				m[columns[i]], _ = time.Parse("2006-01-02", value)
			case "TIME":
				m[columns[i]], _ = time.Parse("2006-01-02 15:04:05", value)
			case "DATETIME":
				m[columns[i]], _ = time.Parse("2006-01-02 15:04:05", value)
			case "TIMESTAMP":
				m[columns[i]], _ = time.Parse("2006-01-02 15:04:05", value)
			case "DECIMAL":
				m[columns[i]], _ = strconv.ParseFloat(value, 32)
			case "FLOAT":
				m[columns[i]], _ = strconv.ParseFloat(value, 64)
			case "DOUBLE":
				m[columns[i]], _ = strconv.ParseFloat(value, 64)
			}
		}
		result = append(result, m)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return result
}

func QuerySingle(s string, args ...interface{}) map[string]interface{} {
	list := QuerySql(s, args...)
	if len(list) == 1 {
		return list[0]
	}
	return nil
}
