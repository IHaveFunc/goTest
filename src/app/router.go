package app

// Route struct defines route pattern rule item.
/**
 * 解读: 应用初始化的时候会定义很多的路由
 * 每个路由又都有相应的函数。所以这里定一个切片。
 */
type Route struct {
	method string    `Request:"method"`
	params []string  `Request:"params"`
	fn     []Handler `Request:"correspond function"`
	// regex  *regexp.Regexp `不知道这个有什么用`
}

const (
	ROUTER_METHOD_GET    = "GET"
	ROUTER_METHOD_POST   = "POST"
	ROUTER_METHOD_PUT    = "PUT"
	ROUTER_METHOD_DELETE = "DELETE"
)

// 路由器实例提供路由器模式和处理程序
type Router struct {
	routeSlice []*Route `Route:"struct"`
}

// Create an object using the Route type
func newRoute() *Route {
	route := new(Route)              // new 一个结构体
	route.params = make([]string, 0) // 请求参数
	return route
}

// Create an object using the Router type
func NewRouter() *Router {
	rt := new(Router)                 // new 一个路由模块
	rt.routeSlice = make([]*Route, 0) // 路由切片
	return rt
}

// Get 请求
func (rt *Router) Get(pattern string, fn ...Handler) {
	route := newRoute()              //
	route.method = ROUTER_METHOD_GET // GET请求
	route.fn = fn                    // 路由对应的函数
	rt.routeSlice = append(rt.routeSlice, route)
}

// 查找确实找到匹配的规则和解析路由URL，返回路由参数和匹配的处理程序。

//

/* Router 模型
Router.routeSlice = [
    {
        Route:{
            method,
            params,
            fn
        }
    },
    ... More
]
*/

// Handler定义了路由处理程序，中间件处理程序类型。
type Handler func()
