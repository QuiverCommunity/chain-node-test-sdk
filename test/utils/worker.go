package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime/debug"
	"time"
)

// This file manages block heights, sending transactions by block height
type Action struct {
	ActionId               string `yaml:"action_id"`
	OffsetHeight           int64  `yaml:"offset_height"`
	FromKey                string `yaml:"from_key"`
	Action                 string `yaml:"action"`
	Param                  string `yaml:"param"`
	ModifyBaseHeightSource string `yaml:"modify_block_height_source"`
}

type Worker struct {
	Action
	RunBlockHeight int64
	Done           bool
}

func GetRunBlockHeight(works []Worker, modSource string) int64 {
	if len(modSource) == 0 {
		return Config.InitialBlockHeight
	}
	for _, work := range works {
		if work.ActionId == modSource {
			return work.RunBlockHeight
		}
	}
	return 0xFFFFFF
}

func RunWorker(todos []Action) (string, error) {
	log := ""

	// initialize works
	works := []Worker{}
	for _, todo := range todos {
		works = append(works, Worker{todo, 0xFFFFFF, false})
	}
	round := 1
	for {
		allDone := true
		for _, work := range works {
			if work.Done == false {
				allDone = false
				break
			}
		}
		if allDone { // when all tasks are done, then break
			break
		}
		// fetch block height regularly
		nodeStatus, nodeQueryLog, queryErr := QueryNodeStatus()
		log += nodeQueryLog
		if queryErr != nil {
			// not able to query node, then break
			return log, queryErr
		}
		nowBlockHeight := nodeStatus.SyncInfo.LatestBlockHeight
		log += fmt.Sprintf("%d", nowBlockHeight)

		// run actions for block height fit actions, and run
		for index, work := range works {
			startBlockHeight := GetRunBlockHeight(works, work.ModifyBaseHeightSource) + work.OffsetHeight
			log += fmt.Sprintf("\nstartBlockHeight: %d, nowBlockHeight: %d", startBlockHeight, nowBlockHeight)
			if work.Done == false && nowBlockHeight >= startBlockHeight {
				works[index].RunBlockHeight = nowBlockHeight
				actionLog, actionErr := RunAction(work.Action.Action, work.Param)
				log += fmt.Sprintf("\nAction log for %s for %s param.\n", work.Action.Action, work.Param)
				log += actionLog
				if actionErr != nil {
					// find one task failure, then break
					log += "\n"
					log += string(debug.Stack())
					return log, actionErr
				} else {
					log += fmt.Sprintf("\n DONE work %s", work.ActionId)
					works[index].Done = true
				}
			}
		}
		worksLog, _ := json.Marshal(works)
		log += fmt.Sprintf("\nwork status: %+v", string(worksLog))
		if round >= 20 {
			return log, errors.New("not finished after 20 round of work")
		}
		// TODO external signal for ending loop, then break
		// TODO should give external server log signal so that users can see real time logs
		// sleep 1s, TODO should be able to configure this by config.yml
		time.Sleep(1 * time.Second)
		round += 1
	}
	return log, nil
}
