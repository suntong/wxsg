package main

import (
	"fmt"

	"github.com/skip2/go-qrcode"
	"github.com/eatMoreApple/openwechat"
)

func main() {
	bot := openwechat.DefaultBot(openwechat.Desktop)
	// 注册登陆二维码回调
	bot.UUIDCallback = ConsoleQrCode

	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() {
			fmt.Println("你收到了一条新的文本消息")
		}
	}

	// 创建热存储容器对象
	reloadStorage := openwechat.NewJsonFileHotReloadStorage("storage.json")

	// 执行热登陆
	if err := bot.HotLogin(reloadStorage); err != nil {
		fmt.Println(err)
		return
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取所有的好友
	friends, err := self.Friends()
	fmt.Println(friends, err)

	// 获取所有的群组
	groups, err := self.Groups()
	fmt.Println(groups, err)

	// 阻塞主goroutine, 知道发生异常或者用户主动退出
	bot.Block()
}


func ConsoleQrCode(uuid string) {
	q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Low)
	fmt.Println(q.ToString(true))
}
