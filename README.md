# kcn
A tool to check service whether need to be killed

### TODO

1. check a file sha. ✅
1. run a script from user ✅
1. check multiple files.
1. run multiple script
1. check a folder

## Usage

```shell
kcn:
  a toole to check service
  need two ENV :
    KCN_CHECKFILE : a file. if file changed, return false.
    KCN_SCRIPT    : a script path. run script.
```

```shell
export KCN_CHECKFILE=`pwd`/configmap.txt
export KCN_SCRIPT=`pwd`/sh.sh
echo "12345 KCN_CHECKFILE" > configmap.txt
echo "echo 'hello KCN_SCRIPT' && pwd" > sh.sh
./kcn
echo $?
```