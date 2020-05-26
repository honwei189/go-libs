System environment

choco install copy
choco install move
or
http://gnuwin32.sourceforge.net/packages/coreutils.htm
https://github.com/PowerShell/Win32-OpenSSH/releases

$GOPATH


Windows (Home, Global):

D:\Windows\go;Y:\prog\go


Windows (Home, User):

%USERPROFILE%\go;Y:\prog\go



Windows (Office, User):

%USERPROFILE%\go;Y:\prog\go


Linux:

vi /etc/profile  or  /root/.profile

GOPATH="/root/go"
export GOPATH="$GOPATH:/prog/go"

After save and quit file, enter "source /etc/profile"


#export GOROOT=/usr/local/go  #設置為go安裝的路徑，有些安裝包會自動設置默認的goroot
#export GOPATH=$HOME/gocode   #默認安裝包的路徑
#export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
#source /etc/profile