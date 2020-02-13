module github.com/lemon-cloud-service/lemon-cloud-wechat/lemon-cloud-wechat-service

go 1.13

require (
	github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components v0.0.0-00010101000000-000000000000
	github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-define v0.0.0-00010101000000-000000000000
	github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-model v0.0.0-20200208132904-9f1e2968067f
	github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils v0.0.0-20200206055610-d888d2092c7c
	github.com/sirupsen/logrus v1.4.2
)

replace github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-define => ../../lemon-cloud-common/lemon-cloud-common-define

replace github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils => ../../lemon-cloud-common/lemon-cloud-common-utils

replace github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-model => ../../lemon-cloud-common/lemon-cloud-common-model

replace github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components => ../../lemon-cloud-common/lemon-cloud-common-components

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
