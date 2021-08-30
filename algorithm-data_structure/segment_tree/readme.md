
线段树: 关注动态数据(动态区间)

应用范围
- 更新: 更新区间中一个元素或者一个区间的值
- 查询: 查询一个区间[i,j]的最大值, 最小值, 或者区间数字和


经典线段树问题: 区间染色

有一面墙, 长度为n, 每次选择一段墙进行染色. 
m次操作后, 我们可以看见多少种颜色?
m次操作后, 我们可以在[i,j]区间内看见多少种颜色?

染色操作(更新区间) 数组O(n)
查询操作(查询区间) 数组O(n)

另一类经典问题: 区间查询
查询一个区间[i,j]的最大值, 最小值, 或者区间数字和.
实质: 基于区间的统计查询

实例:
2017年注册用户中, 消费最高的用户? 消费最少的用户? 学习时间最长的用户?
某个太空区间中天体的总量?



线段树: 节点是区间的树
![](https://raw.githubusercontent.com/y4h2/y4h2.github.io/imagebed/img/blog/20210417170933.png)





线段树不是完全二叉树, 是平衡二叉树




# 例题
- 303. Range Sum Query - Immutable