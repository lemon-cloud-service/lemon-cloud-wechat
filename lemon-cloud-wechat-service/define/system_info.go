package define

import (
	"fmt"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-define/lccd_strings"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-model/lccm_config"
)

const SYSTEM_INFO_NAME string = "Lemon Cloud Wechat Service"
const SYSTEM_INFO_SERVICE_TAG string = "lemon_cloud_wechat"
const SYSTEM_INFO_SERVICE_INTRODUCE string = "Lemon Cloud Wechat Service"
const SYSTEM_INFO_VERSION string = "1.0.0"
const SYSTEM_INFO_SPLIT_LINE string = "====================================================================="

func GetServiceInfo() *lccm_config.ServiceInfo {
	return &lccm_config.ServiceInfo{
		ServiceTag:       SYSTEM_INFO_SERVICE_TAG,
		ServiceName:      SYSTEM_INFO_NAME,
		ServiceIntroduce: SYSTEM_INFO_SERVICE_INTRODUCE,
	}
}

// 打印系统的基础信息，包含LemonCloud字符画和系统名称及版本
func PrintSystemInfo() {
	fmt.Print(lccd_strings.LEMON_CLOUD_ASCII_IMAGE)
	fmt.Print("\n")
	fmt.Println(SYSTEM_INFO_SPLIT_LINE)
	fmt.Printf("Welcome to %v [ver: %v]\n", SYSTEM_INFO_NAME, SYSTEM_INFO_VERSION)
	fmt.Println(SYSTEM_INFO_SPLIT_LINE)
}
