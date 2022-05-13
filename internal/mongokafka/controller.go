package mongokafka

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	service Service
}

func (controller Controller) getKafkalog(c echo.Context) error {
	return c.JSON(http.StatusOK, "success")
}

func (controller Controller) createKafkalog(c echo.Context) error {
	request := new(CreateKafkaLogRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	result, err := controller.service.Create(c.Request().Context(), *request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, result)
}

func NewController(s Service) Controller {
	return Controller{service: s}
}

func RegisterHandler(instance *echo.Echo, controller Controller) {
	instance.GET("/", func(c echo.Context) error {
		c.JSON(http.StatusOK, "kafkalog")
		return nil
	})
	instance.GET("api/kafkalog", controller.getKafkalog)
	instance.POST("api/kafkalog", controller.createKafkalog)
}
