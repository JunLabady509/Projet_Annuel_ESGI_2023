<?php
try{
    $bdd = new PDO('mysql:host=localhost;dbname=kano','root', 'esgi',[PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION]);
}catch(Exception $e){
    die('Erreur : ' . $e->getMessage());
}
?>