# ok0x社区后端源码

[前端源码请查看该工程](https://github.com/OK0X/bbs-front "ok0x - 前端源码") 


## 本地搭建一个 

要求 Go 1.11+

1、下载源码到本地某个目录

```shell
git clone https://github.com/ok0x/bbs-backend
```

2、编译

进入 bbs-backend 项目目录，执行如下命令：

```shell
// unix
make build
// windows
install.bat
```

这样便编译好了 bbs-backend

3、在 bbs-backend 源码中的 bin 目录下应该有了 bbs-backend 可执行文件。

接下来启动 bbs-backend。

```shell
// unix
make start
// windows
start.bat
```

或者

```shell
// unix
bin/studygolang
// windows
bin\studygolang.exe
```

一切顺利的话，bbs-backend 应该就启动了。

4、验证

在浏览器中输入：http://127.0.0.1:8088

应该就能看到了。

接下来你会看到图形化安装界面，一步步照做吧。

* 如果之后有出现页面空白，请查看 error.log 是否有错误

5、问题

1.发布主题时无法选中节点：是因为没有创建子节点，需要先创建一个节点，并添加至少一个子节点。

2.redis报错：未配置redis,配置参考config下面的env.sample.ini配置文件，里面还包括邮箱等许多配置项。

## 参与我们

fork + PR。如果有修改 js 和 css，请执行 gulp （需要先安装 gulp）。注意，Node 版本为：v10.16.2

## 使用该项目搭建的网站

- [ok0x社区(https://b.ok0x.com)
