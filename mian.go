<<<<<<< HEAD
//首先说明,下载完别人的项目如何使用的问题.当然也是自己的项目如何放到哪个目录的问题!!!!1
//如何使用这些下载的自定义哭包文件!!!!!!!!!!!!!!!!!!!!!!!!!!11
//把文件放到goroot目录下.也就是c:/go/src里面 一般是github.com/用户名下面.




package main  // 代码包声明语句。
import "fmt" //系统包用来输出的
import "github.com/algoGuy/EasyMIDI/smf"

func main() {
	// 打印函数调用语句。用于打印输出信息。
	a:=smf.A
	fmt.Println("helloworld")
	fmt.Println(a)
=======
package main  // 代码包声明语句。
import "fmt" //系统包用来输出的
import "./ii"

func main() {
	// 打印函数调用语句。用于打印输出信息。
	//a:=smf.Division{0x23,34}

	fmt.Println("helloworld")
	fmt.Println(ii.A)
>>>>>>> ec1db8e57cead1d8583e9a9b049bd27bddfaed6e
}
