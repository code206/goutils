# Go Util

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/code206/goutils?style=flat-square)

Go一些常用的工具函数收集、实现和整理

- `bytepool` []byte 对象池，回收并复用
- `compress` 获取指定算法压缩后的数据
- `duration` 秒级时长转换为友好的描述
- `fasthttprealip` 从 fasthttp ctx 中获取客户端的请求ip
- `fasthttpserver` 启动 fasthttp server 后，可以同时在 ipv4 和 ipv6 地址上监听端口
- `file` linux 系统中，在不同的文件格式下复制文件
- `hash` md5，sha1，sha256 散列算法
- `http` 从 uri 和 args 获取请求参数
- `math` 绝对值，最大值，最小值
- `now` 时间精度不敏感时，每秒更新当前时间。避免高频场景中time.Now()的压力
- `path` 命令行路径，命令名称，文件/文件夹是否存在
- `slice` 搜索字符串是否和 slice 中对象相匹配
- `string` 字符字节互转，行分割，字符串截取
- `xormdb` xorm 引擎初始化

## Packages

### BytePool

> Package `github.com/code206/goutils/bytepool`
```go
// source at bytepool/bytepool.go
func NewBytePoolCap(maxSize int, width int, capwidth int) (bp *BytePoolCap)
func (bp *BytePoolCap) Get() (b []byte)
func (bp *BytePoolCap) Put(b []byte)
func (bp *BytePoolCap) Width() (n int)
func (bp *BytePoolCap) Len() (n int)
```

### Compress

> Package `github.com/code206/goutils/compress`
```go
// source at compress/gzip.go
func GZipBytes(data []byte) ([]byte, error)
```

### Duration

> Package `github.com/code206/goutils/duration`
```go
// source at duration/second2string.go
func ToSecond(d time.Duration) int64
func SecondToString(t int64) string
```

### FasthttpRealip

> Package `github.com/code206/goutils/fasthttprealip`
```go
// source at fasthttprealip/main.go
func RealIP(ctx *fasthttp.RequestCtx) string
```

### FasthttpServer

> Package `github.com/code206/goutils/fasthttpserver`
```go
// source at fasthttpserver/server.go
func Start(conf *fasthttpserver.Config) error
```

### File

> Package `github.com/code206/goutils/file`
```go
// source at file/copy.go 可以在不同的文件格式下复制文件
func Copy(src, dst string) (int64, error)
```

### Hash

> Package `github.com/code206/goutils/hash`
```go
// source at hash/md5.go
func MD5(s string) string
// source at hash/sha1.go
func Sha1(s string) string
// source at hash/sha256.go
func Sha256(s string) string
```
### Http

> Package `github.com/code206/goutils/http`
```go
// source at http/parse_uri_args.go
func ParseUriArgs(uri string) map[string]string
```

### InSlice

> Package `github.com/code206/goutils/inslice`
```go
// source at slice/in_slice.go
func InSlice(need string, needSlice []string) bool
func InSliceEqualFold(need string, needSlice []string) bool
func InSliceHasPrefix(s string, prefixs []string) bool
func InSliceHasSuffix(s string, suffixs []string) bool
```

### Math

> Package `github.com/code206/goutils/math`
```go
// source at math/abs.go
func Abs(n int64) int64
// source at math/max.go
func Max(x, y int64) int64
// source at math/min.go
func Min(x, y int64) int64
```

### Now

> Package `github.com/code206/goutils/now`
```go
// source at now/now.go
func Now() int64
func NowTime() time.Time
func NowUnixNanoInit()
func NowUnixInit()
```

### Path

> Package `github.com/code206/goutils/path`
```go
// source at path/cmd_path.go
func BinDirPath() (string, error)
func BinName() (string, error)
func BinDirName() (string, error)
// source at path/dir_exist.go
func IsDir(path string) bool
func HasSubDir(path string) (bool, error)
// source at path/file_exist.go
func IsFile(path string) bool
func FileExist(path string) bool
func FileNotExist(path string) bool
// source at path/path_exist.go
func PathExist(path string) bool
func PathNotExist(path string) bool
func PathLinkExist(path string) bool
func PathLinkNotExist(path string) bool
```

### Str

> Package `github.com/code206/goutils/str`
```go
// source at str/b2s.go
func B2S(b []byte) string
// source at str/s2b.go
func S2B(s string) (b []byte)
// source at str/split_lines.go
func SplitLines(str string) []string
// source at str/sub_string.go
func SubString(s string, start, end int) string
// source at str/strings2interfaces.go
func stringsToInterfaces(strings []string) []interface{}
```

### Xormdb

> Package `github.com/code206/goutils/xormdb`
```go
// source at xormdb/init_db.go
func InitDB(conf *Config) (*xorm.Engine, error)
```
