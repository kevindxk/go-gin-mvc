# gin框架开发api接口
## 框架结构搭建
- 初始化项目
    - go mod init project
- 安装依赖
    - go get -u github.com/gin-gonic/gin
    
## gin中整合swagger
- !!! 使用swagger必须和main.go在同一个文件下，而且执行swag init 生成的文件docs文件必须在项目的根目录下
- 安装依赖
  - go get -u github.com/swaggo/swag/cmd/swag
  - 执行命令
    - 在main.go平级下执行命令swag init
    - 下载gin-swagger
      - go get -u github.com/swaggo/gin-swagger
      - go get -u github.com/swaggo/files
    -导入包，在main函数中配置swagger
    - 在函数上写上注释
    - 最后执行swag init


 
