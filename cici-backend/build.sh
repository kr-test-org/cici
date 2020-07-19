#<------------------------------------------------------------------------------->
## 当前可用环境变量如下，你可在构建脚本里进行引用使用
## WORKSPACE  工作目录
## IMAGE      输出镜像名称
## SERVICE    构建的服务名称
#<------------------------------------------------------------------------------->
#!/bin/bash
set -ex

cd $WORKSPACE/cici

make build-backend

tarStamp=`date +%Y%m%d%H%M%S`

PKG_NAME=cici.$tarStamp.tar.gz

cd $WORKSPACE/cici/cici-backend

tar cvf $PKG_NAME cici

echo '#!/usr/bin/expect
set timeout 10
set host [lindex $argv 0]
set username [lindex $argv 1]
set password [lindex $argv 2]
set src_file [lindex $argv 3]
set dest_file [lindex $argv 4]

spawn scp $src_file $username@$host:$dest_file
    expect {
    "(yes/no)?"
    {
    send "yes\n"
    expect "*assword:" { send "$password\n"}
    }
    "*assword:"
    {
    send "$password\n"
    }
}
expect "100%"
expect eof'  > expect_scp
    
chmod +x expect_scp

if [ "$pwd" = "" ] 
then
    exit 1
fi

./expect_scp 10.0.2.17 ubuntu $pwd $WORKSPACE/cici/cici-backend/$PKG_NAME  "~/hhbs/"

#restart service
export SSHPASS=$pwd
sshpass -e ssh ubuntu@10.0.2.17 'cd ~/hhbs && ./restart.sh '$PKG_NAME''