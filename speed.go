package main

import(
    "fmt"
    "time"
)

func getZeroesVar(ary []int) int{
    var counter int
    for _,i:=range ary{
        if i==0{
            counter+=1
        }
    }
    return counter
}

func getZeroesVarReturn(ary []int) *int{
    var counter int
    for _,i:=range ary{
        if i==0{
            counter+=1
        }
    }
    return &counter
}

func getZeroesPointer(ary *[]int) int{
    var counter int
    for _,i:=range *ary{
        if i==0{
            counter+=1
        }
    }
    return counter
}

func getZeroesPointerReturn(ary *[]int) *int{
    var counter int
    for _,i:=range *ary{
        if i==0{
            counter+=1
        }
    }
    return &counter
}

func main(){
    var ary=[]int{0,1,0,0,0,0,1}
    // passing and returning pointer
    now:=time.Now()
    fmt.Println("passing and returning pointer - ",*getZeroesPointerReturn(&ary))
    fmt.Println(time.Since(now))
    // passing array
    now=time.Now()
    fmt.Println("passing array - ",getZeroesVar(ary))
    fmt.Println(time.Since(now))
    // passing array return pointer
    now=time.Now()
    fmt.Println("passing array return pointer - ",*getZeroesVarReturn(ary))
    fmt.Println(time.Since(now))
    // passing pointer to array
    now=time.Now()
    fmt.Println("passing pointer array - ",getZeroesPointer(&ary))
    fmt.Println(time.Since(now))
    // passing and returning pointer
    now=time.Now()
    fmt.Println("passing and returning pointer - ",*getZeroesPointerReturn(&ary))
    fmt.Println(time.Since(now))
    
}