
<?php

// Définir les informations de la leçon en ligne
$onlineLessonData = array(
    'title' => 'Titre de la leçon en ligne',
    'description' => 'Description de la leçon en ligne',
    'price' => 29.99,
    'start_time' => '2023-06-30 10:00:00',
    'end_time' => '2023-06-30 12:00:00',
    'instructor_id' => 1,
    'video_link' => array(
        'https://www.youtube.com/watch?v=video1',
        'https://www.youtube.com/watch?v=video2',
        'https://www.youtube.com/watch?v=video3'
    ),
    'uploaded_time' => '2023-06-26 15:30:00',
    'insight' => 'Informations supplémentaires sur la leçon'
);

// Convertir les données en JSON
$onlineLessonJson = json_encode($onlineLessonData);

// Créer une requête POST pour créer une leçon en ligne
$ch = curl_init('http://localhost:44446/onlinelessons');
curl_setopt($ch, CURLOPT_POST, 1);
curl_setopt($ch, CURLOPT_POSTFIELDS, $onlineLessonJson);
curl_setopt($ch, CURLOPT_HTTPHEADER, array('Content-Type: application/json'));
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);

// Exécuter la requête
$response = curl_exec($ch);

// Vérifier si la requête a réussi
if ($response === false) {
    echo 'Erreur de la requête : ' . curl_error($ch);
} else {
    // Obtenir le code de réponse HTTP
    $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);

    // Traiter la réponse en fonction du code de réponse
    if ($httpCode == 201) {
        echo 'La leçon en ligne a été créée avec succès.';
    } else {
        echo 'Erreur lors de la création de la leçon en ligne : ' . $response;
    }
}

// Fermer la requête
curl_close($ch);
