package model

import (
	"buffers_test/DAL/core"
	"reflect"
)

// Block is model
type block struct {
	baseModel
	ID           int64  `column:"id" pk:"true"`
	Hash         []byte `column:"hash"`
	Data         []byte `column:"data"`
	EcosystemID  int64  `column:"ecosystem_id"`
	KeyID        int64  `column:"key_id"`
	NodePosition int64  `column:"node_position"`
	Time         int64  `column:"time"`
	Tx           int32  `column:"tx"`
}

func NewBlock() block {
	result := block{}
	result.tableName = "block_chain"
	result.isVDE = false
	result.tablePrefix = -1
	return result
}

func (b block) Type() reflect.Type {
	return reflect.TypeOf(b)
}

func (b block) GetCurrent(blockID int64) (block, error) {
	val := caches[reflect.TypeOf(b).String()].GetCurrent(b.tablePrefix, blockID)
	if val == nil {
		return NewBlock(), core.ErrRecordNotFound
	}
	return val.(block), nil
}

func (b block) GetVerified(blockID int64) (block, error) {
	val := caches[reflect.TypeOf(b).String()].GetVerified(b.tablePrefix, blockID)
	if val == nil {
		return NewBlock(), core.ErrRecordNotFound
	}
	return val.(block), nil
}

func (b block) Update(newBlock block) error {
	return caches[reflect.TypeOf(newBlock).String()].Update(newBlock.tablePrefix, newBlock.ID, newBlock)
}

func (b block) Upsert(newBlock block) {
	caches[reflect.TypeOf(b).String()].Push(newBlock.tablePrefix, newBlock.ID, newBlock)
}

func (b block) Delete(newBlock block) error {
	return caches[reflect.TypeOf(b).String()].Delete(newBlock.tablePrefix, newBlock.ID, newBlock)
}

func (b block) UpdateCache() error {
	rows, err := core.DBConn.Query("select * from block_chain;")
	if err != nil {
		return err
	}

	cache := caches[reflect.TypeOf(b).String()]
	for rows.Next() {
		next := NewBlock()
		err := rows.Scan(&next.ID, &next.Hash, &next.Data, &next.EcosystemID, &next.KeyID, &next.NodePosition, &next.Time, &next.Tx)
		if err != nil {
			return err
		}
		cache.Push(-1, next.ID, next)
	}
	return nil
}

// // TableName returns name of table
// func (Block) TableName() string {
// 	return "block_chain"
// }

// // Create is creating record of model
// func (b *Block) Create(transaction *DbTransaction) error {
// 	return GetDB(transaction).Create(b).Error
// }

// // Get is retrieving model from database
// func (b *Block) Get(blockID int64) (bool, error) {
// 	return isFound(DBConn.Where("id = ?", blockID).First(b))
// }

// // GetMaxBlock returns last block existence
// func (b *Block) GetMaxBlock() (bool, error) {
// 	return isFound(DBConn.Last(b))
// }

// // GetBlockchain is retrieving chain of blocks from database
// func GetBlockchain(startBlockID int64, endblockID int64) ([]Block, error) {
// 	var err error
// 	blockchain := new([]Block)
// 	if endblockID > 0 {
// 		err = DBConn.Model(&Block{}).Order("id asc").Where("id > ? AND id <= ?", startBlockID, endblockID).Find(&blockchain).Error
// 	} else {
// 		err = DBConn.Model(&Block{}).Order("id asc").Where("id > ?", startBlockID).Find(&blockchain).Error
// 	}
// 	if err != nil {
// 		return nil, err
// 	}
// 	return *blockchain, nil
// }

// // GetBlocks is retrieving limited chain of blocks from database
// func (b *Block) GetBlocks(startFromID int64, limit int32) ([]Block, error) {
// 	var err error
// 	blockchain := new([]Block)
// 	if startFromID > 0 {
// 		err = DBConn.Order("id desc").Limit(limit).Where("id > ?", startFromID).Find(&blockchain).Error
// 	} else {
// 		err = DBConn.Order("id desc").Limit(limit).Find(&blockchain).Error
// 	}
// 	return *blockchain, err
// }

// // GetBlocksFrom is retrieving ordered chain of blocks from database
// func (b *Block) GetBlocksFrom(startFromID int64, ordering string) ([]Block, error) {
// 	var err error
// 	blockchain := new([]Block)
// 	err = DBConn.Order("id "+ordering).Where("id > ?", startFromID).Find(&blockchain).Error
// 	return *blockchain, err
// }

// // DeleteById is deleting block by ID
// func (b *Block) DeleteById(transaction *DbTransaction, id int64) error {
// 	return GetDB(transaction).Where("id = ?", id).Delete(Block{}).Error
// }
