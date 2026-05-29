# 获取缺陷模板列表

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_bug_template_list.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [返回结果](#返回结果)
-   [缺陷模板重要字段说明](#缺陷模板重要字段说明)

# 说明

返回符合查询条件的所有缺陷模板

# url

`https://api.tapd.cn/bugs/template_list`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回所有缺陷模板

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

`curl -u 'api_user:api_password' 'https://api.tapd.cn/bugs/template_list?workspace_id=10104801'`

## 返回结果

```
{
    "status": 1,
    "data": [
        {
            "WorkitemTemplate": {
                "id": "1010104801000068639",
                "name": "创建模板",
                "description": "AA",
                "sort": "1",
                "default": "0",
                "creator": "v_xuanfang",
                "editor_type": "1"
            }
        },
        {
            "WorkitemTemplate": {
                "id": "1010104801000031595",
                "name": "System default template",
                "description": "Auto created by the system",
                "sort": "1",
                "default": "1",
                "creator": "SYSTEM",
                "editor_type": "1"
            }
        }
    ],
    "info": "success"
}
```


# 缺陷模板字段说明

## 缺陷模板重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">模板ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">default</td><td style="text-align:center;">是否启用</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">提交人</td></tr></tbody></table>
