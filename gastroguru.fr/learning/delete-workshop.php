<?php

$workshopId = 1;
$curl = curl_init();
curl_setopt($curl, CURLOPT_URL, 'http://localhost:8080/workshops/' . $workshopId);
curl_setopt($curl, CURLOPT_CUSTOMREQUEST, 'DELETE');
curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
$response = curl_exec($curl);
$httpCode = curl_getinfo($curl, CURLINFO_HTTP_CODE);
curl_close($curl);

if ($httpCode === 200) {
    echo 'Atelier supprimé avec succès.';
} elseif ($httpCode === 404) {
    echo 'L\'atelier avec l\'ID ' . $workshopId . ' n\'existe pas.';
} else {
    echo 'Erreur lors de la suppression de l\'atelier : ' . $response;
}
