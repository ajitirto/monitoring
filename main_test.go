package main

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
    gin.SetMode(gin.TestMode)

    r := gin.Default()


    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, Prometheus!")
    })

    r.GET("/get-user", func(c *gin.Context) {
        param := c.DefaultQuery("param", "")
        if param == "error" {
            c.JSON(http.StatusInternalServerError, gin.H{
                "message": internalServerError,
            })
            return
        }
        if param == "not-found" {
            c.JSON(http.StatusNotFound, gin.H{
                "message": notFound,
            })
            return
        }
        c.String(http.StatusOK, "Success Get Users")
    })
    

    r.GET("/get-role", func(c *gin.Context) {
        param := c.DefaultQuery("param", "")
        if param == "error" {
            c.JSON(http.StatusInternalServerError, gin.H{
                "message": internalServerError,
            })
            return
        }
        if param == "not-found" {
            c.JSON(http.StatusNotFound, gin.H{
                "message": notFound,
            })
            return
        }
        c.String(http.StatusOK, "Success Get Roles")
    })

    return r
}

func TestRootRoute(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Equal(t, "Hello, Prometheus!", w.Body.String())
}

func TestGetUserSuccess(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/get-user", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Equal(t, "Success Get Users", w.Body.String())
}

func TestGetUserInternalServerError(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/get-user?param=error", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusInternalServerError, w.Code)
    expectedBody := `{"message":"Internal Server Error"}`
    assert.Contains(t, w.Body.String(), expectedBody)
}

func TestGetUserNotFound(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/get-user?param=not-found", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusNotFound, w.Code)
    expectedBody := `{"message":"Not Found"}`
    assert.Contains(t, w.Body.String(), expectedBody)
}

func TestGetRoleSuccess(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/get-role", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Equal(t, "Success Get Roles", w.Body.String())
}