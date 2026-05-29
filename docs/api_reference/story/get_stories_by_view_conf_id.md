# 获取视图对应的需求列表

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_stories_by_view_conf_id.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目下的需求数据](#获取项目下的需求数据)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [获取视图 1010104801030259563 的数据，并根据创建时间过滤出 2024 之后的需求](#获取视图-1010104801030259563-的数据-并根据创建时间过滤出-2024-之后的需求)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [需求重要字段说明](#需求重要字段说明)
-   [需求优先级(priority)取值字段说明](#需求优先级-priority-取值字段说明)

# 说明

返回视图下最新的30条需求（分页显示，默认一页30条）

# url

`https://api.tapd.cn/stories/get_stories_by_view_conf_id`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">view_conf_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">视图ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">current_user</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">当前登录用户视图</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">需求字段</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">需求字段均可参与基于视图的二次过滤。具体字段及规则可以参考：<a href="/document/api-doc/API文档/api_reference/story/get_stories.html" class="">获取需求</a></td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下的需求数据

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/stories/get_stories_by_view_conf_id?view_conf_id=1010104801030259563&workspace_id=10104801'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Story": {
                "id": "1010104801855983211",
                "name": "sss1",
                "priority": "",
                "status": "developing",
                "owner": "anyechen;",
                "iteration_id": "1010104801000708781",
                "begin": "2022-05-01",
                "due": "2022-05-31",
                "progress": "0"
            }
        },
        {
            "Story": {
                "id": "1010104801855713221",
                "name": "adadf",
                "priority": "",
                "status": "planning",
                "owner": "davidning;anyechen;",
                "iteration_id": "1010104801000708781",
                "begin": "2022-05-01",
                "due": "2022-05-31",
                "progress": "0"
            }
        }
    ],
    "info": {
        "total": 90,
        "current_page": 1,
        "page_size": 30,
        "total_page": 3
    }
}
```


## 获取视图 1010104801030259563 的数据，并根据创建时间过滤出 2024 之后的需求

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/stories/get_stories_by_view_conf_id?view_conf_id=1010104801030259563&workspace_id=10104801&created=>2024-01-01'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Story": {
                "id": "1010104801855983211",
                "name": "sss1",
                "priority": "",
                "status": "developing",
                "owner": "anyechen;",
                "iteration_id": "1010104801000708781",
                "begin": "2022-05-01",
                "due": "2022-05-31",
                "progress": "0"
            }
        }
    ],
    "info": {
        "total": 1,
        "current_page": 1,
        "page_size": 1,
        "total_page": 1
    }
}
```


# 需求字段说明

## 需求重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">priority</td><td style="text-align:center;">优先级</td></tr><tr><td style="text-align:center;">business_value</td><td style="text-align:center;">业务价值</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">状态</td></tr><tr><td style="text-align:center;">version</td><td style="text-align:center;">版本</td></tr><tr><td style="text-align:center;">module</td><td style="text-align:center;">模块</td></tr><tr><td style="text-align:center;">feature</td><td style="text-align:center;">特性</td></tr><tr><td style="text-align:center;">test_focus</td><td style="text-align:center;">测试重点</td></tr><tr><td style="text-align:center;">size</td><td style="text-align:center;">规模</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">处理人</td></tr><tr><td style="text-align:center;">cc</td><td style="text-align:center;">抄送人</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">developer</td><td style="text-align:center;">开发人员</td></tr><tr><td style="text-align:center;">begin</td><td style="text-align:center;">预计开始</td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">预计结束</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">completed</td><td style="text-align:center;">完成时间</td></tr><tr><td style="text-align:center;">iteration_id</td><td style="text-align:center;">迭代ID</td></tr><tr><td style="text-align:center;">template_id</td><td style="text-align:center;">模板ID</td></tr><tr><td style="text-align:center;">effort</td><td style="text-align:center;">预估工时</td></tr><tr><td style="text-align:center;">effort_completed</td><td style="text-align:center;">完成工时</td></tr><tr><td style="text-align:center;">remain</td><td style="text-align:center;">剩余工时</td></tr><tr><td style="text-align:center;">exceed</td><td style="text-align:center;">超出工时</td></tr><tr><td style="text-align:center;">category_id</td><td style="text-align:center;">需求分类</td></tr><tr><td style="text-align:center;">release_id</td><td style="text-align:center;">发布计划</td></tr><tr><td style="text-align:center;">is_archived</td><td style="text-align:center;">是否归档</td></tr><tr><td style="text-align:center;">source</td><td style="text-align:center;">来源</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">类型</td></tr><tr><td style="text-align:center;">parent_id</td><td style="text-align:center;">父需求</td></tr><tr><td style="text-align:center;">children_id</td><td style="text-align:center;">子需求</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">workitem_type_id</td><td style="text-align:center;">需求类别</td></tr><tr><td style="text-align:center;">confidential</td><td style="text-align:center;">是否保密</td></tr><tr><td style="text-align:center;">created_from</td><td style="text-align:center;">需求创建来源（为空时代表web创建）</td></tr></tbody></table>

## 需求优先级(priority)取值字段说明

<table><thead><tr><th style="text-align:center;">取值</th><th style="text-align:center;">字面值</th></tr></thead><tbody><tr><td style="text-align:center;">4</td><td style="text-align:center;">High</td></tr><tr><td style="text-align:center;">3</td><td style="text-align:center;">Middle</td></tr><tr><td style="text-align:center;">2</td><td style="text-align:center;">Low</td></tr><tr><td style="text-align:center;">1</td><td style="text-align:center;">Nice To Have</td></tr></tbody></table>
