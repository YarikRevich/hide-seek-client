package loaderpostprocessor

//Wrapper for calling loader post processors
func ApplyPostProcessors(c ...func()){
	for _, v := range c{
		v()
	}
}