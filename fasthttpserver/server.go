package fasthttpserver

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
)

const (
	defaultNetwork     = "tcp"
	defaultAddr        = ":8080"
	defaultReadTimeout = 8 * time.Second
)

type logger interface {
	Printf(format string, v ...interface{})
}

// 配置数据
type Config struct {
	Network          string
	Addr             string
	Reuseport        bool
	GracefulShutdown bool
	Server           *fasthttp.Server
	log              logger
}

// server 配置
type server struct {
	Network          string
	Addr             string
	Reuseport        bool
	GracefulShutdown bool
	fasthttpServer   *fasthttp.Server
	log              logger
}

// 启动
func Start(conf *Config) error {
	serv := new(server)

	if conf.Network == "tcp" || conf.Network == "tcp4" || conf.Network == "tcp6" {
		serv.Network = conf.Network
	} else if conf.Network == "" {
		serv.Network = defaultNetwork
	} else {
		panic("Invalid network: " + conf.Network)
	}

	if conf.Addr == "" {
		serv.Addr = defaultAddr
	} else {
		serv.Addr = conf.Addr
	}

	if conf.log == nil {
		serv.log = logger(log.New(os.Stderr, "", log.LstdFlags))
	} else {
		serv.log = conf.log
	}

	serv.fasthttpServer = conf.Server
	serv.fasthttpServer.Logger = serv.log
	return serv.listenAndServe()
}

// 监听并服务
func (serv *server) listenAndServe() error {
	ln, err := serv.getListener()
	if err != nil {
		return err
	}

	if serv.GracefulShutdown {
		return serv.ServeGracefully(ln)
	} else {
		return serv.Serve(ln)
	}
}

// 监听
func (serv *server) getListener() (net.Listener, error) {
	if serv.Reuseport {
		return reuseport.Listen(serv.Network, serv.Addr)
	} else {
		return net.Listen(serv.Network, serv.Addr)
	}

}

// 服务
func (serv *server) Serve(ln net.Listener) error {
	defer ln.Close()

	serv.Network = ln.Addr().Network()
	serv.Addr = ln.Addr().String()

	serv.log.Printf("Listening on: http://%s/", serv.Addr)

	return serv.fasthttpServer.Serve(ln)
}

// 可优雅关闭的服务
func (serv *server) ServeGracefully(ln net.Listener) error {
	if serv.fasthttpServer.ReadTimeout <= 0 {
		serv.fasthttpServer.ReadTimeout = defaultReadTimeout
	}

	listenErr := make(chan error, 1)

	go func() {
		listenErr <- serv.Serve(ln)
	}()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-listenErr:
		return err
	case <-osSignals:
		serv.log.Printf("Shutdown signal received")

		if err := serv.fasthttpServer.Shutdown(); err != nil {
			return err
		}

		serv.log.Printf("Server gracefully stopped")
	}

	return nil
}
