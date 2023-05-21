<?php
session_start();

// Connexion à la base de données (remplacez les valeurs par les vôtres)

try{
    $bdd = new PDO('mysql:host=localhost;dbname=projet2023' ,'root', '',[PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION]);
}catch(Exception $e){
    die('Erreur : ' . $e->getMessage());
}


// Vérification de la connexion
if ($bdd->connect_error) {
    die("Connexion échouée : " . $conn->connect_error);
}

// Récupération des données du formulaire
$name = $_POST['name'];
$email = $_POST['email'];
$date = $_POST['date'];
$time = $_POST['time'];
$guests = $_POST['guests'];

// Requête d'insertion des données dans la table "reservations" (adaptée à votre structure de base de données)
$sql = "INSERT INTO reservations (name, email, date, time, guests) VALUES ('$name', '$email', '$date', '$time', '$guests')";

if ($bdd->query($sql) === TRUE) {
    echo "Réservation effectuée avec succès !";
} else {
    echo "Erreur lors de la réservation : " . $bdd->error;
}

$bdd->close();
?>
