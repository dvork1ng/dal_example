package buffers

// import (
// 	"sync"
// )

// type mapKeys struct {
// 	keys sync.Map
// }

// func (bk *mapKeys) GetKey(tablePrefix int64, id int64) (key Key, found bool, err error) {
// 	result := Key{}
// 	m, ok := bk.keys.Load(tablePrefix)
// 	if !ok {
// 		return result, false, nil
// 	}

// 	t := m.(sync.Map)

// 	inner, ok := t.Load(id)
// 	if !ok {
// 		return result, false, nil
// 	}

// 	result = inner.(KeyWithHistory).Key
// 	return result, true, nil
// }

// func (bk *mapKeys) UpdateKey(tablePrefix int64, id int64, key Key) (found bool, err error) {
// 	m, ok := bk.keys.Load(tablePrefix)
// 	if !ok {
// 		return false, nil
// 	}

// 	t := m.(sync.Map)

// 	val, ok := t.Load(id)
// 	if !ok {
// 		return false, nil
// 	}

// 	kwh := val.(KeyWithHistory)
// 	if kwh.Status != New {
// 		kwh.Status = Updated
// 	}
// 	kwh.Key = key
// 	t.Store(id, kwh)

// 	return true, nil
// }

// func (bk *mapKeys) PushKey(tablePrefix int64, id int64, key Key) (found bool, err error) {
// 	m, ok := bk.keys.Load(tablePrefix)
// 	if !ok {
// 		return false, nil
// 	}

// 	t := m.(sync.Map)

// 	kwh := KeyWithHistory{}
// 	kwh.Status = New
// 	kwh.Key = key
// 	kwh.History = Key{}
// 	t.Store(id, kwh)

// 	return true, nil
// }
