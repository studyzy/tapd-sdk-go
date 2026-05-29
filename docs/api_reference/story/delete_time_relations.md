# 批量删除需求前后置关系

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/delete_time_relations.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
    -   [备注和调用推荐](#备注和调用推荐)
-   [按照节点，删除从需求A（1020375552855142317）指向需求B（1020375552855139943）的前后置关系](#按照节点-删除从需求a-1020375552855142317-指向需求b-1020375552855139943-的前后置关系)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
-   [按照ID，删除前后置关系](#按照id-删除前后置关系)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

批量删除需求前后置关系

# url

`https://api.tapd.cn/stories/delete_time_relations`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求参数

<table><thead><tr><th style="text-align:left;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:left;">说明</th></tr></thead><tbody><tr><td style="text-align:left;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:left;">项目ID</td></tr><tr><td style="text-align:left;">relations</td><td style="text-align:center;"><code>否</code></td><td style="text-align:center;">array</td><td style="text-align:left;">所需删除的关系列表（按起点和终点）</td></tr><tr><td style="text-align:left;">relations[0][workitem_id]</td><td style="text-align:center;"><code>否</code></td><td style="text-align:center;">integer</td><td style="text-align:left;">示例, 所需按节点删除的第0条关系的“起点需求id”</td></tr><tr><td style="text-align:left;">relations[0][dst_workitem_id]</td><td style="text-align:center;"><code>否</code></td><td style="text-align:center;">integer</td><td style="text-align:left;">示例, 所需按节点删除的第0条关系的“终点需求id”</td></tr><tr><td style="text-align:left;">relation_ids</td><td style="text-align:center;"><code>否</code></td><td style="text-align:center;">array</td><td style="text-align:left;">所需删除的关系列表（按id）</td></tr><tr><td style="text-align:left;">relation_ids[0]</td><td style="text-align:center;"><code>否</code></td><td style="text-align:center;">integer</td><td style="text-align:left;">所需按id删除的第0条关系的id</td></tr><tr><td style="text-align:left;">current_user</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:left;">执行此操作的用户的nick</td></tr></tbody></table>

# 调用示例及返回结果

### 备注和调用推荐

-   理想情况下，开发者能够提前知道所需删除的关系的id，直接利用“按id删除”模式调用即可。
-   但在一些特殊情况下，难免会出现“只知道要删除从需求A到需求B的关系，却不知道它的具体id是什么”的情况。在这种情况下，如果需要额外获取关系的id，路径比较曲折且不必要。此时则可以在参数中采用“按节点删除”的模式。
-   如果上述两种情况都存在，那么可以在参数里混传两个列表，系统会分别按照id和按照节点查询各自所需删除的实际关系，并将两个列表合并后进行删除。

## 按照节点，删除从需求A（1020375552855142317）指向需求B（1020375552855139943）的前后置关系

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=20375552&relations[0][workitem_id]=1020375552855142317&relations[0][dst_workitem_id]=1020375552855139943&current_user=xinweihe' 'https://api.tapd.cn/stories/delete_time_relations'`

## 按照ID，删除前后置关系

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=20375552&relation_ids[0]=1220375552000009739&current_user=xinweihe' 'https://api.tapd.cn/stories/delete_time_relations'`

### 返回结果

结果中的num代表了实际删除的条数。

```
{
    "status": 1,
    "data": {
        "num": 1,
    },
    "info": "success"
}
```

