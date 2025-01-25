# LabProxy

在公网可达的 vps 上转发你 homolab 的 endpoint，这种方式应该可以解决动态 ip 的问题。

境外 wg 组网真是太不稳🌶！

```bash
./labproxy -endpoint your.homolab.org:homoport
```

Manager API 绝赞编写中，如果 endpoint 为 ip 形式应该可以通过这种形式更新

> 写的很简单，你就说能不能用吧