package xormdb

import (
	"errors"
	"runtime"
	"time"

	"xorm.io/xorm"
	"xorm.io/xorm/caches"
	"xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

var (
	defaultMaxOpenConns   = 6
	defaultCacheQuerySize = 256
	defaultNameMapper     = names.SnakeMapper{}
)

type Config struct {
	DriverName     string // 必填
	DataSourceName string // 必填

	MaxOpenConns    int           // 设置最大打开连接数
	MaxIdleConns    int           // 设置数据库最大闲置连接数
	ConnMaxLifetime time.Duration // 设置数据库可以重用连接的最长时间

	NameMapper      names.Mapper // 指定使用的命名映射策略
	CacheMapper     bool         //是否启用名称映射缓存
	TableNamePrefix string       // 表名称前缀
	TableNameSuffix string       // 表名称后缀

	CacheQuery        bool // 是否启用查询缓存
	CacheQueryMaxSize int  // 最少 defaultCacheQuerySize 条缓存

	Sync       bool          // 是否启用同步
	SyncModels []interface{} // 那些数据结构要同步到表结构

	ShowSQL  bool         // 是否打印 sql 语句
	LogLevel log.LogLevel // LOG_DEBUG LOG_INFO LOG_WARNING LOG_ERR LOG_OFF LOG_UNKNOWN
}

// 初始化 xorm 引擎
// sqlite3 sync2 时，xorm tag 中不能使用 comment('')
// 调用 InitSB 时，需在调用源码文件中 import 对应的数据库，例如：import _ "github.com/mattn/go-sqlite3"
func InitDB(conf *Config) (*xorm.Engine, error) {
	if conf.DriverName == "" || conf.DataSourceName == "" {
		return nil, errors.New("xorm config error")
	}

	// 建立数据库 构建连接字符串："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	//db, err := sql.Open("mysql", "user:password@unix(/tmp/mysql.sock)/test?charset=utf8")
	//db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/test?charset=utf8")   //指定IP和端口
	//db, err := sql.Open("mysql", "user:password@/test?charset=utf8")  //默认方式
	db, err := xorm.NewEngine(conf.DriverName, conf.DataSourceName)
	if err != nil {
		return nil, err
	}

	dbConnSet(db, conf)

	tableNameMapperSet(db, conf)

	dbCacheQuerySet(db, conf)

	if err = dbSyncSet(db, conf); err != nil {
		return nil, err
	}

	// 打印 sql 语句
	db.ShowSQL(conf.ShowSQL)
	db.Logger().SetLevel(conf.LogLevel)

	// 检查数据库连接，并返回
	if err = db.Ping(); err != nil {
		return nil, err
	} else {
		return db, nil
	}
}

func dbConnSet(db *xorm.Engine, conf *Config) {
	cores := runtime.NumCPU()
	n := cores * 2
	if defaultMaxOpenConns < n {
		defaultMaxOpenConns = n
	}

	if conf.MaxOpenConns == 0 {
		conf.MaxOpenConns = defaultMaxOpenConns
	}
	if conf.MaxIdleConns == 0 || conf.MaxIdleConns > conf.MaxOpenConns {
		conf.MaxIdleConns = defaultMaxOpenConns
	}
	if conf.ConnMaxLifetime == 0 {
		conf.ConnMaxLifetime = time.Duration(defaultMaxOpenConns*12) * time.Second
	}

	// === 建立数据库后，对数据库进行设置 ===
	//设置同时打开的连接数(使用中+空闲)的最大限制，将此值设置为小于或等于0表示没有限制
	// 建议设置为 cpu核心数 * 2
	// 默认配置为 5
	db.SetMaxOpenConns(conf.MaxOpenConns)
	// 设置数据库最大闲置连接数
	// 不能大于 SetMaxOpenConns
	// 建议设置等于 SetMaxOpenConns，配合 SetConnMaxLifetime，表示连接池保持一定数量连接，只是定期重启连接
	// 默认配置为 2
	db.SetMaxIdleConns(conf.MaxIdleConns)
	// 设置可重用链接得最大时间长度（为了防止连接不可用，需要定期启用新的连接）
	// 这不是空闲超时。连接将在第一次创建后开始计时，超过指定时长后将无法重用
	// 每秒自动执行一次清除操作，从连接池中删除 “过期” 的连接。
	// 默认配置为 0，表示连接一直保持，不主动关闭
	db.SetConnMaxLifetime(time.Minute)
	return
}

func tableNameMapperSet(db *xorm.Engine, conf *Config) {

	if conf.NameMapper == nil {
		conf.NameMapper = defaultNameMapper
	}
	// 设置是否在内存中缓存曾经映射过的命名映射
	if conf.CacheMapper {
		conf.NameMapper = names.NewCacheMapper(conf.NameMapper)
	}

	var tbMapper names.Mapper = conf.NameMapper
	// 设置数据表名称前缀
	if conf.TableNamePrefix != "" {
		tbMapper = names.NewPrefixMapper(tbMapper, conf.TableNamePrefix)
	}
	// 设置数据表名称后缀
	if conf.TableNameSuffix != "" {
		tbMapper = names.NewSuffixMapper(tbMapper, conf.TableNameSuffix)
	}

	db.SetTableMapper(tbMapper)
	db.SetColumnMapper(conf.NameMapper)
	return
}

func dbCacheQuerySet(db *xorm.Engine, conf *Config) {
	if conf.CacheQuery == false {
		return
	}

	if conf.CacheQueryMaxSize < defaultCacheQuerySize {
		conf.CacheQueryMaxSize = defaultCacheQuerySize
	}
	cacher := caches.NewLRUCacher(caches.NewMemoryStore(), conf.CacheQueryMaxSize)
	db.SetDefaultCacher(cacher)
	return
}

func dbSyncSet(db *xorm.Engine, conf *Config) error {
	if conf.Sync && len(conf.SyncModels) > 0 {
		return db.Sync2(conf.SyncModels...)
	} else {
		return nil
	}
}
