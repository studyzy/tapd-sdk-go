# 批量修改需求的保密信息

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/batch_update_secret_info.html

# 说明

批量修改需求的保密信息

# url

`https://api.tapd.cn/stories/batch_update_secret_info`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

当一次修改涉及20个及以上的需求时，会触发后台异步执行，不会立即生效。

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">story_id_list</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">所需修改的需求id列表, 用"|"隔开多个</td></tr><tr><td style="text-align:center;">secret_scope</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">修改后的保密状态(public:设置为公开 | secret:设置为保密)</td></tr><tr><td style="text-align:center;">allow_list</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">保密白名单(支持填入用户nick和用户组id), 用";"隔开多个</td></tr><tr><td style="text-align:center;">add_participant_fields</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">是否将需求树的参与人动态纳入到保密范围内 (true:纳入|false:不纳入)</td></tr><tr><td style="text-align:center;">operation_type</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">保密范围(上述allow_list)的操作模式(0:默认模式/覆盖模式, 用传入的新名单覆盖旧名单 | 1:新增模式, 在旧名单基础上进行新增此次传入的名单 | 2:删除模式, 从旧名单里删除此次传入的名单)</td></tr><tr><td style="text-align:center;">current_user</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">执行此操作的用户的nick</td></tr></tbody></table>

# 调用示例及返回结果

## 修改需求1010104801871430407和1010104801871430409的保密信息

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10104801&story_id_list=1010104801871430407|1010104801871430409&secret_scope=secret&allow_list=xinweihe;1000000000000000002&add_participant_fields=false&operation_type=0&current_user=xinweihe' 'https://api.tapd.cn/stories/batch_update_secret_info'`

### 返回结果

```json
{
    "status": 1,
    "data": {
        "code": "succeed",
        "msg": "需求可访问人员设置成功"
    },
    "info": "success"
}
```
