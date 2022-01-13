# Set

[![Home](https://godoc.org/github.com/gookit/event?status.svg)](file:///D:/EC-TSJ/Documents/CODE/SOURCE/Go/pkg/lib/cli)
[![Build Status](https://travis-ci.org/gookit/event.svg?branch=master)](https://travis-ci.org/)
[![Coverage Status](https://coveralls.io/repos/github/gookit/event/badge.svg?branch=master)](https://coveralls.io/github/)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/event)](https://goreportcard.com/report/github.com/)

> **[EN README](README.md)**

Set es una librería para manipular la DataStructure set.

## GoDoc

- [godoc for github](https://godoc.org/github.com/)

## Funciones Principales
--- 
Tiene los objetos siguientes:

- Type `struct` ***Element***, para realizar las definiciones.

- type `[]Element` ***Set***  con métodos:
  - *`ToSet() ISet`*
  - *`String() string`*

- Type `ìnterface` ***ISet***, con métodos:
	- *`Add(i Item) bool`*
	- *`Cardinality() int`*
	- *`Clear()`*
	- *`Clone() Set`*
	- *`Contains(i ...Item) bool`*
	- *`Difference(other Set) Set`*
	- *`Equal(other Set) bool`*
	- *`Intersect(other Set) Set`*
	- *`IsProperSubset(other Set) bool`*
	- *`IsProperSuperset(other Set) bool`*
	- *`IsSubset(other Set) bool`*
	- *`IsSuperset(other Set) bool`*
	- *`Each(func(Item) bool)`*
	- *`Iter() <-chan interface{}`*
	- *`Iterator() *Iterator`*
	- *`Remove(i Item)`*
	- *`String() string`*
	- *`SymmetricDifference(other Set) Set`*
	- *`Union(other Set) Set`*
	- *`Pop() Item`*
	- *`PowerSet() Set`*
	- *`CartesianProduct(other Set) Set`*
	- *`ToSlice() []Item`*
	-	*`MarshalJSON() ([]byte, error)`*
	-	*`UnmarshalJSON(b []byte) error`*

- Type `struct` ***Iterator***, con métodos:
  - *`Stop()`*
- Type `map[interface{}]struct{}` ***uSet***, con los métodos de ISet.
- Type  `struct` ***OrderedPair***, con los métodos:
	- *`Equal(other OrderedPair) bool`* 

-Funciones:
  - *`NewSet(...Item) ISet`*



## Ejemplos
```go

	requiredClasses := NewSet()
	requiredClasses.Add("Cooking")
	requiredClasses.Add("English")
	requiredClasses.Add("Math")
	requiredClasses.Add("Biology")

	intClasses := NewSet(215, 1285, 7695, 3455)
	scienceClasses := NewSet("Biology", "Chemistry")
	electiveClasses := NewSet()
	electiveClasses.Add("Welding")
	electiveClasses.Add("Music")
	electiveClasses.Add("Automotive")

	bonusClasses := NewSet()
	bonusClasses.Add("Go Programming")
	bonusClasses.Add("Python Programming")

	//Show me all the available classes I can take
	allClasses := requiredClasses.Union(scienceClasses).Union(electiveClasses).Union(bonusClasses).Union(intClasses)
	fmt.Println(allClasses) //Set{Cooking, English, Math, Chemistry, Welding, Biology, Music, Automotive, Go Programming, Python Programming}

	//Is cooking considered a science class?
	fmt.Println(scienceClasses.Contains("Cooking")) //false

	//Show me all classes that are not science classes, since I hate science.
	fmt.Println(allClasses.Difference(scienceClasses)) //Set{Music, Automotive, Go Programming, Python Programming, Cooking, English, Math, Welding}

	//Which science classes are also required classes?
	fmt.Println(scienceClasses.Intersect(requiredClasses)) //Set{Biology}

	//How many bonus classes do you offer?
	fmt.Println(bonusClasses.Cardinality()) //2
	fmt.Println(intClasses.Cardinality())   //2

	//Do you have the following classes? Welding, Automotive and English?
	fmt.Println(allClasses.IsSuperset(set.NewSet("Welding", "Automotive", "English"))) //true

```
## Notas





<!-- - [gookit/ini](https://github.com/gookit/ini) INI配置读取管理，支持多文件加载，数据覆盖合并, 解析ENV变量, 解析变量引用
-->
## LICENSE

**[MIT](LICENSE)**
