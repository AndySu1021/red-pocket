module red-packet

go 1.16

require (
	github.com/bsm/redislock v0.7.2
	github.com/cenkalti/backoff/v4 v4.1.2
	github.com/gin-contrib/pprof v1.3.0
	github.com/gin-gonic/gin v1.7.7
	github.com/go-redis/redis/v8 v8.11.5
	github.com/golang/mock v1.6.0
	github.com/kr/pretty v0.3.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/rs/xid v1.3.0
	github.com/rs/zerolog v1.26.1
	github.com/spf13/cobra v1.4.0
	github.com/spf13/viper v1.10.1
	go.uber.org/fx v1.17.0
	golang.org/x/sys v0.0.0-20220317061510-51cd9980dadf // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gorm.io/driver/mysql v1.3.2
	gorm.io/driver/sqlite v1.3.1
	gorm.io/gorm v1.23.2
)
