# Pprof-分析golang运行状态
- Go语言内置了获取程序运行数据的工具，包括以下两个标准库
  - runtime/pprof: 采集工具型应用运行数据进行分析
  - net/http/pprof: 采集服务型应用运行时数据进行分析
  - pprof开启后，每隔一段时间(10ms)就会收集当前的堆栈信息，获取各个函数占用的CPU以及内存资源，然后通过对这些采样数据进行分析，形成一个性能分析报告。
- 性能优化主要有一下几个方面：
  - CPU Profile：报告程序的CPU使用情况，按照一定频率去采集应用程序在CPU和寄存器上面的数据。
  - Memory Profile（Heap Profile）:报告程序的内存使用情况。
  - Block Profiling: 报告goroutines不在运行状态的情况，可以用来分析和查找死锁等性能瓶颈。
  - Goroutine Profiling: 报告goroutines的使用情况，有哪些roroutines，它们的调用关系是怎样的。