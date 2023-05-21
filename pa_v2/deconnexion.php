<?php

session_start();

$_SESSION['email'] = "";
$_SESSION['prenom'] = "";
$_SESSION['image'] = "";
session_destroy();
header('location: log_sign.php');
exit;

?>