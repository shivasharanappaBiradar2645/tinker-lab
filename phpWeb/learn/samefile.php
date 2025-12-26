<html>
<body>

<?php

if ($_SERVER["REQUEST_METHOD"] == "POST") {
	echo "<p> shiva </p>";
	echo $_POST["name"];

}

if ($_SERVER["REQUEST_METHOD"] == "GET"){

	print_r($_GET["lang"]);
}

?>

<form method="post" >
<input type="text" name="name">
<input type="submit" >

</form>

<form method="get" >
<input type="checkbox" name="lang[]" value="PHP">php
<input type="checkbox" name="lang[]" value="JS">js
<input type="submit" >
</form>


</body>
</html>
