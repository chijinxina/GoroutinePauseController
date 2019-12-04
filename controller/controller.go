/*
 * @Author: chijinxin
 * @Email: chijinxind@hotmail.com
 * @Date:  5:29 下午
 */
package controller

import (
	"sync/atomic"
)

type GoroutinePauseController struct {
	pauseFlag int32
	pauseChan chan struct{}
}

func NewGoroutinePauseController() *GoroutinePauseController {
	return &GoroutinePauseController{
		pauseFlag: int32(0),
		pauseChan: make(chan struct{}, 1),
	}
}

// 被控协程需要调用此函数 用于暂停时阻塞协程执行
func (p *GoroutinePauseController) WaitScheduleSignal() {
	if atomic.CompareAndSwapInt32(&p.pauseFlag, 1, 2) {
		<-p.pauseChan
	}
}

// 可在任意协程中调用此方法, 使被控协程暂停执行
func (p *GoroutinePauseController) Pause() {
	atomic.CompareAndSwapInt32(&p.pauseFlag, 0, 1)
}

// 可在任意协程中调用此方法, 使被控协程继续执行
func (p *GoroutinePauseController) Resume() {
	if atomic.CompareAndSwapInt32(&p.pauseFlag, 2, 0) {
		p.pauseChan <- struct{}{}
	}
	atomic.CompareAndSwapInt32(&p.pauseFlag, 1, 0)
}
