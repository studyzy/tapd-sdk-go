# 获取单个文档下载链接

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/attachment/documents_down.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [项目文档重要字段说明](#项目文档重要字段说明)

# 说明

获取单个文档下载链接

# url

`https://api.tapd.cn/documents/down`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

每次只能获取一个文档下载链接

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">文档ID</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/documents/down?workspace_id=10104801&id=1010104801001648871'`

# 返回结果

```
{
    "status": 1,
    "data": {
        "Document": {
            "id": "1010104801001648871",
            "workspace_id": "10104801",
            "name": "报告标题.docx",
            "type": "",
            "folder_id": "1010104801000035293",
            "creator": "anyechen",
            "modifier": "anyechen",
            "status": null,
            "created": "2021-12-24 16:40:36",
            "modified": "2021-12-24 16:40:36",
            "download_url": "{{ $page.apiHost }}/10104801/documents/download/1010104801001648871?api_tmp_token=6997a679567ea299251c9d13b3d658a2&client_id=v_xuanfang"
        }
    },
    "info": "success"
}
```


# 项目文档字段说明

## 项目文档重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">文档ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">文档类型</td></tr><tr><td style="text-align:center;">folder_id</td><td style="text-align:center;">文件夹ID</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">modifier</td><td style="text-align:center;">最后修改人</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">download_url</td><td style="text-align:center;">下载链接</td></tr></tbody></table>
