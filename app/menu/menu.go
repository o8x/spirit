package menu

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
)

type Menu struct {
	submenu *menu.Menu
	parent  *menu.MenuItem
}

func New(label string) *Menu {
	submenu := menu.NewMenu()
	item := menu.SubMenu(label, submenu)

	return &Menu{
		submenu: submenu,
		parent:  item,
	}
}

func (m *Menu) Reset() {
	m.parent.SubMenu.Items = nil
}

func (m *Menu) Add(label string, click menu.Callback) *menu.MenuItem {
	item := menu.Text(label, nil, click)
	m.parent.Append(item)
	return item
}

func (m *Menu) AddSubmenu(label string) *menu.Menu {
	return m.submenu.AddSubmenu(label)
}

func (m *Menu) AddText(label string) *menu.MenuItem {
	item := menu.Text(label, nil, nil)
	item.Disabled = true
	m.parent.Append(item)
	return item
}

func (m *Menu) Remove(label string) {
	for i, it := range m.parent.SubMenu.Items {
		if it.Label == label {
			m.parent.SubMenu.Items = append(m.parent.SubMenu.Items[:i], m.parent.SubMenu.Items[i+1:]...)
			return
		}
	}
}

func (m *Menu) BindKey(label string, accelerator *keys.Accelerator) {
	for _, it := range m.parent.SubMenu.Items {
		if it.Label == label {
			it.Accelerator = accelerator
		}
	}
	m.parent.Accelerator = accelerator
}

func (m *Menu) AddSeparator() {
	m.submenu.AddSeparator()
}

func (m *Menu) Build() *menu.MenuItem {
	return m.parent
}
