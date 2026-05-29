# 获取工作流状态中英文名对应关系

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workflow/get_workflow_status_map.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目缺陷工作流状态中英文名对应关系](#获取项目缺陷工作流状态中英文名对应关系)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [获取项目需求工作流状态中英文名对应关系](#获取项目需求工作流状态中英文名对应关系)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

获取工作流状态中英文名对应关系

# url

`https://api.tapd.cn/workflows/status_map`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次只能获取一个项目的工作流状态中英文名对应关系

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">system</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">系统名。取 bug （缺陷的）或者 story（需求的）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">workitem_type_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">需求类别id(获取需求类别ID接口：<a href="/document/api-doc/API文档/api_reference/story/get_workitem_types.html" class="">获取需求类别</a>查询需求状态时需必传</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目缺陷工作流状态中英文名对应关系

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workflows/status_map?system=bug&workspace_id=10158231'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "closed": "已关闭",
        "new": "新",
        "reopened": "重新打开",
        "in_progress": "接受\/处理",
        "resolved": "已解决",
        "rejected": "已拒绝",
        "verified": "已验证"
    },
    "info": "success"
}
```


## 获取项目需求工作流状态中英文名对应关系

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workflows/status_map?system=story&workspace_id=10158231'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "planning": "规划中",
        "developing": "实现中",
        "resolved": "已实现",
        "rejected": "已拒绝"
    },
    "info": "success"
}
```


# 返回字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">键</td><td style="text-align:center;">状态英文名</td></tr><tr><td style="text-align:center;">:----:</td><td style="text-align:center;">:----:</td></tr><tr><td style="text-align:center;">值</td><td style="text-align:center;">状态中文名</td></tr><tr><td style="text-align:center;">:----:</td><td style="text-align:center;">:----:</td></tr></tbody></table>
