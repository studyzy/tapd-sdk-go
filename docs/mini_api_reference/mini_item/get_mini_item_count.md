# 获取工作项数量

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/mini_api_reference/mini_item/get_mini_item_count.html

-   [获取空间下工作项的数量](#获取空间下工作项的数量)
    -   [curl 使用 Basic Auth 鉴权调用示例,](#curl-使用-basic-auth-鉴权调用示例)
    -   [curl 使用 OAuth Access Token 鉴权调用示例](#curl-使用-oauth-access-token-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [获取空间下未归档工作项的数量](#获取空间下未归档工作项的数量)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [curl 使用 OAuth Access Token 鉴权调用示例](#curl-使用-oauth-access-token-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

计算符合查询条件的工作项数量并返回（默认包含已归档工作项）

# url

`https://api.tapd.cn/mini_items/count`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

只返回工作项数量

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">ID</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">priority</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">优先级</td><td style="text-align:center;">支持枚举查询</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态</td><td style="text-align:center;">支持枚举查询</td></tr><tr><td style="text-align:center;">label</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标签查询</td><td style="text-align:center;">支持枚举查询</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">处理人</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">begin</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计开始</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计结束</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">最后修改时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">completed</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">完成时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">progress_manual</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">进度</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">category_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">分组</td><td style="text-align:center;">支持枚举查询</td></tr><tr><td style="text-align:center;">parent_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">父工作项</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">children_id</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">子工作项</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">详细描述</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">空间ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">is_archived</td><td style="text-align:center;">否</td><td style="text-align:center;">boolean</td><td style="text-align:center;">是否归档</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">cus_{$自定义字段别名}</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">缺陷自定义字段值，参数名会由后台自动转义为custom_field_*，如：cus_自定义字段的名称</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">custom_field_*</td><td style="text-align:center;">否</td><td style="text-align:center;">string或者integer</td><td style="text-align:center;">自定义字段参数，具体字段名通过接口 <a href="/document/api-doc/API文档/mini_api_reference/mini_item/get_mini_item_custom_fields_settings.html" class="">获取工作项自定义字段配置</a> 获取</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取空间下工作项的数量

### curl 使用 Basic Auth 鉴权调用示例,

`curl -u 'api_user:api_password' 'https://api.tapd.cn/mini_items/count?workspace_id=69993260'`

### curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' 'https://api.tapd.cn/mini_items/count?workspace_id=69993260'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "count": 104
    },
    "info": "success"
}
```


## 获取空间下未归档工作项的数量

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/mini_items/count?workspace_id=69993260&is_archived=0'`

### curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' 'https://api.tapd.cn/mini_items/count?workspace_id=69993260&is_archived=0'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "count": 93
    },
    "info": "success"
}
```

