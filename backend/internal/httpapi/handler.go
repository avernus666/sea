package httpapi

import (
 "encoding/json"
 "errors"
 "net/http"
 "strings"
 "time"
 "sea/backend/internal/repository"
)

type Handler struct { repo repository.InstitutionRepository; origin string }
func NewHandler(repo repository.InstitutionRepository,origin string)*Handler{return &Handler{repo:repo,origin:origin}}
func(h *Handler)Routes()http.Handler{
 mux:=http.NewServeMux()
 mux.HandleFunc("GET /health",func(w http.ResponseWriter,r *http.Request){writeJSON(w,200,map[string]string{"status":"ok"})})
 mux.HandleFunc("GET /api/v1/institutions",h.list)
 mux.HandleFunc("GET /api/v1/institutions/{id}",h.get)
 return h.security(h.cors(mux))
}
func(h *Handler)list(w http.ResponseWriter,r *http.Request){
 q:=strings.TrimSpace(r.URL.Query().Get("query"));if len(q)>100{writeError(w,400,"query is too long");return}
 data,err:=h.repo.List(r.Context(),q);if err!=nil{writeError(w,500,"could not load institutions");return};writeJSON(w,200,data)
}
func(h *Handler)get(w http.ResponseWriter,r *http.Request){
 id:=strings.TrimSpace(r.PathValue("id"));if id==""||len(id)>100{writeError(w,400,"invalid institution id");return}
 data,err:=h.repo.GetByID(r.Context(),id);if errors.Is(err,repository.ErrNotFound){writeError(w,404,"institution not found");return};if err!=nil{writeError(w,500,"could not load institution");return};writeJSON(w,200,data)
}
func(h *Handler)cors(next http.Handler)http.Handler{return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
 origin:=r.Header.Get("Origin");if origin!=""&&origin==h.origin{w.Header().Set("Access-Control-Allow-Origin",origin);w.Header().Set("Vary","Origin");w.Header().Set("Access-Control-Allow-Methods","GET, OPTIONS");w.Header().Set("Access-Control-Allow-Headers","Content-Type")}
 if r.Method==http.MethodOptions{w.WriteHeader(http.StatusNoContent);return};next.ServeHTTP(w,r)
})}
func(h *Handler)security(next http.Handler)http.Handler{return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
 w.Header().Set("X-Content-Type-Options","nosniff");w.Header().Set("X-Frame-Options","DENY");w.Header().Set("Referrer-Policy","strict-origin-when-cross-origin");w.Header().Set("Cache-Control","no-store");next.ServeHTTP(w,r)
})}
func writeJSON(w http.ResponseWriter,status int,data any){w.Header().Set("Content-Type","application/json; charset=utf-8");w.WriteHeader(status);_=json.NewEncoder(w).Encode(data)}
func writeError(w http.ResponseWriter,status int,message string){writeJSON(w,status,map[string]string{"error":message,"timestamp":time.Now().UTC().Format(time.RFC3339)})}
