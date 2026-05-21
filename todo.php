<?php declare(strict_types=1);

function validate_todo_item(object $item): bool
{
    if (!is_object($item)) {
        echo "no object";
        return false;
    }

    return true;
}
function read_and_validate_todo_json(): array|null
{
    // var_dump($_SERVER);
    $file_name = 'todo.json';
    $myfile = fopen($file_name, "r");
    if (!$myfile) {
        echo "file could not be found";
        return null;
    }
    $content = fread($myfile, filesize($file_name));
    if ($content === false) {
        echo "no content in file";
        return null;
    }
    fclose($myfile);

    $data = json_decode($content);

    if (json_last_error() !== JSON_ERROR_NONE) {
        echo 'error parsing json';
        return null;
    }
    if (!is_array($data)) {
        echo 'no array';
        return null;
    }

    foreach ($data as $item) {
        if (!validate_todo_item($item)) {
            echo 'The json does not match the spec';
            return null;
        }

    }

    return $data;
}
function print_todos_table(array $todos): void
{
    if (empty($todos)) {
        echo 'oh no';
        return;
    }
    ?>
    <table>
        <thead>
            <tr>
                <th>id</th>
                <th>name</th>
            </tr>
        </thead>
        <tbody>
            <?php
            foreach ($todos as $item) {
                ?>
                <tr><?php
                foreach ($item as $x => $y) {
                    echo "<td>$y</td>";
                }
                ?></tr><?php
            }
            ?>
        </tbody>
    </table>
    <?php

}

$todos = read_and_validate_todo_json();
if (empty($todos)) {
    echo "error";
} else {
    print_todos_table($todos);
}
?>