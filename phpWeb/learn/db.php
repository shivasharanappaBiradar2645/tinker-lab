<html>
<body>

<?php
//phpinfo();
// uncoment the mysqli extension
//sudo apt install php8.x-mysql
//Then enable it
//sudo phpenmod mysqli



$conn = mysqli_connect("localhost", "tutorial","123456789","tutorial");

$result = mysqli_query($conn, "CREATE TABLE IF NOT EXISTS user(
id INT AUTO_INCREMENT PRIMARY KEY,
name VARCHAR(20));");

//mysqli_query($conn,"
//INSERT INTO user(name) VALUES ('SHIVA'),('SHIV');");

$result = mysqli_query($conn, "SELECT * FROM user");

while( $row = mysqli_fetch_assoc($result)) {

	echo $row["name"];
}

?>

</body>
</html>
