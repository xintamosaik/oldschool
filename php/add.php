<?php
print_r($_POST);
$label = $_POST["label"];

include("todo.php");

if ($label == "") {
        echo "empty";
} else {
        add_todo($label);
}
show_todos();