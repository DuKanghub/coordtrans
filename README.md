# 使用说明
## 编译
```shell
go build
```
## 使用
```shell
实现一些坐标转换，
目前支持以下坐标系转换：
- GPS坐标与百度系坐标互转：WGS84与BD09(bd09ii)
- GPS坐标与火星系坐标互转：WGS84与GCJ02，即高德坐标或腾讯坐标
- 百度坐标与火星系坐标互转：BD09(bd09ii)与GCJ02

使用方法：
coordtrans [-m <method>] [-f <from>] [-t <to>] [-o <outPut>] [-a <ak>] 经度,纬度
多个坐标用空格分隔，如：经度,纬度 经度,纬度
注意：百度api暂不支持目标坐标系为GPS坐标。

Usage:
  coordtrans [flags]

Flags:
  -k, --ak string       私钥，非必须，默认为空，如果使用百度接口，则必传
      --config string   config file (default is $HOME/.coordtrans.yaml)
  -f, --from int        源坐标系，即传入的坐标系类型。非必须，默认为1，可选值：1, 3, 5 (default 1)
  -h, --help            help for coordtrans
  -m, --method string   接口模式，非必须，默认为mod，可选值：mod、bd (default "mod")
  -o, --output string   将结果保存到指定目录下，非必须，默认为空
  -t, --to int          目标坐标系，即需要转换成的坐标系类型。非必须，默认为5，可选值：1, 3, 5 (default 5)
```
坐标系统数字说明：
```shell
1: WGS84坐标系，国际标准，GPS设备采用的坐标系统（WGS84）
3: GCJ02坐标系，火星坐标系统，即高德地图，腾讯地图等使用的。
5: BD09坐标系，百度坐标系统，即百度地图使用的。
```
使用模块方法：
```shell
coordtrans -f 1 -t 5 116.405467,39.907761
# 多个坐标用空格分隔，如：经度,纬度 经度,纬度
coordtrans -f 1 -t 5 116.405467,39.907761 116.405467,39.907761
```
使用百度接口方法：
```shell
coordtrans -f 1 -t 5 -m bd -k ak 116.405467,39.907761
# ak为百度开发者密钥
```
`-f -t`参数可以省略，默认为1，5。
# 问题须知
- 百度接口暂不支持目标系为WGS84的转换。