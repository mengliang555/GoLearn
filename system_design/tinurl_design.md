# [TinyUrl](https://blog.csdn.net/NXHYD/article/details/122612092)

- 需求分析
    - URL 缩短用于为长 URL 创建较短的别名。我们将这些缩短的别名称为“短链接”。当用户点击这些短链接时，他们会被重定向到原始 URL。短链接节省了很多显示、打印、发送消息或推文时的空间。
    - 功能需求细化
        - 对于给定一个URL，这个系统需要为它生成一个短且唯一的别名alias，称作短链接
        - 当用户点击这个短链接，这个服务应该能够为它重定向到原始链接。
        - 用户可以有选择地去自定义他们URL的短链接。
        - 这些短链接在某个默认是时间段之后就会失效，用户应该能够去设置这个失效时间。
    - 非功能需求细化
        - 系统应该是高度可用的。因为如果我们的服务关闭，所有 URL 重定向都将开始失败。
        - URL 重定向应该以最小的延迟实时发生。
        - 短链接不应该是可猜测的（也就是说有一定规律被预测到的）
    - 扩展需求
        - 重定向次数
        - 服务提供方式
- 接口定义

```golang
type TinyUrlService interface{
    GererateTinyUrl(devKey string，originUrl string,expireDate int64,customAligns ...string)
    DeleteIinyUrl(devKey string，originUrl string)
    RenewalIinyUrl(devKey string，originUrl string,renewalTime int64)
}
```

- 数据估算
    - 每个月为5Million（开始阶段 的url创建和使用
    - Traffic estimates（读/写=100：1）
        - 创建URL的QPS：5000000/30/24/3600 = 1.92 url/s
        - 重定向的QPD：192 url/s
    - Storage estimates
        - 预设存储时间为1month，单个url为500byte
        - space: 5 million * 500byte = 2.3 GB
    - Bandwidth estimates
        - Write QPS = 1.92*500 ~= 1kb/s
        - Read QPS = 192*500 ~= 100kb/s
    - Memory estimates
        - 如果我们遵循 80-20 规则，即 20% 的 URL 产生 80% 的流量，我们希望缓存这 20% 的热门 URLs
        - For Read QPS for day
            - 192 * 3600 * 24 = 1658w 的请求
            - 1658w * 500 byte * 20% ～= 1.54GB

- 数据存储模型定义
    - 数据特点
        - 记录量很大
        - 单挑记录数据量少
        - 记录之间无关联关系
        - 服务的读QPS很高
    - 表设计
        - 设计两张表
            - 记录用户数据
            - 记录URL的映射关系
        - 数据库使用no sql

- 详细算法设计
    - 对用户提供的替代url进行判重
    - 自生成的字符串
        - Encoding actual URL
            - 使用某种算法，编码 or 压缩原始字符串
        - Generating keys offline
            - 离线生成key set。顺序使用即可
- 详细设计
    - 数据分片： 大批量的数据存储（选定策略进行分片）
        - 根据actual url的第一位进行分片
        - 根据actual url的长度进行分片
    - cache：缓存
        - 缓存：20%的热点数据
        - 缓存更新策略：LRU（redis）
        - 缓存副本更新：
            - 懒加载（先cache-》redis-》mysql-》redis-》
    - data analyze
        - 根据IP区分地域
        - 高频url（对应的网站
    - 安全和权限控制
        - 设计用户权限
- 负载均衡
    - client -> server
    - server -> cache
    - server -> mysql
- 过期数据清理
    - 定期数据压缩存档（过期数据）
    - 定时任务清理数据
    - 查询时候软删除过期记录