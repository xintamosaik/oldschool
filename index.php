<!DOCTYPE html>
<html>

<head>
    <title>PHP Test</title>
    <meta name="color-scheme" content="dark light" />
</head>

<body>

    <?php
    function show_new() {
        include('new.php');
    }

    function respond_add() {
        include('add.php');
    }

    function show_edit() {
        include('edit.php');
    }

    function show_todos() {
           include('todo.php');
            $todos = read_and_validate_todo_json();
            if (empty($todos)) {
                echo "error";
            } else {
                print_todos_table($todos);
            }
    }

    $subpage = $_SERVER['REQUEST_URI'];
    echo $subpage;
    switch ($subpage) {
        case '/new':
        case 'new': {
            show_new();
            break;
        }
        case '/add':
        case 'add': {
            respond_add() ;
            show_todos();
            break;
        }
        default: {
            show_todos();
        }

    }


    ?>
</body>

</html>