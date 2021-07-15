package init

import (
	"log"
	"os"
)

func init(){
	if _,err := os.Stat("config") ; err != nil {
		ApplicationDirectoryInitialization()
	}
}

// ApplicationDirectoryInitialization App init
func ApplicationDirectoryInitialization() {
	type pair struct {
		key  string
		desc string
	}

	var dict []pair

	dict = append(dict, pair{"asset", "前端资源目录"})
	dict = append(dict, pair{"asset/css", "前端资源css"})
	dict = append(dict, pair{"asset/icon", "前端资源icon"})
	dict = append(dict, pair{"asset/image", "前端资源image"})
	dict = append(dict, pair{"asset/js", "前端资源js"})
	dict = append(dict, pair{"config", "配置目录"})
	dict = append(dict, pair{"controller", "控制器"})
	dict = append(dict, pair{"database", "数据库"})
	dict = append(dict, pair{"database/sql", "数据库sql语句"})
	dict = append(dict, pair{"log", "日志"})
	dict = append(dict, pair{"middleware", "中间件"})
	dict = append(dict, pair{"model", "模型"})
	dict = append(dict, pair{"restful", "restful接口"})
	dict = append(dict, pair{"restful/v1", "版本1接口"})
	dict = append(dict, pair{"restful/v2", "版本2接口"})
	dict = append(dict, pair{"restful/v3", "版本3接口"})
	dict = append(dict, pair{"router", "url路由"})
	dict = append(dict, pair{"tool", "工具类"})
	dict = append(dict, pair{"validation", "验证"})
	dict = append(dict, pair{"view", "视图"})
	dict = append(dict, pair{"view/template", "视图模板"})
	dict = append(dict, pair{"view/admin", "视图后台"})
	dict = append(dict, pair{"view/html", "视图前端"})

	for _,v:= range dict {
		if os.Mkdir(v.key,os.ModeDir) == nil {
			log.Println(v.key,v.desc,"创建成功")
		}else{
			log.Println(v.key,v.desc,"创建失败")
		}
	}
}

