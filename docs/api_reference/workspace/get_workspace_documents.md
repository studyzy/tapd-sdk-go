# 获取项目文档

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/get_workspace_documents.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [项目文档重要字段说明](#项目文档重要字段说明)

# 说明

获取项目文档

# url

`https://api.tapd.cn/documents/get_workspace_documents`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条，可通过传 limit 参数设置，最大取 200

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/documents/get_workspace_documents?workspace_id=20003271'`

# 返回结果

```
 {
  "status": 1,
  "data": [
    {
      "Document": {
        "id": "1147043561001001330",
        "workspace_id": "47043561",
        "name": "熟悉思维导图",
        "type": "mindmap",
        "folder_id": "1147043561001000694",
        "creator": "TAPD",
        "modifier": "TAPD",
        "status": null,
        "created": "2021-09-09 16:08:52",
        "modified": "2021-09-09 16:08:52"
      }
    },
    {
      "Document": {
        "id": "1147043561001001329",
        "workspace_id": "47043561",
        "name": "文档功能使用秘籍",
        "type": "word",
        "folder_id": "1147043561001000694",
        "creator": "TAPD",
        "modifier": "TAPD",
        "status": null,
        "created": "2021-09-09 16:08:51",
        "modified": "2021-09-09 16:08:51"
      }
    }
  ],
  "info": "success"
}
```


# 项目文档字段说明

## 项目文档重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">文档类型</td></tr><tr><td style="text-align:center;">folder_id</td><td style="text-align:center;">文件夹ID</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">modifier</td><td style="text-align:center;">最后修改人</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr></tbody></table>
