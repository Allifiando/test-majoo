package productHandler

import (
	"net/http"
	"strconv"

	productDomain "test-majoo/src/domain/product"
	helpers "test-majoo/src/helper"
	middlewares "test-majoo/src/middleware"
	Validator "test-majoo/src/pkg/validator"

	"github.com/gin-gonic/gin"
)

func (a *AppHandler) GetListMerchantByUserId(c *gin.Context) {
	errorParams := map[string]interface{}{}
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	search := c.Query("search")
	sort := c.Query("sort")

	User := middlewares.GetUserCustom(c)

	if page == 0 {
		page = helpers.DefaultPage
	}

	limit, offset := helpers.PaginationPageOffset(page, limit)
	data, count, err := a.Product.GetListMerchantByUserId(c, offset, limit, search, sort, int64(User["id"].(float64)))
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

func (a *AppHandler) GetListOutletByMerchantId(c *gin.Context) {
	errorParams := map[string]interface{}{}
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	search := c.Query("search")
	sort := c.Query("sort")
	merchantId, _ := strconv.Atoi(c.Query("merchant_id"))

	if page == 0 {
		page = helpers.DefaultPage
	}

	limit, offset := helpers.PaginationPageOffset(page, limit)
	data, count, err := a.Product.GetListOutletByMerchantId(c, offset, limit, search, sort, int64(merchantId))
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

func (a *AppHandler) GetListProductByOutletId(c *gin.Context) {
	errorParams := map[string]interface{}{}
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	search := c.Query("search")
	sort := c.Query("sort")
	outletId, _ := strconv.Atoi(c.Query("outlet_id"))

	if page == 0 {
		page = helpers.DefaultPage
	}

	limit, offset := helpers.PaginationPageOffset(page, limit)
	data, count, err := a.Product.GetListProductByOutletId(c, offset, limit, search, sort, int64(outletId))
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

func (a *AppHandler) GetMerchantById(c *gin.Context) {
	errorParams := map[string]interface{}{}
	id, _ := strconv.Atoi(c.Query("id"))

	data, err := a.Product.GetMerchantById(c, int64(id))
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

	params := map[string]interface{}{
		"meta":    "success",
		"payload": data,
	}
	c.JSON(http.StatusOK, helpers.OutputAPIResponseWithPayload(params))
}

func (a *AppHandler) GetOutletById(c *gin.Context) {
	errorParams := map[string]interface{}{}
	id, _ := strconv.Atoi(c.Query("id"))

	data, err := a.Product.GetOutletById(c, int64(id))
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

	params := map[string]interface{}{
		"meta":    "success",
		"payload": data,
	}
	c.JSON(http.StatusOK, helpers.OutputAPIResponseWithPayload(params))
}

func (a *AppHandler) GetProductById(c *gin.Context) {
	errorParams := map[string]interface{}{}
	id, _ := strconv.Atoi(c.Query("id"))

	data, err := a.Product.GetProductById(c, int64(id))
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

	params := map[string]interface{}{
		"meta":    "success",
		"payload": data,
	}
	c.JSON(http.StatusOK, helpers.OutputAPIResponseWithPayload(params))
}

func (a *AppHandler) CreateMerchant(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200
	var body productDomain.SetMerchant
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

	if ok, err := Validator.IsRequestValid(body); !ok && err != nil {
		statusCode := http.StatusBadRequest
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(errorParams))
		return
	}

	exist, _ := a.Product.GetMerchantByName(c, body.CompanyName)
	if exist.CompanyName != "" {
		statusCode = 400
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": "Data sudah ada",
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(errorParams))
		return
	}

	err = a.Product.CreateMerchant(c, body)
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

func (a *AppHandler) CreateOutlet(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200
	var body productDomain.SetOutlet
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

	err = a.Product.CreateOutlet(c, body)
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

func (a *AppHandler) CreateProduct(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200
	var body productDomain.SetProduct

	outletForm := c.PostForm("outlet_id")
	nameForm := c.PostForm("name")
	priceForm := c.PostForm("price")
	qtyForm := c.PostForm("qty")
	file, _ := c.FormFile("file")
	isActiveForm := c.PostForm("is_active")

	outletId, _ := strconv.Atoi(outletForm)
	price, _ := strconv.Atoi(priceForm)
	isActive, _ := strconv.ParseBool(isActiveForm)

	//upload
	path := "tmp/" + file.Filename
	err := c.SaveUploadedFile(file, path)
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

	body.OutletId = int64(outletId)
	body.Name = nameForm
	body.Price = float64(price)
	body.Qty = qtyForm
	body.Filename = path
	body.IsActive = isActive

	err = a.Product.CreateProduct(c, body)
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

func (a *AppHandler) UpdateProduct(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200
	id, _ := strconv.Atoi(c.Query("id"))

	var body productDomain.SetProduct
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

	err = a.Product.UpdateProduct(c, body, int64(id))
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

func (a *AppHandler) DeleteProduct(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200
	param := c.Param("id")
	id, _ := strconv.Atoi(param)

	err := a.Product.DeleteProduct(c, int64(id))
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
