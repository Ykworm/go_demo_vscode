package model

type Order struct {
	ID         int
	CustomerID int
	Amount     float64
	Status     string
	Items      []OrderItem
}

type OrderItem struct {
	ProductID int
	Quantity  int
	Price     float64
}

// Product 表示商品，包含折扣信息
type Product struct {
	ID    int    // 魔法变量：名称不明确
	Name  string // 魔法变量：含义模糊
	Price float64
	Dtype string // 应该用枚举：折扣类型用魔法字符串
}
