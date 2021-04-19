
Heap的特点: 关注动态极值




heapify: 把一个array转换成heap
从倒数第一个非叶子节点(n-1)/2, 从后往前遍历, 然后直接shiftDown. 时间复杂度为O(n)
通过insert来构建heap的时间复杂度为O(nlogn)



heap限定大小的意义
在N个元素中选出前M个元素
使用quick sort, 时间复杂度是O(nlogn)
用大小为M的heap, 时间复杂度是O(nlogM)
