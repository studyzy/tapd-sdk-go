# Get User View List

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/user/get_user_view_list.html

- 获取当前用户需求视图列表返回结果
- 需求视图相关重要字段说明



 
# # 说明

 获取当前用户需求视图列表（无分页）

 
# # url

 https://api.tapd.cn/user_oauth/get_user_view_list

 
# # 支持格式

 JSON/XML（默认JSON格式）

 
# # HTTP请求方式

 GET

 
# # 请求数限制

 - 此接口不支持 Basic Auth 调用
- 仅支持用户态 OAuth Access Token 调用
- 一次返回所有符合条件的值

 
# # 请求参数

 <table><thead><tr><th style="text-align:center;">字段名</th> <th style="text-align:center;">必选</th> <th style="text-align:center;">类型及范围</th> <th style="text-align:center;">说明</th> <th style="text-align:center;">特殊规则</th></tr></thead> <tbody><tr><td style="text-align:center;">workspace_id</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">int</td> <td style="text-align:center;">项目ID</td> <td style="text-align:center;">无</td></tr> <tr><td style="text-align:center;">type</td> <td style="text-align:center;"><code>否</code></td> <td style="text-align:center;">string</td> <td style="text-align:center;">对象类型（目前只支持 story）</td> <td style="text-align:center;">无</td></tr></tbody></table>

 
# # 调用示例及返回结果

 
## # 获取当前用户需求视图列表

 
### # 返回结果

 ```
{
    "status": 1,
    "data": {
        "1000000000000000016": {
            "id": "1010104801016212067",
            "title": "所有的",
            "enable": "1",
            "type": "system",
            "default_show": "true",
            "view_id": "1000000000000000016",
            "sort": "1"
        },
        "1000000000000000017": {
            "id": "1010104801016212069",
            "title": "我负责的",
            "enable": "1",
            "type": "system",
            "default_show": "false",
            "view_id": "1000000000000000017",
            "sort": "2"
        },
        "1000000000000000413": {
            "id": "1010104801016212089",
            "title": "我关注的",
            "enable": "1",
            "type": "system",
            "default_show": "false",
            "view_id": "1000000000000000413",
            "sort": "3"
        },
        "1000000000000000018": {
            "id": "1010104801016212073",
            "title": "我创建的",
            "enable": "1",
            "type": "system",
            "default_show": "false",
            "view_id": "1000000000000000018",
            "sort": "4"
        },
        "1000000000000000100": {
            "id": "1010104801016212085",
            "title": "待规划需求Backlog",
            "enable": "1",
            "type": "system",
            "default_show": "false",
            "view_id": "1000000000000000100",
            "sort": "10"
        }
    },
    "info": "success"
}
```

 123456789101112131415161718192021222324252627282930313233343536373839404142434445464748495051
# # 相关字段说明

 
## # 需求视图相关重要字段说明

 <table><thead><tr><th style="text-align:center;">字段</th> <th style="text-align:center;">说明</th></tr></thead> <tbody><tr><td style="text-align:center;">id</td> <td style="text-align:center;">需求id</td></tr> <tr><td style="text-align:center;">title</td> <td style="text-align:center;">标题</td></tr> <tr><td style="text-align:center;">enable</td> <td style="text-align:center;">是否启用</td></tr> <tr><td style="text-align:center;">type</td> <td style="text-align:center;">类型</td></tr> <tr><td style="text-align:center;">default_show</td> <td style="text-align:center;">是否默认显示</td></tr> <tr><td style="text-align:center;">view_id</td> <td style="text-align:center;">视图ID</td></tr> <tr><td style="text-align:center;">sort</td> <td style="text-align:center;">排序</td></tr></tbody></table>
