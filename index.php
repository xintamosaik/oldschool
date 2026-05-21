<?php declare(strict_types=1); ?>
<!DOCTYPE html>
<html>

<head>
    <title>PHP Test</title>
</head>

<body>

    <form action="welcome.php" method="get">
        Name: <input type="text" name="name"><br>
        E-mail: <input type="text" name="email"><br>
        <input type="submit">
    </form>
    <?php
    // var_dump($_SERVER);
    $file_name = 'todo.json';
    $myfile = fopen($file_name, "r") or die("Error: Unable to open file!");
    $content = fread($myfile, filesize($file_name));
    fclose($myfile);
    echo $content;
    ?>
    <br>
    JSON:
    <?php
    $data = json_decode($content);
    $err = json_last_error();
    if (empty($err)) {
        echo 'ok';
    } else {
        echo $err;
    }
    ?>
    <br>
    DATA:
    <?php
    var_dump($data);
    ?>
    <br>
    <?php
    if (is_array($data)) {
        echo 'array';
        foreach ($data as $item) {
            if (is_object($item)) {
                foreach ($item as $x => $y) {
                 echo "$x: $y <br>";
                }
            } else {
                echo "no object";
            }
            ?>
            <br>
            <?php
        }
    } else {
        echo 'no array';
    }
    ?>


</body>

</html>