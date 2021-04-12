package linux_system_source

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"github.com/shirou/gopsutil/mem"

	"github.com/mengjunwei/go_api_project/logger"
	"github.com/mengjunwei/go_api_project/models"
)

var (
	memList [][]string
	lock    sync.Mutex
)

func init() {
	memList = make([][]string, 0)
}

type Service struct {
	LoginUser *models.LoginUser
}

func (s *Service) SetMemory(params *models.SystemSetDTO) (interface{}, error) {
	memUsedFn := func() (float64, uint64, error) {
		memInfo, err := mem.VirtualMemory()
		if err != nil {
			logger.Error(err)
			return 0, 0, err
		}

		memFree := memInfo.Free + memInfo.Buffers + memInfo.Cached
		if memInfo.Available > 0 {
			memFree = memInfo.Available
		}

		memUsed := memInfo.Total - memFree

		pmemUsed := 0.0
		if memInfo.Total != 0 {
			pmemUsed = float64(memUsed) * 100.0 / float64(memInfo.Total)
		}
		return pmemUsed, memInfo.Total, nil
	}

	go func(memValueSet float64) {
		defer func() {
			runtime.GC()
		}()
		t := time.NewTicker(time.Duration(100) * time.Millisecond)
		for {
			select {
			case <-t.C:
				curVal, mTotal, _ := memUsedFn()
				if curVal < memValueSet {
					byteSlice := make([]string, 0, 1<<22)
					for index, item := range byteSlice {
						byteSlice[index] = fmt.Sprintf("%d", time.Now().UnixNano())
						byteSlice = append(byteSlice, item)
					}
					lock.Lock()
					memList = append(memList, byteSlice)
					lock.Unlock()
					continue
				} else {
					setMemIndex := mTotal * uint64(params.Value) / (100 * 1 << 22)
					if len(memList) >= int(setMemIndex) {
						tmpList := make([][]string, 0, int(setMemIndex))
						for index, item := range memList {
							if index < int(setMemIndex) {
								tmpList = append(tmpList, item)
							}
						}
						lock.Lock()
						memList = tmpList
						lock.Unlock()
					}

				}
				logger.Info(curVal)
				return
			}
		}
	}(params.Value)

	return "ok", nil
}

func ReadMemList() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	t := time.NewTicker(time.Duration(1000) * time.Millisecond)
	for {
		select {
		case <-t.C:
			lock.Lock()
			if len(memList) > 0 {
				index := r.Intn(len(memList))
				for _, item := range memList[index] {
					str := fmt.Sprint(item)
					if str == "" {
						logger.Debug(str)
					}
				}
			}
			lock.Unlock()
		}
	}
}
