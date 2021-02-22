package composer

// Type
const (
	LayoutTypeTiling   = "1"
	LayoutTypeFloating = "2"
	LayoutTypeMono     = "0"
)

// Layout ...
type Layout interface {
	new()
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
	*Screen
}

// New ...
func (l *LayoutTiling) New() {

	//l.Windows[0].New("0", "http://localhost:8090", 1000, 500, 0, 0, document)
	l.Master.Width = l.Screen.Size.Width / 2
	l.Master.Height = l.Screen.Size.Height

	l.Slave.Width = l.Screen.Size.Width - l.Master.Width
	l.Slave.Height = l.Screen.Size.Height / (len(l.Screen.Windows[l.ID]) - 1)

	l.Screen.Windows[l.ID][0].New("0", "http://localhost:8090", l.Master.Width, l.Master.Height, 0, 0, l.Screen.Document)
	l.Screen.Windows[l.ID][0].New("1", "http://localhost:8090", l.Slave.Width, l.Slave.Height, l.Master.Width, 0, l.Screen.Document)
}
