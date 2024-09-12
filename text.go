package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Text struct {
	rl.Vector2
	Color    rl.Color
	FontSize float
	Spacing  float
	Font     rl.Font

	txt string
}

func NewText(
	txt string,
	size float,
	color rl.Color,
	pos rl.Vector2,
) Text {
	return Text{
		Vector2:  pos,
		Color:    color,
		txt:      txt,
		FontSize: size,

		// Default values.
		Spacing: 6,
		Font:    rl.GetFontDefault(),
	}
}

func NewTextEx(
	txt string,
	fontSize float,
	color rl.Color,
	font rl.Font,
	pos rl.Vector2,
	spacing float,
) Text {
	return Text{
		txt:      txt,
		FontSize: fontSize,
		Color:    color,
		Font:     font,
		Vector2:  pos,
		Spacing:  spacing,
	}
}

func (t Text) Size() rl.Vector2 {
	return rl.MeasureTextEx(
		t.Font,
		t.txt,
		t.FontSize,
		t.Spacing,
	)
}

func (t Text) Text() string {
	return t.txt
}

func (t *Text) SetText(txt string) {
	t.txt = txt
}

func (t Text) Draw() {
	rl.DrawTextEx(
		t.Font,
		t.txt,
		t.Vector2,
		t.FontSize,
		t.Spacing,
		t.Color,
	)
}
