# 获取发布评审依据

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/get_launch_accessories.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取发布评审依据](#获取发布评审依据)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

获取发布评审依据

# url

`https://api.tapd.cn/launch_accessories`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次只能获取一个评审单的依据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">form_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">评审单ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">评审依据ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created_by</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持时间查询</td></tr></tbody></table>

# 调用示例及返回结果

## 获取发布评审依据

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/launch_accessories?workspace_id=10104801&form_id=1010104801000402051'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "LaunchAccessory": {
                "id": "1010104801000253485",
                "form_id": "1010104801000402051",
                "workspace_id": "10104801",
                "type": "launch_tasks_list",
                "tag": null,
                "title": "任务列表",
                "content": "1010104801500601739",
                "description": null,
                "content_type": "task",
                "created_by": "v_xuanfang",
                "created": "2020-06-11 16:17:56",
                "group_id": null,
                "source": ""
            }
        },
        {
            "LaunchAccessory": {
                "id": "1010104801000253483",
                "form_id": "1010104801000402051",
                "workspace_id": "10104801",
                "type": "launch_bugs_list",
                "tag": null,
                "title": "缺陷列表",
                "content": "1010104801500657621,1010104801500657725",
                "description": null,
                "content_type": "bug",
                "created_by": "v_xuanfang",
                "created": "2020-06-10 19:30:56",
                "group_id": null,
                "source": ""
            }
        },
        {
            "LaunchAccessory": {
                "id": "1010104801000253481",
                "form_id": "1010104801000402051",
                "workspace_id": "10104801",
                "type": "launch_tasks_list",
                "tag": null,
                "title": "任务列表",
                "content": "1010104801854777111,1010104801500605567",
                "description": null,
                "content_type": "task",
                "created_by": "v_xuanfang",
                "created": "2020-06-10 19:30:56",
                "group_id": null,
                "source": ""
            }
        },
        {
            "LaunchAccessory": {
                "id": "1010104801000253479",
                "form_id": "1010104801000402051",
                "workspace_id": "10104801",
                "type": "launch_stories_list",
                "tag": null,
                "title": "需求列表",
                "content": "1010104801854801053,1010104801854801021",
                "description": null,
                "content_type": "story",
                "created_by": "v_xuanfang",
                "created": "2020-06-10 19:30:56",
                "group_id": null,
                "source": ""
            }
        },
        {
            "LaunchAccessory": {
                "id": "1010104801000253477",
                "form_id": "1010104801000402051",
                "workspace_id": "10104801",
                "type": "launch_url",
                "tag": null,
                "title": "URL",
                "content": "baidu.com",
                "description": null,
                "content_type": null,
                "created_by": "v_xuanfang",
                "created": "2020-06-10 19:30:56",
                "group_id": null,
                "source": ""
            }
        },
        {
            "LaunchAccessory": {
                "id": "1010104801000194339",
                "form_id": "1010104801000402051",
                "workspace_id": "10104801",
                "type": "launch_attachment_group",
                "tag": "attachment_group",
                "title": "附件",
                "content": "1010104801000194337",
                "description": "attachement_list",
                "content_type": "attachment_list",
                "created_by": "anyechen",
                "created": "2017-09-01 11:11:27",
                "group_id": null,
                "source": ""
            }
        },
        {
            "LaunchAccessory": {
                "id": "1010104801000194337",
                "form_id": "1010104801000402051",
                "workspace_id": "10104801",
                "type": "launch_attachment",
                "tag": "attachment",
                "title": "raingeek.jpg",
                "content": "/data/home/anyechen/public_html/tapd/app/webroot/files/launch_attachments/10104801/",
                "description": null,
                "content_type": "image/jpeg",
                "created_by": "anyechen",
                "created": "2017-09-01 11:11:27",
                "group_id": "1010104801000194339",
                "source": ""
            }
        }
    ],
    "info": "success"
}
```


# 返回字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">评审依据ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">所属项目ID</td></tr><tr><td style="text-align:center;">form_id</td><td style="text-align:center;">评审单ID</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">类型列表</td></tr><tr><td style="text-align:center;">title</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">tag</td><td style="text-align:center;">附件</td></tr><tr><td style="text-align:center;">content</td><td style="text-align:center;">评审依据数据</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">content_type</td><td style="text-align:center;">数据类型</td></tr><tr><td style="text-align:center;">created_by</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr></tbody></table>
