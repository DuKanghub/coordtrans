/*
Copyright © 2022 DuKang <dukang@dukanghub.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/DuKanghub/coordtrans/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile string
	method	string
	outPut string
	ak		string
	lon 	float64
	lat 	float64
	from int
	to	int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "coordtrans",
	Short: "坐标转换",
	Long: `实现一些坐标转换，
目前支持以下坐标系转换：
- GPS坐标与百度系坐标互转：WGS84与BD09(bd09ii)
- GPS坐标与火星系坐标互转：WGS84与GCJ02，即高德坐标或腾讯坐标
- 百度坐标与火星系坐标互转：BD09(bd09ii)与GCJ02
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		a := pkg.Account{
			Method: method,
			AK: ak,
		}
		transer := pkg.NewTransFormer(a)
		//fmt.Println(coordTransform.WGS84toBD09(116.404, 39.915))
		fromTo := [2]int{from, to}
		switch fromTo {
		case [2]int{1,5}:
			fmt.Println(transer.WGS84toBD09(lon,lat))
		case [2]int{5,1}:
			fmt.Println(transer.BD09toWGS84(lon,lat))
		case [2]int{1,3}:
			fmt.Println(transer.WGS84toGCJ02(lon,lat))
		case [2]int{3,1}:
			fmt.Println(transer.GCJ02toWGS84(lon,lat))
		case [2]int{5,3}:
			fmt.Println(transer.BD09toGCJ02(lon,lat))
		case [2]int{3,5}:
			fmt.Println(transer.GCJ02toBD09(lon,lat))
		default:
			fmt.Println("暂不支持该坐标转换")
		}

		//fmt.Println(transer.WGS84toBD09(116.404, 39.915))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.coordtrans.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "T", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&method, "method", "m", "mod", "接口模式，非必须，默认为mod，可选值：mod、bd")
	rootCmd.PersistentFlags().StringVarP(&ak, "ak", "k", "", "私钥，非必须，默认为空，如果使用百度接口，则必传")
	rootCmd.PersistentFlags().Float64VarP(&lon, "lon", "x", 0, "经度，必须")
	rootCmd.PersistentFlags().Float64VarP(&lat, "lat", "y", 0, "纬度，必须")
	rootCmd.PersistentFlags().IntVarP(&from, "from", "f", 1, "源坐标系，即传入的坐标系类型。非必须，默认为1，可选值：1, 3, 5")
	rootCmd.PersistentFlags().IntVarP(&to, "to", "t", 5, "目标坐标系，即需要转换成的坐标系类型。非必须，默认为5，可选值：1, 3, 5")
	rootCmd.PersistentFlags().StringVarP(&outPut, "output", "o", "", "将结果保存到指定目录下，非必须，默认为空")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".coordtrans" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".coordtrans")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
