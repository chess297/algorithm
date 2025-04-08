// Golang 类型 和接口
// 1. 定义结构体
// 2. 定义接口
// 3. 定义类型别名

package l1_test

import "testing"

// 1. type struct
type Student struct {
	Name string
	Age  int
	Person
	Fly
}

func (s *Student) Say() string {
	return s.Name
}

func TestStruct(t *testing.T) {
	// 两种常用的结构体实例化方式
	s := Student{Name: "Tom", Age: 18} // 结构体实例化
	t.Log(s.Say())
	var s2 Person = &Student{Name: "Tom", Age: 18} // 结构体实例化
	t.Log(s2.Say())
}

// 2. type interface
type Person interface {
	Say() string
}

type Fly interface {
	Fly() string
}

func TestInterface(t *testing.T) {
	var a Person = &Student{Name: "Tom", Age: 18}
	t.Log(a.Say())
}

// 3. type alias
type ChinesePerson = Person

func TestAlias(t *testing.T) {
	var s ChinesePerson = &Student{Name: "小明", Age: 18}
	t.Log(s.Say())
}

// 4. 结构体的类型映射
type GDStudent Student // 类型映射

func TestTypeMap(t *testing.T) {
	s := GDStudent{Name: "Tom", Age: 18}
	// t.Log(s.Say()) // 没有Say方法，会报错
	t.Log(s.Name) // 可以通过Person类型的方法访问Say方法
	t.Log(s.Age)  // 可以通过Person类型的方法访问Say方法
}
