# 获取工作日历设置列表及启用选项

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/get_work_calendar_settings.html

# 说明

获取工作日历设置列表及启用选项

# url

`https://api.tapd.cn/workspaces/get_work_calendar_settings`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目id</td></tr></tbody></table>

# 调用示例及返回结果

## 获取工作日历设置列表及启用选项

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workspaces/get_work_calendar_settings?workspace_id=48464494'`

### 返回结果

```
{
     "status": 1,
    "data": [
        {
            "name": "中国大陆法定工作日",
            "type": "system",
            "enable": true
        },
        {
            "name": "自定义工作日",
            "type": "custom",
            "enable": false
        }
    ],
    "info": "success"
}
```


# 返回字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">name</td><td style="text-align:center;">名称</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">类型</td></tr><tr><td style="text-align:center;">enable</td><td style="text-align:center;">是否启用</td></tr></tbody></table>
