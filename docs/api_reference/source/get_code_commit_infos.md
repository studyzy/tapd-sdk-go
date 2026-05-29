# 获取GIT关联提交数据(GitCommit)

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/source/get_code_commit_infos.html

# 说明

返回某个业务对象(需求、缺陷、任务)的GIT关联提交的信息

# url

`https://api.tapd.cn/code_commit_infos`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">TAPD业务对象类型（story,bug,task）</td></tr><tr><td style="text-align:center;">object_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">int</td><td style="text-align:center;">TAPD业务对象ID</td></tr><tr><td style="text-align:center;">commit_time</td><td style="text-align:center;"><code>否</code></td><td style="text-align:center;">string</td><td style="text-align:center;">提交时间（格式：&gt;HH-mm-dd HH:ii:ss,&lt;HH-mm-dd HH:ii:ss,HH-mm-dd HH:ii:ss~HH-mm-dd HH:ii:ss）</td></tr><tr><td style="text-align:center;">related_type</td><td style="text-align:center;"><code>否</code></td><td style="text-align:center;">string</td><td style="text-align:center;">关联类型值: all所有，branch 分支关联提交，source_code 源码关联提交; 默认为all</td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">int</td><td style="text-align:center;">设置返回数量限制，默认为30</td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">int</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td></tr></tbody></table>

# 调用示例及返回结果

## 获取TAPD业务对象(需求、缺陷、任务)的GIT关联提交的信息

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/code_commit_infos?workspace_id=20358374&type=story&object_id=1020358374854843133'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "id": "1020358374000262989",
            "user_name": "jeffjffang",
            "user_id": "2071532707",
            "hook_user_name": "jeffjffang",
            "commit_id": "111",
            "workspace_id": "20358374",
            "message": "--story=1020358374854843133 SMARTCOMMIT #fix 111222",
            "path": "http://git.code.oa.com/TAPDG/tapd_test/commit/047e5764c392bef48fd0e4176c147c7c30a9f32a",
            "web_url": "http://git.code.oa.com/TAPDG/tapd_test",
            "hook_project_name": "tapd_test",
            "commit_time": "2019-08-26 15:57:04",
            "created": "2020-11-30 14:31:29",
            "ref": "refs/heads/master",
            "ref_status": "0",
            "git_env": "CodeGit",
            "file_commit":"{\"A\":[],\"M\":[\"tapd_interface_test\\/cases\\/\云\端_\公\司\设\置.jmx\"],\"R\":[]}",
            "repo_id": "20045801",
            "branch_id": "20045801/refs/heads/master",
            "file_sort": {
                "tapd_interface_test/cases/云端_公司设置.jmx": 0
            }
        },
        {
            "id": "1020358374000262991",
            "user_name": "jeffjffang",
            "user_id": "2071532707",
            "hook_user_name": "jeffjffang",
            "commit_id": "jeffjffang1231",
            "workspace_id": "20358374",
            "message": "--story=1020358374854843133 SMARTCOMMIT #fix 111222",
            "path": "http://git.code.oa.com/TAPDG/tapd_test/commit/047e5764c392bef48fd0e4176c147c7c30a9f32a",
            "web_url": "http://git.code.oa.com/TAPDG/tapd_test",
            "hook_project_name": "tapd_test",
            "commit_time": "2019-08-26 15:57:04",
            "created": "2020-11-30 14:31:53",
            "ref": "refs/heads/master",
            "ref_status": "0",
            "git_env": "CodeGit",
            "file_commit":"{\"A\":[],\"M\":[\"tapd_interface_test\\/cases\\/\云\端_\公\司\设\置.jmx\"],\"R\":[]}",
            "repo_id": "20045801",
            "branch_id": "20045801/refs/heads/master",
            "file_sort": {
                "tapd_interface_test/cases/云端_公司设置.jmx": 0
            }
        },
        {
            "id": "1020358374000263277",
            "user_name": "jeffjffang",
            "user_id": "2071532707",
            "hook_user_name": "jeffjffang",
            "commit_id": "047e5764c392bef48fd0e4176c147c7c30a9f32a",
            "workspace_id": "20358374",
            "message": "--story=1020358374854843133 SMARTCOMMIT #fix 0000",
            "path": "http://git.code.oa.com/TAPDG/tapd_test/commit/047e5764c392bef48fd0e4176c147c7c30a9f32a",
            "web_url": "http://git.code.oa.com/TAPDG/tapd_test",
            "hook_project_name": "tapd_test",
            "commit_time": "2019-08-26 15:57:04",
            "created": "2020-12-01 12:15:38",
            "ref": "refs/heads/master",
            "ref_status": "0",
            "git_env": "CodeGit",
            "file_commit":"{\"A\":[],\"M\":[\"tapd_interface_test\\/cases\\/\云\端_\公\司\设\置.jmx\"],\"R\":[]}",
            "repo_id": "20045801",
            "branch_id": "20045801/refs/heads/master",
            "file_sort": {
                "tapd_interface_test/cases/云端_公司设置.jmx": 0
            }
        }
    ],
    "info": "success"
}
```

