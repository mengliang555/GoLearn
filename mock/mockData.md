# Mock
- Grpc mock工具
  - [grip mock](https://github.com/tokopedia/gripmock)
    - 原理：
      - 生成相关的pb.go文件
      - 代码生成器，生成模版代码
      - 启动两个服务：
        - mock服务器
        - 加载相关mock数据的服务器