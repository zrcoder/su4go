package cache

import "testing"


type person struct {
	age int
	name string
}

func TestAll(t *testing.T)  {
	cache := NewWithCapacity(2)
	if cache.capacity != 2 {
		t.Error("wrong capacity")
	}

	p1 := person{18, "Lily"}
	p2 := person{20, "Lucy"}
	cache.Add("p1", p1)
	cache.Add("p2", p2)
	t.Log(cache.items)
	cache.Add("kkk", "vvv")
	t.Log(cache.items)

	cache.Remove("kkk")
	t.Log(cache.items)

	cache.Replace("p1", "s1")
	cache.Replace("p2", "s2")
	cache.Replace("xxxx", "xxx")
	t.Log(cache.items)

	if v, foud := cache.Search("p1"); foud {
		t.Log(v)
	}
	if v, foud := cache.Search("p2"); foud {
		t.Log(v)
	}
	if v, foud := cache.Search("kkk"); foud {
		t.Log(v)
	}
	if v, foud := cache.Search("xxxx"); foud {
		t.Log(v)
	}
}
