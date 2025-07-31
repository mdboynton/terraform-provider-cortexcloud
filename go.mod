module github.com/PaloAltoNetworks/terraform-provider-cortexcloud

go 1.24.3

require (
	dario.cat/mergo v1.0.2
	github.com/hashicorp/terraform-plugin-framework v1.14.1
	github.com/hashicorp/terraform-plugin-framework-validators v0.17.0
	github.com/hashicorp/terraform-plugin-log v0.9.0
	github.com/mdboynton/cortex-cloud-go/api v0.0.0-00010101000000-000000000000
	github.com/mdboynton/cortex-cloud-go/appsec v0.0.0-00010101000000-000000000000
	github.com/mdboynton/cortex-cloud-go/cloudonboarding v0.0.0-00010101000000-000000000000
	github.com/mdboynton/cortex-cloud-go/enums v0.0.0-00010101000000-000000000000
	github.com/mdboynton/cortex-cloud-go/log v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/fatih/color v1.16.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-plugin v1.6.2 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/terraform-plugin-go v0.26.0 // indirect
	github.com/hashicorp/terraform-registry-address v0.2.4 // indirect
	github.com/hashicorp/terraform-svchost v0.1.1 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mdboynton/cortex-cloud-go v0.0.0-20250530204549-8c630f4f6da1 // indirect
	github.com/mdboynton/cortex-cloud-go/internal/app v0.0.0-00010101000000-000000000000 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/oklog/run v1.0.0 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
	google.golang.org/grpc v1.69.4 // indirect
	google.golang.org/protobuf v1.36.3 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace github.com/mdboynton/cortex-cloud-go => ../cortex-cloud-go

replace github.com/mdboynton/cortex-cloud-go/internal/app => ../cortex-cloud-go/internal/app

replace github.com/mdboynton/cortex-cloud-go/api => ../cortex-cloud-go/api

replace github.com/mdboynton/cortex-cloud-go/enums => ../cortex-cloud-go/enums

replace github.com/mdboynton/cortex-cloud-go/log => ../cortex-cloud-go/log

replace github.com/mdboynton/cortex-cloud-go/appsec => ../cortex-cloud-go/appsec

replace github.com/mdboynton/cortex-cloud-go/cloudonboarding => ../cortex-cloud-go/cloudonboarding

replace github.com/mdboynton/cortex-cloud-go/xsiam => ../cortex-cloud-go/xsiam
