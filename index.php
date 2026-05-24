<!DOCTYPE html>
<html>

<head>
    <title>PHP Test</title>
    <meta name="color-scheme" content="dark light" />
</head>

<body>
    <a href="/new">new</a>
    <?php
    $subpage = $_SERVER['REQUEST_URI'];
    echo $subpage;
    switch ($subpage) {
        case '/new':
        case 'new': {
            include('new.php');
            break;
        }
        case '/add':
        case 'add': {
            include('add.php');
            include('todo.php');
            $todos = read_and_validate_todo_json();
            if (empty($todos)) {
                echo "error";
            } else {
                print_todos_table($todos);
            }
            break;
        }
        default: {
            include("todo.php");
            $todos = read_and_validate_todo_json();
            if (empty($todos)) {
                echo "error";
            } else {
                print_todos_table($todos);
            }
        }

    }


    ?>
</body>

</html>