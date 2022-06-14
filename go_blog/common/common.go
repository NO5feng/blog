package common

import (
	"encoding/json"
	"go_blog/config"
	"go_blog/models"
	"io/ioutil"
	"log"
	"net/http"
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

func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body) //读取Request的body
	_ = json.Unmarshal(body, &params) //解析json文件body,传递到数组params中
	return params
}

func Success(w http.ResponseWriter, data interface{}) {
	//接受用户名和密码 返回 对应的数据
	var result models.Result //定义一个返回值
	result.Code = 200
	result.Error = ""
	result.Data = data
	resultJson, _ := json.Marshal(result) //将result 转换为json数据类型
	//在Go语言中，客户端请求信息都封装到了Request对象，但是发送给客户端的响应并不是Response对象，而是ResponseWriter
	w.Header().Set("Content-Type", "application/json") //设置响应头,告诉Response,返回的是json数据
	_, err := w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}

func Error(w http.ResponseWriter, err error) {
	//接受用户名和密码 返回 对应的数据
	var result models.Result //定义一个返回值
	result.Code = -808
	result.Error = err.Error()
	resultJson, _ := json.Marshal(result) //将result 转换为json数据类型
	//在Go语言中，客户端请求信息都封装到了Request对象，但是发送给客户端的响应并不是Response对象，而是ResponseWriter
	w.Header().Set("Content-Type", "application/json") //设置响应头,告诉Response,返回的是json数据
	_, err = w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}
