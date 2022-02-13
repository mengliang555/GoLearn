# Golang-Thread-Local
- 实现方法
  - 核心：根据gID，构造相关sync.map[int64]interface{}
  - 