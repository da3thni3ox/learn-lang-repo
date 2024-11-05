package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserOfIdOK(t *testing.T) {
	req := httptest.NewRequest("GET", "/users?userid=0", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(userGetHandler)
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("Статус ответа сервера %d != 200", status)
	}
}

func TestGetUserMoreQuantity(t *testing.T) {
	req := httptest.NewRequest("GET", "/users?userid=10", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(userGetHandler)
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Errorf("Ответ сервера на запрос пользователя с идентификатором выходящем за пределы списка завершился с http status code =%d отличным от 400", status)
	}
}
