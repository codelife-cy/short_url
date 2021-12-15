#### 短域名服务

环境：

1. golang语言开发
2. MySql数据库

配置文件：

```toml
# mysql 配置
[mysql]

  [mysql.base]
    # 连接最大时间
    conn_max_life_time = 60
    # 最大空闲连接数
    max_idle_conn = 60
    # 最大连接数
    max_open_conn = 10
  # mysql 读配置
  [mysql.read]
    addr = "192.168.74.128:3306"
    name = "short_url"
    pass = "root"
    user = "root"
  # mysql写配置
  [mysql.write]
    addr = "192.168.74.128:3306"
    name = "short_url"
    pass = "root"
    user = "root"
# redis 配置
[redis]
  addr = "127.0.0.1:6379"
  db = "0"
  # 最大重试次数
  max_retries = 3
  # 最小空闲连接
  min_idle_conn = 5
  pass = ""
  # 连接池大小
  pool_size = 10

# 服务配置
[server]
    port = ":3030"

[common]
# short urls that will be filtered to use
black_short_urls = ["version","health","short","expand","css","js","fuck","stupid"]

# Base string used to generate short url
base_string = "Ds3K9ZNvWmHcakr1oPnxh4qpMEzAye8wX5IdJ2LFujUgtC07lOTb6GYBQViSfR"

# Short url service domain name. This is used to filter short url loop.
domain_name = "127.0.0.1:3030"

schema = "http"

# 配置域名开始长度（1-62）
domain_length = 5
```

1. 修改数据库配置

```toml
 # mysql 读配置
  [mysql.read]
    addr = "192.168.74.128:3306"
    name = "short_url"
    pass = "root"
    user = "root"
# mysql写配置
  [mysql.write]
    addr = "192.168.74.128:3306"
    name = "short_url"
    pass = "root"
    user = "root"
```

2. 修改服务地址

```toml
domain_name = "127.0.0.1:3030"
#此处为生成的短域名地址部分
```

3. 可动态控制端域名地址长度

```toml
domain_length = 5
# 例如配置为5 ，shortURL:http://127.0.0.1:3030/sDDDH
```

接口示例：

1. 获取端域名

![获取短域名](.\image\获取短域名.png)

2. 使用短域名

![使用短域名URL](.\image\使用短域名URL.png)

![使用短域名URL结果](.\image\使用短域名URL结果.png)

项目结构：

![项目结构](.\image\项目结构.png)
