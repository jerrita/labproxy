# LabProxy

在公网可达的 vps 上转发你 homolab 的 endpoint，这种方式应该可以解决动态 ip 的问题。

境外 wg 组网真是太不稳🌶！

## 使用方法

### 配置文件

创建 `config.json` 文件：

```json
{
  "mappings": [
    {
      "src": 80,
      "dst": 80,
      "endpoint": "google.com"
    },
    {
      "src": 443,
      "dst": 443,
      "endpoint": "google.com"
    }
  ]
}
```

### 本地运行

```bash
./labproxy [config.json]
```

如果不指定配置文件，默认使用 `/etc/labproxy.json`。


## 配置说明

- `src`: 本地监听端口
- `dst`: 目标端口
- `endpoint`: 目标服务器地址

> 写的很简单，你就说能不能用吧