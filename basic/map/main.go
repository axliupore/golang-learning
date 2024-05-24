package main

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)

// CreateMap 创建一个 map
func CreateMap() {
	fmt.Println("===CreateMap===")
	// 1.创建一个没有初始化的 map
	var m1 map[string]int
	// 这里运行时会发生错误
	// m1["one"] = 10
	fmt.Println(m1["one"]) // 这里输出值的默认值

	// 2.创建一个字面量的 map
	m2 := map[string]int{"two": 10}
	// 这里可以直接赋值，不会报错
	m2["three"] = 10
	fmt.Println(m2["three"])

	// 3.使用 make 创建 map，默认为 10 个大小
	m3 := make(map[string]int, 10)
	// 直接赋值即可
	m3["four"] = 10
	fmt.Println(m3["four"])
}

// ExistKey 判断键是否存在
func ExistKey() {
	fmt.Println("===ExistKey===")
	m := make(map[string]int, 10)
	value, exist := m["axliu"]
	if exist {
		fmt.Println("key: ", value)
	}
	m["axliu"] = 10
	value, exist = m["axliu"]
	if exist {
		fmt.Println("key: ", value)
	}
}

// RangeMap 遍历 map
func RangeMap() {
	fmt.Println("===RangeMap===")
	m := make(map[string]int, 10)
	m["one"] = 10
	m["two"] = 20
	m["three"] = 30
	for k, v := range m {
		fmt.Println(k, v)
	}
	// 当只有一个数时，这个变量是 map 的键
	fmt.Println("只有一个变量时：")
	for k := range m {
		fmt.Println(k, m[k])
	}
}

// DeleteMap 删除某个键
func DeleteMap() {
	fmt.Println("===DeleteMap===")
	m := make(map[string]int, 10)
	m["one"] = 10
	m["two"] = 20
	m["three"] = 30
	// 删除 one 这个键
	delete(m, "one")
	// 查看删除之后还是否存在
	v, exist := m["one"]
	if !exist {
		fmt.Println("key not exist")
	} else {
		fmt.Println("key one's value is", v)
	}
}

// SortMap 按照指定顺序遍历 map
func SortMap() {
	fmt.Println("===SortMap===")
	rand.New(rand.NewSource(time.Now().UnixNano())) // 初始化随机种子
	m := make(map[string]int, 20)
	for i := range 20 {
		key := fmt.Sprintf("key%02d", i)
		value := rand.Intn(100) // 生成 0～99 的随机数
		m[key] = value
	}
	// 取出 map 中的所有 key 存入 keys 切片中
	keys := make([]string, 0, 20)
	for key := range m {
		keys = append(keys, key)
	}
	// 对切片进行排序
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println("key:", k, "value:", m[k])
	}
}

func main() {
	CreateMap()
	ExistKey()
	RangeMap()
	DeleteMap()
	SortMap()
}
