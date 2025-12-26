<html>
<body>
<?php

//setcookie("user","shiva", time() + 3600);
echo "data in cookie ";
echo  $_COOKIE["user"]  ;

move_uploaded_file($_FILES["file"]["tmp_name"] , "uploads/" . $_FILES["file"]["name"]);

?>

<form method="post" enctype="multipart/form-data">
<input type="file" name="file">
<input type="submit">
</form>


</body>

</html>
