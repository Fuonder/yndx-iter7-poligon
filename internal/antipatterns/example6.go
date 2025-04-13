package antipatterns

///* плохо */
//switch v {
//case 1:
//fmt.Println(v)
//break
//case 2:
//fmt.Println(v)
//break
//}
//
///* хорошо */
//switch v {
//case 1:
//fmt.Println(v)
//case 2:
//fmt.Println(v)
//}

//switch v {
//case 1:
//fmt.Println(v)
//fallthrough  // выполним код следующего case
//case 2:
//fmt.Println(v)
//}
