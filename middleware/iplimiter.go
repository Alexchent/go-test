package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/netip"
)

// IpLimiter ip限制的中间件，ip的格式为"192.168.1.0/24"的形式，
// 若传入参数为空，则自动添加"10.0.0.0/8"，否则按传入的ip段处理。
// 127.0.0.1会默认添加
func IpLimiter(allowedNetIps []string) gin.HandlerFunc {
	var prefixes []netip.Prefix
	if len(allowedNetIps) == 0 {
		allowedNetIps = []string{"10.0.0.0/8"}
	}
	for _, netIp := range allowedNetIps {
		prefix, err := netip.ParsePrefix(netIp)
		if err != nil {
			continue
		}
		prefixes = append(prefixes, prefix)
	}
	prefixes = append(prefixes, netip.PrefixFrom(netip.MustParseAddr("127.0.0.1"), 32))

	return func(c *gin.Context) {
		clientIp := c.ClientIP()
		ip, err := netip.ParseAddr(clientIp)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		for _, prefix := range prefixes {
			if prefix.Contains(ip) {
				c.Next()
				return
			}
		}
		c.AbortWithStatus(http.StatusForbidden)
	}
}
