# # 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/pipeline/get_third_relations.html

获取TAPD业务对象与构建记录的关联关系信息。

  
## # SDK 方法名

 nodeJs:

 SDK使用方式  (opens new window)

 插件中使用SDK

 方法名::

 ```
getThirdRelations
```

 1 
# # url

 https://api.tapd.cn/third_relations

 
# # HTTP请求方式

 GET

 
# # 支持格式

 JSON/XML（默认JSON格式）

 
# # HTTP请求方式

 GET

 
# # 请求数限制

 一次只能获取一个构建记录与TAPD业务对象的关联关系
通过Git 数据查tapd数据：

 
# # 请求参数

 通过tapd数据查流水线构建记录数据：

 <table><thead><tr><th style="text-align:center;">字段名</th> <th style="text-align:center;">必选</th> <th style="text-align:center;">类型</th> <th style="text-align:center;">说明</th> <th style="text-align:center;">特殊规则</th></tr></thead> <tbody><tr><td style="text-align:center;">source_type</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">string</td> <td style="text-align:center;">Git 资源类型</td> <td style="text-align:center;">目前type可选值：build</td></tr> <tr><td style="text-align:center;">workspace_id</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">int</td> <td style="text-align:center;">tapd项目id</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">tapd_id</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">int</td> <td style="text-align:center;">tapd对象id</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">tapd_type</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">string</td> <td style="text-align:center;">tapd对象类型</td> <td style="text-align:center;">目前可选值：story,task,bug</td></tr></tbody></table>

 
# # 调用示例及返回结果

 
### # curl 调用示例

 curl -u 'api_user:api_password' -d 'source_type=build&tapd_id=1169354332001000176$tapd_type=story' 'https://api.tapd.cn/third_relations'

 
### # 返回结果

 ```
{
    "status": 1,
    "data": [
        {
            "ThirdRelations": {
                "id": "1169354332001000003",
                "workspace_id": "69354332",
                "source_app_id": "1",
                "source_project_id": "111",
                "source_id": "222",
                "source_type": "build",
                "source_data": "",
                "source_tag": "",
                "tapd_id": "1169354332001000176",
                "workitem_type_id": "0",
                "tapd_type": "story",
                "created": "2024-12-05 15:46:33",
                "modified": "2024-12-05 15:46:33",
                "status": "1",
                "operator": "xxx"
            }
        },
        {
            "ThirdRelations": {
                "id": "1169354332001000004",
                "workspace_id": "69354332",
                "source_app_id": "1",
                "source_project_id": "444",
                "source_id": "333",
                "source_type": "build",
                "source_data": "",
                "source_tag": "",
                "tapd_id": "1169354332001000176",
                "workitem_type_id": "0",
                "tapd_type": "story",
                "created": "2024-12-05 15:46:48",
                "modified": "2024-12-05 15:59:03",
                "status": "1",
                "operator": "jin"
            }
        }
    ],
    "info": "success"
}
```

 1234567891011121314151617181920212223242526272829303132333435363738394041424344
# # 字段说明

 
## # 关联关系字段说明

 <table><thead><tr><th style="text-align:center;">字段</th> <th style="text-align:center;">说明</th></tr></thead> <tbody><tr><td style="text-align:center;">id</td> <td style="text-align:center;">关联ID</td></tr> <tr><td style="text-align:center;">workspace_id</td> <td style="text-align:center;">项目ID</td></tr> <tr><td style="text-align:center;">source_project_id</td> <td style="text-align:center;">流水线id</td></tr> <tr><td style="text-align:center;">source_id</td> <td style="text-align:center;">流水线构建记录id</td></tr> <tr><td style="text-align:center;">source_type</td> <td style="text-align:center;">资源类型</td></tr> <tr><td style="text-align:center;">tapd_id</td> <td style="text-align:center;">tapd对象ID</td></tr> <tr><td style="text-align:center;">tapd_type</td> <td style="text-align:center;">tapd对象类型</td></tr> <tr><td style="text-align:center;">created</td> <td style="text-align:center;">创建时间</td></tr> <tr><td style="text-align:center;">modified</td> <td style="text-align:center;">修改时间</td></tr> <tr><td style="text-align:center;">status</td> <td style="text-align:center;">关联状态</td></tr></tbody></table>
