package benchmarks

import "testing"

//MY FUNCTION
//func BenchmarkLoadBalancer(b *testing.B) {
//	b.Run("chanelBalancer", func(b *testing.B) {
//		cons := make([]*Connection, 10)
//		b.ResetTimer()
//		for i := 0; i < b.N; i++ {
//			balancer := NewLoadBalancerChan(cons)
//			balancer.Init()
//			for range len(cons) {
//				balancer.NextConn()
//			}
//		}
//	})
//	b.Run("atomicBalancer", func(b *testing.B) {
//		cons := make([]*Connection, 10)
//		b.ResetTimer()
//		for i := 0; i < b.N; i++ {
//			balancer := NewLoadBalancerAtomic(cons)
//			for range len(cons) {
//				balancer.NextConn()
//			}
//		}
//	})
//	b.Run("mutexBalancer", func(b *testing.B) {
//		cons := make([]*Connection, 10)
//		b.ResetTimer()
//		for i := 0; i < b.N; i++ {
//			balancer := NewLoadBalancerMutex(cons)
//			for range len(cons) {
//				balancer.NextConn()
//			}
//		}
//	})
//}

func BenchmarkLoadBalancer(b *testing.B) {
	const connsN = 100
	const triesN = 1000

	conns := make([]*Connection, 0, connsN)
	for i := 0; i < connsN; i++ {
		conns = append(conns, &Connection{})
	}

	lbChan := NewLoadBalancerChan(conns)
	lbChan.Init()
	defer lbChan.Close()
	lbAtomic := NewLoadBalancerAtomic(conns)
	lbMutex := NewLoadBalancerMutex(conns)

	b.ResetTimer()

	b.Run("chan", func(b *testing.B) {
		for i := 0; i < triesN; i++ {
			lbChan.NextConn()
		}
	})

	b.Run("atomic", func(b *testing.B) {
		for i := 0; i < triesN; i++ {
			lbAtomic.NextConn()
		}
	})

	b.Run("mutex", func(b *testing.B) {
		for i := 0; i < triesN; i++ {
			lbMutex.NextConn()
		}
	})
}
