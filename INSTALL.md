OS: ubuntu 20.04  

apt remove golang-go  
apt remove --auto-remove golang-go  
wget https://golang.org/dl/go1.20.3.linux-amd64.tar.gz  
tar -C /usr/local -xzf go1.20.3.linux-amd64.tar.gz  
export PATH=$PATH:/usr/local/go/bin  
go version  

git clone https://github.com/projectclone/epusdt.git  
cd ~/epusdt/src/  
CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w --extldflags "-static"' -o epusdt  
ls -alht epusdt  

