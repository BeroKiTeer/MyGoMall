package rpc

import "sync"

var (
	once sync.Once
)

func InitClient() {
	once.Do(func() {
		initAuthClient()
		initUserClient()
		initCartClient()
	})
}
