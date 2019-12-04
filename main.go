/*
 * @Author: chijinxin
 * @Email: chijinxind@hotmail.com
 * @Date:  5:34 下午
 */
package main

import (
	"fmt"
	"time"

	"GoroutinePauseController/controller"
)

func main() {
	ctl := controller.NewGoroutinePauseController()
	go func() {
		for {
			fmt.Println("waiting ...")
			ctl.WaitScheduleSignal()
			fmt.Println("pass ... ")
			fmt.Println("running")
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			flag := 0
			fmt.Scanln(&flag)
			if flag == 1 {
				fmt.Println("暂停")
				ctl.Pause()
				fmt.Println("暂停成功")

			} else if flag == 2 {
				fmt.Println("继续")
				ctl.Resume()
				fmt.Println("继续成功")
			}
		}
	}()
	select {}
}
