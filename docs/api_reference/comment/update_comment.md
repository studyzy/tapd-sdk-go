# 更新评论接口

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/comment/update_comment.html

# 说明

更新评论，返回更新评论的数据

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
updateComment
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

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">description</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">内容</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">评论id</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">change_creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">变更人</td><td style="text-align:center;"></td></tr></tbody></table>

### curl 调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=20355782&description=xxxxx&id=1020355782058781915' 'https://api.tapd.cn/comments'`

### 返回结果

```
{
	"status":1,
	"data":{
		"id":"1020355782058781915",
		"title":"\u5728\u72b6\u6001 [\u65b0] \u6dfb\u52a0",
		"description":"xxxxx",
		"author":"v_xuanfang",
		"entry_type":"bug",
		"entry_id":"1020355782500647717",
		"created":"2019-12-24 18:33:53",
		"modified":"2020-01-03 10:14:04",
		"workspace_id":"20355782"
	  }
   },
   "info":"success"
}
```


# 评论字段说明

## 评论重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">评论ID</td></tr><tr><td style="text-align:center;">title</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">内容</td></tr><tr><td style="text-align:center;">author</td><td style="text-align:center;">评论人</td></tr><tr><td style="text-align:center;">entry_type</td><td style="text-align:center;">评论类型（取值： bug、 bug_remark （流转缺陷时候的评论）、 stories、 tasks 。</td></tr><tr><td style="text-align:center;">entry_id</td><td style="text-align:center;">评论所依附的业务对象实体id</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后更改时间</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr></tbody></table>
