# Rainbou

###### Disclaimer: this software is a work in progress. Further updates may [break your workflow](https://xkcd.com/1172/).

**Rainbou** *[<strong>reyn</strong>-boh]* is a command-line tool that generates
themes for a diversity of programs (as long as it has templates written for
these programs).

Given a color scheme and a template file,
rainbou will generate a theme and print it to `stdout` (your
terminal's screen).

## Installation

```bash
go get github.com/tudurom/rainbou
```

## Usage

```bash
─── ./rainbou
Usage: ./rainbou <scheme_file> <template_file>

Where scheme_file is the name or the path of the color scheme file and
template_file is the name or the path of the color scheme file.
```

## Color scheme format

Color scheme files are yaml files that have the following structure:

```yaml
metadata:
  name: foo
  author: bar
colors:
  bg: "000000"
  fg: "ffffff"
  cr: "ffffff"

  00: "aaaaaa"
  08: "bbbbbb"

  01: "cccccc"
  09: "dddddd"

  ...

  07: "eeeeee"
  15: "ffffff"
```

Where `bg` is the background color, `fg` is the foreground color, `cr` is the
color of the cursor and `00` to `15` are terminal ANSI color code numbers.
Colors are written as an HTML color code without the leading `#`.

[Full example](https://github.com/tudurom/rainbou/blob/master/db/colors/thunder.yaml).

## Template file format

Template files are normal [Go template files](https://golang.org/pkg/text/template/). Example:

```
! {{.Metadata.Name}} by {{.Metadata.Author}}

! special
*.foreground:   #{{index .ColorMap "fg" "hex"}}
*.background:   #{{index .ColorMap "bg" "hex"}}
*.cursorColor:  #{{index .ColorMap "cr" "hex"}}

! black
*.color0:       #{{index .ColorMap "00" "hex"}}
*.color8:       #{{index .ColorMap "08" "hex"}}

! red
*.color1:       #{{index .ColorMap "01" "hex"}}
*.color9:       #{{index .ColorMap "09" "hex"}}

! green
*.color2:       #{{index .ColorMap "02" "hex"}}
*.color10:      #{{index .ColorMap "10" "hex"}}

! yellow
*.color3:       #{{index .ColorMap "03" "hex"}}
*.color11:      #{{index .ColorMap "11" "hex"}}

! blue
*.color4:       #{{index .ColorMap "04" "hex"}}
*.color12:      #{{index .ColorMap "12" "hex"}}

! magenta (or orange)
*.color5:       #{{index .ColorMap "05" "hex"}}
*.color13:      #{{index .ColorMap "13" "hex"}}

! cyan
*.color6:       #{{index .ColorMap "06" "hex"}}
*.color14:      #{{index .ColorMap "14" "hex"}}

! white
*.color7:       #{{index .ColorMap "07" "hex"}}
*.color15:      #{{index .ColorMap "15" "hex"}}

! vim: set ft=xdefaults :
```

`{{.Metadata.Name}}` is the name of the color scheme and `{{.Metadata.Author}}`
is the name of the Author.

Colors can be written as:

```
{{index .ColorMap "<code>" "<format>"}}
```

Where `<code>` is a number from `00` to `15` and `<format>` is any of the following:

- `hex` - HTML color code without the leading `#`

- `hexbgr` - like `hex` but with the color components reversed (`123456` ->
`563412`)

- `dhex` - "double" `hex`; like `hex` but with color components doubled
(`123456` -> `121234345656`)

- `rgb` - array consisting of three color components. The values are in the
[0-255] range. To access the `R`
component:

  ```
  {{index .ColorMap "<code>" "rgb" 0}}
  ```

- `srgb` - same as `rgb` but with values ranging from `0` to `1`.
