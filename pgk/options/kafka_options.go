package options

import (
	"github.com/segmentio/kafka-go"
	"time"
)

type WriterOptions struct {
	// 发送消息的最大尝试次数 默认为 10
	MaxAttempts int `mapstructure:"max-attempts"`
	// 副本分区的确认 默认为 -1
	RequiredAcks int `mapstructure:"require-acks"`
	// 设置是否异步发送消息 默认为 false
	Async bool `mapstructure:"async"`
	// 批量发送消息的大小 默认为 100
	BatchSize int `mapstructure:"batch-size"`
	// 批量发送消息的超时时间 默认为 1秒
	BatchTimeout time.Duration `mapstructure:"batch-timeout"`
	// 最大批量发送消息的字节数 默认为 1048576
	BatchBytes int `mapstructure:"batch-byte"`
}

type ReaderOptions struct {
	GroupID   string `mapstructure:"group-id"`
	Partition int    `mapstructure:"partition"`
	// 内部消息队列容量 默认为100
	QueueCapacity int `mapstructure:"queue-capacity"`
	MinBytes      int `mapstructure:"min-bytes"`
	// 单次读取的最大字节数 默认为 1M
	MaxBytes          int           `mapstructure:"max-bytes"`
	MaxWait           time.Duration `mapstructure:"max-wait"`
	ReadBatchTimeout  time.Duration `mapstructure:"read-batch-time"`
	HeartbeatInterval time.Duration `mapstructure:"heartbeat-interval"`
	CommitInterval    time.Duration `mapstructure:"commit-interval"`
	RebalanceTimeout  time.Duration `mapstructure:"rebalance-timeout"`
	StartOffset       int64         `mapstructure:"start-offset"`
	MaxAttempts       int           `mapstructure:"max-attempts"`
}

// KafkaOptions defines options for read cluster.
// Common options for read-go reader and writer.
type KafkaOptions struct {
	// read-go reader and writer common options
	Brokers       []string      `mapstructure:"brokers"`
	Topic         string        `mapstructure:"topic"`
	ClientID      string        `mapstructure:"client-id"`
	Timeout       time.Duration `mapstructure:"timeout"`
	SASLMechanism string        `mapstructure:"mechanism"`
	Username      string        `mapstructure:"username"`
	Password      string        `mapstructure:"password"`
	Algorithm     string        `mapstructure:"algorithm"`
	Compressed    bool          `mapstructure:"compressed"`

	// read-go writer options
	WriterOptions WriterOptions `mapstructure:"writer"`
	// read-go reader options
	ReaderOptions ReaderOptions `mapstructure:"reader"`
}

// NewKafkaOptions create a `zero` value instance.
func NewKafkaOptions() *KafkaOptions {
	return &KafkaOptions{
		Brokers: []string{"127.0.0.1:9092"},
		Timeout: 3 * time.Second,
		WriterOptions: WriterOptions{
			RequiredAcks: 1,
			MaxAttempts:  10,
			Async:        true,
			BatchSize:    100,
			BatchTimeout: 1 * time.Second,
			BatchBytes:   1 * MiB,
		},
		ReaderOptions: ReaderOptions{
			QueueCapacity:     100,
			MinBytes:          1,
			MaxBytes:          1 * MiB,
			MaxWait:           10 * time.Second,
			ReadBatchTimeout:  10 * time.Second,
			HeartbeatInterval: 3 * time.Second,
			CommitInterval:    0 * time.Second,
			RebalanceTimeout:  30 * time.Second,
			StartOffset:       kafka.FirstOffset,
			MaxAttempts:       3,
		},
	}
}
