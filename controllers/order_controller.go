package controllers

import (
    "net/http"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
    "main.go/models"
    "main.go/utils"
    "gorm.io/gorm"
)

type OrderController struct {
    DB *gorm.DB
}

func NewOrderController(db *gorm.DB) *OrderController {
    return &OrderController{DB: db}
}

func (oc *OrderController) CreateOrder(c *gin.Context) {
    var order models.Order
    if err := c.BindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    order.OrderedAt = time.Now().UTC()

    if err := oc.DB.Create(&order).Error; err != nil {
        utils.LogError(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
        return
    }

    c.JSON(http.StatusCreated, order)
}

func (oc *OrderController) GetOrder(c *gin.Context) {
    orderID, err := strconv.ParseUint(c.Param("orderID"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
        return
    }

    var order models.Order
    if err := oc.DB.Preload("Items").First(&order, orderID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
            return
        }
        utils.LogError(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order"})
        return
    }

    c.JSON(http.StatusOK, order)
}

func (oc *OrderController) UpdateOrder(c *gin.Context) {
    orderID, err := strconv.ParseUint(c.Param("orderID"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
        return
    }

    var order models.Order
    if err := oc.DB.First(&order, orderID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
            return
        }
        utils.LogError(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order"})
        return
    }

    var updatedOrder models.Order
    if err := c.BindJSON(&updatedOrder); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    order.CustomerName = updatedOrder.CustomerName
    order.Items = updatedOrder.Items
    order.Status = updatedOrder.Status

    if err := oc.DB.Save(&order).Error; err != nil {
        utils.LogError(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
        return
    }

    c.JSON(http.StatusOK, order)
}

func (oc *OrderController) DeleteOrder(c *gin.Context) {
    orderID, err := strconv.ParseUint(c.Param("orderID"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
        return
    }

    var order models.Order
    if err := oc.DB.First(&order, orderID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
            return
        }
        utils.LogError(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order"})
        return
    }

    if order.ID == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    if err := oc.DB.Delete(&order).Error; err != nil {
        utils.LogError(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
