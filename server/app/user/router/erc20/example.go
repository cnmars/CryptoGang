package erc20

import (
	"github.com/gin-gonic/gin"
)

type Example struct{}

func (e *Example) InitExcelRouter(Router *gin.RouterGroup) {
	// exampleRouter := Router.Group("example")
	// examplelApi := v1.ApiGroupApp.ExampleApiGroup.ExcelApi
	// {
	// 	exampleRouter.POST("importExcel", examplelApi.ImportExcel)          // 导入Excel
	// 	exampleRouter.GET("loadExcel", examplelApi.LoadExcel)               // 加载Excel数据
	// 	exampleRouter.POST("exportExcel", examplelApi.ExportExcel)          // 导出Excel
	// 	exampleRouter.GET("downloadTemplate", examplelApi.DownloadTemplate) // 下载模板文件
	// }
}
