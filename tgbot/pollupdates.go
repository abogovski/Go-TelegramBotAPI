package tgbot

import (
	"log"
	"sync/atomic"
)

// PollUpdatesCB .
func PollUpdatesCB(botAPIURL string, params Params, handleUpdates func([]Update, Integer) (Integer, bool)) (Integer, int, error) {
	pollNext := true
	var offset Integer
	if val, ok := params["offset"]; ok {
		offset = val.(Integer)
	}

	for pollNext {
		updates, status, err := GetUpdates(botAPIURL, params)
		if status == 239 {
			log.Println("RateLimit exceeded")
			return offset, status, err
		} else if err != nil {
			return offset, status, err
		} else {
			offset, pollNext = handleUpdates(updates, offset)
			params["offset"] = offset
		}
	}

	params["limit"] = 1
	params["timeout"] = 0
	_, status, err := GetUpdates(botAPIURL, params)
	if err != nil {
		return offset, status, err
	}
	return offset, status, nil
}

// PollUpdates .
func PollUpdates(botAPIURL string, params Params, output chan<- Update, stop *int32) (Integer, int, error) {
	handleUpdates := func(updates []Update, offset Integer) (Integer, bool) {
		for _, update := range updates {
			if atomic.LoadInt32(stop) > 0 {
				log.Println("PollUpdates received stop.")
				close(output)
				return offset, false
			}
			offset = update.UpdateID + 1
			output <- update
		}
		return offset, true
	}

	return PollUpdatesCB(botAPIURL, params, handleUpdates)
}
