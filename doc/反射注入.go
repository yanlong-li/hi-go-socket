package doc

import "reflect"

// 从字节流中反射出对应的结构体并注入到指定方法中
func Unmarshal(f interface{}) {
	t := reflect.TypeOf(f)
	//构造一个存放函数实参 Value 值的数纽
	in := make([]reflect.Value, t.NumIn())
	// 取出所有需要注入的依赖参数
	for i := 0; i < t.NumIn(); i++ {
		// 获取顺序的 参数
		params := t.In(i)
		// 创建一个reflect.value类型的params需要的指针类型的数据
		elem := reflect.New(params).Elem()
		for k := 0; k < elem.NumField(); k++ {
			field := elem.Field(k)
			switch field.Kind() {
			case reflect.String:
				field.SetString("A")
			case reflect.Int:
				field.SetInt(208)
				// todo 继续成所有类型
			}
		}
		in[i] = elem
	}
	reflect.ValueOf(f).Call(in)
}
