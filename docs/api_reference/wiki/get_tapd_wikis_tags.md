# 获取wiki标签信息

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis_tags.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目下 wiki 的标签信息](#获取项目下-wiki-的标签信息)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

返回符合查询条件的所有Wiki的标签信息（wiki与标签多对多对应）（分页显示，默认一页30条）

# url

`https://api.tapd.cn/tapd_wikis_tags`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

按wiki分页，默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">wiki_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">wiki id</td><td style="text-align:center;">不传取项目下的所有wiki</td></tr><tr><td style="text-align:center;">tag</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标签</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标签创建人nick</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">标签创建时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下 wiki 的标签信息

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/tapd_wikis_tags?workspace_id=10104801'`

### 返回结果

```
{
	"status": 1,
	"data": [
		{
			"Tags": {
				"creator": "huanjinxie",
				"created": "2021-01-07 20:40:05",
				"wiki_id": "1220358527000044697",
				"tag": "首页"
			}
		},
		{
			"Tags": {
				"creator": "huanjinxie",
				"created": "2021-01-07 20:40:05",
				"wiki_id": "1220358527000044697",
				"tag": "测试"
			}
		}
	],
	"info": "success"
}
```

