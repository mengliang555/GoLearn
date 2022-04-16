# Golang Runtime
- GC
  - GC的三种状态
    - _GCoff GC未运行
    - _GCmark 标记中，启用写屏障
    - _GCmarktermination 标记终止，启用写屏障
    ```golang
    const (
        _GCoff             = iota // GC not running; sweeping in background, write barrier disabled
        _GCmark                   // GC marking roots and workbufs: allocate black, write barrier ENABLED
        _GCmarktermination        // GC mark termination: allocate black, P's help GC, write barrier ENABLED
    )
    ```
  - golang支持的三种模式
    - gcBackgroundMode 默认模式，标记与清扫过程都是并发执行的
    - gcForceMode 只在清扫阶段支持并发；
    - gcForceBlockMode GC全程需要STW。
    ```
    const (
        gcBackgroundMode gcMode = iota // concurrent GC and sweep
        gcForceMode                    // stop-the-world GC now, concurrent sweep
        gcForceBlockMode               // stop-the-world GC now and STW sweep (forced by user)
    )
    ```
    - GC过程中的全局变量
      - gcController
        - gcController 实现了GC 控制器，该控制器确定了何时触发并发垃圾回收，以及多少个辅助助手(mutator assists)和后台标记
      - work
        - 主要用来记录GC期间的一些工作统计信息，如根对象的数量及大小、GC的次数，标记字节大小以及全局锁等等；
    - GC过程
      - GC过程将会STW
      - 首先获取上一轮的GC次数，然后调用 gcWaitOnMark() 函数确保上次GC的三个阶段 “清扫终止”、”标记” 和 “标记终止“ 完成
      - 



## 参考文章
- [Runtime: Golang GC源码分析](https://blog.haohtml.com/archives/26358)
- [GC 四个阶段](https://blog.haohtml.com/archives/25372)