package core

import "github.com/blitzdb/blitz/config"

// Evicts the first key it found while iterating the map
// TODO: Make it efficient by doing thorough sampling
func evictFirst() {
	for k := range store {
		delete(store, k)
		return
	}
}

func evictAllkeysRandom(){
	evictCount:=int64(config.EvictionRatio*float64(config.KeysLimit))

	for k:=range store{
		Del(k)
		evictCount--
		if evictCount<=0{
			break
		}
	}
}

// TODO: Make the eviction strategy configuration driven
// TODO: Support multiple eviction strategies
func evict() {
	switch config.EvictionStrategy {
	case "simple-first":
		evictFirst()

	case "allkeys-random":
		evictAllkeysRandom()

	}
}
