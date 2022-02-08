package sources

// const (
// 	GameMap GameRole = "gamemap"
// )

// //Runtime defined metadata model
// type RuntimeDefined struct {
// 	ZoomedScale types.Vec2
// }

// type Transition struct {
// 	StartScale, EndScale types.Vec2
// }

// //All the positioning properties should be in range (0; 100)
// type MetadataModel struct {
// 	Type

// 	Animation Animation

// 	Transition Transition

// 	//Information about metadata file
// 	Info struct {
// 		//Parent file metadata one related to
// 		GameRole

// 		Parent                   string
// 		ScrollableX, ScrollableY bool
// 	}

// 	Size types.Vec2

// 	Margins types.Vec2

// 	Spawns []*server_external.PositionInt

// 	Physics struct {
// 		G float64
// 	}

// 	Buffs struct {
// 		Speed float64
// 	}

// 	Scale types.Vec2

// 	Offset types.Vec2

// 	Text struct {
// 		Symbols  string
// 		Position TextPosition
// 	}

// 	Fonts struct {
// 		FontColor
// 	}

// 	Camera struct {
// 		MaxZoom, MinZoom, InitZoom float64
// 	}

// 	Effects struct {
// 		TraceColorBegin [4]float32
// 		TraceColorEnd   [4]float32
// 	}
// }

// // func (m *MetadataModel) GetRect() image.Rectangle {
// // 	ms := m.GetMargins()
// // 	s := m.GetSize()
// // 	ma := m.GetMargins()
// // 	return image.Rect(int(ma.X), int(ma.Y), int(ms.X+s.X), int(ms.Y+s.Y))
// // }

// func (m *MetadataModel) GetSize() types.Vec2 {
// 	return m.Size
// }

// func (m *MetadataModel) GetMargins() types.Vec2 {
// 	ss := m.GetSize()
// 	sc := m.GetScale()
// 	s := screen.UseScreen()
// 	size := s.GetSize()
// 	r := types.Vec2{X: ((m.Margins.X * size.X) / 100) - (ss.X * sc.X / 2), Y: ((m.Margins.Y * size.Y) / 100) - (ss.Y * sc.Y / 2)}

// 	if m.Type.Contains("scrollable") {
// 		o := m.GetOffset()

// 		if m.Type.Contains("scrollablex") {
// 			r.X += o.X
// 		}
// 		if m.Type.Contains("scrollabley") {
// 			r.Y += o.Y
// 		}
// 	}

// 	return r
// }

// func (m *MetadataModel) GetScale() types.Vec2 {
// 	s := screen.UseScreen()
// 	size := s.GetSize()
// 	// lastSize := s.GetLastSize()
// 	return types.Vec2{X: ((m.Scale.X * size.X) / 100), Y: ((m.Scale.Y * size.Y) / 100)}
// }

// func (m *MetadataModel) GetBuffSpeed() types.Vec2 {
// 	s := screen.UseScreen()
// 	size := s.GetSize()
// 	lastSize := s.GetLastSize()
// 	y := m.Buffs.Speed * (size.Y / lastSize.Y)
// 	x := (m.Buffs.Speed * (size.X / lastSize.X)) - (y / 2)
// 	avg := (y + x) / 2
// 	return types.Vec2{X: avg, Y: avg}
// }

// 		if len(mm.Info.Parent) != 0 {
// 			path := filepath.Join(utils.GetBasePath(strings.Split(path, "dist/metadata/")[1]), mm.Info.Parent)
// 			img := UseSources().Images().GetImage(path)

// 			x, y := img.Size()

// 			if mm.Animation.FrameNum != 0 {
// 				x /= int(mm.Animation.FrameNum)
// 			}

// 			mm.Size.X = float64(x)
// 			mm.Size.Y = float64(y)
// 		}
