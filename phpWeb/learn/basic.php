<?php 


$name = "shiva";
$age = 20;

echo $name; 
echo " is of  age ";
echo  $age;


echo "{$name} is {$age} old";

//constants
define("PI", 3.14 );
echo PI;

const PORT = "8080";
echo PORT;



$day = "sunday";

switch($day){
	case "Mon":echo "Monday";break;
	default:echo "Hakuna matata";
}


for ($i=0; $i<5; $i++) {
	echo $i;
}

while ($i <10) {
	echo $i;
	$i++;
}


function add($a, $b) {

	return $a + $b;
}
echo "\n";
echo add(2,3);


// callback function
function test($cb) {
	$cb();
}

test(function () {
	echo "\n Callback";
});


//arrays

$nums = [1,2,3];
$student = ["name"=>"shiva", "age"=>21];
$matrix = [[1,2,3],[2,3,4]];


foreach ($nums as $n) {
	echo $n;
}	

/* array functions
count(array)
array_push(array,value)
sort(array
*/




$file = fopen("test.txt", "w"); //creates the file
fwrite($file,"Hello world");
fclose($file);

echo file_get_contents("test.txt");

echo "\n";
?>
