package main
import( "os" 
      "fmt"
	"log")

func main(){
	fmt.Println("Hello world")
	portString:=os.Getenv("PORT")
	if portString==""{
		log.Fatal("PORT is not found in env")
	}
	fmt.Println("Port:",portString)
}