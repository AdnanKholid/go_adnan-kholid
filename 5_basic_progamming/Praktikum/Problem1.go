package main

import ("fmt")

func main(){
    const PI = 3.14;

var jrjr float32
var tinggi float32
    fmt.Printf("Input jrjr : " ) 
  
    fmt.Scanf("%f", &jrjr)
    fmt.Printf("Input tinggi : ")  
    fmt.Scanf("\n %f", &tinggi)

var Problem1 float32  
    Problem1 = (2 * PI * jrjr)*(jrjr * tinggi);
    
fmt.Println("Problem1 :",Problem1)

}