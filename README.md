教程源地址：https://github.com/eddycjy/go-gin-example

---

# 目录结构

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

## JWT接口校验获取token

如果接口中没有token参数会报401，需要先获取token：

`http://127.0.0.1:8000/auth?username=test&password=test123456`

后续接口都需要带上token参数。

## Tag标签部分

### 获取全部标签

`GET http://127.0.0.1:8000/api/v1/tags?token=xxxx`

### 新建标签

`POST http://127.0.0.1:8000/api/v1/tags?name=1&state=1&created_by=test&token=xxxx`

### 修改标签

`PUT http://127.0.0.1:8000/api/v1/tags/1?name=edit1&state=0&modified_by=edit1&token=xxxx`

### 删除标签

`DELETE http://127.0.0.1:8000/api/v1/tags/1?token=xxxx`

## Article文章部分

### 新增文章

`POST http://127.0.0.1:8000/api/v1/articles?tag_id=1&title=test1&desc=test-desc&content=test-content&created_by=test-created&state=1&token=xxxx`

### 查询全部文章

`GET http://127.0.0.1:8000/api/v1/articles?token=xxxx`

### 查询指定文章

`GET http://127.0.0.1:8000/api/v1/articles/1?token=xxxx`

### 修改文章

`PUT http://127.0.0.1:8000/api/v1/articles/1?tag_id=1&title=test-edit1&desc=test-desc-edit&content=test-content-edit&modified_by=test-created-edit&state=0&token=xxxx`

### 删除文章

`DELETE http://127.0.0.1:8000/api/v1/articles/1?token=xxxx`