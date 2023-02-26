package main
 
import (
    "encoding/json"
    "fmt"
    "syscall/js"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func goCallback(this js.Value, args []js.Value) any {
    fmt.Println("Hello ", args[0])
    return args[0]
}
 
func main() {
    fmt.Println("Hello, World! (from go)")
    js.Global().Call("updateDOM", "Hello, World")
    js.Global().Set("aBoolean", true)
    js.Global().Set("aString", "hello world")
    js.Global().Set("aNumber", 1)
 
    // Work around for passing structs to JS
    frank := &Person{Name: "Frank", Age: 28}
    jsGlobalSetJSON("anObject", frank)
    c := make(chan struct{}, 0)
    js.Global().Set("aFunction", js.FuncOf(goCallback))
    <-c
}

func jsGlobalSetJSON(s string, person *Person){
    p, err := json.Marshal(person)
    if err != nil {
        fmt.Println(err)
        return
    }
    obj := js.Global().Get("JSON").Call("parse", string(p))
    js.Global().Set(s, obj)
}

