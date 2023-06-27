<?php

$apiUrl = 'http://localhost:44446/homecourses';

// Données du cours à domicile à créer
$data = [
    'title' => 'Nouveau cours à domicile',
    'description' => 'Description du nouveau cours à domicile',
    'price' => 49.99,
    'start_time' => '2023-06-26 09:00:00',
    'end_time' => '2023-06-26 11:00:00',
    'instructor_id' => 1,
    'duration' => 2,
    'address' => 'Adresse du nouveau cours à domicile'
];

// Envoi de la requête POST pour créer le cours à domicile
$ch = curl_init($apiUrl);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($data));
curl_setopt($ch, CURLOPT_HTTPHEADER, ['Content-Type: application/json']);
$response = curl_exec($ch);

// Vérification de la réponse
if ($response === false) {
    echo 'Erreur lors de la création du cours à domicile : ' . curl_error($ch);
} else {
    $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
    if ($httpCode === 201) {
        echo 'Le cours à domicile a été créé avec succès.';
    } else {
        echo 'Erreur lors de la création du cours à domicile. Code HTTP : ' . $httpCode;
    }
}

curl_close($ch);
?>
