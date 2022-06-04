package tools

import (
	"my_dousheng/dao"
	"my_dousheng/router"
	"sync"
)

func Init() {
	sy := sync.WaitGroup{}
	sy.Add(2)
	go func() {
		dao.InitDBTool()
		sy.Done()
	}()
	go func() {
		router.InitRouter()
		sy.Done()
	}()
	sy.Wait()
}
