# 使用并行工作流的需求，获取其节点信息

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_steps.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目下需求](#获取项目下需求)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [字段说明](#字段说明)

# 说明

使用并行工作流的需求，获取其节点信息

# url

`https://api.tapd.cn/stories/get_story_step_list`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

获取指定需求的所有节点列表

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">story_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">需求ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下需求

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/stories/get_story_step_list?workspace_id=70002667&story_id=1070002667006658827'`

### 返回结果

```
{
  "status": 1,
  "data": [
    {
      "WorkitemStepInfo": {
        "id": "1070002667000137213",
        "workspace_id": "70002667",
        "entity_type": "story",
        "workitem_id": "1070002667006658827",
        "step": "step_2970811_1",
        "status": "0",
        "owner": "",
        "begin": null,
        "due": null,
        "effort": "3",
        "iteration_id": "0",
        "begin_time": "2026-01-04 09:37:57",
        "complete_time": "2026-01-04 09:38:23",
        "time_cost": "26",
        "completer": "ocenhu"
      }
    }
  ],
  "info": "success"
}
```


## 字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">step</td><td style="text-align:center;">节点原名</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">节点状态</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">节点负责人</td></tr><tr><td style="text-align:center;">begin</td><td style="text-align:center;">节点预计开始</td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">节点预计结束时间</td></tr><tr><td style="text-align:center;">effort</td><td style="text-align:center;">节点预估工时</td></tr><tr><td style="text-align:center;">iteration_id</td><td style="text-align:center;">节点迭代</td></tr><tr><td style="text-align:center;">begin_time</td><td style="text-align:center;">实际开始时间</td></tr><tr><td style="text-align:center;">complete_time</td><td style="text-align:center;">实际完成时间</td></tr><tr><td style="text-align:center;">time_cost</td><td style="text-align:center;">节点停留时长</td></tr><tr><td style="text-align:center;">completer</td><td style="text-align:center;">操作完成人</td></tr></tbody></table>
