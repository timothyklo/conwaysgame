<!DOCTYPE html>
<html lang="en">

<head>
	<title>Conway</title>

	<script>
		function validateForm() {
			var x = document.forms["myForm"]["size"].value;
			if (x == "") {
				alert("Name must be filled out");
				return false;
			}
			if (x < 1) {
				alert("The world size must be greater than 0");
				return false;
			}
		}
	</script>

</head>

<body>

	<form name="myForm" action="/getsize" onsubmit="return validateForm()" method="post">
		How big should the world be:<input type="number" name="size">
		<input type="submit" value="Size">
	</form>

</body>

</html>