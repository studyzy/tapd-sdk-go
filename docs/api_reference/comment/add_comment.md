# 添加评论接口

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/comment/add_comment.html

# 说明

新建评论，返回新建评论的数据

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
addComment
```


# url

`https://api.tapd.cn/comments`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次插入一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">description</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">内容</td></tr><tr><td style="text-align:center;">author</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">评论人</td></tr><tr><td style="text-align:center;">entry_type</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">评论类型（取值： bug、 bug_remark （流转缺陷时候的评论）、 stories、 tasks、 wiki 。）</td></tr><tr><td style="text-align:center;">entry_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">评论所依附的业务对象实体id</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">root_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">根评论ID</td></tr><tr><td style="text-align:center;">reply_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">评论回复的ID</td></tr></tbody></table>

### curl 调用示例

#### Bug 评论

`curl -u 'api_user:api_password' -d 'workspace_id=20355782&entry_type=bug&entry_id=1020355782500647717&description=ccc&author=xuanfang' 'https://api.tapd.cn/comments'`

#### Wiki 评论

`curl -u 'api_user:api_password' -d 'workspace_id=38122198&entry_type=wiki&entry_id=1138122198001004168&description=测试评论&author=dev' 'https://api.tapd.cn/comments'`

> **注意**：`entry_type=wiki` 虽未在 TAPD 官方文档中列出，但实际 API 支持。Wiki 的 `entry_id` 为 Wiki 页面 ID（如 `1138122198001004168`）。

### 返回结果

```
{
	"status":1,
	"data":{
		"Comment":{
		"id":"1020355782058781915",
		"title":"\u5728\u72b6\u6001 [\u65b0] \u6dfb\u52a0",
		"description":"ccc",
		"author":"v_xuanfang",
		"entry_type":"bug",
		"entry_id":"1020355782500647717",
		"created":"2019-12-24 18:33:53",
		"modified":"2019-12-24 18:33:53",
		"workspace_id":"20355782"
	  }
   },
   "info":"success"
}
```


# 评论字段说明

## 评论重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">评论ID</td></tr><tr><td style="text-align:center;">title</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">内容</td></tr><tr><td style="text-align:center;">author</td><td style="text-align:center;">评论人</td></tr><tr><td style="text-align:center;">entry_type</td><td style="text-align:center;">评论类型（取值： bug、 bug_remark （流转缺陷时候的评论）、 stories、 tasks 。</td></tr><tr><td style="text-align:center;">entry_id</td><td style="text-align:center;">评论所依附的业务对象实体id</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后更改时间</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">root_id</td><td style="text-align:center;">根评论ID</td></tr><tr><td style="text-align:center;">reply_id</td><td style="text-align:center;">评论回复的ID</td></tr></tbody></table>
