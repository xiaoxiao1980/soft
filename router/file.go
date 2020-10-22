package router

import (
	"app/api"

	"github.com/gin-gonic/gin"
)

func InitFileRouter(Router *gin.RouterGroup) {
	FileRouter := Router.Group("file")
	{
		// FileRouter.POST("", api.Upload)
		FileRouter.POST("upload/:type", api.MultiUpload)
		FileRouter.GET("", api.GetAllFile)
		//UserRouter.POST("Find", v2.Find)
		//UserRouter.POST("FindOne", v2.FindOne)
		//UserRouter.POST("FindOneAndDelete", v2.FindOneAndDelete)
		//UserRouter.POST("FindOneAndReplace", v2.FindOneAndReplace)
		//UserRouter.POST("FindOneAndUpdate", v2.FindOneAndUpdate)
		//UserRouter.POST("UpdateOne", v2.UpdateOne)
		//UserRouter.POST("UpdateMany", v2.UpdateMany)
		//UserRouter.POST("DeleteOne", v2.DeleteOne)
		//UserRouter.POST("DeleteMany", v2.DeleteMany)
		//UserRouter.POST("Aggregate", v2.Aggregate)
		//UserRouter.POST("BulkWrite", v2.BulkWrite)
		//UserRouter.POST("CountDocuments", v2.CountDocuments)
		//UserRouter.POST("Distinct", v2.Distinct)
		//UserRouter.POST("EstimatedDocumentCount", v2.EstimatedDocumentCount)
		//UserRouter.POST("ReplaceOne", v2.ReplaceOne)
		//UserRouter.POST("Watch", v2.Watch)
	}
}
