package utils

import (
	"time"
)

// This file manages block heights, sending transactions by block height
type Action struct {
	OffsetHeight int64  `yaml:"offset_height"`
	FromKey      string `yaml:"from_key"`
	Action       string `yaml:"action"`
	Param        string `yaml:"param"`
}

type Worker struct {
	Action
	Done bool
}

func RunWorker(todos []Action) (string, error) {
	log := ""

	// initialize works
	works := []Worker{}
	for _, todo := range todos {
		works = append(works, Worker{todo, false})
	}
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
		nodeStatus, queryErr := QueryNodeStatus()
		if queryErr != nil {
			// not able to query node, then break
			return log, queryErr
		}
		nowBlockHeight := nodeStatus.SyncInfo.LatestBlockHeight

		// run actions for block height fit actions, and run
		for _, work := range works {
			if work.Done == false && nowBlockHeight >= work.OffsetHeight+Config.InitialBlockHeight {
				actionLog, actionErr := RunAction(work.Action.Action, work.FromKey)
				log += actionLog
				if actionErr != nil {
					// find one task failure, then break
					return log, actionErr
				} else {
					work.Done = true
				}
			}
		}
		// TODO external signal for ending loop, then break
		// sleep 1s, TODO should be able to configure this by config.yml
		time.Sleep(1000)
	}
	// check if block height fit for worker
	// run action if fit and collect logs
	// return collected logs and errors
	return log, nil
}
