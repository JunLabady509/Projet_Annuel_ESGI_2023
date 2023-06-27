<?php

$apiUrl = 'http://localhost:8080/homecourses/1'; // Remplacez "1" par l'ID du cours à domicile que vous souhaitez supprimer

// Envoi de la requête DELETE pour supprimer le cours à domicile
$ch = curl_init($apiUrl);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_CUSTOMREQUEST, 'DELETE');
$response = curl_exec($ch);

// Vérification de la réponse
if ($response === false) {
    echo 'Erreur lors de la suppression du cours à domicile : ' . curl_error($ch);
} else {
    $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
    if ($httpCode === 200) {
        echo 'Le cours à domicile a été supprimé avec succès.';
    } else {
        echo 'Erreur lors de la suppression du cours à domicile. Code HTTP : ' . $httpCode;
    }
}

curl_close($ch);
?>
