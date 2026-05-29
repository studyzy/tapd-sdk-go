# 通过工作项短id换长id

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/get_workitems_long_id_by_short_ids.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [根据短id获取业务对象信息](#根据短id获取业务对象信息)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [如返回结果受限，可按如下规则确认数据权限](#如返回结果受限-可按如下规则确认数据权限)

# 说明

通过工作项短id换长id

# url

`https://api.tapd.cn/workspaces/get_workitems_long_id_by_short_ids`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">short_ids</td><td style="text-align:center;"><code>否</code></td><td style="text-align:center;">string</td><td style="text-align:center;">短ID，多个以;分隔</td><td style="text-align:center;">short_ids和long_ids不允许都不传</td></tr><tr><td style="text-align:center;">long_ids</td><td style="text-align:center;"><code>否</code></td><td style="text-align:center;">string</td><td style="text-align:center;">长ID，多个以;分隔</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目id</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">entity_type</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">查找业务对象类型的范围</td><td style="text-align:center;">候选值：story,task,bug</td></tr></tbody></table>

其余支持参数参照各业务对象的获取文档

# 调用示例及返回结果

## 根据短id获取业务对象信息

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workspaces/get_workitems_long_id_by_short_ids?short_ids=1000276;1000277;1000104&workspace_id=48464494&entity_type=story'`

`curl 'https://api.tapd.cn/workspaces/get_workitems_long_id_by_short_ids?short_ids=1000276;1000277;1000104&workspace_id=48464494&entity_type=story&access_token=ACCESS_TOKEN'`

### 返回结果

```
    "status": 1,
    "data": {
        "valid_id_map": [
            {
                "short_id": "1000276",
                "long_id": "1148464494001000276",
                "entity_type": "story",
                "workspace_id": "48464494",
                "company_id": "39418254"
            },
            {
                "short_id": "1000277",
                "long_id": "1148464494001000277",
                "entity_type": "story",
                "workspace_id": "48464494",
                "company_id": "39418254"
            }
        ],
        "invalid_long_ids": [
            "1000104",
            "123213223121000276",
            "1231231231231231000277"
        ],
        "invalid_short_ids": []
    },
    "info": "success"
}
```


# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">valid_id_map</td><td style="text-align:center;">有效的id集合</td></tr><tr><td style="text-align:center;">long_id</td><td style="text-align:center;">长ID</td></tr><tr><td style="text-align:center;">short_ids</td><td style="text-align:center;">短ID</td></tr><tr><td style="text-align:center;">company_id</td><td style="text-align:center;">公司ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目id</td></tr><tr><td style="text-align:center;">entity_type</td><td style="text-align:center;">查找业务对象类型的范围候选值：story,task,bug</td></tr><tr><td style="text-align:center;">invalid_long_ids</td><td style="text-align:center;">无效的长id集合</td></tr><tr><td style="text-align:center;">invalid_short_ids</td><td style="text-align:center;">无效的短id集合</td></tr></tbody></table>

# 数据权限说明

## 如返回结果受限，可按如下规则确认数据权限

-   确认API账号是否有对应的业务对象获取权限
-   确认传入的短id是否是在当前业务对象类型下
