package composer

import (
	"bytes"
	"html/template"
)

const (
	panelhtml = `
	<table style="width: 100%%; background:#333; height:45px;">
    <colgroup>
       <col span="1" style="width: 33%%;">
       <col span="1" style="width: 33%%;">
       <col span="1" style="width: 33%%;">
    </colgroup>  
    <!-- Put <thead>, <tbody>, and <tr>'s here! -->
    <tbody>
        <tr>
            <td><button class="open1" style="width:100%%; height:100%%;" >ciao</button></td>
            <td>70%%</td>
            <td>15%%</td>
        </tr>
    </tbody>
    </table>
    `
	panelbutton = `
	<table style="width: 100%%; background:#333; height:45px;">
    <colgroup>
       <col span="1" style="width: auto;">
       <col span="1" style="width: 100%%;">
       <col span="1" style="width: auto;">
    </colgroup>
    
    
    
    <!-- Put <thead>, <tbody>, and <tr>'s here! -->
    <tbody>
        <tr>
            <td style="text-align:left;"><table><tr>{{range .Left}}<td style="border-right: 1px solid var(--select);">{{ . }}</td>{{end}}</tr></table></td>
            <td style="text-align:center; "><table><tr>{{range .Center}}<td >{{ . }}</td>{{end}}</tr></table></td>
            <td style="text-align:right;"><table><tr>{{range .Right}}<td style="border-left: 1px solid var(--select);">{{ . }}</td>{{end}}</tr></table></td>
        </tr>
    </tbody>
</table>
    `
)

// Element
const (
	PanelElementSession = `
	<button style="width:80px;"><img src="http://localhost:8090/.habitat/assets/icons/interface/32/exit.png"/></button>
    `
	PanelElementUser = `
	<button style="width:80px;"><img src="http://localhost:8090/.habitat/assets/icons/interface/32/user.png"/></button>
    `
	PanelElementBarRight = `
	<button style="width:80px;"><img src="http://localhost:8090/.habitat/assets/icons/arrows/32/bar-right-show.png"/></button>
    `
	PanelElementBarLeft = `
	<button style="width:80px;"><img src="http://localhost:8090/.habitat/assets/icons/arrows/32/bar-left-show.png"/></button>
    `
	PanelElementNotification = `
	
    `
	PanelElementTasklist = `
	<table style="background:#333;">    
    <tbody>
        <tr>
		    {{range .Items}}<td><img class="open{{.ID}}" width="22" src="{{.Icon}}"/></td>{{else}}<div><strong>no rows</strong></div>{{end}}
        </tr>
    </tbody>
    </table>
    `
	PanelElementSeparator = `
	
    `
	PanelElementManagerWin = `
	<table style="background:#333;">    
    <tbody>
        <tr>		    
		    {{range .Items}}<td><img width="22" src="{{.Icon}}" /></td>{{end}}
			<td style="border-left: 1px solid var(--select);"><button style="width:40px;"><img  height="32" src="https://img.icons8.com/pastel-glyph/64/000000/add-layer--v1.png"/></button></div></td>
        </tr>
    </tbody>
    </table>
    `
	PanelElementManagerPlu = `
	<button style="width:40px;"><img height="32" src="https://img.icons8.com/material-sharp/24/000000/plugin.png"/></button>
    `
	PanelElementManagerWidget = `
	<button style="width:40px;"><img height="32" src="https://img.icons8.com/windows/32/000000/delete-widget.png"/></button>
    `
)

// TaskList ...
type TaskList struct {
	ID    string
	Items []Window
}

// Panel ...
type Panel struct {
	ID     string
	Left   []template.HTML
	Center []template.HTML
	Right  []template.HTML
}

// NewManagerWindow ...
func NewManagerWindow(listwindow []Window) template.HTML {

	load, _ := template.New("NewManagerWindow").Parse(PanelElementManagerWin)

	data := TaskList{
		ID:    "10",
		Items: listwindow,
	}
	buf := new(bytes.Buffer)
	load.Execute(buf, data)

	return template.HTML(buf.String())
}

// NewTasklist ...
func NewTasklist(listwindow []Window) template.HTML {

	load, _ := template.New("tasklist").Parse(PanelElementTasklist)

	data := TaskList{
		ID:    "10",
		Items: listwindow,
	}

	buf := new(bytes.Buffer)
	load.Execute(buf, data)

	return template.HTML(buf.String())

	/*
		for _, v := range desk2.Window {


		}*/

}

// NewPanel ...
func NewPanel(id string, left, center, right []template.HTML) template.HTML {

	load, _ := template.New("panel").Parse(panelbutton)

	panel := Panel{
		ID:     id,
		Left:   left,
		Center: center,
		Right:  right,
	}

	buf := new(bytes.Buffer)
	load.Execute(buf, panel)

	return template.HTML(buf.String())
}
