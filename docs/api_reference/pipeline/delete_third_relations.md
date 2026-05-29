# # 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/pipeline/delete_third_relations.html

解除指定构建记录与业务对象关联。接口沿用之前的RESTful风格，请求列表接口后得到关联关系主键ID，以DELETE方式请求此接口进行删除。

  
## # SDK 方法名

 nodeJs:

 SDK使用方式  (opens new window)

 插件中使用SDK

 方法名::

 ```
deleteThirdRelations
```

 1 
# # url

 https://api.tapd.cn/third_relations

 
# # HTTP请求方式

 DELETE

 
# # 请求数限制

 只能传关联列表项的主键ID，一次只解除一个关联

 
# # 请求参数

 <table><thead><tr><th style="text-align:center;">字段名</th> <th style="text-align:center;">必选</th> <th style="text-align:center;">类型及范围</th> <th style="text-align:center;">说明</th> <th style="text-align:center;">特殊规则</th></tr></thead> <tbody><tr><td style="text-align:center;">id</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">integer</td> <td style="text-align:center;">关联ID，列表接口返回的id参数</td> <td style="text-align:center;">无</td></tr> <tr><td style="text-align:center;">operator</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">string</td> <td style="text-align:center;">操作人</td> <td style="text-align:center;">英文名，如dobbyzhang</td></tr></tbody></table>

 
# # 调用示例及返回结果

 curl -H 'Authorization: Bearer Access Token' -X DELETE 'https://api.tapd.cn/third_relations?id=99999999&operator=dobbyzhang'

 
### # 返回结果

 ```
{
  "status": 1,
    "data": {
        "result": true
    },
    "info": "success"
}
```

 1234567
# # 返回字段说明

 <table><thead><tr><th style="text-align:center;">字段</th> <th style="text-align:center;">说明</th></tr></thead> <tbody><tr><td style="text-align:center;">status</td> <td style="text-align:center;">状态</td></tr> <tr><td style="text-align:center;">data</td> <td style="text-align:center;">操作状态</td></tr> <tr><td style="text-align:center;">info</td> <td style="text-align:center;">操作信息</td></tr></tbody></table>
