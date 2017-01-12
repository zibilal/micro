package middleware

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"net"
	"net/http"
	"time"
)

func Logger() echo.MiddlewareFunc {
	name := "manhattan"
	l := logrus.StandardLogger()
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			req := c.Request()
			res := c.Response()

			remoteAddr := req.RemoteAddress()
			if ip := req.Header().Get(echo.HeaderXRealIP); ip != "" {
				remoteAddr = ip
			} else if ip = req.Header().Get(echo.HeaderXForwardedFor); ip != "" {
				remoteAddr = ip
			} else {
				remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
			}

			entry := l.WithFields(logrus.Fields{
				"request": req.URI(),
				"method":  req.Method,
				"remote":  remoteAddr,
			})

			if reqID := req.Header().Get("X-Request-Id"); reqID != "" {
				entry = entry.WithField("request_id", reqID)
			}

			entry.Info("started handling request")
			if err := next(c); err != nil {
				c.Error(err)
			}

			latency := time.Since(start)
			entry.WithFields(logrus.Fields{
				"size":        res.Size(),
				"status":      res.Status(),
				"text_status": http.StatusText(res.Status()),
				"took":        latency,
				fmt.Sprintf("measure#%s.latency", name): latency.Nanoseconds(),
			}).Info("completed handling request")

			return nil
		}
	}
}
