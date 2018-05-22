 package main
 import "fmt"
 const N=13
 var col [N]int
 var d1,d2 [N*2]int
 var ret int
 func dfs(x int){
     if x==N{
         ret++
     }else{
         for i:=0;i<N;i++ {
             if col[i]==0 && d1[i+x]==0 && d2[i-x+N]==0 {
                 col[i],d1[i+x],d2[i-x+N]=1,1,1
                 dfs(x+1)
                 col[i],d1[i+x],d2[i-x+N]=0,0,0
             }
         }
     }
 }
 func main(){
     dfs(0)
     fmt.Println(ret)
 }
