# 获取项目报告

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/report/get_workspace_reports.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目报告数据](#获取项目报告数据)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [项目报告相关重要字段说明](#项目报告相关重要字段说明)
-   [状态(status)取值字段说明](#状态-status-取值字段说明)
-   [报告类型(report\_type)取值字段说明](#报告类型-report-type-取值字段说明)

# 说明

获取项目报告数据

# url

`https://api.tapd.cn/workspace_reports`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;">无</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">ID</td><td style="text-align:center;">无</td></tr><tr><td style="text-align:center;">title</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">标题</td><td style="text-align:center;">无</td></tr><tr><td style="text-align:center;">author</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">创建人</td><td style="text-align:center;">无</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">无</td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目报告数据

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workspace_reports?workspace_id=10104801'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "WorkspaceReport": {
                "id": "1010104801000399873",
                "title": "TAPD 乌云转测试报告",
                "workspace_id": "10104801",
                "report_type": "totest",
                "receiver": "v_zujunlin;",
                "cc": "anyechen;",
                "receiver_organization_ids": null,
                "cc_organization_ids": null,
                "sender": "anyechen",
                "send_time": "2018-09-25 16:32:11",
                "author": "anyechen",
                "created": "2018-09-25 16:32:10",
                "status": "sent",
                "modified": "2018-09-25 16:32:11",
                "last_modify": "anyechen"
            }
        }
    ],
    "info": "success"
}
```


# 项目相关字段说明

## 项目报告相关重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">报告id</td></tr><tr><td style="text-align:center;">title</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目id</td></tr><tr><td style="text-align:center;">report_type</td><td style="text-align:center;">报告类型</td></tr><tr><td style="text-align:center;">receiver</td><td style="text-align:center;">收件人</td></tr><tr><td style="text-align:center;">cc</td><td style="text-align:center;">抄送人</td></tr><tr><td style="text-align:center;">receiver_organization_ids</td><td style="text-align:center;">做为接收人的组织架构</td></tr><tr><td style="text-align:center;">cc_organization_ids</td><td style="text-align:center;">做为抄送人的组织架构</td></tr><tr><td style="text-align:center;">sender</td><td style="text-align:center;">发件人</td></tr><tr><td style="text-align:center;">send_time</td><td style="text-align:center;">发件时间</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">author</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">状态</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">last_modify</td><td style="text-align:center;">最后修改人</td></tr></tbody></table>

## 状态(status)取值字段说明

<table><thead><tr><th style="text-align:center;">取值</th><th style="text-align:center;">字面值</th></tr></thead><tbody><tr><td style="text-align:center;">sent</td><td style="text-align:center;">已发送</td></tr><tr><td style="text-align:center;">draft</td><td style="text-align:center;">草稿</td></tr><tr><td style="text-align:center;">abandon</td><td style="text-align:center;">已删除</td></tr></tbody></table>

## 报告类型(report\_type)取值字段说明

<table><thead><tr><th style="text-align:center;">取值</th><th style="text-align:center;">字面值</th></tr></thead><tbody><tr><td style="text-align:center;">normal</td><td style="text-align:center;">项目进度报告</td></tr><tr><td style="text-align:center;">totest</td><td style="text-align:center;">项目转测试</td></tr><tr><td style="text-align:center;">test</td><td style="text-align:center;">测试报告</td></tr></tbody></table>
