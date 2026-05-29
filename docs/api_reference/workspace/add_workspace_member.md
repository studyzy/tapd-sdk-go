# 添加项目成员

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/add_workspace_member.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [添加 davidning 到项目 10104801，角色为 测试人员 产品人员](#添加-davidning-到项目-10104801-角色为-测试人员-产品人员)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

添加项目成员

# url

`https://api.tapd.cn/workspaces/add_workspace_member`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次添加一位成员。如果项目下已存在该成员，则更新（角色组）

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">待加入的项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">nick</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">用户的英文昵称</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">company_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">成员所在公司ID（云端必填）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">role_ids</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">角色组</td><td style="text-align:center;">多个使用 , 分隔。项目下所有角色取值参考文档：获取用户组ID对照关系接口</td></tr></tbody></table>

# 调用示例及返回结果

## 添加 davidning 到项目 10104801，角色为 测试人员 产品人员

### curl 使用 Basic Auth 鉴权调用示例

`curl -u "api_user:api_password" -d 'workspace_id=10104801&nick=davidning&role_ids=1000000000000000010,1000000000000000015' 'https://api.tapd.cn/workspaces/add_workspace_member'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "success": true
    },
    "info": "add member davidning to 10104801 success"
}
```

