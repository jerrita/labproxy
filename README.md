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

如果不指定配置文件，默认使用 `config.json`。

### Docker 运行

```bash
# 使用预构建的镜像
docker run -d \
  --name labproxy \
  -p 80:80 \
  -p 443:443 \
  -v $(pwd)/config.json:/root/config.json \
  ghcr.io/yourusername/labproxy:latest

# 或者构建本地镜像
docker build -t labproxy .
docker run -d \
  --name labproxy \
  -p 80:80 \
  -p 443:443 \
  -v $(pwd)/config.json:/root/config.json \
  labproxy
```

### Docker Compose

```yaml
version: '3.8'
services:
  labproxy:
    image: ghcr.io/yourusername/labproxy:latest
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./config.json:/root/config.json
    restart: unless-stopped
```

## 自动构建

项目配置了 GitHub Actions，会自动构建 Docker 镜像并推送到 `ghcr.io`：

- 推送到 `main` 或 `master` 分支时构建 `latest` 标签
- 推送标签时构建对应版本号的镜像
- 支持 `linux/amd64` 和 `linux/arm64` 多架构

## 配置说明

- `src`: 本地监听端口
- `dst`: 目标端口
- `endpoint`: 目标服务器地址

> 写的很简单，你就说能不能用吧