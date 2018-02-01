package model

import (
	"buffers_test/DAL/core"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var caches = map[string]core.Cache{}

func InitCache(model core.CoreModel, cache core.Cache) error {
	_, ok := caches[reflect.TypeOf(model).String()]
	if ok {
		return nil
	}

	caches[reflect.TypeOf(model).String()] = cache
	err := model.UpdateCache()
	return err
}

func GenerateInsertQuery(m core.CoreModel) string {
	fields := "("
	vals := "("

	t := reflect.TypeOf(m)
	v := getValue(m)

	for i := 0; i < v.Elem().NumField(); i++ {
		field := t.Field(i).Tag.Get("column")
		if field == "" {
			continue
		}
		fields += field + ","
		vals += getStringValue(t.Field(i), reflect.ValueOf(m).Field(i)) + ","
	}
	fields = strings.TrimRight(fields, ",") + ")"
	vals = strings.TrimRight(vals, ",") + ")"

	result := fmt.Sprintf(`INSERT INTO "%s" %s VALUES %s;`, m.TableName(), fields, vals)

	return result
}

func GenerateUpdateQuery(newValue core.CoreModel, oldValue core.CoreModel) string {
	updateQuery := ""
	whereQuery := ""

	t := reflect.TypeOf(newValue)
	v := getValue(newValue)

	for i := 0; i < v.Elem().NumField(); i++ {
		if t.Field(i).Tag.Get("pk") == "true" {
			whereQuery += fmt.Sprintf("%s=%s", t.Field(i).Tag.Get("column"),
				getStringValue(t.Field(i), reflect.ValueOf(newValue).Field(i)))
			continue
		}

		field := t.Field(i).Tag.Get("column")
		if field != "" {
			if !isEqualFields(t.Field(i).Type.String(), reflect.ValueOf(newValue).Field(i), reflect.ValueOf(oldValue).Field(i)) {
				updateQuery += fmt.Sprintf(`%s=%s,`, field, getStringValue(t.Field(i), reflect.ValueOf(newValue).Field(i)))
			}
		}
	}

	result := ""
	if updateQuery != "" {
		updateQuery = strings.TrimRight(updateQuery, ",")
		result = fmt.Sprintf(`UPDATE "%s" SET %s WHERE %s;`, newValue.TableName(), updateQuery, whereQuery)
	}

	return result
}

func GenerateDeleteQuery(m core.CoreModel) string {
	t := reflect.TypeOf(m)
	v := getValue(m)

	whereQuery := ""
	for i := 0; i < v.Elem().NumField(); i++ {
		if t.Field(i).Tag.Get("pk") == "true" {
			whereQuery += fmt.Sprintf("%s=%s", t.Field(i).Tag.Get("column"),
				getStringValue(t.Field(i), reflect.ValueOf(m).Field(i)))
			break
		}
	}

	result := fmt.Sprintf(`DELETE FROM "%s" WHERE %s;`, m.TableName(), whereQuery)
	return result
}

func GenerateRollbackQuery(ecosystemID int64, blockID int64, newValue core.CoreModel, oldValue core.CoreModel, txHash []byte) string {
	rollbackJSON := ""
	id := ""

	t := reflect.TypeOf(newValue)
	v := getValue(newValue)

	for i := 0; i < v.Elem().NumField(); i++ {
		if t.Field(i).Tag.Get("pk") == "true" {
			id += fmt.Sprintf("%s", getStringValue(t.Field(i), reflect.ValueOf(newValue).Field(i)))
			continue
		}

		field := t.Field(i).Tag.Get("column")
		if field != "" {
			if !isEqualFields(t.Field(i).Type.String(), reflect.ValueOf(newValue).Field(i), reflect.ValueOf(oldValue).Field(i)) {
				rollbackJSON += fmt.Sprintf(`"%s": %s,`, field, getStringValue(t.Field(i), reflect.ValueOf(newValue).Field(i)))
			}
		}
	}

	result := ""
	if rollbackJSON != "" {
		rollbackJSON = fmt.Sprintf("{ %s }", strings.TrimRight(rollbackJSON, ","))
		result = fmt.Sprintf(`INSERT INTO rollback_tx(block_id, tx_hash, table_name, table_id, data) 
		VALUES (%d, decode('%s', 'HEX'), '%s', '%s', '%s')`, blockID, txHash, newValue.TableName(), id, rollbackJSON)
	}

	return result
}

func getValue(m core.CoreModel) reflect.Value {
	switch reflect.TypeOf(m).String() {
	case "model.block":
		v := m.(block)
		return reflect.ValueOf(&v)
	}
	panic("unknown model")
}

func isEqualFields(t string, v reflect.Value, v1 reflect.Value) bool {
	switch t {
	case "int32":
		return v.Interface().(int32) == v1.Interface().(int32)
	case "int64":
		return v.Interface().(int64) == v1.Interface().(int64)
	case "[]uint8":
		return core.IsSlicesEqual(v.Interface().([]uint8), v1.Interface().([]uint8))
	case "string":
		return v.Interface().(string) == v1.Interface().(string)
	}
	return false
}

func getStringValue(f reflect.StructField, v reflect.Value) string {
	switch f.Type.String() {
	case "int32":
		return strconv.FormatInt(int64(v.Interface().(int32)), 10)
	case "int64":
		return strconv.FormatInt(v.Interface().(int64), 10)
	case "[]uint8":
		return fmt.Sprintf(`decode('%s', 'HEX')`, hex.EncodeToString(v.Interface().([]uint8)))
	case "string":
		return fmt.Sprintf(`'%s'`, v.Interface().(string))
	}
	return ""
}
