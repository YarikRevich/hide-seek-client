package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/YarikRevich/hide-seek-client/internal/core/latency"
	"github.com/YarikRevich/hide-seek-client/internal/core/networking"
	"github.com/YarikRevich/hide-seek-client/internal/core/notifications"
	"github.com/YarikRevich/hide-seek-client/internal/core/paths"
	"github.com/YarikRevich/hide-seek-client/internal/core/profiling/runtime"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
	"github.com/YarikRevich/hide-seek-client/internal/loop"
	isconnect "github.com/alimasyhur/is-connect"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"

	"github.com/YarikRevich/hide-seek-client/tools/params"
	"github.com/YarikRevich/hide-seek-client/tools/printer"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
)

func init() {
	flag.Parse()
	paths.InitSystemPaths()
	rand.Seed(time.Now().Unix())

	if isconnect.IsOnline() {
		statemachine.Networking.SetState(statemachine.NETWORKING_ONLINE)
	}

	sr := beep.SampleRate(44100)
	if err := speaker.Init(sr, sr.N(time.Second/10000)); err != nil {
		logrus.Fatal("error happened initializing audio speaker")
	}

	logrus.SetFormatter(new(logrus.JSONFormatter))
	lgf, err := os.OpenFile(filepath.Join(paths.GAME_LOG_DIR, "log.log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	logrus.SetOutput(lgf)
	logrus.SetLevel(logrus.WarnLevel)
	if params.IsDebug() {
		logrus.SetLevel(logrus.DebugLevel)
	}

	printer.PrintCliMessage("HideSeek\nClient!")
	if params.IsDebug() {
		printer.PrintCliMessage("You are in\nDebug mode")
	}
}

func main() {
	if params.IsProfileCPU() {
		runtime.UseProfiler().StartMonitoring()

		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt)
		go func() {
			for range c {
				runtime.UseProfiler().StopMonitoring()
				os.Exit(0)
			}
		}()
	}

	nm := networking.NewNetworkingManager()
	nm.DialWANServer()

	ntm := notifications.NewNotificationManager()
	latency.UseLatency().Timings().ExecEach(latency.Connectivity, time.Second*3, func() {
		if !isconnect.IsOnline() && statemachine.Dial.Check(statemachine.DIAL_WAN) {
			ntm.Write("You are offline, turn on LAN server to play locally!", -1, notifications.Error)
			if statemachine.Layers.Check(statemachine.LAYERS_SETTINGS_MENU) ||
				statemachine.Layers.Check(statemachine.LAYERS_START_MENU) {
				statemachine.Networking.SetState(statemachine.NETWORKING_OFFLINE)
			}
		} else if !nm.IsWANConnected() {
			nm.WAN.Connect()
			ntm.Write("Servers are offline!", -1, notifications.Error)
			statemachine.Networking.SetState(statemachine.NETWORKING_OFFLINE)
		} else {
			statemachine.Networking.SetState(statemachine.NETWORKING_ONLINE)
		}
	})

	sm := screen.NewScreenManager()
	maxSize := sm.GetMaxSize()
	minSize := sm.GetMinSize()

	fmt.Println(maxSize, minSize)
	ebiten.SetWindowSize(int(maxSize.X), int(maxSize.Y))
	ebiten.SetWindowTitle("HideSeek-Client")
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSizeLimits(int(minSize.X), int(minSize.Y), -1, -1)

	log.Fatalln(ebiten.RunGame(loop.New(&loop.LoopOpts{
		ScreenManager:       sm,
		NotificationManager: ntm,
		WorldManager:        world.NewWorldManager(),
		NetworkingManager:   nm,
	})))
}
