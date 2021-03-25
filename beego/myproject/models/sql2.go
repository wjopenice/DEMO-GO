package models
import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type User struct{
	Id int
	Name string `orm:"size(100)"`
}

type Post struct {
	Id      int
	Title   string
	Content string
	Status  int
}

func main(){

	//ORM 以 QuerySeter 来组织查询，每个返回 QuerySeter 的方法都会获得一个新的 QuerySeter 对象
	o := orm.NewOrm()
	// 获取 QuerySeter 对象，user 为表名
	qs1 := o.QueryTable("user")
	// 也可以直接使用 Model 结构体作为表名
	qs2 := o.QueryTable(&User)
	// 也可以直接使用对象作为表名
	user := new(User)
	qs3 := o.QueryTable(user) // 返回 QuerySeter


    //QuerySeter 中用于描述字段和 sql 操作符，使用简单的 expr 查询方法
	qs.Filter("id", 1) // WHERE id = 1
	qs.Filter("profile__age", 18) // WHERE profile.age = 18
	qs.Filter("Profile__Age", 18) // 使用字段名和 Field 名都是允许的
	qs.Filter("profile__age__gt", 18) // WHERE profile.age > 18
	qs.Filter("profile__age__gte", 18) // WHERE profile.age >= 18
	qs.Filter("profile__age__in", 18, 20) // WHERE profile.age IN (18, 20)

	qs.Filter("profile__age__in", 18, 20).Exclude("profile__lt", 1000)
	// WHERE profile.age IN (18, 20) AND NOT profile_id < 1000

	//Operators
	//当前支持的操作符号：
	//
	//exact / iexact 等于
	//contains / icontains 包含
	//gt / gte 大于 / 大于等于
	//lt / lte 小于 / 小于等于
	//startswith / istartswith 以…起始
	//endswith / iendswith 以…结束
	//in
	//isnull
	//后面以 i 开头的表示：大小写不敏感

	//exact
	qs.Filter("name", "slene") // WHERE name = 'slene'
	qs.Filter("name__exact", "slene") // WHERE name = 'slene'
	// 使用 = 匹配，大小写是否敏感取决于数据表使用的 collation
	qs.Filter("profile_id", nil) // WHERE profile_id IS NULL

    //iexact
	qs.Filter("name__iexact", "slene")
	// WHERE name LIKE 'slene'
	// 大小写不敏感，匹配任意 'Slene' 'sLENE'

    //contains
	qs.Filter("name__contains", "slene")
	// WHERE name LIKE BINARY '%slene%'
	// 大小写敏感, 匹配包含 slene 的字符

	//icontains
	qs.Filter("name__icontains", "slene")
	// WHERE name LIKE '%slene%'
	// 大小写不敏感, 匹配任意 'im Slene', 'im sLENE'

    //in
	qs.Filter("age__in", 17, 18, 19, 20)
	// WHERE age IN (17, 18, 19, 20)
	ids:=[]int{17,18,19,20}
	qs.Filter("age__in", ids)
	// WHERE age IN (17, 18, 19, 20)
	// 同上效果

    //gt / gte
	qs.Filter("profile__age__gt", 17)
	// WHERE profile.age > 17
	qs.Filter("profile__age__gte", 18)
	// WHERE profile.age >= 18

	//lt / lte
	qs.Filter("profile__age__lt", 17)
	// WHERE profile.age < 17
	qs.Filter("profile__age__lte", 18)
	// WHERE profile.age <= 18

	//startswith
	qs.Filter("name__startswith", "slene")
	// WHERE name LIKE BINARY 'slene%'
	// 大小写敏感, 匹配以 'slene' 起始的字符串

	//istartswith
	qs.Filter("name__istartswith", "slene")
	// WHERE name LIKE 'slene%'
	// 大小写不敏感, 匹配任意以 'slene', 'Slene' 起始的字符串

	//endswith
	qs.Filter("name__endswith", "slene")
	// WHERE name LIKE BINARY '%slene'
	// 大小写敏感, 匹配以 'slene' 结束的字符串

	//iendswith
	qs.Filter("name__iendswithi", "slene")
	// WHERE name LIKE '%slene'
	// 大小写不敏感, 匹配任意以 'slene', 'Slene' 结束的字符串

	//isnull
	qs.Filter("profile__isnull", true)
	qs.Filter("profile_id__isnull", true)
	// WHERE profile_id IS NULL
	qs.Filter("profile__isnull", false)
	// WHERE profile_id IS NOT NULL

	//Filter
	qs.Filter("profile__isnull", true).Filter("name", "slene")
	// WHERE profile_id IS NULL AND name = 'slene'

	//Exclude
	qs.Exclude("profile__isnull", true).Filter("name", "slene")
	// WHERE NOT profile_id IS NULL AND name = 'slene'

	//SetCond
	cond := orm.NewCondition()
	cond1 := cond.And("profile__isnull", false).AndNot("status__in", 1).Or("profile__age__gt", 2000)
	qs := orm.QueryTable("user")
	qs = qs.SetCond(cond1)
	// WHERE ... AND ... AND NOT ... OR ...
	cond2 := cond.AndCond(cond1).OrCond(cond.And("name", "slene"))
	qs = qs.SetCond(cond2).Count()
	// WHERE (... AND ... AND NOT ... OR ...) OR ( ... )

	//Limit
	var DefaultRowsLimit = 1000 // ORM 默认的 limit 值为 1000
	// 默认情况下 select 查询的最大行数为 1000
	// LIMIT 1000
	qs.Limit(10)
	// LIMIT 10
	qs.Limit(10, 20)
	// LIMIT 10 OFFSET 20 注意跟 SQL 反过来的
	qs.Limit(-1)
	// no limit
	qs.Limit(-1, 100)
	// LIMIT 18446744073709551615 OFFSET 100
	// 18446744073709551615 是 1<<64 - 1 用来指定无 limit 限制 但有 offset 偏移的情况

	//Offset
	qs.Offset(20)
	// LIMIT 1000 OFFSET 20

	//GroupBy
	qs.GroupBy("id", "age")
	// GROUP BY id,age
	qs.OrderBy("id", "-profile__age")
	// ORDER BY id ASC, profile.age DESC
	qs.OrderBy("-profile__age", "profile")
	// ORDER BY profile.age DESC, profile_id ASC

	//ForceIndex
	qs.ForceIndex(`idx_name1`,`idx_name2`)

	//UseIndex
	qs.UseIndex(`idx_name1`,`idx_name2`)

	//IgnoreIndex
	qs.IgnoreIndex(`idx_name1`,`idx_name2`)

	//Distinct
	qs.Distinct()
	// SELECT DISTINCT

	//RelatedSel
	var DefaultRelsDepth = 5 // 默认情况下直接调用 RelatedSel 将进行最大 5 层的关系查询
	qs := o.QueryTable("post")
	qs.RelatedSel()
	// INNER JOIN user ... LEFT OUTER JOIN profile ...
	qs.RelatedSel("user")
	// INNER JOIN user ...
	// 设置 expr 只对设置的字段进行关系查询
	// 对设置 null 属性的 Field 将使用 LEFT OUTER JOIN

	//Count
	cnt, err := o.QueryTable("user").Count() // SELECT COUNT(*) FROM USER
	fmt.Printf("Count Num: %s, %s", cnt, err)

	//Exist
	exist := o.QueryTable("user").Filter("UserName", "Name").Exist()
	fmt.Printf("Is Exist: %s", exist)

	//Update
	num, err := o.QueryTable("user").Filter("name", "slene").Update(orm.Params{
		"name": "astaxie",
	})
	fmt.Printf("Affected Num: %s, %s", num, err)
	// SET name = "astaixe" WHERE name = "slene"
	// 假设 user struct 里有一个 nums int 字段
	num, err := o.QueryTable("user").Update(orm.Params{
		"nums": orm.ColValue(orm.ColAdd, 100),
	})
	// SET nums = nums + 100
	//ColAdd      // 加
	//ColMinus    // 减
	//ColMultiply // 乘
	//ColExcept   // 除

	//Delete
	num, err := o.QueryTable("user").Filter("name", "slene").Delete()
	fmt.Printf("Affected Num: %s, %s", num, err)
	// DELETE FROM user WHERE name = "slene"

	//PrepareInsert
	var users []*User
	...
	qs := o.QueryTable("user")
	i, _ := qs.PrepareInsert()
	for _, user := range users {
		id, err := i.Insert(user)
		if err == nil {
			...
		}
	}
	// PREPARE INSERT INTO user (`name`, ...) VALUES (?, ...)
	// EXECUTE INSERT INTO user (`name`, ...) VALUES ("slene", ...)
	// EXECUTE ...
	// ...
	i.Close() // 别忘记关闭 statement

	//All
	var users []*User
	num, err := o.QueryTable("user").Filter("name", "slene").All(&users)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)
	// 只返回 Id 和 Title
	var posts []Post
	o.QueryTable("post").Filter("Status", 1).All(&posts, "Id", "Title")

	//One
	var user User
	err := o.QueryTable("user").Filter("name", "slene").One(&user)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		fmt.Printf("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		fmt.Printf("Not row found")
	}
	// 只返回 Id 和 Title
	var post Post
	o.QueryTable("post").Filter("Content__istartswith", "prefix string").One(&post, "Id", "Title")

	//Values
	var maps []orm.Params
	num, err := o.QueryTable("user").Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		for _, m := range maps {
			fmt.Println(m["Id"], m["Name"])
		}
	}

	//ValuesList
	var lists []orm.ParamsList
	num, err := o.QueryTable("user").ValuesList(&lists)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		for _, row := range lists {
			fmt.Println(row)
		}
	}
	//当然也可以指定 expr 返回指定的 Field
	num, err := o.QueryTable("user").ValuesList(&lists, "name", "profile__age")

	//ValuesFlat
	num, err := o.QueryTable("user").ValuesFlat(&list, "name")
}

