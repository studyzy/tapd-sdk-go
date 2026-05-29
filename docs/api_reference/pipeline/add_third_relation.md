# # 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/pipeline/add_third_relation.html

创建TAPD业务对象与流水线构建记录关联关系，返回关系数据

  
## # SDK 方法名

 nodeJs:

 SDK使用方式  (opens new window)

 插件中使用SDK

 方法名::

 ```
addThirdRelation
```

 1 
# # url

 https://api.tapd.cn/third_relations

 
# # HTTP请求方式

 POST

 
# # 请求限制

 每次请求操作只允许一个业务对象和一个issue关联

 
# # 请求参数

 <table><thead><tr><th style="text-align:center;">字段名</th> <th style="text-align:center;">必选</th> <th style="text-align:center;">类型</th> <th style="text-align:center;">说明</th> <th style="text-align:center;">特殊规则</th></tr></thead> <tbody><tr><td style="text-align:center;">source_type</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">string</td> <td style="text-align:center;">资源类型</td> <td style="text-align:center;">目前source_type可选值：build</td></tr> <tr><td style="text-align:center;">source_project_id</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">string</td> <td style="text-align:center;">流水线id</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">source_id</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">string</td> <td style="text-align:center;">流水线构建记录id</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">workspace_id</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">int</td> <td style="text-align:center;">tapd项目id</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">tapd_id</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">string</td> <td style="text-align:center;">tapd业务对象id</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">tapd_type</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">string</td> <td style="text-align:center;">tapd业务对象类型</td> <td style="text-align:center;">目前tapd_type可选值：task,story,bug</td></tr> <tr><td style="text-align:center;">operator</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">string</td> <td style="text-align:center;">操作人</td> <td style="text-align:center;">英文名，如jin</td></tr></tbody></table>

 
# # 调用示例及返回结果

 
### # curl 调用示例

 curl -u 'api_user:api_password' -d 'source_type=build&source_project_id=10220750&source_id=89888&workspace_id=20358496&tapd_id=1020358496854819049&tapd_type=story' 'https://api.tapd.cn/third_relations'

 
### # 返回结果

 ```
{
    "status": 1,
    "data": {
        "ThirdRelations": {
            "id": "644",
            "workspace_id": "20358496",
            "source_project_id": "10220750",
            "source_id": "77",
            "source_type": "build",
            "tapd_id": "1020358496854819049",
            "tapd_type": "story",
            "created": "2020-09-01 17:34:54",
            "modified": "2020-09-01 17:34:54",
            "status": "1",
        }
    },
    "info": "success"
}
```

 123456789101112131415161718
# # 字段说明

 
## # 关联关系字段说明

 <table><thead><tr><th style="text-align:center;">字段</th> <th style="text-align:center;">说明</th></tr></thead> <tbody><tr><td style="text-align:center;">id</td> <td style="text-align:center;">关联ID</td></tr> <tr><td style="text-align:center;">workspace_id</td> <td style="text-align:center;">项目ID</td></tr> <tr><td style="text-align:center;">source_project_id</td> <td style="text-align:center;">流水线id</td></tr> <tr><td style="text-align:center;">source_id</td> <td style="text-align:center;">流水线构建记录id</td></tr> <tr><td style="text-align:center;">source_type</td> <td style="text-align:center;">资源类型</td></tr> <tr><td style="text-align:center;">tapd_id</td> <td style="text-align:center;">tapd对象ID</td></tr> <tr><td style="text-align:center;">tapd_type</td> <td style="text-align:center;">tapd对象类型</td></tr> <tr><td style="text-align:center;">created</td> <td style="text-align:center;">创建时间</td></tr> <tr><td style="text-align:center;">modified</td> <td style="text-align:center;">修改时间</td></tr> <tr><td style="text-align:center;">status</td> <td style="text-align:center;">关联状态 1为有效</td></tr></tbody></table>
