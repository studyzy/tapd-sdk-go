# 新建空间

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/mini_api_reference/workspace/create_mini_project.html

# 说明

新建空间

# url

`https://api.tapd.cn/workspaces/create_mini_project`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

-   一次只能创建一个项目

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">company_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">公司ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">空间名</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td><td style="text-align:center;">创建人必须属于当前公司</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">空间描述</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">template_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">模板ID</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

在10000000公司中使用id为57的模板创建一个名为“ApiCreate”的空间，创建人为tapduser

### curl 使用 Basic Auth 鉴权调用示例

`curl -u "api_user:api_password" -d 'company_id=10000000&name=ApiCreate&creator=tapduser&template_id=57' 'https://api.tapd.cn/workspaces/create_mini_project'`

### curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' -d 'company_id=10000000&name=ApiCreate&creator=tapduser&template_id=57' 'https://api.tapd.cn/workspaces/create_mini_project'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "workspace_id": "69999044",
        "workspace_url": "https://tapd.cn/tapd_fe/t/index/69999044"
    },
    "info": "success"
}
```

