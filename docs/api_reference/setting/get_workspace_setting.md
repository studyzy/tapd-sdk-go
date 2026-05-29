# 获取项目配置开关

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/setting/get_workspace_setting.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目配置开关](#获取项目配置开关)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

获取项目配置开关

# url

`https://api.tapd.cn/settings/get_workspace_setting`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认所有

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">配置名称（is_enabled_story_category 是否启用需求分类树，workspace_metrology 工时单位）</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目配置开关

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/settings/get_workspace_setting?workspace_id=10104801&type=workspace_metrology'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "workspace_metrology": "day"
    },
    "info": "success"
}
```


# 参数 type 取值说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">is_enabled_story_category</td><td style="text-align:center;">是否启用需求分类树（1启用，0未启用 ）</td></tr><tr><td style="text-align:center;">workspace_metrology</td><td style="text-align:center;">工时单位（day 天，hour 小时）</td></tr></tbody></table>
