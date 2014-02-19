// main
package main

import (
	"bytes"
	"github.com/ginuerzh/weixin/mp"
	"io/ioutil"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	wx := mp.New("wxb6ebfdef79a09651", "abd7052dc840112a9fbffe39e6fbeaca", "1234567890")
	wx.Init("/")

	wx.HandleFunc(mp.MsgTypeText, func(w mp.MessageReplyer, m *mp.Message) {
		log.Println("receive message:" + m.Content)

		switch mp.MsgType(m.Content) {
		case mp.MsgTypeImage:
			b, err := ioutil.ReadFile("image.jpg")
			if err != nil {
				log.Println(err)
				break
			}
			imageId, err := wx.UploadMedia(mp.MediaImage, "image.jpg", bytes.NewBuffer(b))
			if err != nil {
				log.Println(err)
				break
			}
			if err := w.ReplyImage(imageId); err != nil {
				log.Println(err)
			}
		case mp.MsgTypeNews:
			articles := make([]mp.Article, 2)
			articles[0].Title = "自驾3500公里 来到大理 丽江 感受云南的自然风光"
			articles[0].Description = "自驾3500公里 来到大理 丽江 感受云南的自然风光"
			articles[0].PicUrl = "http://106.187.48.51:8081/6/2710b07cffba_1/640X853.jpg"
			articles[0].Url = "http://club.autohome.com.cn/bbs/threadowner-o-200042-27866947-1.html#pvareaid=101435"

			articles[1].Title = "秀才过双节，看价值两亿的树桩桩！"
			articles[1].Description = "秀才过双节，看价值两亿的树桩桩！"
			articles[1].PicUrl = "http://106.187.48.51:8081/7/2791b941c5bc_1/640X480.jpg"
			articles[1].Url = "http://club.autohome.com.cn/bbs/threadowner-o-200042-27866796-1.html#pvareaid=101435"

			if err := w.ReplyImageText(articles); err != nil {
				log.Println(err)
			}
		case mp.MsgTypeText:
			fallthrough
		default:
			if err := w.ReplyText("今日自驾: " + m.Content); err != nil {
				log.Println(err)
			}
		}
	})
	wx.Run(8801)
}
