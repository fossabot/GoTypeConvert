package typeconv

import (
	"fmt"
	"log"
	"time"
)

func ExampleMapToInterfaceNotUseCast() {
	type User struct {
		Name string
	}
	v := &User{}
	MapToInterface(v, map[string]interface{}{
		"Name": "testUser",
	})
	fmt.Println(v.Name)
	// Output:
	// testUser
}

func ExampleMapToStructInterface() {
	type User struct {
		Name interface{}
		Year interface{}
	}
	v := &User{}
	MapToInterface(v, map[string]interface{}{
		"Name": "testUser",
		"Year": 1999,
	})
	fmt.Println(v.Name)
	fmt.Println(v.Year)
	// Output:
	// testUser
	// 1999
}

func ExampleMapToStructNil() {
	type User struct {
		Name interface{}
		Year interface{}
	}
	v := &User{}
	r := MapToInterface(v, nil)
	fmt.Println(r.IsNil)
	// Output:
	// true
}

func ExampleMapToInterfaceOverrideValue() {
	type User struct {
		Name string
	}
	v := &User{
		Name: "oldUserName",
	}
	MapToInterface(v, map[string]interface{}{
		"Name": "testUser",
	})
	fmt.Println(v.Name)
	// Output:
	// testUser
}

func ExampleMapToInterfaceNotOverrideValue() {
	type User struct {
		Name string
		Addr string
	}
	v := &User{
		Name: "oldUserName",
		Addr: "oldAddr",
	}
	MapToInterface(v, map[string]interface{}{
		"Name": "testUser",
	})
	fmt.Println(v.Addr)
	// Output:
	// oldAddr
}

func ExampleMapToInterfaceInt() {
	type User struct {
		Count int
	}
	v := &User{}
	MapToInterface(v, map[string]interface{}{
		"Count": 100,
	})
	fmt.Println(v.Count)
	// Output:
	// 100
}

func ExampleMapToInterfaceInt64() {
	type User struct {
		Count int64
	}
	v := &User{}
	r := MapToInterface(v, map[string]interface{}{
		"Count": 100,
	})
	fmt.Println(v.Count)
	fmt.Println(r.Error)
	// Output:
	// 0
	// Error MapToInterface: reflect.Set: value of type int is not assignable to type int64
}

func ExampleMapToInterface() {
	type User struct {
		Name    string
		RegTime *time.Time
	}
	t := time.Now()
	v := MapToInterface(&User{}, map[string]interface{}{
		"Name":    "testUser",
		"RegTime": &t,
	})

	// check
	user := v.A.(*User)
	if user.RegTime != &t {
		log.Panic("error ", user.RegTime, t)
	}
	fmt.Println(user.Name)
	// Output:
	// testUser
}

func ExampleMapToInterfaceNilValue() {
	type User struct {
		Name    string
		RegTime *time.Time
	}
	v := MapToInterface(&User{}, map[string]interface{}{
		"Name":    "testUser",
		"RegTime": nil,
	})
	fmt.Println(v.A.(*User).RegTime)
	// Output:
	// <nil>
}

func ExampleMapToInterfaceArray() {
	type User struct {
		Name  string
		Items []string
	}
	v := MapToInterface(&User{}, map[string]interface{}{
		"Name":  "testUser",
		"Items": []string{"A", "B"},
	})
	fmt.Println(v.A.(*User).Items)
	// Output:
	// [A B]
}

func ExampleMapToInterfaceArrayUseRequestData() {
	type User struct {
		Name  string
		Items []string
	}
	o := &User{}
	MapToInterface(o, map[string]interface{}{
		"Name":  "testUser",
		"Items": []string{"A", "B"},
	})
	fmt.Println(o.Name)
	fmt.Println(o.Items)
	// Output:
	// testUser
	// [A B]
}

func ExampleMapToInterfaceArrayInterface() {
	type User struct {
		Name  string
		Items []string
	}
	o := &User{}
	MapToInterface(o, map[string]interface{}{
		"Name":  "testUser",
		"Items": []interface{}{"A", "B"},
	})
	fmt.Println(o.Name)
	fmt.Println(o.Items)
	// Output:
	// testUser
	// [A B]
}

func ExampleMapToInterfaceArrayInterfaceStringInt() {
	type User struct {
		Name  string
		Items []string
	}
	o := &User{}
	MapToInterface(o, map[string]interface{}{
		"Name":  "testUser",
		"Items": []interface{}{"A", "B", 1},
	})
	fmt.Println(o.Name)
	fmt.Println(o.Items)
	// Output:
	// testUser
	// [A B 1]
}
