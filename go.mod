module github.com/zfinn/t4cobra

replace github.com/coreos/etcd => go.etcd.io/etcd v3.3.11+incompatible

require (
	github.com/coreos/etcd v3.3.11+incompatible
	github.com/gogo/protobuf v1.2.0 // indirect
	github.com/sirupsen/logrus v1.3.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.3.1
	golang.org/x/net v0.0.0-20190110200230-915654e7eabc // indirect
	google.golang.org/grpc v1.18.0 // indirect
	gopkg.in/yaml.v2 v2.2.2
)
