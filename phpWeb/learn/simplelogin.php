<html>
<body>

<?php

session_start();

if ($_POST) {
	if($_POST["user"] == "shiva" && $_POST["pass"] == "123"){
		$_SESSION["login"] = true;
		echo "login successfull \n";
	}
	else{
	echo "wrong credentials \n";
}
}

?>

<form method="post">
<input name="user">
<input type="password" name="pass">
<input type="submit">
</form>

</body>
</html>
	
