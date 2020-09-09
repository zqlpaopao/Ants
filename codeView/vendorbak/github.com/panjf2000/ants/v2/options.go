package ants

import "time"

// Option represents the optional function.
type Option func(opts *Options)

/**
* 加载 基础配置项 过期清理时间 最大任务数等
*/
func loadOptions(options ...Option) *Options {
	opts := new(Options)//&ants.Options{ExpiryDuration:0, PreAlloc:false, MaxBlockingTasks:0, Nonblocking:false, PanicHandler:(func(interface {}))(nil), Logger:ants.Logger(nil)}exit status 12
	//options 为[]ants.Option(nil)  不会进行 for循环
	for _, option := range options {
		option(opts)
	}
	return opts
}

// Options contains all options which will be applied when instantiating a ants pool.
type Options struct {
	// ExpiryDuration is a period for the scavenger goroutine to clean up those expired workers,
	// the scavenger scans all workers every `ExpiryDuration` and clean up those workers that haven't been
	// used for more than `ExpiryDuration`.
	/**
	* 定期清楚过期的groutine
	*/
	ExpiryDuration time.Duration

	// PreAlloc indicates whether to make memory pre-allocation when initializing Pool.
	/**
	* PreAlloc表示在初始化池时是否进行内存预分配。
	*/
	PreAlloc bool

	// Max number of goroutine blocking on pool.Submit.
	// 0 (default value) means no such limit.
	/**
	* 最大数量的goroutine阻塞池。提交。
	* 0(默认值)表示没有这种限制。
	*/
	MaxBlockingTasks int

	// When Nonblocking is true, Pool.Submit will never be blocked.
	// ErrPoolOverload will be returned when Pool.Submit cannot be done at once.
	// When Nonblocking is true, MaxBlockingTasks is inoperative.
	/**
	* 当非阻塞为真时，池。提交永远不会被阻止。
	* ErrPoolOverload将在池中返回。不能马上提交。
	* 当非阻塞为真时，MaxBlockingTasks不起作用。
	*/
	Nonblocking bool

	// PanicHandler is used to handle panics from each worker goroutine.
	// if nil, panics will be thrown out again from worker goroutines.
	/**
	* PanicHandler用于处理来自每个工人的恐慌。
	* 如果为零，恐慌将再次从工人goroutines抛出。
	*/
	PanicHandler func(interface{})

	// Logger is the customized logger for logging info, if it is not set,
	// default standard logger from log package is used.
	/**
	* Logger是用于记录信息的自定义日志记录器，如果没有设置，
	* 使用来自日志包的默认标准日志记录器。
	*/
	Logger Logger
}

// WithOptions accepts the whole options config.
func WithOptions(options Options) Option {
	return func(opts *Options) {
		*opts = options
	}
}

// WithExpiryDuration sets up the interval time of cleaning up goroutines.
func WithExpiryDuration(expiryDuration time.Duration) Option {
	return func(opts *Options) {
		opts.ExpiryDuration = expiryDuration
	}
}

// WithPreAlloc indicates whether it should malloc for workers.
func WithPreAlloc(preAlloc bool) Option {
	return func(opts *Options) {
		opts.PreAlloc = preAlloc
	}
}

// WithMaxBlockingTasks sets up the maximum number of goroutines that are blocked when it reaches the capacity of pool.
func WithMaxBlockingTasks(maxBlockingTasks int) Option {
	return func(opts *Options) {
		opts.MaxBlockingTasks = maxBlockingTasks
	}
}

// WithNonblocking indicates that pool will return nil when there is no available workers.
func WithNonblocking(nonblocking bool) Option {
	return func(opts *Options) {
		opts.Nonblocking = nonblocking
	}
}

// WithPanicHandler sets up panic handler.
func WithPanicHandler(panicHandler func(interface{})) Option {
	return func(opts *Options) {
		opts.PanicHandler = panicHandler
	}
}

// WithLogger sets up a customized logger.
func WithLogger(logger Logger) Option {
	return func(opts *Options) {
		opts.Logger = logger
	}
}
