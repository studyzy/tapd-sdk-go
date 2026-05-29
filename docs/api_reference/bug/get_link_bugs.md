# 获取缺陷与其它缺陷的所有关联关系

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_link_bugs.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取缺陷 1010104801086995895 与其它缺陷的所有关联关系](#获取缺陷-1010104801086995895-与其它缺陷的所有关联关系)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [重要字段说明](#重要字段说明)

# 说明

获取缺陷与其它缺陷的所有关联关系（无分页）

# url

`https://api.tapd.cn/bugs/get_link_bugs`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次返回所有符合条件的值

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;">无</td></tr><tr><td style="text-align:center;">bug_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">19位长度的缺陷ID</td><td style="text-align:center;">无</td></tr></tbody></table>

# 调用示例及返回结果

## 获取缺陷 1010104801086995895 与其它缺陷的所有关联关系

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/bugs/get_link_bugs?workspace_id=10104801&bug_id=1010104801086995895'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "type": "repeat",
            "id": "1010104801085894269",
            "workspace_id": "10104801",
            "actas": "target",
            "linked_workspace_id": 10104801,
            "link_id": "1162187798001000534"
        },
        {
            "type": "copy",
            "id": "1010104801085924155",
            "workspace_id": "10104801",
            "actas": "source",
            "linked_workspace_id": 10104801,
            "link_id": "1162187798001000535",
        }
    ],
    "info": "success"
}
```


# 字段说明

## 重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">type</td><td style="text-align:center;">关系类型。sync_copy 为 同步复制，copy 为复制，repeat 为重复，direct_relate 为直接关联，sync_relate 为同步重复</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">关联的缺陷ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">linked_workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">actas</td><td style="text-align:center;">角色。target 为操作发起方</td></tr><tr><td style="text-align:center;">link_id</td><td style="text-align:center;">bug之间关联关系link的id</td></tr></tbody></table>
