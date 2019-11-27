package service

import (
	"fmt"
	"testing"
)

func TestService(t *testing.T) {

	//defaultDataHandlerService := TagModuleDefaultDataHandlerService{}
	dataHandlerService := TagModuleDataHandlerService{}
	//guessHandlerService := GuessLikeModuleDataHandlerService{}

	//go dataHandlerService.TagModuleDataMachiningFacatory()
	//dataHandlerService.TagModuleDataCollect()

	//defaultDataHandlerService.DefaultDataCollect()
	//guessHandlerService.Handler()
	v := int64(1)
	dataHandlerService.SetCollectSign(v)
	fmt.Println(dataHandlerService.IsCollectEnd())
	fmt.Println("www")
}
