# MYSQL
- [What is mysql](https://www.mysql.com)

## [Index](https://dev.mysql.com/doc/refman/8.0/en/create-index.html)
- Index Type
   - Primary key(聚集索引)
      - the value is unique and not null
      - all the index is based the primary key
   - Unique(唯一索引)
      - like Key but not allowed repeated value
   - Key
      - for all type
      - allowed repeated and empty value
   - Full-Text
      - just for test|char|varchar type
   - SpaTial(空间索引)
- Index Struct
  - B+树 

- Principle Of Build Index
  - 最左前缀匹配原则：mysql会一直向右匹配直到遇到范围查询(>、<、between、like)就停止匹配
  - =和in可以乱序，比如a = 1 and b = 2 and c = 3 建立(a,b,c)索引可以任意顺序，mysql的查询优化器会帮你优化成索引可以识别的形式
  - 尽量选择区分度高的列作为索引，区分度的公式是count(distinct col)/count(*)，表示字段不重复的比例
  - 索引列不能参与计算，保持列“干净”
  - 尽量的扩展索引，不要新建索引


- [慢查询优化](https://tech.meituan.com/2014/06/30/mysql-index.html)
  - 步骤链接中
## Database Transaction

## Engine(InnoDB)

## Best practice