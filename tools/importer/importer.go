package importer

func GetHeroImage() pixel.Picture{
	heroimage, err := Utils.LoadImage("testhero.png")
	if err != nil{
		panic(err)
	}
	return heroimage
}

func LoadImage(path string)(pixel.Picture, error){
	//Loads image by passed path in arg.

	file, err := os.Open(path)
	if err != nil{
		panic(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil{
		panic(err)
	}
	return pixel.PictureDataFromImage(img), nil
}

func GetRandomHeroImage(availableHeroImages map[string]*pixel.Sprite)string{
	//Choses random hero image from the map of all the available hero images.

	var imageNames []string
	for key := range availableHeroImages{
		imageNames = append(imageNames, key)
	}
	return imageNames[GetRandNum(len(imageNames))]
}

func GetRandomWeaponImage(availableWeaponImages map[string]*pixel.Sprite)string{
	//Choses random hero image from the map of all the available hero images.

	var imageNames []string
	for key := range availableWeaponImages{
		imageNames = append(imageNames, key)
	}
	return imageNames[GetRandNum(len(imageNames))]
}

func GetAvailableHeroImages()map[string]*pixel.Sprite{
	/* Saves to map all the available hero images in
	   current directory. Choses files with png extension 
	   and if it contains 'hero' suffix
	*/

	listOfHeroes := make(map[string]*pixel.Sprite)
	CommInstanse := exec.Command("ls", "./SysImages/Icons/Heroes")

	result, err := CommInstanse.Output()
	if err != nil{
		panic(err)
	}
	splitedResults := strings.Split(string(result), "\n")
	for _, value := range splitedResults{
		if len(value) > 0{
			if strings.HasSuffix(value, ".png") && strings.Contains(value, "hero"){
				fileName := strings.Split(value, ".")[0]
				image, err := LoadImage(fmt.Sprintf("./SysImages/Icons/Heroes/%s", value))
				if err != nil{
					panic(err)
				}
				sprite := pixel.NewSprite(image, image.Bounds())
				listOfHeroes[fileName] = sprite
			}
		}
	}
	return listOfHeroes
}

func GetAvailableWeaponImages()map[string]*pixel.Sprite{
	/* Saves to map all the available weapon images in
	   current directory. Choses files with png extension 
	   and if it contains 'hero' suffix
	*/

	listOfHeroes := make(map[string]*pixel.Sprite)
	CommInstanse := exec.Command("ls", "./SysImages/Icons/Weapons/")

	result, err := CommInstanse.Output()
	if err != nil{
		panic(err)
	}
	splitedResults := strings.Split(string(result), "\n")
	for _, value := range splitedResults{
		if len(value) > 0{
			if strings.HasSuffix(value, ".png") && strings.Contains(value, "weapon"){
				fileName := strings.Split(value, ".")[0]
				image, err := LoadImage(fmt.Sprintf("./SysImages/Icons/Weapons/%s", value))
				if err != nil{
					panic(err)
				}
				sprite := pixel.NewSprite(image, image.Bounds())
				listOfHeroes[fileName] = sprite
			}
		}
	}
	return listOfHeroes
}

func GetAvailableWeaponIconImages()map[string]*pixel.Sprite{
	/* Saves to map all the available hero images in
	   current directory. Choses files with png extension 
	   and if it contains 'hero' suffix
	*/

	listOfHeroes := make(map[string]*pixel.Sprite)
	CommInstanse := exec.Command("ls", "./SysImages/GameProcess/ElementsPanel")

	result, err := CommInstanse.Output()
	if err != nil{
		panic(err)
	}
	splitedResults := strings.Split(string(result), "\n")
	for _, value := range splitedResults{
		if len(value) > 0{
			if strings.HasSuffix(value, ".png") && strings.Contains(value, "weapon"){
				fileName := strings.Split(value, ".")[0]
				image, err := LoadImage(fmt.Sprintf("./SysImages/GameProcess/ElementsPanel/%s", value))
				if err != nil{
					panic(err)
				}
				sprite := pixel.NewSprite(image, image.Bounds())
				listOfHeroes[fileName] = sprite
			}
		}
	}
	return listOfHeroes
}