package ui

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
	"github.com/sirupsen/logrus"
)

type ContextOpts struct {
	//Other componenets registered to ui manager
	Components map[string]Component
}

type Component interface {
	SetContext(*ContextOpts)
	Update()

	GetTilemap() *sources.Tilemap
	GetPosition() types.Vec2
	GetID() string
	Render(*screen.ScreenManager)
}

type UIManager struct {
	componentsMap map[string]Component
	components    []Component
}

func (uim *UIManager) Clear() {
	uim.componentsMap = make(map[string]Component)
	uim.components = uim.components[:0]
}

func (uim *UIManager) AddComponent(c Component) {
	if _, ok := uim.componentsMap[c.GetID()]; ok {
		logrus.Fatalf("component with such id already exists: %s", c.GetID())
	}

	c.SetContext(&ContextOpts{
		Components: uim.componentsMap,
	})

	uim.componentsMap[c.GetID()] = c
	uim.components = append(uim.components, c)
}

func (uim *UIManager) Update() {
	for _, c := range uim.components {
		c.Update()
	}
}

func (uim *UIManager) Render(s *screen.ScreenManager) {
	for _, c := range uim.components {
		c.Render(s)
	}
}

func NewUIManager() *UIManager {
	return &UIManager{componentsMap: make(map[string]Component)}
}
