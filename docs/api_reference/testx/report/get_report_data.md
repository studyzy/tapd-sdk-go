# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/report/get_report_data.html

获取报告详情数据

# url

`https://api.tapd.cn/api/testx/report/v1/namespaces/{namespace}/reports/{report_uid}/templates/{template_uid}/data`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

GET

# 请求数限制

一次获取一条数据

# 请求参数

## Request


| 参数名称 | 类型 | 含义 |
| --- | --- | --- |
| TemplateUid | string | 模板唯一标识 |
| ReportUid | string | 报告唯一标识 |
| Namespace | string | 命名空间 |

# 返回结果


```json
{
"RequestId": "",
"Error": null,
"Data": "{\"total_case_count\":4,\"succeed_case_rate\":25,\"regression_total_case_count\":4,\"regression_succeed_case_rate\":25,\"regression_not_test_case_count\":1,\"regression_executed_case_rate\":75,\"execution_rate\":75,\"automate_case_count\":0,\"case_automate_rate\":0,\"tested_case_count\":3,\"succeed_case_count\":1,\"regression_succeed_case_count\":1,\"fail_case_count\":1,\"block_case_count\":1,\"retry_case_count\":0,\"ignore_case_count\":0,\"error_case_count\":0,\"unknow_tester_case_count\":0,\"p0_case_count\":0,\"p0_success_case_count\":0,\"p0_success_case_rate\":0,\"total_bug_count\":4,\"solved_serious_bug_count\":0,\"serious_bug_solve_rate\":0,\"bug_solve_rate\":0,\"regression_solved_serious_bug_count\":0,\"regression_total_bug_count\":3,\"regression_serious_bug_solve_rate\":0,\"regression_bug_solve_rate\":0,\"closed_bug_count\":0,\"regression_closed_bug_count\":0,\"serious_bug_count\":0,\"regression_serious_bug_count\":0,\"fatal_bug_count\":0,\"normal_bug_count\":0,\"prompt_bug_count\":0,\"advice_bug_count\":1,\"empty_bug_count\":3,\"bug_state_2_count\":{\"新\":4},\"start_time\":\"2025-07-16 18:01:22\",\"end_time\":\"\",\"spend_time\":\"4d 20h 46m\",\"iteration_name\":\"\",\"tester_list\":[\"test\"],\"omitempty\":null,\"producer_list\":[\"test\"],\"project_list\":null,\"product_list\":null,\"work_schedule\":\"\",\"plan_list\":[{\"uid\":\"152324\",\"folder_uid\":\"152316\",\"name\":\"ivantest2121\",\"case_count\":4,\"progress\":75,\"success_rate\":25,\"state\":\"RUNNING\",\"testers\":[\"734242230\"]}],\"test_result_summary\":\"\",\"total_test_rate\":75,\"problem_risk\":\"\",\"bug_open_img\":{\"tooltip\":{\"trigger\":\"item\"},\"legend\":{\"itemGap\":20,\"orient\":\"vertical\",\"top\":\"center\",\"left\":\"60%\",\"textStyle\":{\"color\":\"#606c80\",\"fontSize\":14},\"formatter\":\"${name}  |  ${percent}%  ${value}\",\"icon\":\"circle\",\"itemWidth\":10,\"itemHeight\":10},\"series\":[{\"left\":\"10%\",\"type\":\"pie\",\"width\":\"30%\",\"radius\":[\"60%\",\"70%\"],\"avoidLabelOverlap\":false,\"label\":{\"show\":false,\"position\":\"center\"},\"labelLine\":{\"show\":false},\"data\":[{\"name\":\"新\",\"value\":4,\"percent\":100}]}],\"xAxis\":{\"type\":\"\",\"data\":null},\"yAxis\":{\"type\":\"\",\"boundaryGap\":null,\"min\":0}},\"bug_tester_img\":{\"tooltip\":{\"show\":true},\"xAxis\":{\"type\":\"category\",\"data\":[\"\",\"test\"]},\"yAxis\":{\"type\":\"value\",\"boundaryGap\":[0.2,0.1],\"min\":0},\"color\":[\"#80b3ff\"],\"grid\":{\"left\":40,\"right\":10,\"top\":30,\"bottom\":30},\"series\":[{\"data\":[1,1,2],\"type\":\"bar\",\"barWidth\":\"auto\",\"barMaxWidth\":24,\"label\":{\"show\":true,\"position\":\"top\"}}]},\"bug_handler_img\":{\"tooltip\":{\"show\":true},\"xAxis\":{\"type\":\"category\",\"data\":[\"test\",\"未分配\"]},\"yAxis\":{\"type\":\"value\",\"boundaryGap\":[0.2,0.1],\"min\":0},\"color\":[\"#80b3ff\"],\"grid\":{\"left\":40,\"right\":10,\"top\":30,\"bottom\":30},\"series\":[{\"data\":[1,3],\"type\":\"bar\",\"barWidth\":\"auto\",\"barMaxWidth\":24,\"label\":{\"show\":true,\"position\":\"top\"}}]},\"bug_level_img\":{\"tooltip\":{\"trigger\":\"item\"},\"color\":[\"#f64d3e\",\"#fba337\",\"#fbd341\",\"#2dd36f\",\"#c5cedb\",\"#e6e9ed\"],\"legend\":{\"itemGap\":20,\"orient\":\"vertical\",\"top\":\"center\",\"left\":\"60%\",\"textStyle\":{\"color\":\"#606c80\",\"fontSize\":14},\"formatter\":\"${name}  |  ${percent}%  ${value}\",\"icon\":\"circle\",\"itemWidth\":10,\"itemHeight\":10,\"data\":[\"建议\",\"空\"]},\"series\":[{\"left\":\"10%\",\"type\":\"pie\",\"width\":\"30%\",\"radius\":[\"60%\",\"70%\"],\"avoidLabelOverlap\":false,\"label\":{\"show\":false,\"position\":\"center\"},\"labelLine\":{\"show\":false},\"data\":[{\"name\":\"建议\",\"value\":1,\"percent\":25},{\"name\":\"空\",\"value\":3,\"percent\":75}]}],\"xAxis\":{\"type\":\"\",\"data\":null},\"yAxis\":{\"type\":\"\",\"boundaryGap\":null,\"min\":0}},\"case_result_img\":{\"tooltip\":{\"trigger\":\"item\"},\"color\":[\"#2dd36f\",\"#f64d3e\",\"#6c6e96\",\"#e6e9ed\"],\"legend\":{\"itemGap\":20,\"orient\":\"vertical\",\"top\":\"center\",\"left\":\"60%\",\"textStyle\":{\"color\":\"#606c80\",\"fontSize\":14},\"formatter\":\"${name}  |  ${percent}%  ${value}\",\"icon\":\"circle\",\"itemWidth\":10,\"itemHeight\":10,\"data\":[\"受阻\",\"失败\",\"通过\",\"未测\"]},\"series\":[{\"left\":\"10%\",\"type\":\"pie\",\"width\":\"30%\",\"radius\":[\"60%\",\"70%\"],\"avoidLabelOverlap\":false,\"label\":{\"show\":false,\"position\":\"center\"},\"labelLine\":{\"show\":false},\"data\":[{\"name\":\"通过\",\"value\":1,\"percent\":25},{\"name\":\"失败\",\"value\":1,\"percent\":25},{\"name\":\"受阻\",\"value\":1,\"percent\":25},{\"name\":\"未测\",\"value\":1,\"percent\":25}]}],\"xAxis\":{\"type\":\"\",\"data\":[]},\"yAxis\":{\"type\":\"\",\"boundaryGap\":[],\"min\":0}},\"case_tester_img\":{\"tooltip\":{\"show\":true},\"xAxis\":{\"type\":\"category\",\"data\":[\"ivanzbzhang\",\"未测\"]},\"yAxis\":{\"type\":\"value\",\"boundaryGap\":[0.2,0.1],\"min\":0},\"color\":[\"#80b3ff\"],\"grid\":{\"left\":40,\"right\":10,\"top\":30,\"bottom\":30},\"series\":[{\"data\":[3,1],\"type\":\"bar\",\"barWidth\":\"auto\",\"barMaxWidth\":24,\"label\":{\"show\":true,\"position\":\"top\"}}]},\"case_type_img\":{\"xAxis\":{\"type\":\"category\",\"data\":[\"手工测试\",\"自动化测试\",\"未测试\"]},\"legend\":{\"top\":\"top\",\"selectedMode\":false,\"icon\":\"circle\",\"itemWidth\":10,\"itemHeight\":10},\"yAxis\":{\"type\":\"value\",\"boundaryGap\":[0.2,0.1],\"min\":0},\"tooltip\":{\"trigger\":\"axis\"},\"color\":[\"#2dd36f\",\"#f64d3e\",\"#6c6e96\",\"#BFD9FF\"],\"grid\":{\"left\":40,\"right\":10,\"top\":30,\"bottom\":30},\"series\":[{\"name\":\"通过\",\"data\":[1,0,0],\"type\":\"bar\",\"barWidth\":24,\"stack\":\"total\",\"label\":{\"show\":false,\"position\":\"top\"},\"markPoint\":{\"data\":null,\"symbolSize\":0,\"symbol\":\"\",\"itemStyle\":{},\"symbolOffset\":null}},{\"name\":\"失败\",\"data\":[1,0,0],\"type\":\"bar\",\"barWidth\":24,\"stack\":\"total\",\"label\":{\"show\":false,\"position\":\"top\"},\"markPoint\":{\"data\":null,\"symbolSize\":0,\"symbol\":\"\",\"itemStyle\":{},\"symbolOffset\":null}},{\"name\":\"受阻\",\"data\":[1,0,0],\"type\":\"bar\",\"barWidth\":24,\"stack\":\"total\",\"label\":{\"show\":false,\"position\":\"top\"},\"markPoint\":{\"data\":null,\"symbolSize\":0,\"symbol\":\"\",\"itemStyle\":{},\"symbolOffset\":null}},{\"name\":\"未测\",\"data\":[0,0,1],\"type\":\"bar\",\"barWidth\":24,\"stack\":\"total\",\"label\":{\"show\":false,\"position\":\"top\"},\"markPoint\":{\"data\":[{\"name\":\"手工测试\",\"value\":3,\"xAxis\":0,\"yAxis\":3},{\"name\":\"自动化测试\",\"value\":0,\"xAxis\":1,\"yAxis\":0},{\"name\":\"未测试\",\"value\":1,\"xAxis\":2,\"yAxis\":1}],\"symbolSize\":20,\"symbol\":\"rect\",\"itemStyle\":{\"color\":\"white\"},\"symbolOffset\":[0,\"-60%\"]}}]},\"tester_case_count_img\":{\"tooltip\":{\"trigger\":\"item\"},\"legend\":{\"itemGap\":20,\"orient\":\"vertical\",\"top\":\"center\",\"left\":\"60%\",\"textStyle\":{\"color\":\"#606c80\",\"fontSize\":14},\"formatter\":\"${name}  |  ${percent}%  ${value}\",\"icon\":\"circle\",\"itemWidth\":10,\"itemHeight\":10,\"data\":[\"未分配\"]},\"series\":[{\"left\":\"10%\",\"type\":\"pie\",\"width\":\"30%\",\"radius\":[\"60%\",\"70%\"],\"avoidLabelOverlap\":false,\"label\":{\"show\":false,\"position\":\"center\"},\"labelLine\":{\"show\":false},\"data\":[{\"name\":\"未分配\",\"value\":4,\"percent\":100}]}]},\"case_execute_details\":[{\"CaseName\":\"2-test2\",\"CaseResult\":\"SUCCEED\",\"Testers\":\"734242230\",\"TestCount\":0,\"Stories\":[{\"Uid\":\"1166136271001007531\",\"Url\":\"66136271\",\"WorkspaceUid\":\"66136271\"}]},{\"CaseName\":\"2-test1\",\"CaseResult\":\"BLOCK\",\"Testers\":\"734242230\",\"TestCount\":0,\"Stories\":[{\"Uid\":\"1166136271001007531\",\"Url\":\"66136271\",\"WorkspaceUid\":\"66136271\"}]},{\"CaseName\":\"1-test2\",\"CaseResult\":\"FAIL\",\"Testers\":\"734242230\",\"TestCount\":0,\"Stories\":[{\"Uid\":\"1166136271001007531\",\"Url\":\"66136271\",\"WorkspaceUid\":\"66136271\"}]},{\"CaseName\":\"1-test1\",\"CaseResult\":\"NONE\",\"Testers\":\"未测\",\"TestCount\":0,\"Stories\":[{\"Uid\":\"1166136271001007531\",\"Url\":\"66136271\",\"WorkspaceUid\":\"66136271\"}]}],\"tester_execute_detail\":null,\"folder_execute_detail\":null,\"folder_execute_progress_detail\":null,\"customize_text\":\"\",\"plan_tapd_list\":[{\"tapd_id\":\"1166136271001007531\",\"namespace\":\"\",\"workspace_id\":\"66136271\",\"tapd_type\":\"STORY\"},{\"tapd_id\":\"1166136271001007530\",\"namespace\":\"\",\"workspace_id\":\"66136271\",\"tapd_type\":\"STORY\"}],\"tapd_list_detail\":null,\"bug_list\":null,\"testcase_repo_list\":null,\"error\":{\"msg\":null,\"module_name\":null},\"template_uid\":\"1742\",\"test_result\":\"\",\"attachments\":null,\"story_list\":[{\"uid\":\"1166136271001007531\",\"workspace_id\":\"66136271\",\"case_pass_rate\":\"25.00%\",\"p0case_pass_rate\":\"\",\"bug_count\":\"3\"},{\"uid\":\"1166136271001007530\",\"workspace_id\":\"66136271\",\"case_pass_rate\":\"\",\"p0case_pass_rate\":\"\",\"bug_count\":\"0\"}]}"
}
```

# 测试报告重要字段说明

## Data


| 参数名称 | 类型 | 含义 |
| --- | --- | --- |
| total_case_count | int32 | 用例总数 |
| succeed_case_rate | float64 | 用例成功率 |
| regression_total_case_count | int32 | 回归用例总数 |
| regression_succeed_case_rate | float64 | 回归用例成功率 |
| regression_not_test_case_count | int32 | 回归用例未执行总数 |
| regression_executed_case_rate | float64 | 回归用例执行进度 |
| execution_rate | float64 | 用例执行进度 |
| automate_case_count | int32 | 自动化用例数 |
| case_automate_rate | float64 | 自动化覆盖率 |
| tested_case_count | int32 | 已测用例 |
| succeed_case_count | int32 | 成功用例 |
| regression_succeed_case_count | int32 | 成功用例 |
| fail_case_count | int32 | 失败用例 |
| block_case_count | int32 | 阻塞用例 |
| retry_case_count | int32 | 重试用例 |
| ignore_case_count | int32 | 忽略用例 |
| error_case_count | int32 | 错误用例 |
| unknow_tester_case_count | int32 | 未分用例 |
| p0_case_count | int32 | P0用例 |
| p0_success_case_count | int32 | P0成功用例 |
| p0_success_case_rate | float64 | P0成功率 |
| total_bug_count | int32 | bug总数 |
| solved_serious_bug_count | int32 | 严重bug解决数 |
| serious_bug_solve_rate | float64 | 严重缺陷解决率 |
| bug_solve_rate | float64 | 缺陷解决率 |
| regression_solved_serious_bug_count | int32 | 回归用例严重bug解决数 |
| regression_total_bug_count | int32 | 回归用例bug总数 |
| regression_serious_bug_solve_rate | float64 | 回归用例严重缺陷解决率 |
| regression_bug_solve_rate | float64 | 回归用例缺陷解决率 |
| closed_bug_count | int32 | 关闭bug |
| regression_closed_bug_count | int32 | 回归用例关闭bug |
| serious_bug_count | int32 | 严重bug |
| regression_serious_bug_count | int32 | 回归用例严重bug |
| fatal_bug_count | int32 | 致命bug |
| normal_bug_count | int32 | 一般bug |
| prompt_bug_count | int32 | 提示bug |
| advice_bug_count | int32 | 建议bug |
| empty_bug_count | int32 | 空白bug |
| bug_state_2_count | map[string]int32 | bug状态对应数量统计 |
| start_time | string | 开始时间 |
| end_time | string | 结束时间 |
| spend_time | string | 执行耗时 |
| iteration_name | string | 迭代名称 |
| tester_list | []string | 测试人员，displayName |
| omitempty | []string | 测试人员，globalKey |
| producer_list | []string | 开发人员 |
| project_list | []string | 项目经理 |
| product_list | []string | 产品经理 |
| work_schedule | string | 工作进度 |
| plan_list | []PlanInfo | 计划信息+计划进度 |
| test_result_summary | string | 测试总结 |
| total_test_rate | float64 | 测试整体进度 |
| problem_risk | string | 关键问题及风险 |
| customize_text | string | 文本（支持Markdown） |
| template_uid | string | 模板唯一标识 |
| test_result | string | 测试结果 |
| attachments | []Attachment | 附件 |
| story_list | []StoryItem | 需求列表 |