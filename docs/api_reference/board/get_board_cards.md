# 获取看板工作项接口

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/board/get_board_cards.html

# 说明

返回符合查询条件的所有看板工作项（分页显示，默认一页30条）

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
getBoardCards
```


# url

`https://api.tapd.cn/board_cards`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">工作项ID</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">b_board_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">看板ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">b_column_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">板块ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">负责人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">cc</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">参与人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">工作项标题</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">begin</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">开始时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">截止时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">b_label</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">标签ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

`curl -u 'api_user:api_password' 'https://api.tapd.cn/board_cards?workspace_id=20355782'`

## 返回结果

```
{
	status: 1,
	data: [
	    {
			BoardCard: {
				id: "1020355782500624163",
				name: "('测试2')",
				description: "546546541davsavddv按房产税大V是",
				created: "2019-09-11 14:39:45",
				modified: "2019-09-11 14:42:41",
				workspace_id: "20355782",
				owner: "null;",
				due: "2019-09-19",
				b_label: "[0]",
				b_sort: "2",
				b_board_id: "1020355782000010725",
				b_column_id: "1020355782000045509",
				status: "open",
				cc: "null",
				begin: "2019-09-11"
			}
		}
	],
	info: "success"
}
```


# 看板工作项字段说明

## 看板工作项重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">工作项ID</td></tr><tr><td style="text-align:center;">b_board_id</td><td style="text-align:center;">看板ID</td></tr><tr><td style="text-align:center;">b_column_id</td><td style="text-align:center;">板块ID</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">负责人</td></tr><tr><td style="text-align:center;">cc</td><td style="text-align:center;">参与人</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">状态</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">工作项标题</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">begin</td><td style="text-align:center;">开始时间</td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">截止时间</td></tr><tr><td style="text-align:center;">b_label</td><td style="text-align:center;">标签ID</td></tr></tbody></table>
