module github.com/wrs-news/bff-api-getaway

go 1.17

require (
	github.com/99designs/gqlgen v0.17.2
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0
	github.com/vektah/gqlparser/v2 v2.4.1
	github.com/wrs-news/golang-proto v0.3.5
)

require github.com/duckpie/cherry v0.1.1

require (
	github.com/asaskevich/govalidator v0.0.0-20200108200545-475eaeb16496 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f // indirect
	golang.org/x/sys v0.0.0-20211019181941-9d821ace8654 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

require (
	github.com/BurntSushi/toml v1.0.0
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hashicorp/golang-lru v0.5.0 // indirect
	github.com/mitchellh/mapstructure v1.2.3 // indirect
	github.com/spf13/cobra v1.4.0
	github.com/stretchr/testify v1.7.1
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.45.0
)

replace google.golang.org/genproto => ./libs/genproto
