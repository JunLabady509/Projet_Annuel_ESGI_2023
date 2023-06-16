<?php

// Connexion à la base de données
try{
    $bdd = new PDO('mysql:host=localhost;dbname=projet2023' ,'root', '',[PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION]);
}catch(Exception $e){
    die('Erreur : ' . $e->getMessage());
}
$bdd->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);

  if ($_SERVER['REQUEST_METHOD'] === 'POST') {

    // Récupérer les données du formulaire
    $nom = $_POST['name'];
    $email = $_POST['email'];
    $date = $_POST['date'];
    $heure = $_POST['time'];
    $nombrePersonnes = $_POST['guests'];
    // Préparer la requête d'insertion
    $requete = $bdd->prepare("INSERT INTO RESERVATIONS (nom, email, date, heure, nombre_personnes) VALUES (:nom, :email, :date, :heure, :nombre_personnes)");

    // Exécuter la requête avec les valeurs du formulaire
    $requete->execute(array(
      'nom' => $nom,
      'email' => $email,
      'date' => $date,
      'heure' => $heure,
      'nombre_personnes' => $nombrePersonnes,
    ));

    // Afficher un message de confirmation
    echo "Réservation effectuée avec succès !";
  }
