package test

import (
	"bufio"
	"fmt"
	"github.com/sonhineboy/gsadmin/service/app"
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/repositorys"
	"github.com/sonhineboy/gsadmin/service/app/requests"
	"github.com/sonhineboy/gsadmin/service/global"
	"github.com/sonhineboy/gsadminGen"
	"github.com/sonhineboy/gsadminGen/pkg"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestStringTo(t *testing.T) {

	b := sync.WaitGroup{}
	a := "xxx"
	b.Add(1)
	go func() {
		fmt.Println(a)
		b.Done()
	}()

	b.Wait()
}

type b struct {
	Name string
}

func (receiver b) GetC() {

	fmt.Printf("%p", &receiver)
	receiver.Name = "xcc"
}

func (receiver b) GetName() string {
	fmt.Printf("%p", &receiver)
	return receiver.Name

}

func newB() *b {
	return &b{}
}

func TestP(t *testing.T) {

	role := models.Role{}

	var a *int
	var c *int
	var b int
	b = 3
	a = &b
	c = a
	fmt.Println(a, "---", c)

	//model.ID = 45
	model, _ := reflect.TypeOf(role).FieldByName("GsModel")

	f := model.Type.Field(0)

	fmt.Println(model.Type.NumField())
	fmt.Println(f.Tag.Get("gorm"), f.Tag.Get("json"))

	fmt.Println(reflect.TypeOf(&model))
}

func TestVersion(t *testing.T) {

	str := "he,llo, , world! Hello 你好"
	m := make(map[string]int)

	// s为临时字符串，用于拼接每个字母，组成单词
	s := ""

	//  temp用于统计每个单词中的字母的个数
	temp := 0

	// 遍历这个str，获取每个字符char
	for _, char := range str {
		// 当char是26个字母的大写或小写
		if (char >= 65 && char <= 90) || (char >= 97 && char <= 122) {
			temp++
			s += string(char)
		} else { // 如果遍历到的字符 不是26个字母
			if temp != 0 { // 此处判断用于防止将非26个字母的字符加入到map中
				m[s]++ // 将这个单词添加到map中
				s = "" // 将临时字符串s清空
			}
			temp = 0 // 每个单词中字母的个数清理
		}
	}

	// 防止当字符串str的最后一个字符为字母时，不会将最后一个单词添加进map
	if s != "" {
		m[s]++
	}

	for k, v := range m {
		fmt.Printf("%s:%d\n", k, v)
	}

}

// deleteMarkOfWord 删除单词末尾的标点符号
func deleteMarkOfWord(word string) string {
	// 从单词末端开始遍历，删除标点符号
	for i := len(word) - 1; i >= 0; i-- {
		if (word[i] >= 97 && word[i] <= 122) || (word[i] >= 65 && word[i] <= 90) {
			break
		} else {
			word = word[:len(word)-1]
		}
	}
	return word
}

func TestSocket(t *testing.T) {
	//print(test(3, 6))

	a := make([]int, 2, 5)
	fmt.Println(a[0:])
}

func test(b int, c int) (a int) {
	a = b + c
	return
}

func TestSlice(t *testing.T) {
	defer func() {

		fmt.Println("defer :-->1")
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	defer func() {
		fmt.Println("defer:--->3")
	}()

	car := make([]int, 0, 10)

	for i := 0; i <= 9; i++ {
		car = append(car, i)
	}

	//car = nil
	fmt.Println("长度--》", len(car))

}

func TestDemo(t *testing.T) {

	id, err := strconv.Atoi("asfd")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(id)
}

func TestGenModel(t *testing.T) {

	/**

	default:""
	describe:""
	index:"Null"
	isNull:true
	json:"content"
	name:"content"
	primary:false
	transform:""
	type:"longtext"

	*/

	fields := []pkg.Field{
		{
			Name:      "userName",
			Json:      "user_name",
			Default:   "",
			Describe:  "用户名",
			Primary:   false,
			Index:     "UNIQUE",
			IsNull:    false,
			Type:      "varchar",
			Transform: "用户名",
		},
		{
			Name:      "age",
			Json:      "age",
			Default:   "0",
			Describe:  "年龄",
			Primary:   false,
			Index:     "UNIQUE",
			IsNull:    true,
			Type:      "int",
			Transform: "年龄",
		},
	}

	v := pkg.TableModal{
		Name:   "user_member",
		Fields: fields,
	}

	var err error

	fileName := "user_member"

	err = gsadminGen.GenController("../app/controllers/demo/"+gsadminGen.UnderToConvertSoreLow(fileName)+"Controller.go", v, "demo")

	if err != nil {
		t.Fatal(err.Error())
	}

	err = gsadminGen.GenModel("../app/models/"+gsadminGen.UnderToConvertSoreLow(fileName)+".go", v)

	if err != nil {
		t.Fatal(err.Error())
	}

	err = gsadminGen.GenRequest("../app/requests/"+gsadminGen.UnderToConvertSoreLow(fileName)+"Request.go", v)

	if err != nil {
		t.Fatal(err.Error())
	}

	err = gsadminGen.GenRepository("../app/repositorys/"+gsadminGen.UnderToConvertSoreLow(fileName)+"Repository.go", v)

	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestHasSlice(t *testing.T) {
	//fmt.Println(dir)
}

func TestRead(t *testing.T) {

	// 打开文件
	file, err := os.Open("orm_test.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 创建一个新的 scanner
	scanner := bufio.NewScanner(file)

	i := 1

	// 按行读取文件
	for scanner.Scan() {
		line := scanner.Text()     // 获取当前行的内容
		fmt.Println(i, "  ", line) // 打印当前行
		i++
	}

	// 检查读取过程中是否出现错误
	if err := scanner.Err(); err != nil {
		fmt.Println(i, "  ", err)
	}

}

func TestZap(t *testing.T) {

	zapcore.NewTee()
}

func TestFor(t *testing.T) {

	ints := make(map[int]int)

	for i := 0; i < 3; i++ {
		ints[i] = i
	}

	for i2, i3 := range ints {
		c := i2
		d := i3
		go func() {
			fmt.Println(c, d)
		}()
	}

	time.Sleep(5 * time.Second)

}

func TestConfig(t *testing.T) {

	app.TestLoad()
	fmt.Println(global.Config.Db.TablePrefix)
}

func TestAddMenu(t *testing.T) {

	app.TestLoad()
	re := &repositorys.SystemMenuRepository{MenuModel: models.AdminMenu{}}
	// {"icon": "el-icon-menu", "type": "menu", "title": "生成模块"}
	db, _ := re.Add(requests.MenuPost{
		Component: "demo",
		Name:      "demo", //别名，权限判断唯一标识
		ParentId:  0,
		Path:      "/demo",
		Meta: map[string]interface{}{
			"icon":  "el-icon-menu",
			"type":  "menu",
			"title": "演示菜单",
		},
		ApiList: []map[string]string{
			{"code": "get", "url": "/api/demo/index"},
		},
	})

	if db.Error != nil {
		t.Errorf("db-->%v", db.Error)
	}

}

func TestDb(t *testing.T) {
	app.TestLoad()

	var db *gorm.DB

	fmt.Println(db == nil)

}

func TestReg(t *testing.T) {
	reg := regexp.MustCompile("^gc_")

	fmt.Println(reg.ReplaceAllString("gc_gcaa", ""))
}

func TestViper(t *testing.T) {

	fmt.Printf("xxx %s", "asfs")
}
func TestG2(t *testing.T) {
	c := make(chan int, 20)
	c <- 1
	close(c)
	cv, ok := <-c
	fmt.Println(cv, ok)
	cv, ok = <-c
	fmt.Println(cv, ok)

	t.Fatal("asdfasd")

}
func TestG(t *testing.T) {
	c := make(chan int, 20)
	test5(c)
	for {
		select {
		case c, ok := <-c:
			fmt.Println(c, ok)
		default:
			fmt.Println("--默认")
		}
		time.Sleep(time.Second)
	}

}

func test5(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

// test6 aabb
func test6(a int) {
	fmt.Println(a)
}

func TestReflect(t *testing.T) {
	type myTest struct {
		Name string
		Age  string
	}
	myUser := &myTest{}
	var reV reflect.Value
	if reflect.ValueOf(myUser).Kind() == reflect.Ptr {
		reV = reflect.ValueOf(myUser).Elem()
	} else {
		reV = reflect.ValueOf(myUser)
	}

	if reV.Kind() == reflect.Struct {
		fieldNum := reV.NumField()
		for i := 0; i < fieldNum; i++ {
			reV.Field(i).Set(reflect.ValueOf("1"))
		}
	}
	fmt.Println(reV)

	fmt.Println(reflect.ValueOf(test6).Type().In(0).Kind())

	c := map[string]string{
		"a": "b",
		"c": "d",
	}

	cMap := reflect.ValueOf(c)
	cMapKeys := cMap.MapKeys()
	for i, _ := range cMapKeys {
		fmt.Println(cMap.MapIndex(cMapKeys[i]))
	}

}

func TestChan(t *testing.T) {

	ch := make(chan int)

	str := "hello word!"
	go printStr(str, "A", ch)
	go printStr(str, "B", ch)
	ch <- 0
	time.Sleep(time.Second * 2)
}

func printStr(str, name string, ch chan int) {

loop:
	for {
		select {
		case index, ok := <-ch:
			if ok && index <= len(str)-1 {
				fmt.Println(name, "-->", string(str[index]))
				index++
				ch <- index
			} else {
				if ok {
					close(ch)
				}
				break loop
			}
		}
	}
	fmt.Println("\n 退出")
}

func IsClosed(ch <-chan int) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}
