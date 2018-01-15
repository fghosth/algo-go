# 使用
```
package main
import (
"github.com/jvole/hashring"
"fmt"
)

func main (){

  nodeWeight := make(map[string]int)
    for i := 0; i < 10; i++ {
		si := fmt.Sprintf("%d", i)
    nodeWeight["211.104.10."+si] = 2
	}
    vitualSpots := 0
    hash := hashring.NewHashRing(vitualSpots)


    //add nodes
    hash.AddNodes(nodeWeight)

    // for _, v := range hash.GetNodes(){
    // 		fmt.Println(" nodeInfo:", v)
    // 	}
    //remove node
    // hash.RemoveNode("node3")


    //add node
    // hash.AddNode("node3", 3)


    //get key's node

    ipMap := make(map[string]int, 0)
    node0 :=make(map[string]string,0)//第二个节点所有数据
    	for i := 0; i < 1000; i++ {
    		si := fmt.Sprintf("key.%d", i)
    		k := hash.GetNode(si)
        if k=="211.104.10.2" {
    		    sii := fmt.Sprintf("%d", i)
            node0["211.104.10.2-"+sii] = si
        }
    		if _, ok := ipMap[k]; ok {
    			ipMap[k] += 1
    		} else {
    			ipMap[k] = 1
    		}
    	}

       for k, v := range ipMap {
       		fmt.Println("Node IP:", k, " count:", v)
       	}
      //for k,v := range node0{
        //fmt.Println(k+"===="+v)
      //}

}
```
