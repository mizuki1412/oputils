# oputils

运维工具

# mod_waster

消耗cpu、内存、存储。

原理上就是简单的对一块大数组进行反复的赋值，达到消耗cpu和内存的目的。

run `./waster -h` for help

```shell
## 占用20%的总内存，占用一核cpu
./waster --mem=20

## 占用2核
./waster --core=2

## 生成20GB的文件 /demo/test.bin， 生成完后，将执行内存和cpu占用
./waster --mem=2 --spath=/demo/test.bin --ssize=20
```