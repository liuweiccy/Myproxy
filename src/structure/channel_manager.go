package structure

import (
    "sync"
    "util"
)

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

// 存储已经链接的通道
func (manager *ChannelManager) PutChannel(channel *Channel) {
    manager.mutex.Lock()
    defer manager.mutex.Unlock()
    manager.channels = append(manager.channels, *channel)
    manager.mapSrc[channel.SrcUrl()] = channel
    manager.mapDst[channel.DstUrl()] = channel
}

// 删除指定链接通道
func (manager *ChannelManager) DeleteChannel(channel *Channel) {
    manager.mutex.Lock()
    defer manager.mutex.Unlock()
    index := util.SliceIndex(manager.channels, *channel)
    if index >= 0 {
        manager.channels = append(manager.channels[:index], manager.channels[index + 1:]...)
        delete(manager.mapSrc, channel.SrcUrl())
        delete(manager.mapDst, channel.DstUrl())
    }
}
func (manager *ChannelManager) Clean() {
    for _, channel := range manager.channels {
        delete(manager.mapSrc, channel.SrcUrl())
        delete(manager.mapDst, channel.DstUrl())
        channel.Close()
    }
    manager.channels = manager.channels[:0]
}

func (manager *ChannelManager)GetChannels() []Channel {
    return manager.channels
}
