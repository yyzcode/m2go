package view

const model = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.title}}</title>
	<link rel="stylesheet" href="/static/css/reset.css">
</head>
<body>
{{range $_,$v:=.tables}}
<p>
{{$v.Name}}
</p>
{{end}}
</body>
</html>
`
