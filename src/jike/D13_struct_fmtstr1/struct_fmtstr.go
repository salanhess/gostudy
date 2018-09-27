package main

//(Need wechat Login)https://time.geekbang.org/column/article/18035

import (
	"fmt"
	"log"
)

type AnimalCategory struct {
	phylum  string // 门。
	class   string // 纲。
	order   string // 目。
	family  string // 科。
	genus   string // 属。
	species string // 种。
}
type Animal struct {
	scientificName string // 学名。
	AnimalCategory        // 动物基本分类。
}

func (ac AnimalCategory) String() string {
	return fmt.Sprintf("[AnimalCategory]####rewrite String methd####,AnimalCategory is %s %s %s\n", ac.class, ac.order, ac.family)
}

func (a Animal) String() string {
	return fmt.Sprintf("[Animal]***rewrite String methd####,scientificName is %s AnimalCategory is %s %s %s\n", a.scientificName, a.class, a.order, a.family)
}

func (a Animal) Category() string {
	return a.AnimalCategory.String()
}

/*
值方法的接收者是该方法所属的那个类型值的一个副本。我们在该方法内对该副本
的修改一般都不会体现在原值上，除非这个类型本身是某个引用类型（比如切片或
字典）的别名类型。
*/
func (a Animal) ModifyCategoryCopy() {
	a.AnimalCategory.class = "ModifyCategoryCopy"
}

/*
而指针方法的接收者，是该方法所属的那个基本类型值的指针值的一个副本。我们
在这样的方法内对该副本指向的值进行修改，却一定会体现在原值上。
*/

func (a *Animal) ModifyCategoryRef() {
	a.AnimalCategory.class = "ModifyCategoryCopy"
}

func main() {
	ac := AnimalCategory{class: "Niao", order: "qiake", family: "feifei"}
	fmt.Printf("%s\n", ac)
	/*
	   嵌入字段AnimalCategory的String方法被“屏蔽”了。
	   注意，只要名称相同，无论这两个方法的签名是否一致，被嵌入类型的方法都会“屏蔽”掉嵌入字段的同名方法
	*/
	a := Animal{scientificName: "craolbigsb", AnimalCategory: ac}
	fmt.Printf("%s\n", a)
	fmt.Printf("Animal Category method is %s", a.Category())
	a.ModifyCategoryCopy()
	fmt.Printf("[AfterModifyCopy]Animal Category method is %s", a.Category())
	a.ModifyCategoryRef()
	fmt.Printf("[AfterModifyRef]Animal Category method is %s", a.Category())
	// test code
	ax := struct{}{}
	bx := struct{}{}
	log.Println(ax == bx)            // true
	log.Printf("%p, %p\n", &ax, &bx) // 0x7bb7f8, 0x7bb7f8
}

/* OutPut:
[root@izhp3c4nnfdk1zyhek9f3ez D13_struct_fmtstr1]# go run struct_fmtstr.go
[AnimalCategory]####rewrite String methd####,AnimalCategory is Niao qiake feifei

[Animal]***rewrite String methd####,scientificName is craolbigsb AnimalCategory is Niao qiake feifei

Animal Category method is [AnimalCategory]####rewrite String methd####,AnimalCategory is Niao qiake feifei
[AfterModifyCopy]Animal Category method is [AnimalCategory]####rewrite String methd####,AnimalCategory is Niao qiake feifei
[AfterModifyRef]Animal Category method is [AnimalCategory]####rewrite String methd####,AnimalCategory is ModifyCategoryCopy qiake feifei
2018/09/27 23:41:13 true
2018/09/27 23:41:13 0x5571b8, 0x5571b8
*/
