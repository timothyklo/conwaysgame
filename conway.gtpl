<!DOCTYPE html>
<html lang="en">

<head>
	<title>Conway</title>
</head>

<body>

	<form action="/runlife" method="post">
		<input type="hidden" name="world" value={{.}}>
		<input type="submit" value="Run">
	</form>

	<ul style="list-style: none;">
		{{ range $key, $value := . }}
		<li>{{ $value }}</li>
		{{ end }}
	</ul>

	<a href="/getsize">Restart</a>

</body>
</html>
