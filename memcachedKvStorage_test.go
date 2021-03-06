package storage

import (
	"reflect"
	"strconv"
	"testing"
)

type T struct {
	A int
}

func TestGetSet(t *testing.T) {
	tt := T{1}
	mcStorage := NewMcStorage([]string{"localhost:12000"}, "test", 0, reflect.TypeOf(&tt))
	mcStorage.Set("1", tt)
	res, _ := mcStorage.Get("1")
	defer mcStorage.Delete("1")
	if reflect.TypeOf(res) != reflect.TypeOf(&tt) {
		t.Error("res type is not T")
	}
	ttRes := res.(*T)
	if ttRes.A != tt.A {
		t.Error("res A field is not equals tt field")
	}
}

func TestMultiGetSet(t *testing.T) {
	tt := T{1}
	mcStorage := NewMcStorage([]string{"localhost:12000"}, "test", 0, reflect.TypeOf(&tt))
	valueMap := make(map[interface{}]interface{})
	keys := make([]interface{}, 10)
	for i := 0; i < 10; i++ {
		keys[i] = strconv.Itoa(i)
		valueMap[strconv.Itoa(i)] = T{i}
		defer mcStorage.Delete(strconv.Itoa(i))
	}
	mcStorage.MultiSet(valueMap)
	res, _ := mcStorage.MultiGet(keys)
	for k, v := range res {
		if reflect.TypeOf(v) != reflect.TypeOf(&tt) {
			t.Error("res type is not T")
		}
		kint, err := strconv.Atoi(k.(string))
		if err != nil {
			t.Error("key %s is not int ", k)
		}
		vT := v.(*T)
		if kint != vT.A {
			t.Error("value should be %s,while it is %s", kint, vT.A)
		}
	}
}

func TestGetSetDelete(t *testing.T) {
	tt := T{1}
	mcStorage := NewMcStorage([]string{"localhost:12000"}, "test", 0, reflect.TypeOf(&tt))
	mcStorage.Set("1", tt)
	res, _ := mcStorage.Get("1")
	if reflect.TypeOf(res) != reflect.TypeOf(&tt) {
		t.Error("res type is not T")
	}
	ttRes := res.(*T)
	if ttRes.A != tt.A {
		t.Error("res A field is not equals tt field")
	}
	mcStorage.Delete("1")
	res, _ = mcStorage.Get("1")
	if res != nil {
		t.Error("res should be nil ,after delete")
	}
}
