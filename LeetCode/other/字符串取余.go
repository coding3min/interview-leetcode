package main
import "fmt"
func main(){
    var str string
    for{
        n,_ := fmt.Scan(&str)
        if n == 0{
            break
        }else{
            splitStr(str,8)
        }
    }
}
func splitStr(str string,n int){
    zeroStr := "00000000"
    for len(str)>8{
        fmt.Println(str[:8])
        str = str[8:]
    }
    fmt.Println(str+zeroStr[:8-len(str)])
}