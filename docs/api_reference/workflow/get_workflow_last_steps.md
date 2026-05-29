# 获取工作流结束状态

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workflow/get_workflow_last_steps.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目缺陷工作流结束状态中英文名对应关系](#获取项目缺陷工作流结束状态中英文名对应关系)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [获取项目需求工作流结束状态中英文名对应关系](#获取项目需求工作流结束状态中英文名对应关系)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

获取工作流结束状态

# url

`https://api.tapd.cn/workflows/last_steps`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次只能获取一个项目的工作流结束状态

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">system</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">系统名。取 bug （缺陷的）或者 story（需求的）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">workitem_type_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">需求类别id（不传则取项目下所有工作流的结束状态）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">节点类型，仅并行工作流需区分。status 状态，step 并行工作流节点。默认只返回结束状态。若需要同时返回结束状态和结束节点，支持数组type[]=status&amp;type[]=step</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目缺陷工作流结束状态中英文名对应关系

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workflows/last_steps?system=bug&workspace_id=10104801'`

`curl 'https://api.tapd.cn/workflows/last_steps?system=bug&workspace_id=10104801&access_token=ACCESS_TOKEN'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "closed": "已关闭"
    },
    "info": "success"
}
```


## 获取项目需求工作流结束状态中英文名对应关系

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workflows/status_map?system=story&workspace_id=10104801'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "resolved": "已实现x",
        "rejected": "已拒绝"
    },
    "info": "success"
}
```


# 返回字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">键</td><td style="text-align:center;">状态英文名</td></tr><tr><td style="text-align:center;">:----:</td><td style="text-align:center;">:----:</td></tr><tr><td style="text-align:center;">值</td><td style="text-align:center;">状态中文名</td></tr><tr><td style="text-align:center;">:----:</td><td style="text-align:center;">:----:</td></tr></tbody></table>
