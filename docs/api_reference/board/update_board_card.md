# 更新看板工作项

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/board/update_board_card.html

# 说明

更新看板工作项，返回更新看板工作项的数据

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
updateBoardCard
```


# url

`https://api.tapd.cn/board_cards`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次修改一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">是</td><td style="text-align:center;">int</td><td style="text-align:center;">工作项ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">工作项标题</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">负责人</td></tr><tr><td style="text-align:center;">cc</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">参与人</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态</td></tr><tr><td style="text-align:center;">begin</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">开始时间</td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">截止时间</td></tr><tr><td style="text-align:center;">b_label</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">标签ID</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">详细描述</td></tr></tbody></table>

# 调用示例及返回结果

`curl -u 'api_user:api_password' -d 'workspace_id=20355782&id=1020355782500624255&description=更新内容' 'https://api.tapd.cn/board_cards'`

## 返回结果

```
{
	status: 1,
	data: [
	    {
			BoardCard: {
				"id":"1020355782500624255",
				"name":"test1",
				"description":"内容更新",
				"created":"2019-09-17 15:58:43",
				"modified":"2019-09-17 16:01:03",
				"workspace_id":"20355782",
				"owner":null,
				"due":null,
				"b_label":"[0]",
				"b_sort":"3",
				"b_board_id":"1020355782000010725",
				"b_column_id":"1020355782000045509",
				"status":"open",
				"cc":null,
				"begin":null
			}
		}
	],
	info: "success"
}
```


# 看板工作项字段说明

## 看板工作项重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">工作项ID</td></tr><tr><td style="text-align:center;">b_board_id</td><td style="text-align:center;">看板ID</td></tr><tr><td style="text-align:center;">b_column_id</td><td style="text-align:center;">板块ID</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">负责人</td></tr><tr><td style="text-align:center;">cc</td><td style="text-align:center;">参与人</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">状态</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">工作项标题</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">begin</td><td style="text-align:center;">开始时间</td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">截止时间</td></tr><tr><td style="text-align:center;">b_label</td><td style="text-align:center;">标签ID</td></tr></tbody></table>
