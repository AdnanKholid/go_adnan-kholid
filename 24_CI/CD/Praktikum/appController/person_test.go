package appController

import (
	"bytes"
	"echo-api"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestPersonAddValid(t *testing.T) {
	e, pc := initTestEcho()

	// compose request
	newPerson, err := json.Marshal(map[string]string{
		"name":     "ncek",
		"email":    "nceki@gmail.com",
		"password": "root",
	})
	if err != nil {
		t.Errorf("marshalling new person failed")
	}
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(newPerson))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/persons")

	// send request
	if err = pc.Add(c); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	// compare status
	if rec.Code != 200 {
		t.Errorf("should return 200, get: %d", rec.Code)
	}

	// compare response
	var p appModel.Person
	if err = json.Unmarshal(rec.Body.Bytes(), &p); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}
	expectedName := "ncek"
	if p.Name != expectedName {
		t.Errorf("person name should be %s, get: %s", expectedName, p.Name)
	}
	expectedEmail := "nceki@gmail.com"
	if p.Email != expectedEmail {
		t.Errorf("person email should be %s, get: %s", expectedEmail, p.Email)
	}
	expectedPassword := "root"
	if p.Password != expectedPassword {
		t.Errorf("person pasword should be %s, get: %s", expectedPassword, p.Password)
	}
}

func TestPersonLogin(t *testing.T) {
	e, pc := initTestEcho()
	person1 := appModel.Person{Name: "ncek", Email: "nceki@gmail.com", Password: "root"}
	person1.ID = uint(1)
	pc.model.Add(person1)

	// login request
	loginInfo, err := json.Marshal(LoginInfo{Email: "nceki@gmail.com", Password: "root"})
	if err != nil {
		t.Errorf("marshalling new person failed")
	}
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(loginInfo))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/login")

	// send request
	if err := pc.Login(c); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	// compare status
	if rec.Code != 200 {
		t.Errorf("should return 200, get: %d", rec.Code)
	}

	// compare response
	var p appModel.Person
	if err := json.Unmarshal(rec.Body.Bytes(), &p); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}
	if p.Token == "" {
		t.Errorf("token expected")
	}
}

func TestGetAll(t *testing.T) {
	e, pc := initTestEcho()
	person1 := appModel.Person{Name: "ncek", Email: "nceki@gmail.com", Password: "root"}
	person1.ID = uint(1)
	pc.model.Add(person1)

	// login request
	loginInfo, err := json.Marshal(LoginInfo{Email: "nceki@gmail.com", Password: "root"})
	if err != nil {
		t.Errorf("marshalling new person failed")
	}
	loginReq := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(loginInfo))
	loginReq.Header.Set("Content-Type", "application/json")
	loginRec := httptest.NewRecorder()
	loginContext := e.NewContext(loginReq, loginRec)
	loginContext.SetPath("/login")

	// send request
	if err := pc.Login(loginContext); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	// compare status
	if loginRec.Code != 200 {
		t.Errorf("should return 200, get: %d", loginRec.Code)
	}

	// compare response
	var p appModel.Person
	if err := json.Unmarshal(loginRec.Body.Bytes(), &p); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}
	if p.Token == "" {
		t.Errorf("token expected")
	}
	token := p.Token

	// get all request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath("/persons")

	// get all
	if err := middleware.JWT([]byte(pc.jwtSecret))(pc.GetAll)(context); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	// compare status
	if rec.Code != 200 {
		t.Errorf("should return 200, get: %d", rec.Code)
	}

	var pList []appModel.Person
	if err := json.Unmarshal(rec.Body.Bytes(), &pList); err != nil {
		t.Errorf("unmarshalling returned person list failed")
	}

	expectedPListLength := 2
	if len(pList) != expectedPListLength {
		t.Errorf("expecting pList's length to be %d, get %d, data: %#v", expectedPListLength, len(pList), pList)
	}
}
