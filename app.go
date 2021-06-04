package main
import (
  "fmt"
  "math/rand"
  "time"
  "sort"
  "bufio"
  "os"
  "strings"
  "strconv"
)
type Product struct {
  Name     string
  Category string
  Price    int
}
func randomInt(min int, max int) int {
  return rand.Intn(max-min+1) + min
}
func main() {
  //create 20 object
  fmt.Println("---------------create 20 object---------------")
  category := [4]string{"fashion", "electronics", "sport", "food"}
  products := [20]Product{}
  rand.Seed(time.Now().UnixNano())
  for i := 0; i < len(products); i++ {
    products[i] = Product{
      fmt.Sprintf("%s%d", "Products",i),
      category[rand.Intn(len(category))],
      randomInt(100, 200),
    }
  }
  for _, product := range products {
    fmt.Println(product)
  }

//Sort Price
sortedProducts := products[:]
sort.Slice(sortedProducts, func(i, j int) bool {
  return sortedProducts[i].Price > sortedProducts[j].Price
})

//find top 5 Highest Price
fmt.Println("------------Top 5 Highest Price--------------")
for i := 0; i < 5; i++ {
  fmt.Println(sortedProducts[i])
}
  
//Count product by Category
fmt.Println("---------------Count product by Category-------------")
  categoryCount := map[string]int{}
  for _, product := range products {
    if _, ok := categoryCount[product.Category]; ok {
      categoryCount[product.Category]++
    } else {
      categoryCount[product.Category] = 0
    }
  }
  for key, value := range categoryCount {
    fmt.Println(key, value)
  }

  //Search by Name
fmt.Println("----------------Search by Name------------------")

  readerName := bufio.NewReader(os.Stdin)
    fmt.Print("Enter Product name: ")
    text, _, _ := readerName.ReadLine()
    s := string(text)
    for _, product := range products {
        if product.Name == s {
            fmt.Println(product)
        }
    }
    //Search by Category
  fmt.Println("--------------Search by Category-------------")

  readerCategory := bufio.NewReader(os.Stdin)
  fmt.Print("Enter Category name: ")
  texts, _, _ := readerCategory.ReadLine()
  b := string(texts)
  for _, product := range products {
      if product.Category == b {
          fmt.Println(product)
      }
  }


  //Search beetween Price
  fmt.Println("--------------Search by Price-------------")
  //get min
  readerMin := bufio.NewReader(os.Stdin)
  fmt.Print("Enter min number: ")
  min, _ := readerMin.ReadString('\n')
  min = strings.Replace(min,"\n","",-1)

  // convert string to int
  minPrice, e := strconv.Atoi(min)
  if e != nil {
fmt.Println(min,"is not number")
  }

  //get max
  readerMax := bufio.NewReader(os.Stdin)

  fmt.Print("Enter max number: ")
  max, _ := readerMax.ReadString('\n')
  // remove newline
  max = strings.Replace(max,"\n","",-1)

  // convert string variable to int variable
  maxPrice, e := strconv.Atoi(max)
  if e != nil {
fmt.Println(max,"is not number")
  }

  for _, product := range products {

      if product.Price < maxPrice && product.Price > minPrice {
          fmt.Println(product)
      }
  }

}
