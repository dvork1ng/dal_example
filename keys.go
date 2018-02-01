package main

// type KeyWithHistory struct {
// 	Key     Key
// 	History Key
// 	Status  recordStatus
// }

// func NewBufferedKeys() *bufferedKeys {
// 	return &bufferedKeys{keys: make(map[int64]map[int64]*KeyWithHistory)}
// }

// type bufferedKeys struct {
// 	keys        map[int64]map[int64]*KeyWithHistory
// 	rwMutex     sync.RWMutex
// 	updateMutex sync.Mutex
// }

// func loadEcosystemKeys(ecosystemID int64) (*[]Key, error) {
// 	var keys []Key
// 	err := DBConn.Raw(fmt.Sprintf(`select * from "%d_keys";`, ecosystemID)).Scan(&keys).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &keys, nil
// }

// // Key is model
// type Key struct {
// 	tableName string
// 	ID        int64  `gorm:"primary_key;not null"`
// 	PublicKey []byte `gorm:"column:pub;not null"`
// 	Amount    string `gorm:"not null"`
// }

// // SetTablePrefix is setting table prefix
// func (m *Key) SetTablePrefix(prefix int64) *Key {
// 	if prefix == 0 {
// 		prefix = 1
// 	}
// 	m.tableName = fmt.Sprintf("%d_keys", prefix)
// 	return m
// }

// // TableName returns name of table
// func (m Key) TableName() string {
// 	return m.tableName
// }

// // Get is retrieving model from database
// func (m *Key) Get(wallet int64) error {
// 	return DBConn.Where("id = ?", wallet).First(m).Error
// }
