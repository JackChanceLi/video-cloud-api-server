package main
//模拟使用通道的方式实现令牌桶算法，保证不会因为连接过多导致服务器压力过大
import (
	"log"
)

//定义令牌桶结构
type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

//令牌桶构造函数
func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {  //表示已经超过了最大连接
		log.Printf("Reached the rate limitation.")
		return false
	}
	//令牌桶中添加新的连接
	cl.bucket <- 1
	log.Printf("Successfully got connection")
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	//将通道中的bucket取出，相当于释放一个连接供新的请求使用
	c := <-cl.bucket
	log.Printf("New connection coming: %d", c)
}