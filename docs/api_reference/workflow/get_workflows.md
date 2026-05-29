# 获取项目下的工作流列表

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workflow/get_workflows.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目下的模块信息](#获取项目下的模块信息)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [模块重要字段说明](#模块重要字段说明)

# 说明

返回项目下的工作流列表（分页显示，默认一页30条）

# url

`https://api.tapd.cn/workflows`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">id</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">详细描述</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">system_name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">系统名。取 bugtrace （缺陷的）或者 story（需求的）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下的模块信息

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workflows?workspace_id=10104801'`

`curl 'https://api.tapd.cn/workflows?workspace_id=10104801&access_token=ACCESS_TOKEN'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Workflow": {
                "id": "1010104801000128627",
                "workspace_id": "10104801",
                "name": "bbb",
                "description": "",
                "system_name": "story",
                "is_default": "0",
                "created": "2021-01-26 16:42:33",
                "creator": "anyechen",
                "modified": "2021-01-26 16:43:28",
                "modifier": "anyechen",
                "type": "classic"
            }
        },
        {
            "Workflow": {
                "id": "1010104801000050043",
                "workspace_id": "10104801",
                "name": "system workflow",
                "description": "系统默认工作流",
                "system_name": "story",
                "is_default": "1",
                "created": "2017-08-10 10:17:45",
                "creator": "",
                "modified": "2020-11-27 06:13:13",
                "modifier": "v_xuanfang",
                "type": "classic"
            }
        },
        {
            "Workflow": {
                "id": "1010104801000050003",
                "workspace_id": "10104801",
                "name": "system workflow",
                "description": "111\"&amp;gt;</textarea>",
                "system_name": "bugtrace",
                "is_default": "1",
                "created": "2017-06-30 15:54:38",
                "creator": "v_xuanfang",
                "modified": "2020-11-27 06:13:11",
                "modifier": "v_xuanfang",
                "type": "bpm"
            }
        }
    ],
    "info": "success"
}
```


# 模块字段说明

## 模块重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">id</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">system_name</td><td style="text-align:center;">应用类型</td></tr><tr><td style="text-align:center;">is_default</td><td style="text-align:center;">是否默认</td></tr><tr><td style="text-align:center;">modifier</td><td style="text-align:center;">最后修改人</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">工作流类型，bpm 并行工作流，classic 串行工作流</td></tr></tbody></table>
