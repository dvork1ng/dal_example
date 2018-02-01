package model

import "strconv"

type baseModel struct {
	tableName   string
	tablePrefix int64
	isVDE       bool
}

func (bm baseModel) TableName() string {
	if bm.isVDE {
		return strconv.FormatInt(bm.tablePrefix, 10) + "_vde_" + bm.tableName
	}
	if bm.tablePrefix > 0 {
		return strconv.FormatInt(bm.tablePrefix, 10) + "_" + bm.tableName
	}
	return bm.tableName
}

func (bm baseModel) IsEcosystemModel() bool {
	return bm.tablePrefix > 0
}
