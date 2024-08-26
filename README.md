# kratos 实战项目

## 初始化项目工具和包
```
make init
```

### 可选项（window cmd 配置）
windows系统需要使用powershell7版本，mac或linux不需要关心此项，不然相关开发命令执行会报错，你也可以手动修改makefile脚本去适配windows
* [在 Windows 上安装 PowerShell](https://learn.microsoft.com/zh-cn/powershell/scripting/install/installing-powershell-on-windows?view=powershell-7.4)
* [定制主题](https://ohmyposh.dev/docs/installation/windows)


## 开发

### 主要的命令
* 安装开发过程中需要的工具
```
make init
```
* 安装相关依赖
```
make generate
```
* 运行
```
kratos run
```

* 其他一些命令
```
# 根据proto生成配置文件代码
make config 

# 生成新的服务
kratos new app/user --nomod

# 生成新的proto文件
kratos proto add api/helloworld/v1/demo.proto

# 根据proto生成接口相关代码
make api 

# 新增表
ent new 表名

# 修改表结构后，执行如下命令生成代码
go generate ./ent
```

## 资源
* [kratos](https://go-kratos.dev/docs/)
* [kratos-use 文章](https://blog.csdn.net/ghhg521/article/details/141471093)
