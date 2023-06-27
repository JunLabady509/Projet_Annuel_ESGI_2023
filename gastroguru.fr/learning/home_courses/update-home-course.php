<?php

$apiUrl = 'http://localhost:8080/homecourses/1'; // Remplacez "1" par l'ID du cours à domicile que vous souhaitez mettre à jour

// Données du cours à domicile à mettre à jour
$data = [
    'title' => 'Nouveau titre du cours à domicile',
    'description' => 'Nouvelle description du cours à domicile',
    'price' => 59.99,
    'start_time' => '2023-06-26T10:00:00Z',
    'end_time' => '2023-06-26T12:00:00Z',
    'instructor_id' => 2,
    'duration' => '2h',
    'address' => 'Nouvelle adresse du cours à domicile'
];

// Envoi de la requête PUT pour mettre à jour le cours à domicile
$ch = curl_init($apiUrl);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_CUSTOMREQUEST, 'PUT');
curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($data));
curl_setopt($ch, CURLOPT_HTTPHEADER, ['Content-Type: application/json']);
$response = curl_exec($ch);

// Vérification de la réponse
if ($response === false) {
    echo 'Erreur lors de la mise à jour du cours à domicile : ' . curl_error($ch);
} else {
    $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
    if ($httpCode === 200) {
        echo 'Le cours à domicile a été mis à jour avec succès.';
    } else {
        echo 'Erreur lors de la mise à jour du cours à domicile. Code HTTP : ' . $httpCode;
    }
}

curl_close($ch);
?>
