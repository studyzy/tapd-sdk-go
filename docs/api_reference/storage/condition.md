# Condition

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/storage/condition.html

公共存储查询、更新、删除接口支持传递condition参数做数据的过滤，目前支持以下条件配置：

 1. =, >, >=, <, <=, !=
2. IN, NOT IN
3. LIKE, NOT LIKE
4. AND, OR 多级嵌套

 
## # 基础条件配置

 所有基础条件，均可以使用以下语法表示。其中：

 - field1, field2为条件对应的字段名，他们之间是AND的关系
- op是当前字段对应的条件，取值范围为=, >, >=, <, <=, !=, IN, NOT IN, LIKE, NOT LIKE
- val为条件对应的值，如果是op为IN或者NOT IN，要求为数组，其他均为数字或字符串

 例如下面的条件等价于WHERE field1='tapd' AND field2 IN (138, 156)。

 ```
"condition": {
    "field1": {
        "op": "=",
        "val": "tapd"
    },
    "field2": {
        "op": "IN",
        "val": [138, 156]
    }
}
```

 12345678910
## # 便捷使用

 对常用的条件=和IN提供了便捷的配置方式。

 condition对象中使用键值对，其中键为条件对应的字段名，值如果是常量，会构造=条件，值如果是数组类型，会构造IN条件。

 例如下面的条件等价于WHERE field1='tapd' AND field2 IN (138, 156)。

 ```
"condition": {
    "field1": "tapd",
    "field2": [138, 156]
}
```

 1234
## # 嵌套条件

 如果希望构造更复杂的条件，包括OR、多级嵌套，需要使用如下语法。

 - AND： 多个子条件之间使用AND连接
- OR： 多字子条件之间使用OR连接

 
### # AND

 下面的条件配置等价于基础条件配置中的示例，看起来似乎变得更复杂，但可能会使构造条件的过程更加清晰。

 ```
"condition": {
    "AND": [
        {
            "field1": "tapd"
        },
        {
            "field2": [138, 156]
        }
    ]
}
```

 12345678910
### # OR

 下面的条件配置等价于WHERE field1="tapd" OR field2 IN (138, 156)。

 ```
"condition": {
    "OR": [
        {
            "field1": "tapd"
        },
        {
            "field2": [138, 156]
        }
    ]
}
```

 12345678910
### # 复杂嵌套

 更复杂的场景是， 可能需要构造由多个OR、AND嵌套组成的条件，例如WHERE (field1="tapd" OR field2 IN (138, 156) OR (field3 LIKE "tapd%" AND field4>100)) AND field5!=0可以使用如下的条件表示。

 ```
"condition": {
    "OR": [
        {
            "field1": "tapd"
        },
        {
            "field2": [138, 156]
        },
        {
            "field3": {
                "op": "LIKE",
                "val": "tapd%"
            },
            "field4": {
                "op": ">",
                "val": "100"
            }
        }
    ],
    "field5": {
        "op": "!=",
        "val": 0
    }
}
```

 123456789101112131415161718192021222324
