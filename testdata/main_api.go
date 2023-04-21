// Code generated by happy. DO NOT EDIT.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (h *Service) HandlerOptions(r *http.Request) map[string]string {
	parts := []string{}
	for i, p := range strings.Split(r.URL.Path, "/") {
		if i == 0 || p != "" {
			parts = append(parts, p)
		}
	}
	var params []string
	_ = params
	switch parts[0] {
	case "":
		if len(parts) == 1 {
			return nil
		}
		switch parts[1] {
		case "users":
			if len(parts) == 2 {
				switch r.Method { // Leaf
				case "POST":
					return map[string]string{"authenticated": ""}
				}
				return nil
			}
			return nil
		}
	}
	return nil
}

func (h *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	var res any
	parts := []string{}
	for i, p := range strings.Split(r.URL.Path, "/") {
		if i == 0 || p != "" {
			parts = append(parts, p)
		}
	}
	var params []string
	_ = params
	switch parts[0] {
	case "":
		if len(parts) == 1 {
			return
		}
		switch parts[1] {
		case "users":
			if len(parts) == 2 {
				switch r.Method { // Leaf
				case "POST":
					var param1 User
					if err := json.NewDecoder(r.Body).Decode(&param1); err != nil {
						http.Error(w, fmt.Sprintf("Failed to decode request body: %s", err), http.StatusBadRequest)
						return
					}
					err = h.CreateUser(r, param1)
					goto matched
				case "GET":
					var param0 Paginate
					if err := decodePaginate(r.URL.Query(), &param0); err != nil {
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}
					res, err = h.ListUsers(param0)
					goto matched
				}
				return
			}
			switch parts[2] {
			default: // Parameter :id
				params = append(params, parts[2])
				if len(parts) == 3 {
					switch r.Method { // Leaf
					case "GET":
						var param0 ID
						if err := param0.UnmarshalText([]byte(params[0])); err != nil {
							http.Error(w, "id: "+err.Error(), http.StatusBadRequest)
							return
						}
						res, err = h.GetUser(param0)
						goto matched
					}
					return
				}
				switch parts[3] {
				case "avatar":
					if len(parts) == 4 {
						switch r.Method { // Leaf
						case "GET":
							var param0 ID
							if err := param0.UnmarshalText([]byte(params[0])); err != nil {
								http.Error(w, "id: "+err.Error(), http.StatusBadRequest)
								return
							}
							res, err = h.GetAvatar(param0)
							goto matched
						}
						return
					}
					return
				}
			}
		case "shutdown":
			if len(parts) == 2 {
				switch r.Method { // Leaf
				case "POST":
					h.Shutdown(w)
					return
					goto matched
				}
				return
			}
			return
		}
	}
	// No match but we don't return a 404 here, to allow the default handler to take control.
	return
matched:

	// Handle errors
	if err != nil {
		if herr, ok := err.(http.Handler); ok {
			herr.ServeHTTP(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Handle response
	switch res := res.(type) {
	case nil:
		w.WriteHeader(http.StatusNoContent)
	case string:
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, res)
	case []byte:
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(res)
	case *http.Response:
		headers := w.Header()
		for k, v := range res.Header {
			headers[k] = v
		}
		w.WriteHeader(res.StatusCode)
		_, _ = io.Copy(w, res.Body)
	case io.ReadCloser:
		w.Header().Set("Content-Type", "application/octet-stream")
		_, _ = io.Copy(w, res)
		res.Close()
	case io.Reader:
		w.Header().Set("Content-Type", "application/octet-stream")
		_, _ = io.Copy(w, res)
	default:
		data, err := json.Marshal(res)
		if err != nil {
			http.Error(w, `failed to encode response: `+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", strconv.Itoa(len(data)))
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
func decodePaginate(p url.Values, out *Paginate) (err error) {
	if q, ok := p["p"]; ok {
		if out.Page, err = strconv.Atoi(q[len(q)-1]); err != nil {
			return fmt.Errorf("failed to decode query parameter \"p\" as int: %w", err)
		}
	}
	if q, ok := p["s"]; ok {
		if out.Size, err = strconv.Atoi(q[len(q)-1]); err != nil {
			return fmt.Errorf("failed to decode query parameter \"s\" as int: %w", err)
		}
	}
	if q, ok := p["sparse"]; ok {
		out.Sparse = new(bool)
		if *out.Sparse, err = strconv.ParseBool(q[len(q)-1]); err != nil {
			return fmt.Errorf("failed to decode query parameter \"sparse\" as bool: %w", err)
		}
	}
	return nil
}
