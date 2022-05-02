# SystemDesign
## 4S分析方法
- Scenario（场景）
  - 功能设计（需要哪些功能
    - 核心功能-》主流程设计实现
    - 扩展功能-》旁路功能（可不分析
  - 访问量（并发访问量，短期内用户数据量
    - QPS分析
      - 并发用户数量
      - 读QPS
      - 写QPS
    - QPS与服务器资源
      - 一台Web Server承受量约为 1K的QPS（考虑到逻辑处理时间以及数据库查询的瓶颈）
      - 一台SQL Database承受量约为 1K的QPS（如果JOIN和INDEX query比较多的话，这个值会更小）
      - 一台 NoSQL Database (Cassandra) 约承受量是 10k 的 QPS
      - 一台 NoSQL Database (Memcached) 约承受量是 1M 的 QPS
- Service（服务）
  - 逻辑整合：根据具体功能设计case by case
- Storage（存储）
  - 根据每个服务的数据特性选择合适的存储结构，然后细化数据表结构
  - 存储类型
    - 数据库系统
    - 文件系统
    - 缓存系统
- Scale（扩展）
  - 优化系统缺陷
    - 更详细的功能设计
    - 异常边界数据处理
  - 系统维护方案
    - 系统鲁棒性
    - 系统扩展性
---
## 这类问题思考的注意点
- Ask before design. 
  - 问清楚再动手设计，不要一上来就冲着一个巨牛的方案去设计
- No more no less. 
  - 不要总想着设计最牛的系统，要设计够用的系统
- Work solution first. 
  - 先设计一个基本能工作的系统，然后再逐步优化
- Analysis is important than solution. 
  - 系统设计没有标准答案，记住答案是没用的，通过分析过程展示知识储备，权衡各种设计方式的利弊。

---
## 可以考虑的技术方案
- Cache：缓存，万金油，哪里不行优先考虑
- Queue：消息队列，常见使用Linkedin的kafka
- Synchronized：批处理＋异步，减少系统IO瓶颈
- Load Balance: 负载均衡，可以使用一致性hash技术做到尽量少的数据迁移
- Parallelization：并行计算，比如MapReduce
- Replication：提高可靠性，如HDFS，基于位置感知的多块拷贝
- Partition：数据库sharding，通过hash取摸
---
## 三高特点
- 高性能架构设计： 熟悉系统常见性能优化手段比如引入 读写分离、缓存、负载均衡、异步 等等。
- 高可用架构设计 ：CAP理论和BASE理论、通过集群来提高系统整体稳定性、超时和重试机制、应对接口级故障：降级、熔断、限流、排队。
- 高扩展架构设计 ：说白了就是懂得如何拆分系统。你按照不同的思路来拆分软件系统，就会得到不同的架构。
---
## 技术选型
- 根据业务选择合适的组件
---
## 设计步骤
- 需求澄清（Requirments Clarification）
- 系统接口定义 （System Interface Definition） 
- 粗略估算 (Back-of-the-envelope Estimation )
- 定义数据模型 （Defining Data Model） 
- 高级定义 (High Level Design)
- 细节设计（Detailed Design） 
- 识别并解决瓶颈（Indentitying and Resovling Bottlenecks）