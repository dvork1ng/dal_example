package main

import (
	"buffers_test/DAL"
	"buffers_test/DAL/model"
	"fmt"
)

func main() {
	err := DAL.InitializeDAL()
	if err != nil {
		fmt.Println("dal init error ", err)
	}

	block, err := model.NewBlock().GetCurrent(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(block)
	}

	block, err = model.NewBlock().GetCurrent(15)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(block)
	}
}

// type Buffer interface {
// 	GetKey(ecosystemID int64, keyID int64) (Key, bool, error)
// 	UpdateKey(ecosystemID int64, keyID int64, key Key) (bool, error)
// 	PushKey(ecosystemID int64, keyID int64, key Key) (bool, error)
// }

// var k []Key
// var bk *bufferedKeys
// var bm *buffKeys

// const testsCount = 100
// const cyclesCount = 10000

// func main() {
// 	runtime.GOMAXPROCS(8)
// 	err := GormInit("localhost", 5432, "postgres", "postgres", "apla")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	k = make([]Key, testsCount*cyclesCount)
// 	for i := 0; i < testsCount*cyclesCount; i++ {
// 		temp := Key{}
// 		temp.ID = int64(i)
// 		temp.Amount = "1000"
// 		temp.PublicKey = []byte("lwkda;lsdad")
// 		k = append(k, temp)
// 	}

// 	runtime.GOMAXPROCS(8)

// 	fmt.Println("------------------start rw mutex------------------")
// 	bk := NewBufferedKeys()
// 	TestAll(bk)
// 	fmt.Println("------------------stop rw mutex------------------")

// 	fmt.Println("------------------start mutex------------------")
// 	bm := NewKeys()
// 	TestAll(bm)
// 	fmt.Println("------------------stop mutex------------------")

// 	fmt.Println("------------------start syncMap------------------")
// 	bsm := &mapKeys{}
// 	TestAll(bsm)
// 	fmt.Println("------------------stop syncMap------------------")
// }

// func TestAll(buffer Buffer) {
// 	var pushTimes time.Time
// 	for i := 0; i < testsCount; i++ {
// 		startTime := time.Now()
// 		TestPush(buffer)
// 		endTime := time.Now()
// 		pushTimes = pushTimes.Add(endTime.Sub(startTime))
// 	}

// 	fmt.Println("time for push: \t\t", pushTimes.Nanosecond()/100, "ns")

// 	var updateTimes time.Time
// 	for i := 0; i < testsCount; i++ {
// 		startTime := time.Now()
// 		TestUpdate(buffer)
// 		endTime := time.Now()
// 		updateTimes = updateTimes.Add(endTime.Sub(startTime))
// 	}
// 	fmt.Println("time for update: \t", updateTimes.Nanosecond()/100, "ns")

// 	var getTimes time.Time
// 	for i := 0; i < testsCount; i++ {
// 		startTime := time.Now()
// 		TestGet(buffer)
// 		endTime := time.Now()
// 		getTimes = getTimes.Add(endTime.Sub(startTime))
// 	}
// 	fmt.Println("time for get: \t\t", getTimes.Nanosecond()/100, "ns")

// 	var randomTimes time.Time
// 	for i := 0; i < testsCount; i++ {
// 		startTime := time.Now()
// 		TestRandom(buffer)
// 		endTime := time.Now()
// 		randomTimes = randomTimes.Add(endTime.Sub(startTime))
// 	}
// 	fmt.Println("time for random: \t", randomTimes.Nanosecond()/100, "ns")
// }

// func TestPush(buffer Buffer) {
// 	var i int64
// 	var wg sync.WaitGroup
// 	wg.Add(cyclesCount)
// 	for i = 0; i < cyclesCount; i++ {
// 		go func() {
// 			buffer.PushKey(1, k[i].ID, k[i])
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// }

// func TestUpdate(buffer Buffer) {
// 	var i int64
// 	var wg sync.WaitGroup
// 	wg.Add(cyclesCount)
// 	for i = 0; i < cyclesCount; i++ {
// 		go func() {
// 			buffer.UpdateKey(1, k[i].ID, k[i])
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// }

// func TestGet(buffer Buffer) {
// 	var i int64
// 	var wg sync.WaitGroup
// 	wg.Add(cyclesCount)

// 	for i = 0; i < cyclesCount; i++ {
// 		go func() {
// 			buffer.GetKey(1, k[i].ID)
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// }

// func TestRandom(buffer Buffer) {
// 	var i int64
// 	var wg sync.WaitGroup
// 	wg.Add(cyclesCount)

// 	for i = 0; i < cyclesCount/2; i++ {
// 		go func() {
// 			buffer.GetKey(1, k[i].ID)
// 			wg.Done()
// 		}()
// 	}
// 	for i = 0; i < cyclesCount/2; i++ {
// 		go func() {
// 			buffer.UpdateKey(1, k[i].ID, k[i])
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// }
