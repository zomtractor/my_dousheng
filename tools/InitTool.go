package tools

import (
	"my_dousheng/controller"
	"my_dousheng/dao"
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
		controller.InitRouter()
		sy.Done()
	}()
	sy.Wait()
}
