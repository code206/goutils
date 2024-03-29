# Go Util

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/code206/goutils?style=flat-square)

Go一些常用的工具函数收集、实现和整理

- `bytepool` []byte 对象池，回收并复用
- `compress` 获取指定算法压缩后的数据
- `duration` 秒级时长转换为友好的描述
- `fasthttprealip` 从 fasthttp ctx 中获取客户端的请求ip
- `fasthttpserver` 启动 fasthttp server 后，可以同时在 ipv4 和 ipv6 地址上监听端口
- `file` linux 系统中，在不同的文件格式下复制文件
- `goadmin` go-admin中用到的方法
- `hash` md5，sha1，sha256 散列算法
- `http` 从 uri 和 args 获取请求参数
- `inslice` 搜索字符串是否和 slice 中对象相匹配
- `math` 绝对值，最大值，最小值
- `now` 时间精度不敏感时，每秒更新当前时间。避免高频场景中time.Now()的压力
- `now_unix` 时间精度不敏感时，每秒更新当前时间。避免高频场景中time.Now()的压力
- `now_unix_nano` 时间精度不敏感时，每100毫秒更新当前时间。避免高频场景中time.Now()的压力
- `path` 命令行路径，命令名称，文件/文件夹是否存在
- `str` 字符字节互转，行分割，字符串截取
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

### CompressFunc

> Package `github.com/code206/goutils/compressfunc`
```go
// source at compressfunc/gzip.go
func GZipBytes(data []byte) ([]byte, error)
```

### Convert

> Package `github.com/code206/goutils/convert`
```go
// source at convert/b2i64.go
func BytesToInt64(buf []byte) int64
// source at convert/b2s.go
func B2S(b []byte) string
// source at convert/i642b.go
func Int64ToBytes(i int64) []byte
// source at convert/s2b.go
func S2B(s string) (b []byte)
// source at convert/strings2interfaces.go
func strings2Interfaces(strings []string) []interface{}
// source at convert/interface2string.go
func Interface2String(i interface{}) string
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
// source at file/delete.go 删除文件，可以多个字符拼接地址
func DeleteFile(elem ...string) error
```

### GoAdmin

> Package `github.com/code206/goutils/goadmin`
```go
// source at goadmin/create-data.go
func CreateData(values form.Values, fields ...string) (dialect.H, error)
// source at goadmin/params-validator.go
func ParamsValidator(values form.Values, pvr []ParamsValidateRule) error
// source at goadmin/move-upload-file.go
func MoveUploadFile(values form.Values, mfp *MoveFuncParam) (string, error)
func GeneratePaths(fileName string, mfp *MoveFuncParam) (string, string, error)
```

### HashFunc

> Package `github.com/code206/goutils/hashfunc`
```go
// source at hashfunc/md5.go
func S2Md5(s string) string
func B2Md5(b []byte) string
// source at hashfunc/sha1.go
func S2Sha1(s string) string
func B2Sha1(b []byte) string
// source at hashfunc/sha256.go
func S2Sha256(s string) string
func B2Sha256(b []byte) string
```
### HttpFunc

> Package `github.com/code206/goutils/httpfunc`
```go
// source at httpfunc/parse_uri_args.go
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

### MathFunc

> Package `github.com/code206/goutils/mathfunc`
```go
// source at mathfunc/abs.go
func Abs(n int64) int64
// source at mathfunc/max.go
func Max(x, y int64) int64
// source at mathfunc/min.go
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

### NowUnix

> Package `github.com/code206/goutils/nowunix`
```go
// source at nowunix/now.go
func Now() int64
func NowTime() time.Time
```

### NowUnixNano

> Package `github.com/code206/goutils/nowunixnano`
```go
// source at nowunixnano/now.go
func Now() int64
func NowTime() time.Time
```

### PathFunc

> Package `github.com/code206/goutils/pathfunc`
```go
// source at pathfunc/cmd_path.go
func BinDirPath() (string, error)
func BinName() (string, error)
func BinDirName() (string, error)
// source at pathfunc/dir_exist.go
func IsDir(path string) bool
func HasSubDir(path string) (bool, error)
// source at pathfunc/file_exist.go
func IsFile(path string) bool
func FileExist(path string) bool
func FileNotExist(path string) bool
// source at pathfunc/path_exist.go
func PathExist(path string) bool
func PathNotExist(path string) bool
func PathLinkExist(path string) bool
func PathLinkNotExist(path string) bool
// source at pathfunc/path_ext.go
func TrimPathExt(path string) string
func CutPathExt(path string) (string, string)
```

### PC

> Package `github.com/code206/goutils/pc`
```go
// source at pc/pc.go
func NewPool(channelLength int, consumerNumber int, produceFunc func(chan<- interface{}), consumeFunc func(interface{})) *Pool
// source at pc/pc.go
func (p *Pool) Run()
```

### Str

> Package `github.com/code206/goutils/str`
```go
// source at str/split_lines.go
func SplitLines(str string) []string
// source at str/sub_string.go
func SubString(s string, start, end int) string
```

### Xormdb

> Package `github.com/code206/goutils/xormdb`
```go
// source at xormdb/init_db.go
func InitDB(conf *Config) (*xorm.Engine, error)
```
