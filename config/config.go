package config

var Host string = "0.0.0.0"
var Port int = 7379
var KeysLimit int = 5

// var EvictionStrategy string = "simple-first"
var EvictionStrategy string = "allkeys-random"

var AOFFile string ="./blitzdb.aof"

// will evict EvictionRatio of keys when eviction occurs
var EvictionRatio float64=0.40