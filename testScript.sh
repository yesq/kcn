set -ux

export KCN_CHECKFILE=`pwd`/configmap.txt
export KCN_SCRIPT=`pwd`/sh.sh
echo "12345 KCN_CHECKFILE" > configmap.txt
echo "echo 'hello KCN_SCRIPT' && pwd" > sh.sh
./kcn
echo $?

echo "abcde KCN_CHECKFILE" > configmap.txt
./kcn
echo $?

echo "echo 'hello KCN_SCRIPT' && false" > sh.sh
rm -f data.json
./kcn
echo $?