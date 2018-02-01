package model

import "testing"
import "strings"

func TestInsertQuery(t *testing.T) {
	block := NewBlock()
	block.ID = 1123
	block.Hash = []byte("hello")
	block.Data = []byte("hello")
	block.EcosystemID = 12
	block.KeyID = 12345
	block.NodePosition = 11
	block.Time = 31
	block.Tx = 32

	s := GenerateInsertQuery(block)
	testQuery := `INSERT INTO "block_chain" (id,hash,data,ecosystem_id,key_id,node_position,time,tx) VALUES (1123,'68656c6c6f','68656c6c6f',12,12345,11,31,32);`
	if strings.Compare(s, testQuery) != 0 {
		t.Error("incorrect query: ", s)
		t.Error("  correct query: ", testQuery)
	}
}

func TestUpdateQuery(t *testing.T) {
	block := NewBlock()
	block.ID = 1123
	block.Hash = []byte("bye")
	block.Data = []byte("hello")
	block.EcosystemID = 12
	block.KeyID = 12345
	block.NodePosition = 11
	block.Time = 31
	block.Tx = 32

	newBlock := NewBlock()
	newBlock.ID = 1123
	newBlock.Hash = []byte("hello")
	newBlock.Data = []byte("hello")
	newBlock.EcosystemID = 124
	newBlock.KeyID = 12345
	newBlock.NodePosition = 11
	newBlock.Time = 31
	newBlock.Tx = 32

	testQuery := `UPDATE "block_chain" SET hash='68656c6c6f',ecosystem_id=124 WHERE id=1123;`
	s := GenerateUpdateQuery(newBlock, block)
	if strings.Compare(s, testQuery) != 0 {
		t.Error("incorrect query: ", s)
		t.Error("  correct query: ", testQuery)
	}
}

func TestDeleteQuery(t *testing.T) {
	block := NewBlock()
	block.ID = 1123

	s := GenerateDeleteQuery(block)
	testQuery := `DELETE FROM "block_chain" WHERE id=1123;`
	if strings.Compare(s, testQuery) != 0 {
		t.Error("incorrect query: ", s)
		t.Error("  correct query: ", testQuery)
	}
}
