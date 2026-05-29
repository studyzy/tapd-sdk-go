# Get User Todo Bug

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/user/get_user_todo_bug.html

- 返回用户待办的缺陷返回结果
- 缺陷重要字段说明
- 缺陷优先级(priority)字段说明
- 缺陷严重程度(severity)字段说明
- 缺陷解决方法(resolution)字段说明



 
# # 说明

 返回用户待办的缺陷（分页显示，默认一页30条）

 
# # url

 https://api.tapd.cn/user_oauth/get_user_todo_bug

 
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

 
## # 返回用户待办的缺陷

 
### # 返回结果

 ```
{
    "status": 1,
    "data": [
        {
            "Bug": {
                "id": "1010158231500706019",
                "name": "复制的xxxxxxx",
                "owner": "anyechen;"
            }
        },
        {
            "Bug": {
                "id": "1010158231500687937",
                "name": "story_created_by_api",
                "owner": "anyechen;"
            }
        },
        {
            "Bug": {
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
# # 缺陷字段说明

 
## # 缺陷重要字段说明

 <table><thead><tr><th style="text-align:center;">字段</th> <th style="text-align:center;">说明</th></tr></thead> <tbody><tr><td style="text-align:center;">id</td> <td style="text-align:center;">ID</td></tr> <tr><td style="text-align:center;">title</td> <td style="text-align:center;">标题</td></tr> <tr><td style="text-align:center;">priority</td> <td style="text-align:center;">优先级</td></tr> <tr><td style="text-align:center;">severity</td> <td style="text-align:center;">严重程度</td></tr> <tr><td style="text-align:center;">status</td> <td style="text-align:center;">状态</td></tr> <tr><td style="text-align:center;">iteration_id</td> <td style="text-align:center;">迭代</td></tr> <tr><td style="text-align:center;">module</td> <td style="text-align:center;">模块</td></tr> <tr><td style="text-align:center;">feature</td> <td style="text-align:center;">特性</td></tr> <tr><td style="text-align:center;">release_id</td> <td style="text-align:center;">发布计划</td></tr> <tr><td style="text-align:center;">version_report</td> <td style="text-align:center;">发现版本</td></tr> <tr><td style="text-align:center;">version_test</td> <td style="text-align:center;">验证版本</td></tr> <tr><td style="text-align:center;">version_fix</td> <td style="text-align:center;">合入版本</td></tr> <tr><td style="text-align:center;">version_close</td> <td style="text-align:center;">关闭版本</td></tr> <tr><td style="text-align:center;">baseline_find</td> <td style="text-align:center;">发现基线</td></tr> <tr><td style="text-align:center;">baseline_join</td> <td style="text-align:center;">合入基线</td></tr> <tr><td style="text-align:center;">baseline_test</td> <td style="text-align:center;">验证基线</td></tr> <tr><td style="text-align:center;">baseline_close</td> <td style="text-align:center;">关闭基线</td></tr> <tr><td style="text-align:center;">current_owner</td> <td style="text-align:center;">处理人</td></tr> <tr><td style="text-align:center;">cc</td> <td style="text-align:center;">抄送人</td></tr> <tr><td style="text-align:center;">reporter</td> <td style="text-align:center;">创建人</td></tr> <tr><td style="text-align:center;">participator</td> <td style="text-align:center;">参与人</td></tr> <tr><td style="text-align:center;">te</td> <td style="text-align:center;">测试人员</td></tr> <tr><td style="text-align:center;">de</td> <td style="text-align:center;">开发人员</td></tr> <tr><td style="text-align:center;">auditer</td> <td style="text-align:center;">审核人</td></tr> <tr><td style="text-align:center;">confirmer</td> <td style="text-align:center;">验证人</td></tr> <tr><td style="text-align:center;">fixer</td> <td style="text-align:center;">修复人</td></tr> <tr><td style="text-align:center;">closer</td> <td style="text-align:center;">关闭人</td></tr> <tr><td style="text-align:center;">lastmodify</td> <td style="text-align:center;">最后修改人</td></tr> <tr><td style="text-align:center;">created</td> <td style="text-align:center;">创建时间</td></tr> <tr><td style="text-align:center;">in_progress_time</td> <td style="text-align:center;">接受处理时间</td></tr> <tr><td style="text-align:center;">resolved</td> <td style="text-align:center;">解决时间</td></tr> <tr><td style="text-align:center;">verify_time</td> <td style="text-align:center;">验证时间</td></tr> <tr><td style="text-align:center;">closed</td> <td style="text-align:center;">关闭时间</td></tr> <tr><td style="text-align:center;">reject_time</td> <td style="text-align:center;">拒绝时间</td></tr> <tr><td style="text-align:center;">modified</td> <td style="text-align:center;">最后修改时间</td></tr> <tr><td style="text-align:center;">begin</td> <td style="text-align:center;">预计开始</td></tr> <tr><td style="text-align:center;">due</td> <td style="text-align:center;">预计结束</td></tr> <tr><td style="text-align:center;">assigned_time</td> <td style="text-align:center;">分配时间</td></tr> <tr><td style="text-align:center;">deadline</td> <td style="text-align:center;">解决期限</td></tr> <tr><td style="text-align:center;">flows</td> <td style="text-align:center;">流转过的状态</td></tr> <tr><td style="text-align:center;">os</td> <td style="text-align:center;">操作系统</td></tr> <tr><td style="text-align:center;">platform</td> <td style="text-align:center;">软件平台</td></tr> <tr><td style="text-align:center;">testmode</td> <td style="text-align:center;">测试方式</td></tr> <tr><td style="text-align:center;">testphase</td> <td style="text-align:center;">测试阶段</td></tr> <tr><td style="text-align:center;">testtype</td> <td style="text-align:center;">测试类型</td></tr> <tr><td style="text-align:center;">source</td> <td style="text-align:center;">缺陷根源</td></tr> <tr><td style="text-align:center;">bugtype</td> <td style="text-align:center;">缺陷类型</td></tr> <tr><td style="text-align:center;">frequency</td> <td style="text-align:center;">重现规律</td></tr> <tr><td style="text-align:center;">originphase</td> <td style="text-align:center;">发现阶段</td></tr> <tr><td style="text-align:center;">sourcephase</td> <td style="text-align:center;">引入阶段</td></tr> <tr><td style="text-align:center;">resolution</td> <td style="text-align:center;">解决方法</td></tr> <tr><td style="text-align:center;">estimate</td> <td style="text-align:center;">预计解决时间</td></tr> <tr><td style="text-align:center;">description</td> <td style="text-align:center;">详细描述</td></tr> <tr><td style="text-align:center;">workspace_id</td> <td style="text-align:center;">项目ID</td></tr> <tr><td style="text-align:center;">story_id</td> <td style="text-align:center;">需求ID（当前缺陷已转需求时才会有值）</td></tr></tbody></table>

 
## # 缺陷优先级(priority)字段说明

 <table><thead><tr><th style="text-align:center;">取值</th> <th style="text-align:center;">字面值</th></tr></thead> <tbody><tr><td style="text-align:center;">urgent</td> <td style="text-align:center;">紧急</td></tr> <tr><td style="text-align:center;">high</td> <td style="text-align:center;">高</td></tr> <tr><td style="text-align:center;">medium</td> <td style="text-align:center;">中</td></tr> <tr><td style="text-align:center;">low</td> <td style="text-align:center;">低</td></tr> <tr><td style="text-align:center;">insignificant</td> <td style="text-align:center;">无关紧要</td></tr></tbody></table>

 
## # 缺陷严重程度(severity)字段说明

 <table><thead><tr><th style="text-align:center;">取值</th> <th style="text-align:center;">字面值</th></tr></thead> <tbody><tr><td style="text-align:center;">fatal</td> <td style="text-align:center;">致命</td></tr> <tr><td style="text-align:center;">serious</td> <td style="text-align:center;">严重</td></tr> <tr><td style="text-align:center;">normal</td> <td style="text-align:center;">一般</td></tr> <tr><td style="text-align:center;">prompt</td> <td style="text-align:center;">提示</td></tr> <tr><td style="text-align:center;">advice</td> <td style="text-align:center;">建议</td></tr></tbody></table>

 
## # 缺陷解决方法(resolution)字段说明

 <table><thead><tr><th style="text-align:center;">取值</th> <th style="text-align:center;">字面值</th></tr></thead> <tbody><tr><td style="text-align:center;">ignore</td> <td style="text-align:center;">无需解决</td></tr> <tr><td style="text-align:center;">fix later</td> <td style="text-align:center;">延期解决</td></tr> <tr><td style="text-align:center;">failed</td> <td style="text-align:center;">无法重现</td></tr> <tr><td style="text-align:center;">external</td> <td style="text-align:center;">外部原因</td></tr> <tr><td style="text-align:center;">duplicated</td> <td style="text-align:center;">重复</td></tr> <tr><td style="text-align:center;">intentional</td> <td style="text-align:center;">设计如此</td></tr> <tr><td style="text-align:center;">unclear</td> <td style="text-align:center;">问题描述不准确</td></tr> <tr><td style="text-align:center;">hold</td> <td style="text-align:center;">挂起</td></tr> <tr><td style="text-align:center;">feature</td> <td style="text-align:center;">需求变更</td></tr> <tr><td style="text-align:center;">fixed</td> <td style="text-align:center;">已解决</td></tr> <tr><td style="text-align:center;">transferred to story</td> <td style="text-align:center;">已转需求</td></tr></tbody></table>
