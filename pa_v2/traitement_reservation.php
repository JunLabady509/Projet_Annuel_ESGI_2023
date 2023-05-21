<?php
session_start();

// Récupération des données du formulaire
$nom = $_POST['nom'];
$email = $_POST['email'];
$date = $_POST['date'];

//connexion base de donnée
try{
    $bdd = new PDO('mysql:host=localhost;dbname=projet2023' ,'root', '',[PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION]);
}catch(Exception $e){
    die('Erreur : ' . $e->getMessage());
}


//e
$bdd = "INSERT INTO reservations (nom, email, date) VALUES ('$nom', '$email', '$date')";

if ($bdd->query($bdd) === TRUE) {
    echo "Réservation effectuée avec succès.";
} else {
    echo "Erreur lors de la réservation : " . $bdd->error;
}

?>
