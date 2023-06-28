module github.com/meshplus/premo

go 1.15

require (
	github.com/Rican7/retry v0.3.1
	github.com/cheynewallace/tabby v1.1.1
	github.com/cloudflare/cfssl v1.6.1 // indirect
	github.com/coreos/etcd v3.3.27+incompatible
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f // indirect
	github.com/ethereum/go-ethereum v1.10.4
	github.com/gin-gonic/gin v1.9.1
	github.com/gobuffalo/packd v1.0.1
	github.com/gobuffalo/packr/v2 v2.8.3
	github.com/hyperledger/fabric-lib-go v1.0.0 // indirect
	github.com/hyperledger/fabric-sdk-go v1.0.0-alpha5
	github.com/jiuhuche120/spin v0.0.0-20220224085816-ab11eb007ff8
	github.com/libp2p/go-libp2p-core v0.6.1
	github.com/looplab/fsm v0.3.0
	github.com/meshplus/bitxhub-core v1.3.1-0.20220810080906-e29855b55db6
	github.com/meshplus/bitxhub-kit v1.2.1-0.20220610092251-68e35ef37b7a
	github.com/meshplus/bitxhub-model v1.2.1-0.20220811053903-7cd82d10a8ec
	github.com/meshplus/go-bitxhub-client v1.4.1-0.20220719031106-9a0498f18141
	github.com/mitchellh/go-homedir v1.1.0
	github.com/shirou/gopsutil v3.21.11+incompatible
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.3.0 // indirect
	github.com/spf13/viper v1.10.1
	github.com/stretchr/testify v1.8.3
	github.com/urfave/cli/v2 v2.3.0
	github.com/wcharczuk/go-chart/v2 v2.1.0
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	golang.org/x/image v0.0.0-20220413100746-70e8d0d3baa9 // indirect
)

replace github.com/golang/protobuf => github.com/golang/protobuf v1.3.2

replace google.golang.org/grpc => google.golang.org/grpc v1.33.0

replace google.golang.org/genproto => google.golang.org/genproto v0.0.0-20200218151345-dad8c97a84f5

replace github.com/ugorji/go => github.com/ugorji/go v1.2.11
