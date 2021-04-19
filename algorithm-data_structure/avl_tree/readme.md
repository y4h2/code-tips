

平衡二叉树: 任意一个节点, 左右子树高度差不超过1

1. 标注节点高度
2. 计算平衡因子


节点需要维护高度值. 在Add操作时, 更新高度.



什么时候维护平衡?
加入节点后, 向上回溯维护平衡性

左子树高度 - 右子树高度 > 1: 右旋转
![](https://raw.githubusercontent.com/y4h2/y4h2.github.io/imagebed/img/blog/20210418163332.png)
![](https://raw.githubusercontent.com/y4h2/y4h2.github.io/imagebed/img/blog/20210418163526.png)



倾斜
LL/RR
![](https://raw.githubusercontent.com/y4h2/y4h2.github.io/imagebed/img/blog/20210418195306.png)

LR
![](https://raw.githubusercontent.com/y4h2/y4h2.github.io/imagebed/img/blog/20210418200149.png)
对于LR, 先对x进行左旋转, 转化成LL
![](https://raw.githubusercontent.com/y4h2/y4h2.github.io/imagebed/img/blog/20210418200331.png)


RL
![](https://raw.githubusercontent.com/y4h2/y4h2.github.io/imagebed/img/blog/20210418200400.png)
对于RL, 先对x进行右旋转, 转化成RR
![](https://raw.githubusercontent.com/y4h2/y4h2.github.io/imagebed/img/blog/20210418200444.png)
 



删除节点
删除节点直接删除指定数值的节点, 把该节点的右子树中的最小值节点作为后继节点放入原来节点的位置, 删除完之后按照add的逻辑去平衡二叉树