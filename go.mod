module github.com/brocaar/loraserver

go 1.12

require (
	github.com/brocaar/lorawan v0.0.0-20190523144945-4c051b1fa597
	github.com/eclipse/paho.mqtt.golang v1.2.0
	github.com/garyburd/redigo v1.6.0
	github.com/gobuffalo/packr v1.25.0 // indirect
	github.com/golang/protobuf v1.3.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.1.1
	github.com/pkg/errors v0.8.1
	github.com/rubenv/sql-migrate v0.0.0-20190327083759-54bad0a9b051
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.4.2
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a
	github.com/spf13/cobra v0.0.4
	github.com/spf13/viper v1.4.0
	github.com/ziutek/mymysql v1.5.4 // indirect
	golang.org/x/net v0.0.0-20190603091049-60506f45cf65
	golang.org/x/sys v0.0.0-20190602015325-4c4f7f33c9ed // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/grpc v1.21.1
	gopkg.in/gorp.v1 v1.7.2 // indirect
)

replace github.com/brocaar/lorawan => github.com/hobairiku/lorawan v0.0.1-riv-custom
