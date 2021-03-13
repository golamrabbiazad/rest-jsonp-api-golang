package main

import (
	"context"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

const URL string = "https://jsonplaceholder.typicode.com"

type postsResource struct{}

func (rs postsResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.List)
	r.Post("/", rs.Create)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(PostCtx)
		r.Get("/", rs.Get)
		r.Put("/", rs.Update)
		r.Delete("/", rs.Delete)
	})

	return r
}

// Request Handler - GET /posts - Read a list of posts.
func (rs postsResource) List(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(URL + "/posts")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Request Handler - POST /posts - Create a new post.
func (rs postsResource) Create(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post(URL+"/posts", "application/json", r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PostCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (rs postsResource) Get(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)

	resp, err := http.Get(URL + "/posts/" + id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rs postsResource) Update(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
	client := &http.Client{}

	req, err := http.NewRequest("PUT", URL+"/posts/"+id, r.Body)
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rs postsResource) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", URL+"/posts/"+id, nil)
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
