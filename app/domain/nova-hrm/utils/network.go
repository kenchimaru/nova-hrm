package utils

import (
	"net"
	"net/http"
)

func GetClientIP(r *http.Request) string {
	// Check if Cloudflare has set the CF-Connecting-IP header
	if ip := r.Header.Get("CF-Connecting-IP"); ip != "" {
		return ip
	}

	// Check if X-Forwarded-For header is set
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		// X-Forwarded-For can contain multiple IPs, take the first one
		ip := xff
		if commaIndex := len(xff); commaIndex != -1 {
			ip = xff[:commaIndex]
		}
		return ip
	}

	// Check if X-Real-IP header is set (used by some proxies)
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}

	// Fallback: Use the remote address from the connection
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr // return as is if parsing fails
	}
	return ip
}
