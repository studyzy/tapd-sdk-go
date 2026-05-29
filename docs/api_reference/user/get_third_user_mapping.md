# 获取用户在三方系统映射的userId

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/user/get_third_user_mapping.html

# 说明

获取用户关联的第三方系统的user\_id和类型

# url

`https://api.tapd.cn/users/get_third_user_mapping`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">公司ID</td></tr><tr><td style="text-align:center;">user_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">用户ID</td></tr></tbody></table>

# 调用示例及返回结果

## 获取用户个人配置

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/users/get_third_user_mapping?workspace_id=48464494&user_id=1223'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "third_partys": {
            [
                "third_party_id":"123",
                "third_party_type":"qywx"
            ]
        },
       
    },
    "info": "success"
}
```


# 返回字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">third_party_id</td><td style="text-align:center;">第三方系统人员ID</td></tr><tr><td style="text-align:center;">third_party_type</td><td style="text-align:center;">系统类别</td></tr></tbody></table>
