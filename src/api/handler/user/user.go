package userHandler

import (
	"fmt"
	"net/http"
	"strconv"

	userDomain "test-majoo/src/domain/user"
	helpers "test-majoo/src/helper"
	middlewares "test-majoo/src/middleware"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (a *AppHandler) GetListUser(c *gin.Context) {
	errorParams := map[string]interface{}{}
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	search := c.Query("search")
	sort := c.Query("sort")

	if page == 0 {
		page = helpers.DefaultPage
	}

	limit, offset := helpers.PaginationPageOffset(page, limit)
	data, count, err := a.Entity.GetListUser(c, offset, limit, search, sort)
	if err != nil {
		statusCode := http.StatusBadRequest
		errorParams["meta"] = map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(errorParams))
		return
	}

	pagination := helpers.PaginationRes(page, count, limit)
	params := map[string]interface{}{
		"payload": data,
		"meta":    pagination,
	}
	c.JSON(http.StatusOK, helpers.OutputAPIResponseWithPayload(params))
}

func (a *AppHandler) Login(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200
	var body userDomain.Login
	err := c.ShouldBindJSON(&body)
	if err != nil {
		statusCode = 406
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(errorParams))
		return
	}

	data, err := a.Entity.Login(c, body)
	if err != nil {
		statusCode = 400
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(errorParams))
		return
	}

	// Compare Password
	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(body.Password))
	if err != nil {
		statusCode = 400
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(errorParams))
		return
	}

	token, refreshToken, err := middlewares.CreateToken(data)
	if err != nil {
		statusCode = 400
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(errorParams))
		return
	}

	data.Password = ""

	params := map[string]interface{}{
		"meta": map[string]interface{}{
			"token":        token,
			"refreshToken": refreshToken,
		},
		"payload": data,
	}
	c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(params))
}

func (a *AppHandler) Create(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200
	var body userDomain.SetUser
	err := c.ShouldBindJSON(&body)
	if err != nil {
		statusCode = 406
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(errorParams))
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		fmt.Println(err)
	}
	body.Password = string(bytes)

	err = a.Entity.Create(c, body)
	if err != nil {
		statusCode = 400
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(errorParams))
		return
	}

	params := map[string]interface{}{
		"meta":    "success",
		"payload": body,
	}
	c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(params))
}

func (a *AppHandler) Update(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200
	param := c.Param("id")
	id, _ := strconv.Atoi(param)

	var body userDomain.SetUser
	err := c.ShouldBindJSON(&body)
	if err != nil {
		statusCode = 406
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(errorParams))
		return
	}

	err = a.Entity.Update(c, body, id)
	if err != nil {
		statusCode = 400
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(errorParams))
		return
	}

	params := map[string]interface{}{
		"meta":    "success",
		"payload": body,
	}
	c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(params))
}

func (a *AppHandler) Delete(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200
	param := c.Param("id")
	id, _ := strconv.Atoi(param)

	err := a.Entity.Delete(c, id)
	if err != nil {
		statusCode = 400
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(errorParams))
		return
	}

	params := map[string]interface{}{
		"meta": "success",
	}
	c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(params))
}
