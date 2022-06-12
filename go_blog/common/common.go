package common

import (
	"go_blog/config"
	"go_blog/models"
	"sync"
)

var Template models.HtmlTemplate // 将模板地址定义为 Template

func LoadTemplate() {
	w := sync.WaitGroup{} // 声明一个等待组，然后给等待组的计数器 +1
	w.Add(1)
	go func() {
		var err error
		// 将目录下template的地址传到InitTemplate函数中
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}
		w.Done()
	}()
	w.Wait()
}
