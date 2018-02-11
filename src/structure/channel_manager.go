package structure

import "sync"

type ChannelManager struct {
	channels     []Channel
	mapSrc 		 map[string]*Channel
	mapDst		 map[string]*Channel
	mutex 		 *sync.Mutex
}

func (manager *ChannelManager) Init() {
	manager.channels = make([]Channel, 0)
	manager.mapSrc = make(map[string]*Channel)
	manager.mapDst = make(map[string]*Channel)
	manager.mutex = new(sync.Mutex)
}
