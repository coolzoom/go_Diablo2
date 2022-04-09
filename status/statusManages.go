package status

type StatusManage struct {
	Flg             bool
	ChangeScenceFlg bool
	DoorCountFlg    bool
	LoadingFlg      bool
	MusicIsPlay     bool
	OpenBag         bool
	OpenMiniPanel   bool
	CalculateEnd    bool
	UIOFFSETX       int
	ShadowOffsetX   int
	ShadowOffsetY   int
	PLAYERCENTERX   int64
	PLAYERCENTERY   int64
}

func NewStatusManage() *StatusManage {
	n := &StatusManage{
		Flg:             false,
		ChangeScenceFlg: false,
		DoorCountFlg:    false,
		LoadingFlg:      false,
		MusicIsPlay:     false,
		OpenBag:         false,
		OpenMiniPanel:   false,
		CalculateEnd:    false,
		UIOFFSETX:       0,
		ShadowOffsetX:   -350,
		ShadowOffsetY:   365,
		PLAYERCENTERX:   388, //LAYOUTX/2
		PLAYERCENTERY:   242, //LAYOUTY/2
	}
	return n
}