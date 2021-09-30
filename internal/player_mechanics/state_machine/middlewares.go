package statemachine

import "github.com/YarikRevich/HideSeek-Client/internal/buffers/text"

func CleanBuffers(){
	text.UseBuffer().Clean()
}

func UseMiddlewares(){
	CleanBuffers()
}