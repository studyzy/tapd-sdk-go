# 获取当前用户信息

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/user/get_users_info.html

# 说明

获取当前用户信息（无分页）

# url

`https://api.tapd.cn/users/info`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

仅支持用户态 OAuth Access Token 调用 一次返回所有符合条件的值

# 请求参数

无需额外参数，通过 OAuth Access Token 识别当前用户。

# 调用示例及返回结果

## 获取当前用户信息

### 返回结果

```json
{
    "status": 1,
    "data": {
        "id": "6081",
        "nick": "robertyang",
        "name": "杨晓俊",
        "avatar": "http://tiger.oa.com/0/users/avatar/6081/jpg/0/large",
        "enabled": "1",
        "status_id": "1",
        "status_name": "在职"
    },
    "info": "success"
}
```
