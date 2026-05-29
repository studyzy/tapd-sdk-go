# 获取发布评审数量接口

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/get_launch_forms_count.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目下发布评审数量](#获取项目下发布评审数量)
    -   [curl 调用示例](#curl-调用示例)
    -   [返回结果](#返回结果)

# 说明

计算符合查询条件的发布评审数量并返回

# url

`https://api.tapd.cn/launch_forms/count`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

只返回发布评审数量

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">发布评审ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">title</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">version_type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">版本类型</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">baseline</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">基线</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">release_model</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">发布模块</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">roadmap_version</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">路标版本</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">release_type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">发布类型</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">change_type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">变更类型</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">signed_by</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">签发人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">archived_by</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">发布确认人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">cc</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">抄送人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">change_notifier</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">变更通知人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下发布评审数量

### curl 调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/launch_forms/count?workspace_id=20003271'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "count": 1
    },
    "info": "success"
}
```

