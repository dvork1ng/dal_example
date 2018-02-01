package DAL

import (
	"buffers_test/DAL/buffers"
	"buffers_test/DAL/core"
	"buffers_test/DAL/model"
)

func InitializeDAL() error {
	err := core.InitDatabase("localhost", 5432, "postgres", "postgres", "apla")
	if err != nil {
		return err
	}

	mutexCache := buffers.NewMutexCache()

	err = model.InitCache(model.NewBlock(), mutexCache)
	if err != nil {
		return err
	}
	return nil
}
