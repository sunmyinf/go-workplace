package webhook

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/sunmyinf/go-workplace/decode"
)

// Server is to launch server to serve workplace webhook callback request.
type Server struct {
	Secret            string
	AccessToken       string
	VerificationToken string
	objectHandlers    map[decode.Object]func(decode.RequestBody) error
	mux               *http.ServeMux
}

// NewServer return workplace webhook server instance.
// Handler has been registered to it as root '/' pattern by default.
func NewServer(secret, accessToken, verificationToken string) *Server {
	ws := &Server{
		Secret:            secret,
		AccessToken:       accessToken,
		VerificationToken: verificationToken,
		mux:               http.NewServeMux(),
	}

	rootHandlerFunc := func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if err := r.ParseForm(); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if r.FormValue("hub.mode") == "subscription" && r.FormValue("hub.verify_token") == verificationToken {
				w.Write([]byte(r.FormValue("hub.challenge")))
			} else {
				w.WriteHeader(http.StatusForbidden)
			}
		case http.MethodPost:
			rb, err := parsePostRequestBody(r)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			handler, exist := ws.objectHandlers[rb.Data[0].Object]
			if !exist {
				// if object handler not registered, return ok status.
				w.WriteHeader(http.StatusOK)
				return
			}
			if err := handler(rb); err != nil {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		default:
			w.WriteHeader(http.StatusForbidden)
		}
		return
	}

	// Workplace webhook gets root to verify server.
	// And posts root to callback.
	ws.mux.Handle("/", ws.verifySignatureMiddleware(http.HandlerFunc(rootHandlerFunc)))
	return ws
}

// HandleObjectFunc registers handler by object to Server instance.
// If handler of specified object was registered, override it by new one.
func (ws *Server) HandleObjectFunc(object decode.Object, objectHandler func(decode.RequestBody) error) {
	ws.objectHandlers[object] = objectHandler
}

// HandleFunc registers the handler function for the given pattern.
func (ws *Server) HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	ws.mux.HandleFunc(pattern, handler)
}

// ListenAndServe listens on the TCP network address srv.Addr and then
// calls Serve to handle requests on incoming connections.
func (ws *Server) ListenAndServe(addr string) error {
	server := &http.Server{Addr: addr, Handler: ws.mux}
	return server.ListenAndServe()
}

func (ws *Server) verifySignatureMiddleware(nextFunc http.HandlerFunc) http.Handler {
	next := http.Handler(nextFunc)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bufBody := bytes.Buffer{}
		if _, err := bufBody.ReadFrom(r.Body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := verifySignature(r.Header.Get("X-Hub-Signature"), ws.Secret, bufBody.Bytes()); err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func verifySignature(sig, secret string, key []byte) error {
	if sig == "" {
		return errors.New("error: signature is empty")
	}

	elements := strings.Split(sig, "=")
	if len(elements) < 2 {
		return errors.New("errors: invalid signature")
	}

	signatureHash := elements[1]
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(secret))
	expectedHash := hex.EncodeToString(mac.Sum(nil))

	if signatureHash != expectedHash {
		return errors.New("error: signature hash do not match expected hash")
	}
	return nil
}

func parsePostRequestBody(r *http.Request) (decode.RequestBody, error) {
	request := decode.RequestBody{}
	bufBody := bytes.Buffer{}
	if _, err := bufBody.ReadFrom(r.Body); err != nil {
		return request, err
	}

	err := json.Unmarshal(bufBody.Bytes(), &request)
	return request, err
}
