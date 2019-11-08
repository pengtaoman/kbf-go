package main

//func main() {
//	util.Logger()
//	util.RenderFromFileDemo()
//}

//func main() {
//	cmds := [][]string{
//		{"delete", "deployment","web-ui"},
//		{"deployment","mnist-service-local"},
//		{"delete", "svc","web-ui"},
//		{"delete", "svc","mnist-service-local"},
//		{"configmap","mnist-deploy-config"},
//	}
//
//	for cc := range cmds {
//		println(cc)
//	}
//
//	println(">>>>>>>>>>>>>>>>")
//	for i := 0; i < len(cmds); i++ {
//		pp(cmds[i]...)
//	}
// }

func pp(args ...string) {
    for aa := range args {
    	println("!!!!!!!!!!!!!!!!!!!!!!" + args[aa])
	}
}