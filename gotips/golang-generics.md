# golang泛型

- 喜大普奔！在golang1.18版本 ，引入了泛型
- [官方指引](https://go.dev/doc/tutorial/generics)
- [WithMethod](https://tip.golang.org/ref/spec#IdentifierList)
    - 概念介绍
        - 泛型（Generics）是指在定义函数、接口或类的时候，不预先指定具体的类型，而在使用的时候再指定类型的一种特性
        - ~表示了底层为相关类型的所有集合
        - 类型参数定义
  ```
    [P any]
    [S interface{ ~[]byte|string }]
    [S ~[]E, E any]
    [P Constraint[int]]
    [_ any]
  ```
- 实现方式
  - 
- tips
    - 常用类型
        - comparable(不能作为value or variable使用)
        - any(interface{})
- best practice
    - [Daily Try](../go_try/go_generics.go)