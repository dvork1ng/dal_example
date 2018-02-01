package buffers

import (
	"buffers_test/DAL/core"
	"sync"
)

func NewMutexCache() *mutexCache {
	return &mutexCache{vals: make(map[int64]map[int64]cachedVal, 0), flushed: true}
}

type mutexCache struct {
	vals    map[int64]map[int64]cachedVal
	mutex   sync.Mutex
	flushed bool
}

func (mc *mutexCache) GetCurrent(tablePrefix int64, id int64) core.CoreModel {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	_, ok := mc.vals[tablePrefix]
	if !ok {
		return nil
	}

	_, ok = mc.vals[tablePrefix][id]
	if !ok {
		return nil
	}

	return mc.vals[tablePrefix][id].current
}

func (mc *mutexCache) GetVerified(tablePrefix int64, id int64) core.CoreModel {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	_, ok := mc.vals[tablePrefix]
	if !ok {
		return nil
	}

	_, ok = mc.vals[tablePrefix][id]
	if !ok {
		return nil
	}

	return mc.vals[tablePrefix][id].verified
}

func (mc *mutexCache) Update(tablePrefix int64, id int64, val core.CoreModel) error {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	_, ok := mc.vals[tablePrefix]
	if !ok {
		return core.ErrEcosystemNotFound
	}

	_, ok = mc.vals[tablePrefix]
	if !ok {
		return core.ErrRecordNotFound
	}

	cachedVal := cachedVal{current: val, verified: mc.vals[tablePrefix][id].verified}
	if mc.vals[tablePrefix][id].status != new {
		cachedVal.status = updated
	} else {
		cachedVal.status = mc.vals[tablePrefix][id].status
	}

	mc.vals[tablePrefix][id] = cachedVal
	mc.flushed = false
	return nil
}

func (mc *mutexCache) Insert(tablePrefix int64, id int64, val core.CoreModel) error {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	_, ok := mc.vals[tablePrefix]
	if !ok {
		mc.vals[tablePrefix] = make(map[int64]cachedVal)
	}

	if _, ok := mc.vals[tablePrefix][id]; ok {
		return core.ErrRecordAlreadyExists
	}

	mc.vals[tablePrefix][id] = cachedVal{current: val, verified: nil, status: new}
	mc.flushed = false
	return nil
}

func (mc *mutexCache) Push(tablePrefix int64, id int64, val core.CoreModel) {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	_, ok := mc.vals[tablePrefix]
	if !ok {
		mc.vals[tablePrefix] = make(map[int64]cachedVal)
	}

	mc.vals[tablePrefix][id] = cachedVal{current: val, verified: nil, status: new}
	mc.flushed = false
}

func (mc *mutexCache) Delete(tablePrefix int64, id int64, val core.CoreModel) error {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	_, ok := mc.vals[tablePrefix]
	if !ok {
		return core.ErrRecordNotFound
	}

	if _, ok := mc.vals[tablePrefix][id]; ok {
		return core.ErrRecordNotFound
	}

	mc.vals[tablePrefix][id] = cachedVal{current: nil, verified: mc.vals[tablePrefix][id].verified, status: deleted}
	mc.flushed = false
	return nil
}

func (mc *mutexCache) Flush(blockID int64) error {
	updateQueries := ""
	insertQueries := ""
	for ecosystemID, table := range mc.vals {
		for _, key := range table {
			if key.status == new {
				updateQueries += key.current.GenerateInsertQuery()
			} else if key.status == updated {
				insertQueries += key.current.GenerateUpdateQuery(&key.verified)
			}
			updateQueries += key.current.GenerateRollbackQueries(blockID, ecosystemID)
			key.verified = key.current
			key.status = original
		}
	}
	insertQueries += updateQueries
	if len(insertQueries) > 0 {
		_, err := core.DBConn.Exec(insertQueries)
		if err != nil {
			return err
		}
	}
	mc.flushed = true

	return nil
}
