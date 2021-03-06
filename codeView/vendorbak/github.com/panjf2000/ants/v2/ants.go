// MIT License

// Copyright (c) 2018 Andy Pan

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package ants

import (
	"errors"
	"log"
	"math"
	"os"
	"runtime"
	"time"
)

const (
	// DefaultAntsPoolSize is the default capacity for a default goroutine pool.
	/**
	 * DefaultAntsPoolSize是默认goroutine池的默认容量。 2^32 -1
	*/
	DefaultAntsPoolSize = math.MaxInt32

	// DefaultCleanIntervalTime is the interval time to clean up goroutines.
	/**
	* DefaultCleanIntervalTime是清除goroutines的间隔时间。
	*/
	DefaultCleanIntervalTime = time.Second
)

const (
	// OPENED represents that the pool is opened.
	/**
	* open表示池已打开。
	*/
	OPENED = iota

	// CLOSED represents that the pool is closed.
	/**
	* open表示池已关闭。
	 */
	CLOSED
)

var (
	// Error types for the Ants API.
	//---------------------------------------------------------------------------

	// ErrInvalidPoolSize will be returned when setting a negative number as pool capacity, this error will be only used
	// by pool with func because pool without func can be infinite by setting up a negative capacity.
	/**
	* 当设置一个负数作为池容量时，将返回ErrInvalidPoolSize，此错误将仅被使用
	* 通过有func的池，因为没有func的池通过设置一个负容量可以是无限的。
	*/
	ErrInvalidPoolSize = errors.New("invalid size for pool")

	// ErrLackPoolFunc will be returned when invokers don't provide function for pool.
	/**
	* 当调用器没有为池提供函数时，将返回ErrLackPoolFunc。
	*/
	ErrLackPoolFunc = errors.New("must provide function for pool")

	// ErrInvalidPoolExpiry will be returned when setting a negative number as the periodic duration to purge goroutines.
	/**
	* 当设置一个负数作为清除goroutines的周期时间时，将返回ErrInvalidPoolExpiry。
	*/
	ErrInvalidPoolExpiry = errors.New("invalid expiry for pool")

	// ErrPoolClosed will be returned when submitting task to a closed pool.
	/**
	* ErrPoolClosed将在向已关闭的池提交任务时返回。
	*/
	ErrPoolClosed = errors.New("this pool has been closed")

	// ErrPoolOverload will be returned when the pool is full and no workers available.
	/**
	*  ErrPoolOverload将在池满且没有worker可用时返回。
	*/
	ErrPoolOverload = errors.New("too many goroutines blocked on submit or Nonblocking is set")

	// ErrInvalidPreAllocSize will be returned when trying to set up a negative capacity under PreAlloc mode.
	/**
	*  当尝试在预分配模式下设置负容量时，将返回ErrInvalidPreAllocSize。
	*/
	ErrInvalidPreAllocSize = errors.New("can not set up a negative capacity under PreAlloc mode")

	//---------------------------------------------------------------------------

	// workerChanCap determines whether the channel of a worker should be a buffered channel
	// to get the best performance. Inspired by fasthttp at
	// https://github.com/valyala/fasthttp/blob/master/workerpool.go#L139
	/**
	* 检测当前最大核心数
	*/
	workerChanCap = func() int {
		// Use blocking channel if GOMAXPROCS=1.
		// This switches context from sender to receiver immediately,
		// which results in higher performance (under go1.5 at least).
		if runtime.GOMAXPROCS(0) == 1 {
			return 0
		}

		// Use non-blocking workerChan if GOMAXPROCS>1,
		// since otherwise the sender might be dragged down if the receiver is CPU-bound.
		return 1
	}()

	defaultLogger = Logger(log.New(os.Stderr, "", log.LstdFlags))

	// Init a instance pool when importing ants.
	/**
	* 在导入ant时初始化一个实例池。此时 变量，会优先执行
	*/
	defaultAntsPool, _ = NewPool(DefaultAntsPoolSize)
)

// Logger is used for logging formatted messages.
type Logger interface {
	// Printf must have the same semantics as log.Printf.
	Printf(format string, args ...interface{})
}

// Submit submits a task to pool.
func Submit(task func()) error {
	return defaultAntsPool.Submit(task)
}

// Running returns the number of the currently running goroutines.
/**
* 获取当前的协程数
*/
func Running() int {
	return defaultAntsPool.Running()
}

// Cap returns the capacity of this default pool.
/**
* 协程池的大小容量
*/
func Cap() int {
	return defaultAntsPool.Cap()
}

// Free returns the available goroutines to work.
/**
* 有多少可以工作的协程
**/
func Free() int {
	return defaultAntsPool.Free()
}

// Release Closes the default pool.
/**
* 优雅关闭协程池
*/
func Release() {
	defaultAntsPool.Release()
}

// Reboot reboots the default pool.
/**
* 重制协程池
**/
func Reboot() {
	defaultAntsPool.Reboot()
}
