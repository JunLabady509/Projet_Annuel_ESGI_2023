<?php

$workshopId = 1;
$data = [
    'title' => 'Nouveau titre'
];

$curl = curl_init();
curl_setopt($curl, CURLOPT_URL, 'http://localhost:8080/workshops/' . $workshopId);
curl_setopt($curl, CURLOPT_CUSTOMREQUEST, 'PUT');
curl_setopt($curl, CURLOPT_POSTFIELDS, json_encode($data));
curl_setopt($curl, CURLOPT_HTTPHEADER, ['Content-Type: application/json']);
curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
$response = curl_exec($curl);
$httpCode = curl_getinfo($curl, CURLINFO_HTTP_CODE);
curl_close($curl);

if ($httpCode === 200) {
    echo 'Atelier mis à jour avec succès.';
} elseif ($httpCode === 404) {
    echo 'L\'atelier avec l\'ID ' . $workshopId . ' n\'existe pas.';
} else {
    echo 'Erreur lors de la mise à jour de l\'atelier : ' . $response;
}
