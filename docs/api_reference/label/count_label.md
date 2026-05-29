# # 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/label/count_label.html

计算符合查询条件的标签数量并返回

  
# # url

 https://api.tapd.cn/label/count

 
# # 支持格式

 JSON/XML（默认JSON格式）

 
# # HTTP请求方式

 GET

 
# # 请求数限制

 只返回标签数量

 
# # 请求参数

 <table><thead><tr><th style="text-align:center;">字段名</th> <th style="text-align:center;">必选</th> <th style="text-align:center;">类型及范围</th> <th style="text-align:center;">说明</th> <th style="text-align:center;">特殊规则</th></tr></thead> <tbody><tr><td style="text-align:center;">workspace_id</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">integer</td> <td style="text-align:center;">项目ID</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">id</td> <td style="text-align:center;">否</td> <td style="text-align:center;">integer</td> <td style="text-align:center;">id</td> <td style="text-align:center;">支持多ID查询</td></tr> <tr><td style="text-align:center;">name</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">标签名称</td> <td style="text-align:center;">支持模糊匹配</td></tr> <tr><td style="text-align:center;">creator</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">创建人</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">created</td> <td style="text-align:center;">否</td> <td style="text-align:center;">datetime</td> <td style="text-align:center;">创建时间</td> <td style="text-align:center;">支持时间查询</td></tr></tbody></table>

 
# # 调用示例及返回结果

 
## # 获取项目下标签的数量

 
### # curl 使用 Basic Auth 鉴权调用示例

 curl -u 'api_user:api_password' 'https://api.tapd.cn/label/count?workspace_id=10158231'

 curl 'https://api.tapd.cn/label/count?workspace_id=10158231&access_token=ACCESS_TOKEN'

 
### # 返回结果

 ```
{
    "status": 1,
    "data": {
        "count": 1
    },
    "info": "success"
}
```

 1234567
