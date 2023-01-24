module github.com/Canto-Network/canto/v3

go 1.16

require (
	github.com/armon/go-metrics v0.4.0
	github.com/cosmos/cosmos-sdk v0.46.8
	github.com/cosmos/go-bip39 v1.0.0
	github.com/cosmos/ibc-go/v3 v3.0.0
	github.com/ethereum/go-ethereum v1.10.17
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/onsi/ginkgo/v2 v2.1.4
	github.com/onsi/gomega v1.20.0
	github.com/pkg/errors v0.9.1
	github.com/rakyll/statik v0.1.7
	github.com/spf13/cast v1.5.0
	github.com/spf13/cobra v1.6.0
	github.com/stretchr/testify v1.8.1
	github.com/tendermint/tendermint v0.34.24
	github.com/tendermint/tm-db v0.6.7
	github.com/tharsis/ethermint v0.14.0
	go.opencensus.io v0.23.0
	google.golang.org/genproto v0.0.0-20221014213838-99cd37c6964a
	google.golang.org/grpc v1.50.1
	google.golang.org/protobuf v1.28.2-0.20220831092852-f930b1dc76e8
)

require (
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/prometheus/tsdb v0.10.0 // indirect
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/rjeczalik/notify v0.9.2 // indirect
	gopkg.in/yaml.v2 v2.4.0
	nhooyr.io/websocket v1.8.7 // indirect
)

replace (
	github.com/99designs/keyring => github.com/cosmos/keyring v1.1.7-0.20210622111912-ef00f8ac3d76
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/tharsis/ethermint => github.com/Canto-Network/ethermint v0.0.0-beta
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
