<html>
<body>

<?php
session_start();

/*
$_SESSION["time"] = time();

while ( 1 ) {
if ( time() - $_SESSION["time"] > 10) {

	session_destroy();
	echo "Session expired";
	exit();
}
} */

$_SESSION["count"] = ($_SESSION["count"] ?? 0) + 1;

echo $_SESSION["count"];

$hash = password_hash("secret", PASSWORD_DEFAULT);
echo "\n password matched: ";
echo password_verify("secret", $hash);


?>



<p>Session </p>

</body>
</html>
