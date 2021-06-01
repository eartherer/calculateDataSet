package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var resloverPool map[string]DataSetReslover

func main() {

	resloverPool = initilizeResloverPool()

	r := gin.Default()
	r.POST("/calculate", calculateDataSetHandler(&resloverPool))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func initilizeResloverPool() map[string]DataSetReslover {
	result := make(map[string]DataSetReslover)
	result["ExampleReslover"] = &ExampleReslover{}
	return result
}

func calculateDataSetHandler(resolverPool *map[string]DataSetReslover) gin.HandlerFunc {
	return func(c *gin.Context) {

		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		var datasetInput DataSetInput

		err = json.Unmarshal(jsonData, &datasetInput)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		resolver, ok := (*resolverPool)[datasetInput.ResolverName]
		if !ok {
			c.JSON(http.StatusBadRequest, "Resolver not found")
			return
		}

		if !resolver.isValid(datasetInput) {
			c.JSON(http.StatusBadRequest, "DataSet Input is invalid")
			return
		}

		result, err := resolver.resloveProblem(datasetInput)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(200, result)
	}
}

type DataSetInput struct {
	ResolverName string
	Size         int
	Input        []string
	FindPosition []int
}
type ResloveResult struct {
	DataSetInput
	Result map[int]string
}

type DataSetReslover interface {
	isValid(DataSetInput) bool
	resloveProblem(DataSetInput) (ResloveResult, error)
}
