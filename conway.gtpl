<!DOCTYPE html>
<html lang="en">

<head>
	<title>Conway</title>
</head>

<body>

	<ul style="list-style: none;">
		{{ range $key, $value := . }}
		<li>{{ $value }}</li>
		{{ end }}
	</ul>

</body>

</html>
