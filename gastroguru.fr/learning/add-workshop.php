<?php
session_start();

$data = [
    'title' => 'Atelier de cuisine',
    'description' => 'Apprenez les bases de la cuisine française',
    'price' => 49.99,
    'start_time' => '2023-07-01 09:00:00',
    'end_time' => '2023-07-01 12:00:00',
    'instructor_id' => 1234,
    'capacity' => 20,
    'place' => 'Cuisine Academy'
];

$curl = curl_init();
curl_setopt($curl, CURLOPT_URL, 'http://localhost:44446/workshops');
curl_setopt($curl, CURLOPT_POST, true);
curl_setopt($curl, CURLOPT_POSTFIELDS, json_encode($data));
curl_setopt($curl, CURLOPT_HTTPHEADER, ['Content-Type: application/json']);
curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
$response = curl_exec($curl);
$httpCode = curl_getinfo($curl, CURLINFO_HTTP_CODE);
curl_close($curl);

if ($httpCode === 201) {
    echo 'Atelier créé avec succès.';
    header('Location: ../courses.php');

} else {
    echo 'Erreur lors de la création de l\'atelier : ' . $response;
}
