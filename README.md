# starlark helper

## 数据类型

根据数据目标的不同，提供三种数据类型

+ Function 函数，实现了 Callable
+ Struct 结构体，实现了 HasAttrs
+ Module 模块，基于 starlarkstruct.Module 封装

用户也可以定义自己的类型，并实现 CanHelp 接口来提供更多帮助信息