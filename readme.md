## alns求解TSP问题

### 1. 前置知识
- TSP问题
- 邻域算法
- go语言

### 2. 程序运行方法
#### 2.1 tsp文件与坐标系
在参考的文章中，两个算例涉及到两种坐标系，china34算例使用的是WGS坐标系，即经纬度坐标，xqf131使用的是欧式平面坐标

在运行程序时，使用 `-filepath`指定tsp数据源文件路径，如数据源当前文件夹的`test.tsp`文件则为 `-filepath ./test.tsp`，使用`-coor`指定坐标系名称，目前只支持经纬度坐标和欧式平面坐标，如坐标系为欧式平面坐标，则`-coor EUC`，若为经纬度坐标，则`-coor WGS`

> 注：tsp文件前6行为基础信息，默认从第七行读取坐标数据，第一个数据为id，第二个数据为x坐标，第三个数据为y坐标，使用空格分开，如下所示：
```bash
NAME: china34
TYPE: TSP
COMMENT: 34 locations in China
DIMENSION: 34
EDGE_WEIGHT_TYPE: WGS
NODE_COORD_SECTION
1 101.74 36.56
2 103.73 36.03
...
```

#### 2.2 运行程序
```bash
go run main.go -filepath ./benchmark/china34.tsp -coor WGS
```

### 3. 运行效果
1. 运行效果受实验参数影响
对于china34这个算例，迭代1000次仅需`**100ms**`左右   
对于xqf131这个算例，迭代1000次仅需`**2s**`左右

> 注：可以通过调整参数来改良算法的效果

参考：
[1] [着实不错的自适应大邻域搜索算法ALNS](https://mp.weixin.qq.com/s/5iyY6BfozekY6_VAfKX-EQ)