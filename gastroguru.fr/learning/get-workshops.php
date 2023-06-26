<?php

$curl = curl_init();
curl_setopt($curl, CURLOPT_URL, 'http://localhost:8080/workshops');
curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
$response = curl_exec($curl);
$httpCode = curl_getinfo($curl, CURLINFO_HTTP_CODE);
curl_close($curl);

if ($httpCode === 200) {
    $workshops = json_decode($response, true)['workshops'];
    foreach ($workshops as $workshop) {
        echo 'ID : ' . $workshop['id'] . ', Titre : ' . $workshop['title'] . PHP_EOL;
    }
} else {
    echo 'Erreur lors de la récupération des ateliers : ' . $response;
}
