module github.com/PaloAltoNetworks/terraform-provider-cortexcloud

go 1.24.3

require (
	dario.cat/mergo v1.0.2
	github.com/hashicorp/terraform-plugin-framework v1.15.1
	github.com/hashicorp/terraform-plugin-framework-validators v0.17.0
	github.com/hashicorp/terraform-plugin-go v0.28.0
	github.com/hashicorp/terraform-plugin-log v0.9.0
	github.com/hashicorp/terraform-plugin-testing v1.13.3
	github.com/mdboynton/cortex-cloud-go/api v0.0.0-00010101000000-000000000000
	github.com/mdboynton/cortex-cloud-go/appsec v0.0.0-00010101000000-000000000000
	github.com/mdboynton/cortex-cloud-go/cloudonboarding v0.0.0-00010101000000-000000000000
	github.com/mdboynton/cortex-cloud-go/enums v0.0.0-00010101000000-000000000000
	github.com/mdboynton/cortex-cloud-go/log v0.0.0-00010101000000-000000000000
	github.com/mdboynton/cortex-cloud-go/platform v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/ProtonMail/go-crypto v1.1.6 // indirect
	github.com/agext/levenshtein v1.2.2 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/cloudflare/circl v1.6.1 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-checkpoint v0.5.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-cty v1.5.0 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-plugin v1.6.3 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/hashicorp/hc-install v0.9.2 // indirect
	github.com/hashicorp/hcl/v2 v2.23.0 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/terraform-exec v0.23.0 // indirect
	github.com/hashicorp/terraform-json v0.25.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.37.0 // indirect
	github.com/hashicorp/terraform-registry-address v0.2.5 // indirect
	github.com/hashicorp/terraform-svchost v0.1.1 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mdboynton/cortex-cloud-go v0.0.0-20250530204549-8c630f4f6da1 // indirect
	github.com/mdboynton/cortex-cloud-go/errors v0.0.0-00010101000000-000000000000 // indirect
	github.com/mdboynton/cortex-cloud-go/internal/app v0.0.0-00010101000000-000000000000 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/oklog/run v1.0.0 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/zclconf/go-cty v1.16.3 // indirect
	golang.org/x/crypto v0.39.0 // indirect
	golang.org/x/mod v0.25.0 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sync v0.15.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	golang.org/x/tools v0.33.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
	google.golang.org/grpc v1.72.1 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace github.com/mdboynton/cortex-cloud-go => ../cortex-cloud-go

replace github.com/mdboynton/cortex-cloud-go/internal/app => ../cortex-cloud-go/internal/app

replace github.com/mdboynton/cortex-cloud-go/api => ../cortex-cloud-go/api

replace github.com/mdboynton/cortex-cloud-go/errors => ../cortex-cloud-go/errors

replace github.com/mdboynton/cortex-cloud-go/enums => ../cortex-cloud-go/enums

replace github.com/mdboynton/cortex-cloud-go/log => ../cortex-cloud-go/log

replace github.com/mdboynton/cortex-cloud-go/appsec => ../cortex-cloud-go/appsec

replace github.com/mdboynton/cortex-cloud-go/cloudonboarding => ../cortex-cloud-go/cloudonboarding

replace github.com/mdboynton/cortex-cloud-go/platform => ../cortex-cloud-go/platform
