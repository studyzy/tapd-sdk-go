# 关联缺陷

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/link_bugs.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [关联缺陷 1162187798001003385 与其它缺陷1162187798001000128,1162187798001000127,1162187798001000126](#关联缺陷-1162187798001003385-与其它缺陷1162187798001000128-1162187798001000127-1162187798001000126)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

关联缺陷

# url

`https://api.tapd.cn/bugs/link_bugs`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次返回所有符合条件的值

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;">无</td></tr><tr><td style="text-align:center;">bug_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">19位长度的缺陷ID（原始缺陷）</td><td style="text-align:center;">无</td></tr><tr><td style="text-align:center;">relate_bugs</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">多个bug_id（关联缺陷）使用“，”隔开</td><td style="text-align:center;">无</td></tr></tbody></table>

# 调用示例及返回结果

## 关联缺陷 1162187798001003385 与其它缺陷1162187798001000128,1162187798001000127,1162187798001000126

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/bugs/link_bugs --data '{"workspace_id": 62187798,"bug_id": 1162187798001003385,"relate_bugs":"1162187798001000128,1162187798001000127,1162187798001000126"}'`

### 返回结果

```
{
    "status": 1,
    "data": true,
    "info": "success"
}
```

