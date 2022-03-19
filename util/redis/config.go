package redis

// Config ...
type Config struct {
	ClusterMode     bool     `mapstructure:"cluster_mode"`
	Addresses       []string `mapstructure:"addresses"`
	Password        string   `mapstructure:"password"`
	MaxRetries      int      `mapstructure:"max_retries"`
	PoolSizePerNode int      `mapstructure:"pool_size_per_node"`
	DB              int      `mapstructure:"db"`
}
