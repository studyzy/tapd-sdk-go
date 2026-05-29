# 项目集批量关联/取消关联、修改授权范围项目

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/program/program_relate_workspace.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目成员列表](#获取项目成员列表)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

项目集批量**关联/取消关联**、**修改授权范围**项目

# url

`https://api.tapd.cn/program/program_relate_workspace`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目集 id</td></tr><tr><td style="text-align:center;">action</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">可选值：bind/unbind，分表表示关联、取消关联</td></tr><tr><td style="text-align:center;">relate_workspace_ids</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;"><strong>关联/取消关联</strong>的项目id</td></tr><tr><td style="text-align:center;">auth</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">项目授权，unbind时本参数不生效，非必填，默认为全部，可选值[create_story,create_bug,edit_story,edit_bug]</td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目成员列表

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=xxx&relate_workspace_ids=id1,id2,id3&action=bind"}' 'https://api.tapd.cn/program/program_relate_workspace'`

### 返回结果

```
{"status":1,"data":"ok","info":"success"}
```

