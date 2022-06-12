package models

import (
	"io"
	"log"
	"text/template"
	"time"
)

type TemplateBlog struct {
	*template.Template // 取得模板地址
}
type HtmlTemplate struct {
	Category   TemplateBlog
	Custom     TemplateBlog
	Detail     TemplateBlog
	Index      TemplateBlog
	Login      TemplateBlog
	Pigeonhole TemplateBlog
	Writing    TemplateBlog
}

func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
	}
}

func (t *TemplateBlog) WriteError(w io.Writer, err error) {
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
			return
		}
	}
}
func InitTemplate(templateDir string) (HtmlTemplate, error) {
	tp, err := readTemplate(
		[]string{"category", "custom", "detail", "index", "login", "pigeonhole", "writing"},
		templateDir,
	)
	var htmlTemplate HtmlTemplate
	if err != nil {
		return htmlTemplate, err
	}
	htmlTemplate.Category = tp[0]
	htmlTemplate.Custom = tp[1]
	htmlTemplate.Detail = tp[2]
	htmlTemplate.Index = tp[3]
	htmlTemplate.Login = tp[4]
	htmlTemplate.Pigeonhole = tp[5]
	htmlTemplate.Writing = tp[6]
	return htmlTemplate, nil
}

func IsODD(num int) bool {
	return num%2 == 0
}
func GetNextName(strs []string, index int) string {
	return strs[index+1]
}
func Date(layout string) string {
	return time.Now().Format(layout)
}
func DateDay(date time.Time) string {
	// 格式化操作
	return date.Format("2006-01-02 15:04:05")
}
func readTemplate(templates []string, templateDir string) ([]TemplateBlog, error) {
	// 传入想要读取的 模板:html文件名 和 目录地址，返回
	var tbs []TemplateBlog
	for _, view := range templates {
		viewName := view + ".html"
		t := template.New(viewName)
		//访问博客首页模板的时候，因为有多个模板的嵌套，解析文件的时候，需要将其涉及到的所有模板都进行解析
		home := templateDir + "home.html"
		header := templateDir + "layout/header.html"
		footer := templateDir + "layout/footer.html"
		personal := templateDir + "layout/personal.html"
		post := templateDir + "layout/post-list.html"
		pagination := templateDir + "layout/pagination.html"
		// 补充文件内未定义的函数方法
		t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date, "dateDay": DateDay})
		// ParseFiles 解析模板
		t, err := t.ParseFiles(templateDir+viewName, home, header, footer, personal, post, pagination)
		if err != nil {
			log.Println("解析模板出错：", err)
			return nil, err
		}
		var tb TemplateBlog
		tb.Template = t
		tbs = append(tbs, tb)
	}
	println(tbs)
	return tbs, nil
}
