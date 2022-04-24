package engine

import (
	"embed"
	"game/fonts"
	"game/layout"
	"game/mapCreator/mapItems"
	"game/maps"
	"game/music"
	"game/role"
	"game/status"
	"game/tools"
	"runtime"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//配置信息
const (
	SCREENWIDTH  int = 490
	SCREENHEIGHT int = 300
	WEOFFSETX    int = 127
	WEOFFSETY    int = 14
)

type Game struct {
	count, countForMap int
	player             *role.Player         //玩家
	maps               *maps.MapBase        //地图
	mapManage          *mapItems.MapItems   //地图动画等管理
	ui                 *layout.UI           //UI
	music              music.MusicInterface //音乐
	status             *status.StatusManage //状态管理器
	font_style         *fonts.FontBase      //字体
}

var (
	counts      int = 0
	countsFor20 int = 0
	countsFor12 int = 0
	frameNums   int = 4
	frameSpeed  int = 5
	mouseX      int
	mouseY      int
	newPath     []uint8
	turnLoop    uint8 = 0
)

//go:embed resource
var asset embed.FS

//GameEngine
func NewGame() *Game {
	//statueManage
	sta := status.NewStatusManage()
	//Map
	m := maps.NewMap(&asset, sta)
	//Player  设置初始状态和坐标
	r := role.NewPlayer(5280, 1880, tools.IDLE, 0, 0, 0, &asset, m, sta)
	//字体
	f := fonts.NewFont(&asset)
	//UI
	u := layout.NewUI(&asset, sta, m, f)
	//BGM
	bgm := music.NewMusicBGM(&asset)
	//场景动画
	object := mapItems.New(&asset, sta)
	gameEngine := &Game{
		count:       0,
		countForMap: 0,
		player:      r,
		maps:        m,
		ui:          u,
		music:       bgm,
		status:      sta,
		mapManage:   object,
		font_style:  f,
	}
	//启动游戏
	gameEngine.StartEngine()
	return gameEngine
}

//引擎启动
func (g *Game) StartEngine() {
	//隐藏鼠标系统的ICON
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	w := sync.WaitGroup{}
	w.Add(1)
	//UI Init
	go func() {
		//加载字体
		g.font_style.LoadFont("resource/font/pf_normal.ttf")
		g.ui.LoadGameLoginImages()
		runtime.GC()
		w.Done()
	}()
	w.Wait()
	go func() {
		runtime.GC()
	}()
}

func (g *Game) Update() error {
	mouseX, mouseY = ebiten.CursorPosition()
	//切换场景逻辑
	if !g.status.ChangeScenceFlg {
		switch g.status.CurrentGameScence {
		case tools.GAMESCENESTART:
			//进入游戏场景逻辑
			g.changeScenceGameUpdate()
		case tools.GAMESCENEOPENDOOR:
			//游戏加载逻辑
			g.ChangeScenceOpenDoorUpdate()
		case tools.GAMESCENESELECTROLE:
			//进入游戏选择界面逻辑
			g.ChangeScenceSelectUpdate()
		default:
			//进入游戏登录界面逻辑
			g.ChangeScenceLoginUpdate()
		}
	}
	//全屏显示控制
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		i := ebiten.IsFullscreen()
		ebiten.SetFullscreen(!i)
	}
	//Debug 信息显示控制
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.status.DisPlayDebugInfo = !g.status.DisPlayDebugInfo
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	//判断是否切换场景
	if !g.status.ChangeScenceFlg {
		switch g.status.CurrentGameScence {
		case tools.GAMESCENESTART:
			g.ChangeScenceGameDraw(screen)
		case tools.GAMESCENESELECTROLE:
			g.ChangeScenceSelectDraw(screen)
		case tools.GAMESCENEOPENDOOR:
			g.ChangeScenceOpenDoorDraw(screen)
		default:
			g.ChangeScenceLoginDraw(screen)
		}
	}
	//绘制鼠标ICON
	g.ui.DrawMouseIcon(screen, mouseX, mouseY)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return tools.LAYOUTX, tools.LAYOUTY
}
