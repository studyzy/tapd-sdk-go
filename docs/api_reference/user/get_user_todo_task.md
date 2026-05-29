# Get User Todo Task

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/user/get_user_todo_task.html

- 返回用户待办的任务返回结果
- 任务重要字段说明
- 任务状态(status)取值字段说明
- 任务优先级(priority)取值字段说明



 
# # 说明

 返回用户待办的任务（分页显示，默认一页30条）

 
# # url

 https://api.tapd.cn/user_oauth/get_user_todo_task

 
# # 支持格式

 JSON/XML（默认JSON格式）

 
# # HTTP请求方式

 GET

 
# # 请求数限制

 - 仅支持用户态 OAuth Access Token 调用
- 默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

 
# # 请求参数

 <table><thead><tr><th style="text-align:center;">字段名</th> <th style="text-align:center;">必选</th> <th style="text-align:center;">类型及范围</th> <th style="text-align:center;">说明</th> <th style="text-align:center;">特殊规则</th></tr></thead> <tbody><tr><td style="text-align:center;">workspace_id</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">integer</td> <td style="text-align:center;">项目ID</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">user</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">用户</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">limit</td> <td style="text-align:center;">否</td> <td style="text-align:center;">integer</td> <td style="text-align:center;">设置返回数量限制，默认为30</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">page</td> <td style="text-align:center;">否</td> <td style="text-align:center;">integer</td> <td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">order</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td> <td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr> <tr><td style="text-align:center;">fields</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td> <td style="text-align:center;"></td></tr></tbody></table>

 
# # 调用示例及返回结果

 
## # 返回用户待办的任务

 
### # 返回结果

 ```
{
    "status": 1,
    "data": [
        {
            "Task": {
                "id": "1010158231500706019",
                "name": "复制的xxxxxxx",
                "owner": "anyechen;"
            }
        },
        {
            "Task": {
                "id": "1010158231500687937",
                "name": "story_created_by_api",
                "owner": "anyechen;"
            }
        },
        {
            "Task": {
                "id": "1010158231500625827",
                "name": "官网登录页改造",
                "owner": "anyechen;"
            }
        }
    ],
    "info": "success"
}
```

 123456789101112131415161718192021222324252627
# # 任务字段说明

 
## # 任务重要字段说明

 <table><thead><tr><th style="text-align:center;">字段</th> <th style="text-align:center;">说明</th></tr></thead> <tbody><tr><td style="text-align:center;">id</td> <td style="text-align:center;">id</td></tr> <tr><td style="text-align:center;">name</td> <td style="text-align:center;">任务标题</td></tr> <tr><td style="text-align:center;">description</td> <td style="text-align:center;">任务详细描述</td></tr> <tr><td style="text-align:center;">workspace_id</td> <td style="text-align:center;">项目ID</td></tr> <tr><td style="text-align:center;">creator</td> <td style="text-align:center;">创建人</td></tr> <tr><td style="text-align:center;">created</td> <td style="text-align:center;">创建时间</td></tr> <tr><td style="text-align:center;">modified</td> <td style="text-align:center;">最后修改时间</td></tr> <tr><td style="text-align:center;">status</td> <td style="text-align:center;">状态</td></tr> <tr><td style="text-align:center;">owner</td> <td style="text-align:center;">任务当前处理人</td></tr> <tr><td style="text-align:center;">cc</td> <td style="text-align:center;">抄送人</td></tr> <tr><td style="text-align:center;">begin</td> <td style="text-align:center;">预计开始</td></tr> <tr><td style="text-align:center;">due</td> <td style="text-align:center;">预计结束</td></tr> <tr><td style="text-align:center;">release_id</td> <td style="text-align:center;">发布计划ID</td></tr> <tr><td style="text-align:center;">story_id</td> <td style="text-align:center;">关联需求的ID</td></tr> <tr><td style="text-align:center;">iteration_id</td> <td style="text-align:center;">所属迭代的ID</td></tr> <tr><td style="text-align:center;">priority</td> <td style="text-align:center;">优先级</td></tr> <tr><td style="text-align:center;">progress</td> <td style="text-align:center;">进度</td></tr> <tr><td style="text-align:center;">completed</td> <td style="text-align:center;">完成时间</td></tr> <tr><td style="text-align:center;">effort_completed</td> <td style="text-align:center;">完成工时</td></tr> <tr><td style="text-align:center;">exceed</td> <td style="text-align:center;">超出工时</td></tr> <tr><td style="text-align:center;">remain</td> <td style="text-align:center;">剩余工时</td></tr> <tr><td style="text-align:center;">effort</td> <td style="text-align:center;">预估工时</td></tr></tbody></table>

 
## # 任务状态(status)取值字段说明

 <table><thead><tr><th style="text-align:center;">取值</th> <th style="text-align:center;">字面值</th></tr></thead> <tbody><tr><td style="text-align:center;">open</td> <td style="text-align:center;">未开始</td></tr> <tr><td style="text-align:center;">progressing</td> <td style="text-align:center;">进行中</td></tr> <tr><td style="text-align:center;">done</td> <td style="text-align:center;">已完成</td></tr></tbody></table>

 
## # 任务优先级(priority)取值字段说明

 <table><thead><tr><th style="text-align:center;">取值</th> <th style="text-align:center;">字面值</th></tr></thead> <tbody><tr><td style="text-align:center;">4</td> <td style="text-align:center;">High</td></tr> <tr><td style="text-align:center;">3</td> <td style="text-align:center;">Middle</td></tr> <tr><td style="text-align:center;">2</td> <td style="text-align:center;">Low</td></tr> <tr><td style="text-align:center;">1</td> <td style="text-align:center;">Nice To Have</td></tr></tbody></table>
