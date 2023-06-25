<?php
session_start();

$url = 'http://localhost:44446/auth'; // URL de votre point de terminaison d'authentification
$email = $_POST['email']; // Récupérer l'adresse e-mail depuis le formulaire de connexion
$password = $_POST['password']; // Récupérer le mot de passe depuis le formulaire de connexion

$data = http_build_query(array(
  'email' => $email,
  'password' => $password
));

$options = array(
  'http' => array(
      'method' => 'POST',
      'header' => 'Content-type: application/x-www-form-urlencoded',
      'content' => $data
  )
);

$context = stream_context_create($options);
$response = file_get_contents($url, false, $context);

if ($response === false) {
    // Gérer les erreurs de requête
    echo "Une erreur s'est produite lors de la requête d'authentification.";
    exit();
} else {
    $responseData = json_decode($response, true);

    // Vérifier si le token d'accès est présent dans la réponse
    if (isset($responseData['token'])) {
        $accessToken = $responseData['token'];

        // Stocker le token d'accès dans une session pour maintenir la connexion
        $_SESSION['access_token'] = $accessToken;

        // Rediriger vers la page principale ou une autre page sécurisée
        header('Location: index.html');
        exit();
    } else {
        // Gérer les erreurs d'authentification
        echo "Adresse e-mail ou mot de passe incorrect.";
        exit();
    }
}
?>
