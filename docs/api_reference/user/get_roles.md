# 获取角色ID对照关系

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/user/get_roles.html

# 说明

获取角色ID对照关系（无分页）

# url

`https://api.tapd.cn/roles`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次返回所有符合条件的值 一次只能查一个项目

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目 id</td><td style="text-align:center;">无</td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目角色ID对照关系

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/roles?workspace_id=10104801'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "1000000000000000002": "管理员",
        "1000000000000000009": "开发人员",
        "1000000000000000010": "测试人员",
        "1000000000000000015": "产品人员",
        "1000000000000000014": "项目经理  ",
        "1000000000000000011": "质量管理者",
        "1000000000000000013": "UI人员",
        "1000000000000000006": "关注人",
        "1000000000000000016": "高层管理者",
        "1000000000000000088": "运维人员",
        "1000000000000000018": "访客",
        "1000000000000000017": "其他"
    },
    "info": "success"
}
```

