package main 
import "fmt"
func main(){
b1:= struct{
Title string
Author string 
Price float32
Pages int
}{
	Title: "English for today",
	Author: "Thinod Wjdin",
	Price: 357.25,
	Pages: 3254,
}
fmt.Println(b1)
fmt.Println(b1.Title)
fmt.Println(b1.Author)
fmt.Println(b1.Price)
fmt.Println(b1.Pages)
}