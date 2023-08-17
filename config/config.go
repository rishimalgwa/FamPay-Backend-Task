package config

import "github.com/spf13/viper"

var (
	YT_API_KEY1 = ""
	YT_API_KEY2 = ""
	YT_API_KEY3 = ""
	YTApiKeys   = []string{}
)

func LoadConfig() {
	YT_API_KEY1 = viper.GetString("YT_API_KEY1")
	YT_API_KEY2 = viper.GetString("YT_API_KEY2")
	//YT_API_KEY3 = viper.GetString("YT_API_KEY3")
	YTApiKeys = append(YTApiKeys, YT_API_KEY1)
	YTApiKeys = append(YTApiKeys, YT_API_KEY2)
	//YTApiKeys = append(YTApiKeys, YT_API_KEY3)
}
