package structure

import "sync"

type ChannelManager struct {
	channels     []Channel
	mapSrc 		 map[string]*Channel
	mapDst		 map[string]*Channel
	mutex 		 *sync.Mutex
}

func (manager *ChannelManager) Init() {

}
