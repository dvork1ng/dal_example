package buffers

// import (
// 	"bytes"
// 	"encoding/hex"
// 	"fmt"
// )

// func loadVal(ecosystemID int64, keyID int64) (Key, error) {
// 	key := Key{}
// 	k := key.SetTablePrefix(ecosystemID)
// 	err := k.Get(keyID)
// 	return *k, err
// }

// func (bk *bufferedKeys) GetKey(tablePrefix int64, id int64) (key Key, found bool, err error) {
// 	result := Key{}
// 	bk.rwMutex.RLock()
// 	defer bk.rwMutex.RUnlock()
// 	_, ok := bk.keys[tablePrefix]
// 	if !ok {
// 		return result, false, nil
// 	}

// 	_, ok = bk.keys[tablePrefix][id]
// 	if !ok {
// 		return result, false, nil
// 	}

// 	result = bk.keys[tablePrefix][id].Key
// 	return result, true, nil
// }

// func (bk *bufferedKeys) UpdateKey(tablePrefix int64, id int64, key Key) (found bool, err error) {
// 	bk.rwMutex.RLock()
// 	_, ok := bk.keys[tablePrefix]
// 	if !ok {
// 		return false, nil
// 	}

// 	_, ok = bk.keys[tablePrefix]
// 	if !ok {
// 		return false, nil
// 	}
// 	bk.rwMutex.RUnlock()

// 	bk.rwMutex.Lock()
// 	bk.keys[tablePrefix][id].Key = key
// 	if bk.keys[tablePrefix][id].Status != New {
// 		bk.keys[tablePrefix][id].Status = Updated
// 	}
// 	bk.rwMutex.Unlock()
// 	return true, nil
// }

// func (bk *bufferedKeys) PushKey(tablePrefix int64, id int64, key Key) (found bool, err error) {
// 	bk.rwMutex.RLock()
// 	_, ok := bk.keys[tablePrefix]
// 	if !ok {
// 		bk.keys[tablePrefix] = make(map[int64]*KeyWithHistory)
// 	}
// 	bk.rwMutex.RUnlock()

// 	bk.rwMutex.Lock()
// 	bk.keys[tablePrefix][id] = &KeyWithHistory{Key: key, History: Key{}, Status: New}
// 	bk.rwMutex.Unlock()
// 	return true, nil
// }

// func (k Key) GenerateUpdateSQL(ecosystemID int64, oldKey *Key) string {
// 	updateQuery := fmt.Sprintf(` UPDATE "%d_keys" SET`, ecosystemID)
// 	if !bytes.Equal(k.PublicKey, oldKey.PublicKey) {
// 		updateQuery += fmt.Sprintf(` pub = '%s',`, hex.EncodeToString(k.PublicKey))
// 	}
// 	if k.Amount != oldKey.Amount {
// 		updateQuery += fmt.Sprintf(` amount = %s`, k.Amount)
// 	}
// 	updateQuery += fmt.Sprintf(` where id = %d;`, k.ID)
// 	return updateQuery
// }

// func (k Key) GenerateInsertSQL(ecosystemID int64) string {
// 	return fmt.Sprintf(`INSERT INTO "%d_keys" VALUES (%d, '%s', %s);`,
// 		ecosystemID, k.ID, hex.EncodeToString(k.PublicKey), k.Amount)
// }

// func (bk *bufferedKeys) Flush( /*transaction *DbTransaction,*/ blockID int64) error {
// 	updateQueries := ""
// 	insertQueries := ""
// 	for ecosystemID, table := range bk.keys {
// 		for _, key := range table {
// 			if key.Status == New {
// 				updateQueries += key.Key.GenerateInsertSQL(ecosystemID)
// 			} else if key.Status == Updated {
// 				insertQueries += key.Key.GenerateUpdateSQL(ecosystemID, &key.History)
// 			}
// 			updateQueries += key.generateRollbackQueries(blockID, ecosystemID)
// 			key.History = key.Key
// 			key.Status = Original
// 		}
// 	}
// 	insertQueries += updateQueries
// 	if len(insertQueries) > 0 {
// 		err := DBConn.Exec(insertQueries).Error
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (kwh *KeyWithHistory) generateRollbackQueries(blockID int64, ecosystemID int64) string {
// 	result := ""
// 	if kwh.Status == New {
// 		result += fmt.Sprintf(`INSERT INTO rollback_tx(block_id, tx_hash, table_name, table_id, data)
// 		VALUES (%d, decode('', 'HEX'), '%d_keys', '%d', '');`, blockID, ecosystemID, kwh.Key.ID)
// 	} else if kwh.Status == Updated {
// 		fields := "{"
// 		if kwh.Key.Amount != kwh.History.Amount {
// 			fields += fmt.Sprintf(`"amount": %s`, kwh.History.Amount)
// 		}
// 		if !sliceEqual(kwh.Key.PublicKey, kwh.History.PublicKey) {
// 			fields += fmt.Sprintf(`, "pub": %s`, hex.EncodeToString(kwh.History.PublicKey))
// 		}
// 		fields += "}"
// 		result += fmt.Sprintf(`INSERT INTO rollback_tx(block_id, tx_hash, table_name, table_id, data)
// 			VALUES (%d, decode('', 'HEX'), '%d_keys', '%d', '%s');`, blockID, ecosystemID, kwh.Key.ID, fields)
// 	}
// 	return result
// }
