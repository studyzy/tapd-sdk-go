# 复制需求类别接口

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/setting/copy_workitem_type_setting.html

# 说明

复制配置好的需求类别，返回需求类别配置

# url

`https://api.tapd.cn/stories/copy_workitem_type_setting`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次只能复制一个类别到一个项目

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">src_workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">源项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">src_workitem_type_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">源需求类型ID</td><td style="text-align:center;">当前需求类型名称必须在目标项目不存在</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 复制需求类别到目标项目

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10104801&src_workitem_type_id=1020358527000067037&src_workspace_id=755' 'https://api.tapd.cn/stories/copy_workitem_type_setting'`

### 返回结果

```json
{
    "status": 1,
    "data": {
        "WorkitemType": {
            "id": "1010104801000061611",
            "workspace_id": "10104801",
            "entity_type": "story",
            "name": "技术需求",
            "english_name": "TSTORY",
            "status": "3",
            "color": "#5c88c5",
            "workflow_id": "1010104801001087491",
            "creator": "v_xuanfang",
            "created": "2020-12-02 15:47:02",
            "modified_by": "v_xuanfang",
            "modified": "2021-01-21 10:49:57"
        }
    },
    "info": "success"
}
```
