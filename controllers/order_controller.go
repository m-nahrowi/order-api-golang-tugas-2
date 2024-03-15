package controllers

import (
	"net/http"
	"strconv"
	"time"

	// "log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/models"
	"main.go/utils"
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

// update 1
// func (oc *OrderController) UpdateOrder(c *gin.Context) {
//     orderID, err := strconv.ParseUint(c.Param("orderID"), 10, 64)
//     if err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
//         return
//     }

//     var order models.Order
//     if err := oc.DB.Preload("Items").First(&order, orderID).Error; err != nil {
//         if err == gorm.ErrRecordNotFound {
//             c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
//             return
//         }
//         utils.LogError(err)
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order"})
//         return
//     }

//     var updatedOrder models.Order
//     if err := c.BindJSON(&updatedOrder); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }

//     // Memperbarui nama pelanggan dan status pesanan
//     order.CustomerName = updatedOrder.CustomerName
//     order.Status = updatedOrder.Status

//     // Memperbarui item-item yang ada
//     for i := range updatedOrder.Items {
//         // Cari item berdasarkan ID
//         var existingItem models.Item
//         if err := oc.DB.First(&existingItem, updatedOrder.Items[i].ID).Error; err != nil {
//             if err == gorm.ErrRecordNotFound {
//                 // Jika item tidak ditemukan, tambahkan sebagai item baru
//                 order.Items = append(order.Items, updatedOrder.Items[i])
//             } else {
//                 utils.LogError(err)
//                 c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch item"})
//                 return
//             }
//         } else {
//             // Jika item ditemukan, perbarui item yang ada
//             existingItem.ItemCode = updatedOrder.Items[i].ItemCode
//             existingItem.Description = updatedOrder.Items[i].Description
//             existingItem.Quantity = updatedOrder.Items[i].Quantity
//             order.Items[i] = existingItem
//         }
//     }

//     // Simpan pesanan yang telah diperbarui
//     if err := oc.DB.Save(&order).Error; err != nil {
//         utils.LogError(err)
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
//         return
//     }

//     c.JSON(http.StatusOK, order)
// }

// update 2
// func (oc *OrderController) UpdateOrder(c *gin.Context) {
//     orderID, err := strconv.ParseUint(c.Param("orderID"), 10, 64)
//     if err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
//         return
//     }

//     var order models.Order
//     if err := oc.DB.Preload("Items").First(&order, orderID).Error; err != nil {
//         if err == gorm.ErrRecordNotFound {
//             c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
//             return
//         }
//         utils.LogError(err)
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order"})
//         return
//     }

//     var updatedOrder models.Order
//     if err := c.BindJSON(&updatedOrder); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }

//     // Memperbarui nama pelanggan dan status pesanan
//     order.CustomerName = updatedOrder.CustomerName
//     order.Status = updatedOrder.Status

//     // Memperbarui item-item yang ada
//     updatedItems := updatedOrder.Items

//     // Membuat map untuk menampung item yang sudah ada
//     existingItemsMap := make(map[uint]models.Item)
//     for _, existingItem := range order.Items {
//         existingItemsMap[existingItem.ID] = existingItem
//     }

//     // Memperbarui item yang ada dan menambahkan item baru
//     for _, updatedItem := range updatedItems {
//         if existingItem, ok := existingItemsMap[updatedItem.ID]; ok {
//             // Jika item ditemukan, perbarui item yang ada
//             existingItem.ItemCode = updatedItem.ItemCode
//             existingItem.Description = updatedItem.Description
//             existingItem.Quantity = updatedItem.Quantity
//         } else {
//             // Jika item tidak ditemukan, tambahkan sebagai item baru
//             order.Items = append(order.Items, updatedItem)
//         }
//     }

//     // Simpan pesanan yang telah diperbarui
//     if err := oc.DB.Save(&order).Error; err != nil {
//         utils.LogError(err)
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
//         return
//     }

//     c.JSON(http.StatusOK, order)
// }

// update 3

// func (oc *OrderController) UpdateOrder(c *gin.Context) {
//     orderID, err := strconv.ParseUint(c.Param("orderID"), 10, 64)
//     if err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
//         return
//     }

//     var order models.Order
//     if err := oc.DB.Preload("Items").First(&order, orderID).Error; err != nil {
//         if err == gorm.ErrRecordNotFound {
//             c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
//             return
//         }
//         utils.LogError(err)
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order"})
//         return
//     }

//     var updatedOrder models.Order
//     if err := c.BindJSON(&updatedOrder); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }

//     // Memperbarui nama pelanggan dan status pesanan
//     order.CustomerName = updatedOrder.CustomerName
//     order.Status = updatedOrder.Status

//     // Menyimpan item-item yang sudah diproses untuk dihapus
//     itemIDsToDelete := make(map[uint]bool)

//     // Memperbarui item-item yang ada
//     for i := range updatedOrder.Items {
//         // Cari item berdasarkan ID
//         var existingItem models.Item
//         if err := oc.DB.First(&existingItem, updatedOrder.Items[i].ID).Error; err != nil {
//             if err == gorm.ErrRecordNotFound {
//                 // Jika item tidak ditemukan, tambahkan sebagai item baru
//                 order.Items = append(order.Items, updatedOrder.Items[i])
//             } else {
//                 utils.LogError(err)
//                 c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch item"})
//                 return
//             }
//         } else {
//             // Jika item ditemukan, perbarui item yang ada
//             existingItem.ItemCode = updatedOrder.Items[i].ItemCode
//             existingItem.Description = updatedOrder.Items[i].Description
//             existingItem.Quantity = updatedOrder.Items[i].Quantity
//             order.Items[i] = existingItem

//             // Tandai item yang akan dihapus dari pesanan
//             itemIDsToDelete[existingItem.ID] = true
//         }
//     }

//     // Hapus item yang tidak ada dalam payload
//     for i := range order.Items {
//         if _, exists := itemIDsToDelete[order.Items[i].ID]; !exists {
//             if err := oc.DB.Delete(&order.Items[i]).Error; err != nil {
//                 utils.LogError(err)
//                 c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
//                 return
//             }
//         }
//     }

//     // Simpan pesanan yang telah diperbarui
//     if err := oc.DB.Update(&order).Error; err != nil {
//         utils.LogError(err)
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
//         return
//     }

//     c.JSON(http.StatusOK, order)
// }

// update 4
func (oc *OrderController) UpdateOrder(c *gin.Context) {
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

	var updatedOrder models.Order
	if err := c.BindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Menghapus semua item yang terkait dengan pesanan
	if err := oc.DB.Where("order_id = ?", order.ID).Delete(&models.Item{}).Error; err != nil {
		utils.LogError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete items"})
		return
	}

	// Memperbarui nama pelanggan dan status pesanan
	order.CustomerName = updatedOrder.CustomerName
	order.Status = updatedOrder.Status

	// Memperbarui item-item yang ada
	for _, updatedItem := range updatedOrder.Items {
		// Tambahkan kembali item yang diperbarui ke dalam pesanan
		// order.Items = append(order.Items, updatedItem)
		for i, updatedOrderItem := range order.Items {
			if updatedOrderItem.ID == updatedItem.ID {
				order.Items[i] = updatedItem
			} 
            // else {
			// 	order.Items = append(order.Items, updatedItem)
			// }
		}
	}

	// Simpan pesanan yang telah diperbarui
	if err := oc.DB.Save(&order).Error; err != nil {
		utils.LogError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// delete 2
func (oc *OrderController) DeleteOrder(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("orderID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	// Cek apakah pesanan yang akan dihapus ada
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

	// Menghapus semua item yang terhubung dengan pesanan
	if err := oc.DB.Where("order_id = ?", orderID).Delete(&models.Item{}).Error; err != nil {
		utils.LogError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete items"})
		return
	}

	// Menghapus pesanan setelah item-item terhubung dihapus
	if err := oc.DB.Delete(&order).Error; err != nil {
		utils.LogError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

// get All data
func (oc *OrderController) GetAllOrder(c *gin.Context) {
	var orders []models.Order

	// Memuat item-item terkait dengan pesanan
	if err := oc.DB.Preload("Items").Find(&orders).Error; err != nil {
		utils.LogError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
