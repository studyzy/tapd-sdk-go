# 获取需求与其它需求的所有关联关系

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_link_stories.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [SDK 方法名](#sdk-方法名)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [需求重要字段说明](#需求重要字段说明)

# 说明

获取需求与其它需求的所有关联关系（无分页）

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
getLinkStories
```


# url

`https://api.tapd.cn/stories/get_link_stories`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次返回所有符合条件的值

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;">无</td></tr><tr><td style="text-align:center;">story_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">19位长度的需求ID</td><td style="text-align:center;">无</td></tr></tbody></table>

# 调用示例及返回结果

获取需求 1010158231500691431 与其它需求的所有关联关系

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/stories/get_link_stories?story_id=1010158231500691431&workspace_id=10158231'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "type": "derivation",
            "id": "1010158231500691433",
            "story_id": "1010158231500691431",
            "workspace_id": "10158231",
            "actas": "target",
            "created": "2019-08-01 16:32:22",
            "modified": "2022-03-18 18:54:25",
            "linked_workspace_id": 10158231
        },
        {
            "type": "sync_copy",
            "id": "1010158231500691437",
            "story_id": "1010158231500691431",
            "workspace_id": "10158231",
            "actas": "target",
            "created": "2019-08-01 16:33:06",
            "modified": "2024-05-10 11:55:03",
            "linked_workspace_id": 10158231
        },
        {
            "type": "copy",
            "id": "1010104801500691441",
            "story_id": "1010158231500691431",
            "workspace_id": "10158231",
            "actas": "target",
            "created": "2019-08-01 16:33:32",
            "modified": "2022-03-18 18:54:25",
            "linked_workspace_id": 10104801
        },
        {
            "type": "direct_relate",
            "id": "1000000755500691185",
            "story_id": "1010158231500691431",
            "workspace_id": "10158231",
            "actas": "target",
            "created": "2019-08-01 16:33:48",
            "modified": "2022-03-18 18:54:25",
            "linked_workspace_id": 755
        }
    ],
    "info": "success"
}
```


# 需求字段说明

## 需求重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">type</td><td style="text-align:center;">关系类型。sync_copy 为 同步复制，copy 为复制，derivation 为派生（父子关系），direct_relate 为直接关联，sync_relate 为关联同步</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">关联的需求ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">story_id</td><td style="text-align:center;">原需求ID</td></tr><tr><td style="text-align:center;">linked_workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">actas</td><td style="text-align:center;">角色。target 为操作发起方</td></tr></tbody></table>
