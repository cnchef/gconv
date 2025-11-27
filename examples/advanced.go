package main

import (
	"fmt"
	"github.com/cnchef/gconv"
)

// User 用户结构体
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}

// Product 产品结构体
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	fmt.Println("=== gconv 高级使用示例 ===\n")

	// 示例 1: Map 转 Struct
	fmt.Println("1. Map 转 Struct:")
	userMap := map[string]any{
		"id":        123,
		"name":      "张三",
		"email":     "zhangsan@example.com",
		"is_active": true,
	}
	user := gconv.Cast[User](userMap)
	fmt.Printf("   输入: %+v\n", userMap)
	fmt.Printf("   输出: %+v\n\n", user)

	// 示例 2: JSON 动态类型转换
	fmt.Println("2. 处理 JSON 动态数据:")
	jsonData := map[string]any{
		"count":  "42",        // 字符串数字
		"price":  "99.99",     // 字符串浮点数
		"active": "true",      // 字符串布尔值
		"name":   123,         // 数字转字符串
	}
	fmt.Printf("   count (string->int): %d\n", gconv.Cast[int](jsonData["count"]))
	fmt.Printf("   price (string->float64): %.2f\n", gconv.Cast[float64](jsonData["price"]))
	fmt.Printf("   active (string->bool): %t\n", gconv.Cast[bool](jsonData["active"]))
	fmt.Printf("   name (int->string): %s\n\n", gconv.Cast[string](jsonData["name"]))

	// 示例 3: 切片数据处理
	fmt.Println("3. 切片数据处理:")
	mixedSlice := []any{1, "2", 3.0, "4.5", true}
	fmt.Printf("   原始切片: %v\n", mixedSlice)
	
	var intSlice []int
	for _, item := range mixedSlice {
		intSlice = append(intSlice, gconv.Cast[int](item))
	}
	fmt.Printf("   转为 int 切片: %v\n\n", intSlice)

	// 示例 4: API 响应数据转换
	fmt.Println("4. 模拟 API 响应数据转换:")
	apiResponse := map[string]any{
		"code": "200",
		"data": map[string]any{
			"id":    "1001",
			"name":  "笔记本电脑",
			"price": "5999.00",
		},
		"success": "true",
	}

	code := gconv.Cast[int](apiResponse["code"])
	success := gconv.Cast[bool](apiResponse["success"])
	productData := gconv.Cast[map[string]any](apiResponse["data"])
	
	product := Product{
		ID:    gconv.Cast[int](productData["id"]),
		Name:  gconv.Cast[string](productData["name"]),
		Price: gconv.Cast[float64](productData["price"]),
	}

	fmt.Printf("   响应码: %d\n", code)
	fmt.Printf("   是否成功: %t\n", success)
	fmt.Printf("   产品信息: %+v\n\n", product)

	// 示例 5: 数据清洗与类型统一
	fmt.Println("5. 数据清洗与类型统一:")
	dirtyData := []any{
		"100",
		200,
		"300.50",
		nil,
		"invalid",
		true,
	}

	cleanedNumbers := make([]float64, 0)
	for i, item := range dirtyData {
		num := gconv.Cast[float64](item)
		cleanedNumbers = append(cleanedNumbers, num)
		fmt.Printf("   [%d] %v (%T) -> %.2f\n", i, item, item, num)
	}
	fmt.Printf("   清洗后的数据: %v\n\n", cleanedNumbers)

	// 示例 6: 嵌套结构转换
	fmt.Println("6. 嵌套结构转换:")
	type Order struct {
		OrderID  int     `json:"order_id"`
		User     User    `json:"user"`
		Product  Product `json:"product"`
		Quantity int     `json:"quantity"`
	}

	orderMap := map[string]any{
		"order_id": 10001,
		"quantity": "5",
		"user": map[string]any{
			"id":        1,
			"name":      "李四",
			"email":     "lisi@example.com",
			"is_active": true,
		},
		"product": map[string]any{
			"id":    2001,
			"name":  "键盘",
			"price": "299.00",
		},
	}

	order := gconv.Cast[Order](orderMap)
	fmt.Printf("   订单信息: %+v\n", order)
	fmt.Printf("   用户: %+v\n", order.User)
	fmt.Printf("   产品: %+v\n\n", order.Product)

	// 示例 7: 安全的类型转换（避免 panic）
	fmt.Println("7. 安全的类型转换:")
	var nilValue any = nil
	var complexType any = make(chan int)
	
	fmt.Printf("   nil -> int: %d\n", gconv.Cast[int](nilValue))
	fmt.Printf("   nil -> string: '%s'\n", gconv.Cast[string](nilValue))
	fmt.Printf("   nil -> bool: %t\n", gconv.Cast[bool](nilValue))
	fmt.Printf("   chan -> int: %d (不会panic)\n\n", gconv.Cast[int](complexType))

	// 示例 8: 批量数据转换
	fmt.Println("8. 批量数据转换:")
	users := []map[string]any{
		{"id": "1", "name": "用户1", "email": "user1@test.com", "is_active": "true"},
		{"id": "2", "name": "用户2", "email": "user2@test.com", "is_active": "false"},
		{"id": "3", "name": "用户3", "email": "user3@test.com", "is_active": "1"},
	}

	var userList []User
	for _, userMap := range users {
		userList = append(userList, gconv.Cast[User](userMap))
	}

	fmt.Println("   批量转换结果:")
	for i, u := range userList {
		fmt.Printf("   [%d] %+v\n", i+1, u)
	}

	fmt.Println("\n=== 示例结束 ===")
}

