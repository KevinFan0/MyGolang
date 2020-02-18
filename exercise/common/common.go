package common
import (
	"time"
	"os"
	"os/signal"
	"errors"
)

type Runner struct {
	tasks		[]func(int)				// 要执行的任务
	complete	chan error				// 用于通知任务全部完成
	timeout		<- chan time.Time		// 这些任务在多久内完成
	interrupt	chan os.Signal			// 可以控制强制终止的信号
}

var ErrTimeOut = errors.New("执行者执行超时")
var ErrInterrupt = errors.New("执行者被中断")

// 定义工厂函数New，返回我们需要的Runner
func New(tm time.Duration) *Runner {
	// tasks没有初始化，默认就是零值nil
	return &Runner{
		complete: make(chan error),
		timeout: time.After(tm),
		interrupt: make(chan os.Signal, 1),
	}
}

//将需要执行的任务，添加到Runner里
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// 执行任务，执行的过程中接收到中断信号时，返回中断错误
// 如果任务全部执行完，还没有接收到中断信号，则返回nil
func (r *Runner) run()  error { 
	for id, task := range r.tasks {
		if r.isInterrupt(){
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

// 检查是否接收到中断信号
func (r *Runner)  isInterrupt() bool {
	select {
	case <- r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

// 开始执行所以任务，并且监视通道事件
func (r *Runner) Start()  error {
	// 希望接收哪些系统信号
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()
	select {
	case err := <- r.complete:
		return err
	case <- r.timeout:
		return ErrTimeOut
	}
}