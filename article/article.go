package main

func main() {

}

/*
四种常见的 POST 提交数据方式（application/x-www-form-urlencoded，multipart/form-data，application/json，text/xml）
https://blog.csdn.net/xiao__jia__jia/article/details/79357274


深度解密 Go 语言之 sync.Pool
https://www.cnblogs.com/qcrao-2018/p/12736031.html

再有人问你分布式事务，把这篇扔给他
https://juejin.cn/post/6844903647197806605

本地事务和分布式事务工作实践
https://www.cnblogs.com/gjhjoy/p/3513679.html

Go Slice Tricks Cheat Sheet
https://ueokande.github.io/go-slice-tricks/

Go 语言设计与实现
https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-reflect/

从头到尾彻底理解KMP（2014年8月22日版）
https://blog.csdn.net/v_july_v/article/details/7041827

Redis集群方案对比：Codis、Twemproxy、Redis Cluster
https://cloud.tencent.com/developer/article/1701574

从演进式角度看消息队列
https://mp.weixin.qq.com/s/2NoRkIKG0IFcoI-nZienAQ

什么是云原生
https://juejin.cn/post/6844904197859590151#heading-0

DDD 模式从天书到实践
https://zhuanlan.zhihu.com/p/91525839

 Raft 算法
https://github.com/maemual/raft-zh_cn

动画演示
http://thesecretlivesofdata.com/raft/
场景测试：
https://raft.github.io/

一文搞懂Raft算法
https://www.cnblogs.com/xybaby/p/10124083.html#_label_4


Redis的字典渐进式扩容与ConcurrentHashMap的扩容策略比较
https://blog.csdn.net/wangmaohong0717/article/details/84611426

https://blog.csdn.net/ko0491/article/details/101265953
redis - set key value [expiration EX seconds|PX milliseconds] [NX|XX]  分布式锁

Golang实现请求限流的几种办法
https://blog.csdn.net/micl200110041/article/details/82013032

RBAC模型：基于用户-角色-权限控制的一些思考
http://www.woshipm.com/pd/1150093.html


Casbin是一个强大的、高效的开源访问控制框架，其权限管理机制支持多种访问控制模型。
https://casbin.org/docs/zh-CN/middlewares
https://www.kancloud.cn/oldlei/casbin/1289450

Golang Version Manager go版本管理
https://github.com/voidint/g


终于有人把Elasticsearch原理讲透了！
https://zhuanlan.zhihu.com/p/62892586

什么是Service Mesh
https://zhuanlan.zhihu.com/p/61901608

正确入门Service Mesh：起源、发展和现状  阿里技术
https://zhuanlan.zhihu.com/p/164730174

HTTP请求幂等性
https://blog.csdn.net/kepengs/article/details/82706721

火了 2 年的服务网格究竟给微服务带来了什么？
https://zhuanlan.zhihu.com/p/159434170

MySQL索引解析（联合索引/最左前缀/覆盖索引/索引下推）
https://www.cnblogs.com/kkbill/p/11354685.html#1-%E7%B4%A2%E5%BC%95%E5%9F%BA%E7%A1%80


【Java面试+Java后端技术学习指南】：一份通向理想互联网公司的面试指南，
包括 Java，技术面试必备基础知识、Leetcode、计算机操作系统、计算机网络、
系统设计、分布式、数据库（MySQL、Redis）、Java 项目实战等


https://github.com/OUYANGSIHAI/JavaInterview
//面试官问你B树和B+树，就把这篇文章丢给他
https://segmentfault.com/a/1190000020416577


有关于mysql 时间作为条件查询，索引失效的问题。
https://blog.csdn.net/qq_34051366/article/details/106881302?utm_medium=distribute.pc_relevant.none-task-blog-baidujs_title-0&spm=1001.2101.3001.4242

集合取并集用归并排序

一个数组的元素为1~n，在o(n)的时间内对数组排序
https://blog.csdn.net/pengchengliu/article/details/90696007


go map 实现
go mysql 对于时间字段的索引 在查询条件超过一定范围时会失效
go gc会有停顿现象 gc实现

mysql 所有的索引都是存储的主键索引来实现


https://blog.csdn.net/hebtu666/article/details/84254502
二叉树最长路径


块存储、文件存储、对象存储这三者的本质差别是什么？
https://www.zhihu.com/question/21536660/answer/1159036357

文件存储、块存储还是对象存储？
https://www.redhat.com/zh/topics/data-storage/file-block-object-storage
数据存储
https://www.alibabacloud.com/zh/knowledge/data-storage?spm=a2796.255188.1920029590.2.4bd51cdbzbry3a

对象存储，为什么那么火？
https://zhuanlan.zhihu.com/p/166289089

到底什么是CDN？
https://zhuanlan.zhihu.com/p/338951935

如何长时间保存重要数据？
https://www.zhihu.com/question/313837243/answer/660457814

mall项目是一套电商系统，包括前台商城系统及后台管理系统，基于SpringBoot+MyBatis实现，采用Docker容器化部署。
前台商城系统包含首页门户、商品推荐、商品搜索、商品展示、购物车、订单流程、会员中心、客户服务、帮助中心等模块。
后台管理系统包含商品管理、订单管理、会员管理、促销管理、运营管理、内容管理、统计报表、财务管理、权限管理、设置等模块。
https://github.com/macrozheng/mall


https://blog.csdn.net/tianlongtc/article/details/80163661  go socket编程（详细）
TCPIP原理及HTTP和TCPIP与socket的关系https://blog.csdn.net/wteruiycbqqvwt/article/details/89377615
5种网络IO模型（有图，很清楚） https://www.cnblogs.com/findumars/p/6361627.html


一名分布式存储工程师的技能树是怎样的？
https://www.zhihu.com/question/43687427/answer/96677826

MapReduce的基本工作原理
https://blog.csdn.net/fanxin_i/article/details/80388221

Google-MapReduce-GFS-BigTable（Google三大论文）
https://github.com/houmingyuan/Google-MapReduce-GFS-BigTable.git

学习分布式系统需要怎样的知识？
https://www.zhihu.com/question/23645117/answer/124708083

了解数据库内部的学习材料列表
https://github.com/pingcap/awesome-database-learning

「分布式前沿技术专题」系列文章，
https://www.infoq.cn/theme/48

https://zhuanlan.zhihu.com/p/55964292
一篇文章讲透分布式存储
雷神项目，翻译 mit 6.824 2020
https://github.com/ivanallen/thor.git

大型分布式存储方案MinIO介绍，看完你就懂了！
https://cloud.tencent.com/developer/article/1786777

开源分布式对象存储-MinIO
https://zhuanlan.zhihu.com/p/103803549

MinIO资源
http://resources.minio.org.cn/

大规模高性能分布式存储系统设计与实现
https://www.bilibili.com/video/BV1ii4y1s7DA?p=4&spm_id_from=pageDriver

6.824 分布式系统 中文字幕（2020年）
https://www.bilibili.com/video/BV16j411f7F4?from=search&seid=15945026381133569492

2020 MIT 6.824 分布式系统
https://www.bilibili.com/video/BV1R7411t71W?p=1

磁盘I/O那些事
https://tech.meituan.com/2017/05/19/about-desk-io.html

SQL,NoSQL和NewSQL的区别
https://www.cnblogs.com/klb561/p/12064126.html

MySQL学习笔记之四：并发控制和事务机制
https://www.huaweicloud.com/articles/9e5d93420423f749b8f8de4b18a58ee1.html

浅谈“HTAP”
http://www.360doc.com/content/20/0913/22/32351304_935484857.shtml

MVCC多版本并发控制
https://www.jianshu.com/p/8845ddca3b23

RTO和RPO
https://blog.csdn.net/louisjh/article/details/100736389

浅谈数据库并发控制 - 锁和 MVCC
https://draveness.me/database-concurrency-control/

https://zhuanlan.zhihu.com/p/68792989
深度解密Go语言之context

*/
