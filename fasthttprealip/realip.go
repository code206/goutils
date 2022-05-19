package fasthttprealip

import (
	"errors"
	"net"
	"strings"

	"github.com/valyala/fasthttp"
)

// realip ip内网地址端变量
var cidrs []*net.IPNet

// realip 转换ip内网地址段为 net.IPNet 类型
func init() {
	maxCidrBlocks := []string{
		"10.0.0.0/8",     // 24-bit block
		"100.64.0.0/10",  // 22-bit block
		"127.0.0.1/8",    // localhost
		"169.254.0.0/16", // link local address
		"172.16.0.0/12",  // 20-bit block
		"192.0.0.0/24 ",  // 8-bit block
		"192.168.0.0/16", // 16-bit block
		"198.18.0.0/15",  // 17-bit block
		"::1/128",        // localhost IPv6
		"fc00::/7",       // unique local address IPv6
		"fe80::/10",      // link local address IPv6
	}

	cidrs = make([]*net.IPNet, len(maxCidrBlocks))
	for i, maxCidrBlock := range maxCidrBlocks {
		_, cidr, _ := net.ParseCIDR(maxCidrBlock)
		cidrs[i] = cidr
	}
}

// isLocalAddress works by checking if the address is under private CIDR blocks.
// List of private CIDR blocks can be seen on :
//
// https://en.wikipedia.org/wiki/Private_network
//
// https://en.wikipedia.org/wiki/Link-local_address
// realip 检查提交的地址是否是内网地址
func isPrivateAddress(address string) (bool, error) {
	ipAddress := net.ParseIP(address)
	if ipAddress == nil {
		return false, errors.New("address is not valid")
	}

	for i := range cidrs {
		if cidrs[i].Contains(ipAddress) {
			return true, nil
		}
	}

	return false, nil
}

// FromRequest return client's real public IP address from http request headers.
// 获取真实的客户端ip
func RealIP(ctx *fasthttp.RequestCtx) string {
	// 有 cloudflare 的真实ip请求头
	if ip := string(ctx.Request.Header.Peek("CF-Connecting-IP")); ip != "" {
		isPrivate, err := isPrivateAddress(ip)
		if !isPrivate && err == nil {
			return ip
		}
	}

	// 有 X-Forwarded-For 请求头
	if ips := string(ctx.Request.Header.Peek("X-Forwarded-For")); len(ips) > 0 {
		// Check list of IP in X-Forwarded-For and return the first global address
		for _, ip := range strings.Split(ips, ",") {
			ip = strings.TrimSpace(ip)
			isPrivate, err := isPrivateAddress(ip)
			if !isPrivate && err == nil {
				return ip
			}
		}
	}

	return ctx.RemoteIP().String()
}
