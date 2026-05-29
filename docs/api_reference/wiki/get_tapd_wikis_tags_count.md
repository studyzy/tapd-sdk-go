# 获取wiki标签信息数量

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis_tags_count.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目下的wiki标签信息数量](#获取项目下的wiki标签信息数量)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

返回符合查询条件的所有Wiki的标签信息（wiki与标签的映射关系）数量

# url

`https://api.tapd.cn/tapd_wikis_tags/count`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">wiki_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">wiki id</td><td style="text-align:center;">不传取项目下的所有wiki</td></tr><tr><td style="text-align:center;">tag</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标签</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标签创建人nick</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">标签创建时间</td><td style="text-align:center;">支持时间查询</td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下的wiki标签信息数量

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/tapd_wikis_tags/count?workspace_id=10104801'`

### 返回结果

```
{
	"status": 1,
	"data": {
		"count": 2
	},
	"info": "success"
}
```

