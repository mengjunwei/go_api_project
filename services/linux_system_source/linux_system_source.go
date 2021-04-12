package linux_system_source

import (
	"runtime"
	"time"

	"github.com/shirou/gopsutil/mem"

	"github.com/mengjunwei/go_api_project/logger"
	"github.com/mengjunwei/go_api_project/models"
)

var memList [][]byte

func init() {
	memList = make([][]byte, 0)
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
					byteSlice := make([]byte, 1<<26)
					memList = append(memList, byteSlice)
					continue
				} else {
					setMemIndex := mTotal * uint64(params.Value) / (100 * 1 << 26)
					if len(memList) >= int(setMemIndex) {
						tmpList := make([][]byte, 0, int(setMemIndex))
						for index, item := range memList {
							if index < int(setMemIndex) {
								tmpList = append(tmpList, item)
							}
						}
						memList = tmpList
					}

				}
				logger.Info(curVal)
				return
			}
		}
	}(params.Value)

	return "ok", nil
}
