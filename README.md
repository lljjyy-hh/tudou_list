# tudou_list
tudou_list

参考资料：
> [服务器结构](https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html)
> [服务器结构1](https://www.dudley.codes/posts/2020.05.19-golang-structure-web-servers/)
> [服务器结构demo](https://github.com/dudleycodes/golang-microservice-structure)


&nbsp;
---
##### 22.03.24  
+ 新建database包，对gorm二次封装。采用custom_logger，后续可将运行时的报错和sql语句写进数据库里。  
+ 新建target实体和model，做repository层的作用，负责与database层与数据库的交互，后续拟采用server层实现具体业务逻辑。
+ 初步独立出handler包，后续实现路由功能。  

&nbsp;
---
##### 22.03.25  
老样子，分controler、service、repository三层。实现service层Target功能。  

&nbsp;
---
##### 22.03.27  
实现target查询、更新、插入功能。  
