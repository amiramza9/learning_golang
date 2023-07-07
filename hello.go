package main
import(
    "fmt"
    //"os"
    "reflect"
    "strconv"
    //"bufio"
    //"log"
    
)

func main(){
    fmt.Println(reflect.TypeOf('6'))
    /*fmt.Println("Enter Your name! ")
    reader := bufio.NewReader(os.Stdin)
    name,err := reader.ReadString('\n');
    if err==nil {
        fmt.Println("Hello ",name)
} else {
    log.Fatal(err)
}*/
//knowing the types
       /* fmt.Println(reflect.TypeOf(true))
        fmt.Println(reflect.TypeOf(nil))
        fmt.Println(reflect.TypeOf(9090099999))
        fmt.Println(reflect.TypeOf(9.90))
        fmt.Println(reflect.TypeOf(0xaced))
        fmt.Println(reflect.TypeOf(' '))
        fmt.Println(reflect.TypeOf("tyr"))
        */
    v1 :=9.8
    fmt.Println(int(v1),reflect.TypeOf(int(v1)))
    v2 :="490080"
    v3,err :=strconv.Atoi(v2) //returns two values on of which is errs
    fmt.Println(v3,err,reflect.TypeOf(v3))
    v4 :=strconv.Itoa(v3) //returns single value
    fmt.Println(v4,reflect.TypeOf(v4))

    // or
    v5 :="908.7"
    if v6,err := strconv.ParseFloat(v5,64); err==nil { // v6 only accessible in this if block
        fmt.Println(v6)
    }
   // fmt.Println(v6) // inaccessibility/undefined error

    }
    
