<html>
<p>=====结构体渲染=====</p>
{{ .People.Name }}
{{ .People.Age }}

<p>=====数组渲染:一=====</p>
{{ range $k,$v := .arr }}
<p>{{ $k }},{{ $v }}</p>
{{ end }}

<p>=====数组渲染:二=====</p>
{{ range .arr }}
<p>{{ . }}</p>
{{ end }}

<p>=====map渲染=====</p>
{{ .m.name }}
</html>