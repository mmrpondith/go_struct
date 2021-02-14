//struct 
//a struct is a collection of fields or properties.
// user defined data type
package main 
import "fmt"
type Book struct{
Title string
Author string 
Price float32
Pages int
}
func main(){
var b1 Book 
b1.Title="English for today"
b1.Author="Thinod Wjdin"
b1.Price=357.25
b1.Pages=3254
fmt.Println(b1)
fmt.Println(b1.Title)
fmt.Println(b1.Author)
fmt.Println(b1.Price)
fmt.Println(b1.Pages)
}