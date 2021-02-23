package composer

// Type
const (
	LayoutTypeTiling   = 0
	LayoutTypeFloating = 1
	LayoutTypeMono     = 2
	LayoutTypeDock     = 3
)

// Layout ...
type Layout interface {
	New()
	Refresh()
	Show()
	Hide()
}

// LayoutSpace ...
type LayoutSpace struct {
	Top    int
	Bottom int
	Left   int
	Right  int
}

// LayoutTiling ...
type LayoutTiling struct {
	ID     string
	Space  LayoutSpace
	Master WindowSize
	Slave  WindowSize
	*WorkSpace
}

// LayoutFloating ...
type LayoutFloating struct {
	ID     string
	Space  LayoutSpace
	Full   WindowSize
	Normal []WindowSize
	*WorkSpace
}

// New ...
func (l *LayoutTiling) New() {

	l.Master.Width = l.WorkSpace.Size.Width / 2
	l.Master.Height = l.WorkSpace.Size.Height

	l.Slave.Width = l.WorkSpace.Size.Width - l.Master.Width
	l.Slave.Height = l.WorkSpace.Size.Height / (len(l.WorkSpace.Windows[l.ID]) - 1)

	for k, v := range l.WorkSpace.Windows[l.ID] {

		if k == 0 {
			v.New(v.Src, l.Master.Width, l.Master.Height, 0, 0, l.WorkSpace.Document)
		} else {
			v.New(v.Src, l.Slave.Width, l.Slave.Height, l.Master.Width, l.Slave.Height*(k-1), l.WorkSpace.Document)
		}
	}

}

// Refresh ...
func (l *LayoutTiling) Refresh() {

}

// Hide ...
func (l *LayoutTiling) Hide() {

}

// Show ...
func (l *LayoutTiling) Show() {

}
