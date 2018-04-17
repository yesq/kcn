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

### Use with k8s

可以把二进制放在应用容器里。也可以在 pod 里另起一个容器。

```
    spec:
      containers:
      - name: nginx
        image: kcn:v0.0.1
        env:
+       - name: KCN_CHECKFILE
+         value: "/cfg"
        ports:
        - containerPort: 80
        command: ["/bin/sh"]
        #args: ["-c", "echo '123' > /cfg && sleep 10 && echo 'abc' > /cfg && sleep 100000"]
        args: ["-c", "echo '123' > /cfg && sleep 100000"]
+        livenessProbe:
+          exec:
+            command:
+            - kcn
+          initialDelaySeconds: 5
+          periodSeconds: 5
```