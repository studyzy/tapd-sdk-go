# 创建需求与缺陷关联关系

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/create_story_bug.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [创建 需求1010104801864772561 与 缺陷1010104801085924155 关联关系](#创建-需求1010104801864772561-与-缺陷1010104801085924155-关联关系)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)

# 说明

创建需求与缺陷关联关系

# URL

`https://api.tapd.cn/relations`

# 支持格式

JSON/XML (默认JSON格式)

# HTTP请求方式

POST

# 请求数限制

一次插入一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">所属TAPD项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">source_type</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">关联关系源对象类型（story、bug）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">target_type</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">关联关系目标对象类型（story、bug）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">source_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">关联关系源对象id</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">target_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">关联关系目标对象id</td><td style="text-align:center;"></td></tr></tbody></table>

## 创建 需求1010104801864772561 与 缺陷1010104801085924155 关联关系

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10104801&source_type=bug&source_id=1010104801085924155&target_type=story&target_id=1010104801864772561' 'https://api.tapd.cn/relations'`

`curl -u 'Authorization: Bearer ACCESS_TOKEN' -d 'workspace_id=10104801&source_type=bug&source_id=1010104801085924155&target_type=story&target_id=1010104801864772561' 'https://api.tapd.cn/relations'`

# 返回结果

以返回结果为准

```
{
    "status": 1,
    "data": {
        "Relation": {
            "id": "22265547",
            "workspace_id": "10104801",
            "source_type": "story",
            "source_id": "1010104801864772561",
            "target_type": "bug",
            "target_id": "1010104801085924155",
            "modified": "2021-05-20 15:10:59",
            "created": "2021-05-20 15:10:59"
        }
    },
    "info": "success"
}
```


# 返回结果字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">主键ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">source_type</td><td style="text-align:center;">关联关系源对象类型</td></tr><tr><td style="text-align:center;">source_id</td><td style="text-align:center;">关联关系源对象id</td></tr><tr><td style="text-align:center;">target_type</td><td style="text-align:center;">关联关系目标对象类型</td></tr><tr><td style="text-align:center;">target_id</td><td style="text-align:center;">关联关系目标对象id</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr></tbody></table>
