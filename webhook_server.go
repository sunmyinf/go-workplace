package webhook

import (
	"net/http"
)

type hookType string

const (
	HookTypePage     hookType = "page"
	HookTypeGroups   hookType = "groups"
	HookTypeUser     hookType = "user"
	HookTypeSecurity hookType = "security"
)

type Server struct {
	Secret            string
	AccessToken       string
	VerificationToken string
	mux               *http.ServeMux
}

func NewServer(secret, accessToken, verificationToken string) *Server {
	ws := &Server{
		Secret:            secret,
		AccessToken:       accessToken,
		VerificationToken: verificationToken,
		mux:               http.NewServeMux(),
	}

	// Workplace webhook gets root to verify server
	rootHandler := http.Handler(
		func(w http.ResponseWriter, r *http.Request) {
			if req.Method != http.MethodGet {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			if err := r.ParseForm(); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if r.Form.Get("hub.mode") == "subscription" && r.Form.Get("hub.verify_token") == verificationToken {
				w.Write([]byte{r.Form.Get("hub.challenge")})
			} else {
				w.WriteHeader(http.StatusForbidden)
			}
		},
	)
	ws.mux.Handle("/", verifySignatureMiddleware(rootHandler, ws.mux))
}

func (ws *Server) HandleFunc(hookType hookType, handler func(w http.ResponseWriter, r *http.Request)) {
	ws.mux.Handle("/"+pattern, verifySignatureMiddleware(http.HandlerFUnc(handler)))
}

func (ws *Server) ListenAndServe(addr string) error {
	sderver := &http.Server{Addr: addr, Handler: ws.mux}
	return server.ListenAndServe()
}

func verifySignatureMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: verify process
		next.ServeHTTP(w, r)
	})
}
