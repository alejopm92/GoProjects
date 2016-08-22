
package main
import	"github.com/kataras/iris"

func	main()	{
iris.Get("/hi",	hi)
iris.Listen(":9090")
}

func	hi(ctx	*iris.Context){
  ctx.Render("templates/hi.html",	struct	{	Name	string	}{	Name:	"iris"	})
}
