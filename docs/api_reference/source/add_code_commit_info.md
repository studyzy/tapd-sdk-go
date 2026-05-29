# 保存Commit提交数据

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/source/add_code_commit_info.html

# 说明

保存并关联Commit数据，接口会解析commit message，绑定到对应的 需求、缺陷、任务上

# url

`https://api.tapd.cn/code_commit_infos`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次插入一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">commit_id</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">提交ID</td></tr><tr><td style="text-align:center;">author</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">代码提交人</td></tr><tr><td style="text-align:center;">message</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">提交信息</td></tr><tr><td style="text-align:center;">files</td><td style="text-align:center;">是</td><td style="text-align:center;">array</td><td style="text-align:center;">变更文件</td></tr><tr><td style="text-align:center;">repo</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">仓库名</td></tr><tr><td style="text-align:center;">repo_id</td><td style="text-align:center;">是</td><td style="text-align:center;">string/integer</td><td style="text-align:center;">仓库ID</td></tr><tr><td style="text-align:center;">commit_time</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">提交时间，2021-10-25 00:00:00</td></tr><tr><td style="text-align:center;">git_env</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">信息来源，支持类型为:github gitlab svn p4 默认gitlab</td></tr><tr><td style="text-align:center;">repo_url</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">仓库链接，用户界面仓库跳转，留空则不提供跳转能力</td></tr><tr><td style="text-align:center;">commit_url</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">提交链接，用户界面上commit提交跳转，留空则不提供跳转能力</td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d '{"workspace_id":20375571,"commit_id":"zxxxxx","commit_time":"2019-07-22 19:11:11","author":"terrysxu","repo":"repos/xxx_proj","repo_id":"abcd1234-avcd-1234-avcd-1234abcdefgh","message":"--story=854927829 ASA-c2s","files":["U xxx.php","A xxx.js","M xxx.html"]}' 'https://api.tapd.cn/code_commit_infos'`

## 返回结果

```
{
    "status": 1,
    "data": {
        "id": "1020375571000321465",
        "hook_user_name": "terrysxu",
        "commit_id": "xxxxx",
        "user_name": "",
        "workspace_id": 20375571,
        "message": "--story=854927829 ASA-c2s",
        "path": "",
        "web_url": "",
        "hook_project_name": "repos/xxx_proj",
        "commit_time": "2019-07-22 19:11:11 +0800 (Mon, 22 Jul 2019)",
        "ref": "",
        "git_env": "CodeGit",
        "file_commit": "{\"U\":[\"xxx.php\"],\"A\":[\"xxx.js\"],\"M\":[\"xxx.html\"]}",
        "repo_id": "abcd1234-avcd-1234-avcd-1234abcdefgh",
        "related": [
            {
                "type": "story",
                "object_id": "1020375571854927829",
                "commit_id": "1020375571000321465",
                "workspace_id": 20375571,
                "code": "SUCCESS"
            }
        ]
    },
    "info": "success"
}
```


# 字段说明

## 重要字段说明

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">commit_id</td><td style="text-align:center;">提交的版本ID</td></tr><tr><td style="text-align:center;">files</td><td style="text-align:center;">变更文件，数组格式。通过空格分隔操作符和文件名。A表示添加的文件，U表示更新的文件，D表示删除的文件</td></tr></tbody></table>
