<?php
session_start();
if (isset($_POST['submit'])) {
    $name = $_POST['name'];
    $first_name = $_POST['first_name'];
    $address = $_POST['address'];
    $email = $_POST['email'];
    $password = $_POST['password'];
    $confirm_password = $_POST['confirm_password'];
    $phone = $_POST['phone'];

    $url = "http://localhost:44446/users";
    $ch = curl_init($url);

    if ($password != $confirm_password) {
        // Les mots de passe ne correspondent pas
        $_SESSION['error_message'] = "Les mots de passe ne correspondent pas.";
        header('Location: register.php');
        exit();
    }

    $data = array(
        'name' => $name,
        'first_name' => $first_name,
        'address' => $address,
        'email' => $email,
        'password' => $password,
        'phone' => $phone,
    );

    // Vérifier
    $payload = json_encode($data);

    curl_setopt($ch, CURLOPT_POSTFIELDS, $payload);
    curl_setopt($ch, CURLOPT_HTTPHEADER, array('Content-Type:application/json'));
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);

    $result = curl_exec($ch);
    $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
    curl_close($ch);

    // Vous pouvez maintenant gérer la réponse ici.
    // Par exemple, rediriger l'utilisateur ou afficher un message.

    if ($httpCode == 201) {
      // Rediriger l'utilisateur vers la page de connexion
      header('Location: login.html');
      exit();
    } else {
        // Afficher un message d'erreur
        if (!isset($_SESSION['error_message'])) {
            $_SESSION['error_message'] = "Une erreur s'est produite lors de l'inscription.";
        }
        header('Location: register.php');
        exit();
    }
}
?>
