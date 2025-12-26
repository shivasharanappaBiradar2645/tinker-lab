<html>
<body>
<?php

try {
	if (1) {
		throw new Exception("error occured");
	}

} catch (Exception $e) {

	echo $e->getMessage();
	
}
echo "\n";
set_error_handler( function( $errno, $errstr) {
	echo "Error: $errstr";
});

echo $x;

?>





</body>
</html>
