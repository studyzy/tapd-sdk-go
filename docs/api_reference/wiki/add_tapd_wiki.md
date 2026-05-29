# 创建 wiki

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/add_tapd_wiki.html

# 说明

新建Wiki，返回新建Wiki的数据

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
addTapdWiki
```


# url

`https://api.tapd.cn/tapd_wikis`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次插入一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">name</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">markdown_description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">Markdown</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">富文本</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">note</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">备注</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">parent_wiki_id</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">父wiki名</td></tr></tbody></table>

# 调用示例及返回结果

## 在项目下创建 wiki

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'name=test111&description=xxxxxxx&workspace_id=10104801&creator=v_xuanfang' 'https://api.tapd.cn/tapd_wikis'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "Wiki": {
            "id": "1210104801000043897",
            "name": "test111",
            "workspace_id": "10104801",
            "description": "xxxxxxx",
            "markdown_description": "",
            "is_rich": "1",
            "parent_wiki_id": "0",
            "note": "",
            "view_count": "0",
            "created": "2020-08-26 10:15:28",
            "creator": "v_xuanfang",
            "modified": "2020-08-26 10:15:28",
            "modifier": "v_xuanfang"
        }
    },
    "info": "success"
}

# Wiki字段说明
## Wiki重要字段说明
|字段|说明|
|:----:|:----:|
| id | ID |
| name | 标题 |
| description | 富文本 |
| markdown_description | Markdown |
| parent_wiki_id | 父wiki ID |
| author | 修改人 |
| creator | 创建人 |
| note | 备注 |
| view_count | 浏览量 |
| created | 创建时间 |
| modified | 最后修改时间 |
| modifier | 最后修改人 |
| workspace_id | 项目ID |


```

