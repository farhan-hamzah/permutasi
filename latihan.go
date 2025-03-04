package main
import "fmt"
func permutasi(n int, r int)int {
    var i, result, j, k, l int
    j = n - r
    k = 1
    l = 1
    i = n
    for i >= 1{
        l *= i
        i--
    }
    for j >= 1 {
        k *= j
        j--
    }
    result = l/k
    return result
}

func main(){
    var d1, d2 int
    fmt.Scan(&d1, &d2)
    permutasi(d1, d2)
    fmt.Print(permutasi(d1, d2))
}
