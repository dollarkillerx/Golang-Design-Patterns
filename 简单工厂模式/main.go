/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 19:52 2019-09-16
 */
package main

import "fmt"

func main() {

}


type Api interface {
	Say(name string) string
}

func NewApi(t int) Api {
	switch t {
	case 1:
		return nil
	case 2:
		return nil
	default:
		return nil
	}
}

type hiAPI struct {

}

func (*hiAPI) Say(name string) string{
	return fmt.Sprintf("Hi,%s",name)
}

type helloApi struct {

}

func (*helloApi) Say(name string) string {
	return fmt.Sprintf("Hello,%s",name)
}





