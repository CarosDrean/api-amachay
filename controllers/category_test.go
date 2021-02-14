package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/CarosDrean/api-amachay/storage/mocks"
	"github.com/CarosDrean/api-amachay/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCategory_Create_wrong_structure(t *testing.T) {
	data := []byte(`{"name": 123}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	e := echo.New()
	ctx := e.NewContext(r, w)
	h := NewCategory(&mocks.CategoryMockOK{})
	err := h.Create(ctx)
	if err != nil {
		t.Errorf("no se esperaba error, pero se obtuvo: %v", err)
	}
	if w.Code != http.StatusBadRequest {
		t.Errorf("Codigo de estado: se esperaba %d, se obtuvo %d", http.StatusBadRequest, w.Code)
	}

	resp := utils.Response{}
	err = json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		t.Errorf("No se pudo hacer unmarchall al body %v", err)
	}
	wantMessage := "La categoria no tiene la estructura correcta"
	if resp.Message != wantMessage {
		t.Errorf("La respuesta no es la esperada, se obtuvo %q, se esperaba %q", resp.Message, wantMessage)
	}
}

func TestCategory_Create_wrong_storage(t *testing.T) {
	data := []byte(`{"name": "Oficina"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	e := echo.New()
	ctx := e.NewContext(r, w)
	h := NewCategory(&mocks.CategoryMockError{})
	err := h.Create(ctx)
	if err != nil {
		t.Errorf("no se esperaba error, pero se obtuvo: %v", err)
	}

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Codigo de estado: se esperaba %d, se obtuvo %d", http.StatusInternalServerError, w.Code)
	}

	resp := utils.Response{}
	err = json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		t.Errorf("No se pudo hacer unmarchall al body %v", err)
	}
	wantMessage := "Hubo un error al crear la categoria"
	if resp.Message != wantMessage {
		t.Errorf("La respuesta no es la esperada, se obtuvo %q, se esperaba %q", resp.Message, wantMessage)
	}
}

func TestCategory_Create(t *testing.T) {
	data := []byte(`{"name": "Oficina"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	e := echo.New()
	ctx := e.NewContext(r, w)
	h := NewCategory(&mocks.CategoryMockOK{})
	err := h.Create(ctx)
	if err != nil {
		t.Errorf("no se esperaba error, pero se obtuvo: %v", err)
	}

	if w.Code != http.StatusCreated {
		t.Errorf("Codigo de estado: se esperaba %d, se obtuvo %d", http.StatusCreated, w.Code)
	}

	resp := utils.Response{}
	err = json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		t.Errorf("No se pudo hacer unmarchall al body %v", err)
	}
	wantMessage := "Categoria creada correctamente"
	if resp.Message != wantMessage {
		t.Errorf("La respuesta no es la esperada, se obtuvo %q, se esperaba %q", resp.Message, wantMessage)
	}
}
