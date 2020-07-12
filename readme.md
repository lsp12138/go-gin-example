教程源地址：https://github.com/eddycjy/go-gin-example

---

#目录结构

- conf：用于存储配置文件
- doc：文档
- middleware：应用中间件
- models：应用数据库模型
- pkg：第三方包
    - e：api错误码
    - setting：调用配置
    - util：工具包
- routers 路由逻辑处理
- runtime：应用运行时数据

# 测试api

## 文章标签部分

### 获取全部标签

`GET http://127.0.0.1:8000/api/v1/tags`

### 新建标签

`POST http://127.0.0.1:8000/api/v1/tags?name=1&state=1&created_by=test`

### 修改标签

`PUT http://127.0.0.1:8000/api/v1/tags/1?name=edit1&state=0&modified_by=edit1`

### 删除标签

`DELETE http://127.0.0.1:8000/api/v1/tags/1`