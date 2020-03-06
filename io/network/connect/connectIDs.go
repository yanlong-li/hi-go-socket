package connect

import "sync"

//获取ID的互斥锁
var getIdLock sync.Mutex

//自动序列ID
var autoSequenceID uint64 = 0

// 空闲ID--被销毁的可再次使用的ID
var idleSequenceIDs []uint64

// 获得自动序列ID
func GetAutoSequenceID() uint64 {

	// 使用锁防止并发情况出现ID重复问题
	// 添加锁
	getIdLock.Lock()
	//解除锁
	defer getIdLock.Unlock()

	var ID uint64
	if len(idleSequenceIDs) >= 1 {
		// 取出第一个值
		ID = idleSequenceIDs[0]
		idleSequenceIDs = idleSequenceIDs[1:]

	} else {
		// 取出序列ID
		ID = autoSequenceID
		// 序列ID自加 1
		autoSequenceID++
	}

	//返回ID
	return ID
}

// 添加一个空闲ID
func AddIdleSequenceId(ID uint64) {
	// 添加锁
	getIdLock.Lock()
	//解除锁
	defer getIdLock.Unlock()

	idleSequenceIDs = append(idleSequenceIDs, ID)

}
