<?php

$apiUrl = 'http://localhost:44446/events';

// Données de l'événement à créer
$data = array(
    'title' => 'Nouvel événement',
    'description' => 'Description de l\'événement',
    'date' => '2023-06-28T10:00:00Z',
    'duration' => 120,
    'location' => 'Lieu de l\'événement',
    'capacity' => 50,
    'price' => 19.99,
    'image' => 'image.jpg',
    'type' => 'Type d\'événement'
);

// Convertir les données en JSON
$jsonData = json_encode($data);

// Configuration de la requête curl
$ch = curl_init($apiUrl);
curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_POSTFIELDS, $jsonData);
curl_setopt($ch, CURLOPT_HTTPHEADER, array('Content-Type: application/json'));
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);

// Exécuter la requête
$response = curl_exec($ch);

// Vérifier les erreurs
if(curl_errno($ch)) {
    echo 'Erreur curl : ' . curl_error($ch);
}

// Fermer la requête curl
curl_close($ch);

// Afficher la réponse
echo $response;
