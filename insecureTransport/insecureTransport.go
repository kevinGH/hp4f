package insecureTransport

import "net/http"

func HttpListenAndServer(port string, handler http.Handler) error {
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		return err
	}

	return nil
}
