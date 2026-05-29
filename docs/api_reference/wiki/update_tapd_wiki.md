# 更新 wiki

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/update_tapd_wiki.html

# 说明

更新Wiki，返回Wiki更新后的数据

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
updateTapdWiki
```


# url

`https://api.tapd.cn/tapd_wikis`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

每次只允许更新一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">ID</td></tr><tr><td style="text-align:center;">markdown_description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">Markdown</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">富文本</td></tr><tr><td style="text-align:center;">note</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">备注</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">parent_wiki_id</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">父wiki ID</td></tr></tbody></table>

# 调用示例及返回结果

## 更新标题为 test111 的内容

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'name=test111&description=内容被更新&workspace_id=10104801&id=1210104801000043897' 'https://api.tapd.cn/tapd_wikis'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "Wiki": {
            "id": "1210104801000043897",
            "name": "test111",
            "workspace_id": "10104801",
            "description": "内容被更新",
            "markdown_description": "",
            "is_rich": "1",
            "parent_wiki_id": "0",
            "note": "",
            "view_count": "1",
            "created": "2020-08-26 10:15:28",
            "creator": "v_xuanfang",
            "modified": "2020-08-26 10:30:11",
            "modifier": "dev"
        }
    },
    "info": "success"
}
```


# Wiki字段说明

## Wiki重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">富文本</td></tr><tr><td style="text-align:center;">markdown_description</td><td style="text-align:center;">Markdown</td></tr><tr><td style="text-align:center;">parent_wiki_id</td><td style="text-align:center;">父wiki ID</td></tr><tr><td style="text-align:center;">author</td><td style="text-align:center;">修改人</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">note</td><td style="text-align:center;">备注</td></tr><tr><td style="text-align:center;">view_count</td><td style="text-align:center;">浏览量</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">modifier</td><td style="text-align:center;">最后修改人</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr></tbody></table>
