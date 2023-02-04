rm -rf Indra
git clone https://github.com/Vntvxyz/Indra
cd Indra
go mod init http 
go get github.com/valyala/fasthttp
go build
