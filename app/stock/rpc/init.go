package rpc

import "sync"

var (
	once sync.Once
)

func InitClient() {
	once.Do(func() {
		initOrderClient()
	})
}
