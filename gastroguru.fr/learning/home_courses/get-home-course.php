<?php

$apiUrl = 'http://localhost:8080/homecourses/1'; // Remplacez "1" par l'ID du cours à domicile que vous souhaitez récupérer

// Envoi de la requête GET pour récupérer le cours à domicile
$ch = curl_init($apiUrl);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
$response = curl_exec($ch);

// Vérification de la réponse
if ($response === false) {
    echo 'Erreur lors de la récupération du cours à domicile : ' . curl_error($ch);
} else {
    $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
    if ($httpCode === 200) {
        $homeCourse = json_decode($response, true);
        echo 'ID du cours à domicile : ' . $homeCourse['homecourse']['id'] . '<br>';
        echo 'Titre du cours à domicile : ' . $homeCourse['homecourse']['title'] . '<br>';
        // Autres informations du cours à domicile...
    } else {
        echo 'Erreur lors de la récupération du cours à domicile. Code HTTP : ' . $httpCode;
    }
}

curl_close($ch);
?>
