package helper

import (
	"strings"
	"time"

	"management/internal/utilities/slices"
)

func GetDateTimeSQL(datetime string)string{

	dt,_:= time.Parse(time.RFC3339,datetime);
	mysqlFormat := dt.Format("2006-01-02 15:04:05")
	return mysqlFormat

}


func GetInsertStatement(table string, params map[string]interface{}, fields []string) (string, []interface{}) {
	strField := strings.Join(fields, ",")
	bindStr, insertValue := MakeBindStatement(fields, params, false)
	str := "INSERT INTO " + table + "(" + strField + ") VALUES (" + bindStr + ")"
	return str, insertValue
}


func GetEditStatement(table string, params map[string]interface{}, ignoredKey []string, condKey []string) (string, []interface{}) {
	bindStr, updateValue, condStr := MakeEditStatement(params, ignoredKey, condKey)
	str := "UPDATE " + table + " SET " + bindStr + " WHERE " + condStr
	return str, updateValue
}

func MakeBindStatement(fields []string, params map[string]interface{}, ignoreEmpty bool) (string, []interface{}) {
	a := make([]string, len(fields))
	var insertValue []interface{}
	for k, v := range fields {
		if ignoreEmpty == true {
			if v != "" {
				a[k] = "?"
				insertValue = append(insertValue, params[v])
			}
		} else {
			a[k] = "?"
			insertValue = append(insertValue, params[v])
		}
	}
	return strings.Join(a, ","), insertValue
}

func MakeCondStatement(fields []string, params map[string]interface{}) (string, []interface{}) {
	str := ""
	var insertValue []interface{}
	for _, v := range fields {
		name := strings.Split(v, ".")
		field := name[0]
		if len(name) > 1 {
			field = name[1]
		}
		if params[field] != "" {
			if str != "" {
				str += " AND " + field + " = ?"
			} else {
				str += field + " = ?"
			}
			insertValue = append(insertValue, params[field])
		}
	}
	return str, insertValue
}

func MakeEditStatement(params map[string]interface{}, ignoredKey []string, condKey []string) (string, []interface{}, string) {
	arr := []string{}
	condStr := ""

	var updateValue []interface{}
	var condValue []interface{}

	for k, v := range params {
		if !slices.ContainString(ignoredKey, k) {
			arr = append(arr, k+"=?")
			updateValue = append(updateValue, v)
		}

		if slices.ContainString(condKey, k) {
			condValue = append(condValue, v)
			if condStr != "" {
				condStr += " AND " + k + " = ?"
			} else {
				condStr += k + " = ?"
			}
		}
	}

	if len(condValue) > 0 {
		updateValue = append(updateValue, condValue...)
	}

	return strings.Join(arr, ","), updateValue, condStr
}


