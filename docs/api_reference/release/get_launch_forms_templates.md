# 获取发布评审模板

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/get_launch_forms_templates.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取发布评审模板](#获取发布评审模板)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

获取发布评审模板

# url

`https://api.tapd.cn/launch_forms/templates`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认获取所有数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取发布评审模板

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/launch_forms/templates?workspace_id=20042301'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "template": {
                "id": "1120042301001000009",
                "name": "系统默认流程"
            }
        },
        {
            "template": {
                "id": "1120042301001000076",
                "name": "广告歌"
            }
        },
        {
            "template": {
                "id": "1120042301001000077",
                "name": "广告歌"
            }
        },
        {
            "template": {
                "id": "1120042301001000078",
                "name": "广告歌"
            }
        },
        {
            "template": {
                "id": "1120042301001000079",
                "name": "迭代"
            }
        }
    ],
    "info": "success"
}
```


# 返回字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">name</td><td style="text-align:center;">模板名称</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">模板ID</td></tr></tbody></table>
