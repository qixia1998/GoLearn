# 如果一个 Mutex 已经被一个 goroutine 获取了锁，其它等待中的 goroutine 们只能一直等待。那么，等这个锁释放后，等待中的 goroutine 中哪一个会优先获取 Mutex 呢？

等待的 goroutine 们是以 <b>FIFO</b> 排队的

1）当 Mutex 处于正常模式时，若此时没有新的 goroutine 与队头 goroutine 竞争，则队友 goroutine 获得。若有新的 goroutine 竞争大概新的goroutine 获得。
2）当队头 goroutine 竞争锁失败 1ms 后，它会将  Mutex 调整为饥饿模式。进入饥饿模式后，锁的所有权会直接从解锁 goroutine 移交给队头 goroutine，此时新来的 goroutine 直接放入队尾。
3）当一个 goroutine 获取锁后，如果发现自己满足下列条件中的任何一个 #1 它是队列中最后一个 #2 它等待锁的时间少于 1ms， 则将锁切换回正常模式
