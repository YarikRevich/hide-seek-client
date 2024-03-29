package sourceupgrader

// type metadataUpgrader struct{}

// type MetadataUpgrader interface {
// 	Upgrade()
// }

// func (mu *metadataUpgrader) upgradeWithTranslation(mc *sources.ModelCombination) {
// 	m := events.UseEvents().Mouse()
// 	if mc.Origin.Info.ScrollableX {
// 		if mc.Origin.Margins.LeftMargin+m.OffsetX != mc.Modified.Margins.LeftMargin {
// 			mc.Modified.Margins.LeftMargin = mc.Origin.Margins.LeftMargin + m.OffsetX
// 		}
// 	}

// 	if mc.Origin.Info.ScrollableY {
// 		if mc.Origin.Margins.TopMargin+m.OffsetY != mc.Modified.Margins.TopMargin {
// 			mc.Modified.Margins.TopMargin = mc.Origin.Margins.TopMargin + m.OffsetY
// 		}
// 	}
// }

// func (mu *metadataUpgrader) upgradeWithScale(mc *sources.ModelCombination) {
// s := screen.UseScreen()
// screenWidth, screenHeight := s.GetSize()
// screenLastWidth, screenLastHeight := s.GetLastSize()

// mc.Modified.Scale.X = (mc.Modified.Scale.X * screenWidth) / screenLastWidth
// mc.Modified.Scale.Y = (mc.Modified.Scale.Y * screenHeight) / screenLastHeight

// if mc.Origin.Size.Width*mc.Modified.Scale.X != mc.Modified.Size.Width {
// 	mc.Modified.Size.Width = mc.Origin.Size.Width * mc.Modified.Scale.X
// }

// if mc.Origin.Size.Height*mc.Modified.Scale.Y != mc.Modified.Size.Height {
// 	mc.Modified.Size.Height = mc.Origin.Size.Height * mc.Modified.Scale.Y
// }

// if mc.Origin.Buffs.Speed.X*mc.Modified.Scale.X != mc.Modified.Buffs.Speed.X {
// 	mc.Modified.Buffs.Speed.X = mc.Origin.Buffs.Speed.X * mc.Modified.Scale.X
// }

// if mc.Origin.Buffs.Speed.Y*mc.Modified.Scale.Y != mc.Modified.Buffs.Speed.Y {
// 	mc.Modified.Buffs.Speed.Y = mc.Origin.Buffs.Speed.Y * mc.Modified.Scale.Y
// }

// if mc.Origin.Offset.X*mc.Modified.Scale.X != mc.Modified.Offset.X {
// 	mc.Modified.Offset.X = mc.Origin.Offset.X * mc.Modified.Scale.X
// }

// if mc.Origin.Offset.Y*mc.Modified.Scale.Y != mc.Modified.Offset.Y {
// 	mc.Modified.Offset.Y = mc.Origin.Offset.Y * mc.Modified.Scale.Y
// }

// if mc.Origin.Info.GameRole == sources.GameMap {
// 	if screenWidth > mc.Origin.Size.Width {
// 		if m := mc.Origin.Size.Width/screenWidth + mc.Modified.Scale.X; m != mc.Modified.Scale.X {
// 			mc.Modified.Scale.X = m
// 		}
// 	} else {
// 		if m := screenWidth/mc.Origin.Size.Width + mc.Modified.Scale.X; m != mc.Modified.Scale.X {
// 			mc.Modified.Scale.X = m
// 		}
// 	}

// 	if screenHeight > mc.Origin.Size.Height {
// 		if m := mc.Origin.Size.Height/screenHeight + mc.Modified.Scale.Y; m != mc.Modified.Scale.Y {
// 			mc.Modified.Scale.Y = m
// 		}
// 	} else {
// 		if m := screenHeight / mc.Origin.Size.Height; m != mc.Modified.Scale.Y {
// 			mc.Modified.Scale.Y = m
// 		}
// 	}
// }
// }
// func (mu *metadataUpgrader) upgradeWithZoom(mc *sources.ModelCombination) {
// c := world.UseWorld().GetCamera()

// switch mc.Origin.Info.GameRole {
// case sources.GameMap:
// if m := (mc.Modified.Scale.X) / 100 * c.Zoom; m != mc.Modified.RuntimeDefined.ZoomedScale.X {
// 	mc.Modified.RuntimeDefined.ZoomedScale.X = m
// }

// if m := mc.Modified.Scale.Y / 100 * c.Zoom; m != mc.Modified.RuntimeDefined.ZoomedScale.Y {
// 	mc.Modified.RuntimeDefined.ZoomedScale.Y = m
// }
// default:

// }
// }

// //Upgrades metadata with set upgraders
// func (mu *metadataUpgrader) Upgrade() {
// 	metadata := sources.UseSources().Metadata().Collection
// 	for _, v := range metadata {
// 		if screen.UseScreen().IsResized() {
// 			mu.upgradeWithScale(v)
// 		}
// 		mu.upgradeWithTranslation(v)
// 		mu.upgradeWithZoom(v)
// 	}
// 	colliders := sources.UseSources().Colliders().Collection
// 	for _, v := range colliders {
// 		if screen.UseScreen().IsResized() {
// 			mu.upgradeWithScale(v)
// 		}
// 		mu.upgradeWithTranslation(v)
// 		mu.upgradeWithZoom(v)
// 	}
// }

// func NewUpgrader() MetadataUpgrader {
// 	return new(metadataUpgrader)
// }
