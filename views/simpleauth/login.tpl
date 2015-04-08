<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body>
	{{if .flash.error}}
		<h3>{{.flash.error}}</h3>
	{{end}}
	<form post="/login" method="post">
		{{.xsrfdata}}
		User Name: <input type="text" name="username" />
		Password: <input type="password" name="password" />
		<input type="submit" value="Login" />
	</form>
</body>
</html>