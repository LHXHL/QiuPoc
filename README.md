# QiuPoc

**整理的一些poc，实现自动化检测利用，支持单个目标和批量检测，高亮输出。**

```go
git clone https://github.com/LHXHL/QiuPoc.git
go build -o main main.go
./main -h 使用说明
```

### Pocs(慢慢更新)

**poc均来自互联网**

```go
./main -show
```

![](https://p.ipic.vip/x7rwjc.png)

### 单个目标检测

```go
./main -u http://127.0.0.1 -exp n序号
```

![image-20230822151535962](https://p.ipic.vip/cm447o.png)

### 批量目标检测

```go
./main -f urls.txt -exp n序号
```

**url需要保持http||https://的格式，这部分还没来得及改善**

![image-20230822151818257](https://p.ipic.vip/o2kcnm.png)

