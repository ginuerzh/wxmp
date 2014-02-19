// main
package main

import (
	"github.com/ginuerzh/weixin/mp"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	wx := mp.New("wxb6ebfdef79a09651", "abd7052dc840112a9fbffe39e6fbeaca", "1234567890")
	wx.Init("/")

	wx.HandleFunc(mp.MsgTypeText, func(reply mp.MessageReplyer, m *mp.Message) {
		log.Println("receive message:" + m.Content)
		if err := reply.ReplyText("今日自驾: " + m.Content); err != nil {
			log.Println(err)
		}
	})
	wx.Run(8081)
}
