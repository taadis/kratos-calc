# kratos-calc

基于 kratos 的计算器服务

- `kratos new kratos-calc`
- 修改 api.proto, 定义加减乘除基本函数 add/subtract/multiply/divide
- 在 api 目录下执行 `go generate` 命令重新生成 api.pb.go 和 api.bm.go 文件


此时运行 `kratos run` 命令会报错如下:

```
cmd>kratos run
# kratos-calc/api
..\api\client.go:15:68: undefined: DemoClient
..\api\client.go:21:9: undefined: NewDemoClient
panic: exit status 2
```

修改为我们新生成的 pb.go 中的函数

## 修改 service.go

实现 api.pb.go 中定义的接口

## 修改 grpc/server.go

注册新的服务实现

## 修改 http/server.go

注册新的服务实现

## 测试

可以通过 curl/postman/浏览器 访问以下 url 来查看

- [http://localhost:8000/calc/add?a=1&b=2](http://localhost:8000/calc/add?a=1&b=2)
- [http://localhost:8000/calc/subtract?a=1&b=2](http://localhost:8000/calc/subtract?a=1&b=2)
- [http://localhost:8000/calc//multiply?a=1&b=2](http://localhost:8000/calc//multiply?a=1&b=2)

## 搞定
