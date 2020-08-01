教程源地址：https://github.com/eddycjy/go-gin-example

---

# TODO

- 模糊查询
- 查询时过滤软删除的项

# 目录结构

- conf：配置文件
- docs：由swag init生成的文档
- middleware：应用中间件
- models：应用数据库模型
- pkg：第三方包
    - e：api错误码
    - util：工具包
- routers 路由逻辑处理
- runtime：应用运行时数据

# swagger地址

`http://localhost:8000/swagger/index.html`

# 其它说明

如果开启了jwt功能，其余接口中没有token参数会报401，需要把前边获取的token加上。