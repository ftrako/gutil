package gutils

import (
	"bytes"
	"net"
	"net/http"
	"strconv"
	"strings"
)

// 如果使用了NGINX代理，需要在nginx中设置IP字段
// proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
func GetRealIP(r *http.Request) string {
	ip := strings.Replace(r.Header.Get("X-Forwarded-For"), " ", "", -1)
	if ip != "" {
		if strings.Contains(ip, `,`) {
			return strings.Split(ip, `,`)[0]
		}
		return ip
	}

	// 获取客户端IP地址
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}
	return ip
}

func IPStr2Int(ipstr string) int {
	ipSegs := strings.Split(ipstr, ".")
	var ipInt = 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}

func IPInt2Str(ipInt int) string {
	ipSegs := make([]string, 4)
	var len = len(ipSegs)
	buffer := bytes.NewBufferString("")
	for i := 0; i < len; i++ {
		tempInt := ipInt & 0xFF
		ipSegs[len-i-1] = strconv.Itoa(tempInt)
		ipInt = ipInt >> 8
	}
	for i := 0; i < len; i++ {
		buffer.WriteString(ipSegs[i])
		if i < len-1 {
			buffer.WriteString(".")
		}
	}
	return buffer.String()
}

// 判断是否是IP地址
func IsIPAddress(ip string) bool {
	ip = strings.TrimSpace(ip)
	address := net.ParseIP(ip)
	return address != nil
}

// 判断是否是内网IP
func IsLanIP(ip string) bool {
	isInnerIp := false
	ipNum := getIpNum(ip)
	/**
	  私有IP：A类  10.0.0.0-10.255.255.255
	         B类  172.16.0.0-172.31.255.255
	         C类  192.168.0.0-192.168.255.255
	  当然，还有127这个网段是环回地址
	  **/
	aBegin := getIpNum("10.0.0.0")
	aEnd := getIpNum("10.255.255.255")
	bBegin := getIpNum("172.16.0.0")
	bEnd := getIpNum("172.31.255.255")
	cBegin := getIpNum("192.168.0.0")
	cEnd := getIpNum("192.168.255.255")
	isInnerIp = isLanNum(ipNum, aBegin, aEnd) || isLanNum(ipNum, bBegin, bEnd) || isLanNum(ipNum, cBegin, cEnd) || strings.Contains(ip, "127.0.0.1")
	return isInnerIp
}

func getIpNum(ip string) int64 {
	ips := strings.Split(ip, ".")
	a, _ := strconv.Atoi(ips[0])
	b, _ := strconv.Atoi(ips[1])
	c, _ := strconv.Atoi(ips[2])
	d, _ := strconv.Atoi(ips[3])

	ipNum := int64(a*256*256*256 + b*256*256 + c*256 + d)
	return ipNum
}

func isLanNum(userIp, begin, end int64) bool {
	return (userIp >= begin) && (userIp <= end)
}
