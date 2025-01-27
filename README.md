# todo_list

# 项目初始化

go mod init todo_list 初始化一个新的Go模块，当前目录创建go.mod文件。

go mod tidy 扫描项目的源文件，找出所有的导入语句，包括直接和间接导入的包，自动维护go.mod文件中的依赖信息。

# 项目配置和数据库连接

go get gopkg.in/ini.v1 获取用于解析 INI 格式配置文件的 Go 库。

go get github.com/jinzhu/gorm 下载并安装gorm包

go get github.com/gin-gonic/gin 下载并安装gin包

go get github.com/jinzhu/gorm/dialects/mysql 下载并安装gorm库及其 MySQL 方言支持包