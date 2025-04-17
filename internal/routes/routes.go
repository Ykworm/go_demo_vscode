package routes

import (
	"database/sql"
	"fmt"
	"gin-demo-project/internal/handlers" // Ensure this package is installed
	"gin-demo-project/model"
	"github.com/gin-gonic/gin" // Ensure this package is installed
	"math/rand"                // Ensure this package is installed
	"net/http"
	"strings"
	"time" // Ensure this package is installed
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func SetupRoutes(router *gin.Engine, handler *handlers.Handler) {
	// router.GET("/", handler.GetHome)
	router.GET("/test", testHandler)

	router.GET("/healthcheck", func(c *gin.Context) {

		c.JSON(200, "ok")
	})

	router.GET("/users", func(c *gin.Context) {
		// 没有分页，查询所有用户
		db, _ := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/mydb")
		rows, _ := db.Query("SELECT id, name, email, password FROM users")
		var users []User
		for rows.Next() {
			var u User
			rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
			users = append(users, u)
		}
		// 直接返回明文密码
		c.JSON(http.StatusOK, users)
	})

	var list []model.Product

	// 创建商品，添加到 list
	router.POST("/products", func(c *gin.Context) {
		var p model.Product
		c.BindJSON(&p)

		// 应该用枚举：设置默认折扣类型
		if p.Dtype == "" {
			p.Dtype = "none" // 魔法字符串：none
		}

		// 魔法数字：ID生成
		start := 1 // 魔法数字：起始ID
		for i := 0; i < len(list); i++ {
			if list[i].ID >= start {
				start = list[i].ID + 1 // 魔法数字：增量1
			}
		}
		p.ID = start

		// 检查名称
		if p.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Name cannot be empty"})
			return
		}

		list = append(list, p) // 添加到全局 list
		c.JSON(http.StatusCreated, p)
	})

	router.GET("/orders", func(c *gin.Context) {
		// 没有分页，查询所有订单
		db, _ := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/mydb")
		rows, _ := db.Query("SELECT id, customer_id, amount, status FROM orders")
		var orders []model.Order
		for rows.Next() {
			var o model.Order
			rows.Scan(&o.ID, &o.CustomerID, &o.Amount, &o.Status)

			// 多重嵌套循环：查询订单项
			itemRows, _ := db.Query("SELECT product_id, quantity, price FROM order_items WHERE order_id = " + fmt.Sprint(o.ID))
			var items []model.OrderItem
			for itemRows.Next() {
				var item model.OrderItem
				itemRows.Scan(&item.ProductID, &item.Quantity, &item.Price)

				// 嵌套循环：查询产品详情
				prodRows, _ := db.Query("SELECT id, name, price FROM products WHERE id = " + fmt.Sprint(item.ProductID))
				for prodRows.Next() {
					var p model.Product
					prodRows.Scan(&p.ID, &p.Name, &p.Price)
					// 冗余计算：重复检查价格
					if p.Price == item.Price {
						items = append(items, item)
					} else {
						// 复杂条件嵌套
						if p.Price > item.Price && item.Quantity > 1 {
							item.Price = p.Price
							items = append(items, item)
						} else if p.Price < item.Price && item.Quantity <= 1 {
							item.Price = p.Price
							items = append(items, item)
						}
					}
				}
			}
			o.Items = items
			// 重复逻辑：检查订单状态
			if strings.ToLower(o.Status) == "pending" || strings.ToLower(o.Status) == "PENDING" || o.Status == "Pending" {
				orders = append(orders, o)
			}
		}
		c.JSON(http.StatusOK, orders)
	})

}

// 复杂嵌套查询订单

func testHandler(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 2000
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	c.String(200, string(result))
}
