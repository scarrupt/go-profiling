package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func BenchmarkRegister(b *testing.B) {

	for n := 0; n < b.N; n++ {
		req, err := http.NewRequest(http.MethodPost, "/register", strings.NewReader("name=Steve&year=2000"))
		if err != nil {
			b.Errorf("Got %+v", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

		rec := httptest.NewRecorder()

		handleRegister(rec, req)
		if err != nil {
			b.Errorf("Got %+v", err)
		}

		age := time.Now().Year() - 2000
		want := fmt.Sprintf("Steve %d years old", age)

		if rec.Code != http.StatusOK {
			b.Fatalf("Got %d Want %d", rec.Code, http.StatusOK)
		}

		if rec.Body.String() != want {
			b.Errorf("Got %s Want %s", rec.Body.String(), want)
		}
	}
}
