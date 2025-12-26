<?php 
//php -S localhost:8000

session_start();


if ($_POST) {

	if (empty($_POST['name']) ) {
        echo "Name required";
	exit();

}
	echo $_POST['name'];
	$_SESSION['user'] = $_POST['name'];

}

if ($_GET) {
	echo $_GET["name"];
}


if ( !empty($_SESSION['user'])) {
echo "\n accessing session variable ";
echo $_SESSION['user'];
}
session_destroy();

/*
mysql

$conn = mysqli_connect("url","user,"pass","db");
mysqli_query($conn, query);

*/
?>

