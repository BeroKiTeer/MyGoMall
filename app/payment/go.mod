module payment

go 1.23.2

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

replace github.com/BeroKiTeer/MyGoMall/common => ../../common

require (
	github.com/BeroKiTeer/MyGoMall/common v0.0.0-00010101000000-000000000000
	github.com/cloudwego/kitex v0.12.2
	github.com/durango/go-credit-card v0.0.0-20220404131259-a9e175ba4082
	github.com/google/uuid v1.6.0
	github.com/kr/pretty v0.3.1
	github.com/redis/go-redis/v9 v9.7.1
	github.com/streadway/amqp v1.1.0
	gopkg.in/validator.v2 v2.0.1
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.12
)

require (
	github.com/bytedance/gopkg v0.1.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudwego/fastpb v0.0.5 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/protobuf v1.36.3 // indirect
)
