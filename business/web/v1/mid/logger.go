package mid

import (
	"context"
	"net/http"

	"github.com/PhyoYazar/service/foundation/web"
	"go.uber.org/zap"
)

func Logger(log *zap.SugaredLogger) web.Middleware {
	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			// v := web.GetValues(ctx)

			// path := r.URL.Path
			// if r.URL.RawQuery != "" {
			// 	path = fmt.Sprintf("%s?%s", path, r.URL.RawQuery)
			// }

			// log.Infow("request started", "trace_id", v.TraceID, "method", r.Method, "path", path,
			// 	"remoteaddr", r.RemoteAddr)

			log.Infow("request started",  "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)

			err := handler(ctx, w, r)

			log.Infow("request completed","method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr,)

			// log.Infow("request completed", "trace_id", v.TraceID, "method", r.Method, "path", path,
			// 	"remoteaddr", r.RemoteAddr, "statuscode", v.StatusCode, "since", time.Since(v.Now))

			return err
		}

		return h
	}

	return m
}