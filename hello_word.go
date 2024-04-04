// go mod init hello_word
// go get -u github.com/go-sql-driver/mysql
// go mod tidy --untuk menyesuaikan dependence
// go run / go build hello_word.go

package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// object 2
// type Person struct {
// 	Name string
// 	Age  int
// }

func main() {
	fmt.Println("hello word")

	// hasil, genap := bagiDua(7)
	// fmt.Println("Hasil pembagian:", hasil)
	// fmt.Println("Genap:", genap)

	// hasil, v_a, v_b := perkalian(10, 5)
	// fmt.Println("hasil perkalian :", hasil)
	// fmt.Println("variabel a :", v_a)
	// fmt.Println("variabel b :", v_b)

	// array
	// a := []int{1, 2, 3, 4}
	// fmt.Println("variabel b :", a[1])
	// for key, _ := range a {
	// 	fmt.Println(key)
	// }
	// for _, value := range a {
	// 	fmt.Println(value)
	// }

	// slice
	// a := map[int]int{
	// 	1: 1,
	// }
	// a := map[string]int{
	// 	"saya": 1,
	// }
	// fmt.Println("variabel b :", a["saya"])

	// object 1
	// var people []interface{}
	// people = append(people, map[string]interface{}{"name": "John", "age": 30})
	// people = append(people, map[string]interface{}{"name": "Alice", "age": 25})
	// people = append(people, map[string]interface{}{"name": "Bob", "age": 35})

	// for _, person := range people {
	// 	fmt.Println("Name:", person.(map[string]interface{})["name"])
	// 	fmt.Println("Age:", person.(map[string]interface{})["age"])
	// 	fmt.Println("---")
	// }

	// object 2
	// people := []Person{
	// 	{Name: "John", Age: 30},
	// 	{Name: "Alice", Age: 25},
	// 	{Name: "Bob", Age: 35},
	// }

	// for _, person := range people {
	// 	fmt.Println("Name:", person.Name)
	// 	fmt.Println("Age:", person.Age)
	// 	fmt.Println("---")
	// }

	// for i := 1; i < 5; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(1)
	// 	} else {
	// 		fmt.Println(0)
	// 	}
	// }

	// database
	// user := "root"
	// password := ""
	// hostname := "localhost"
	// port := "3306"
	// dbname := "coca_cola"

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, hostname, port, dbname)

	// db, err := sql.Open("mysql", dsn)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer db.Close() //tutup koneksi

	// rows, err := db.Query("SELECT userid,nama FROM mobile_agent")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var userid int
	// 	var nama string
	// 	err := rows.Scan(&userid, &nama)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	fmt.Println("ID:", userid, "Nama:", nama)
	// }

}

func bagiDua(d int) (int, bool) {
	hasil := d / 2
	sisa := d % 2
	return hasil, sisa == 0
}

func perkalian(a int, b int) (int, int, int) {
	return a * b, a, b
}
