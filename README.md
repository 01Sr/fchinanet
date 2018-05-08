---
typora-copy-images-to: ipic
---

Fchinanet
===

[![GitHub repo size in bytes](https://img.shields.io/github/repo-size/01sr/fchinanet.svg)](https://github.com/01sr/fchinanet)
[![Github All Releases](https://img.shields.io/github/downloads/01sr/fchinanet/total.svg)](http://github.com/01sr/fchinanet/releases) 
[![GitHub release](https://img.shields.io/github/release/01sr/fchinanet.svg)](http://github.com/01sr/fchinanet/releases)
[![GitHub issues](https://img.shields.io/github/issues/01sr/fchinanet.svg)](https://github.com/01sr/fchinanet/issues)

Fchinanet是一款用于江苏、安徽等地电信校园上网的工具，接口由逆向掌上大学所得，如果你的所在地的电信宽带是使用掌上大学APP扫PC端二维码登录，那么你可以尝试下本工具。

### 使用方式

**[界面程序](https://github.com/01Sr/fchinanetUI/releases) (https://github.com/01Sr/fchinanetUI/releases)** **[win用户操作参考](https://github.com/01Sr/fchinanet/issues/9)**

- 帮助 `./fchinanet -h`

![mage-20180408152436](http://osxhu29uq.bkt.clouddn.com/img/2018-04-08-image-201804081524360.png)

- 简单的登录示例 `./fchinanet -a account -p password `

![mage-20180408152819](http://osxhu29uq.bkt.clouddn.com/img/2018-04-08-image-201804081528194.png)

- 下线`./fchinanet -a account -p passwd -b 0`

![mage-20180408201008](http://osxhu29uq.bkt.clouddn.com/img/2018-04-08-2018-04-08-image-201804082010087.png)

- **在线设备列表、多设备登录**等更多功能请参见第一条的帮助


- 你可以在windows\linux\macos上创建脚本来方便的执行命令, 如果你愿意可以用crontab命令设置定时执行或加入开机启动项

### 编译

本工具用Go编写，本地Go版本为1.8.3。编译前你需要准备：

- 安装Go 1.8+(低版本未测试，低版本部分ARM架构的编译不支持)
- 配置Go环境变量
- 安装依赖: `go get github.com/fatih/color`
- 编译所需平台的可执行文件，可参考 [build.sh](https://github.com/01Sr/fchinanet/blob/master/build.sh)

### 其他版本(不保证同步更新)

[Android](https://github.com/01Sr/FChinanetAndroid) 

[Node.js](https://github.com/Anapopo/FChinaNet.js)

[Shell](https://github.com/Anapopo/FChinaNet.sh)，此版本比较轻量，适合路由器上使用，**另外感谢 [Anapopo](https://github.com/Anapopo)老哥提供这个版本**

