<?php
$routes = array(
        "new" =>  'new.php',
        'edit' => 'edit.php',
        'add' => 'add.php'
);



$subpage = ltrim($_SERVER['REQUEST_URI'], '/');

$route_is_set_in_hashmap = array_key_exists($subpage, $routes);
if ($route_is_set_in_hashmap) {
        echo "A";
        $file = $routes[$subpage];
        include($file);

} else {
        include("todo.php");
        show_todos();
}


?>