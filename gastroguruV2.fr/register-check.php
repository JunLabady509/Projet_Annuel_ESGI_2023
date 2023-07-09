<?php
session_start();

if (isset($_POST['submit'])) {
    $uploadDir = 'uploads/';

    if ($_FILES['profile_photo']['error'] === UPLOAD_ERR_OK) {
        $tmpName = $_FILES['profile_photo']['tmp_name'];
        $fileName = $_FILES['profile_photo']['name'];
        $filePath = $uploadDir . $fileName;

        if (move_uploaded_file($tmpName, $filePath)) {
            $message = 'Image uploadée avec succès.';
        } else {
            $message = 'Erreur lors de l\'upload.';
        }
    }

    $url = "http://localhost:44446/users";
    $ch = curl_init($url);

    $last_name = $_POST['last_name'];
    $first_name = $_POST['first_name'];
    $address = $_POST['address'];
    $email = $_POST['email'];
    $password = $_POST['password'];
    $confirm_password = $_POST['confirm_password'];
    $phone = $_POST['phone'];

    if ($password != $confirm_password) {
        // Les mots de passe ne correspondent pas
        $_SESSION['error_message'] = "Les mots de passe ne correspondent pas.";
        header('Location: register.php');
        exit();
    }

    $data = array(
        'last_name' => $last_name,
        'first_name' => $first_name,
        'address' => $address,
        'email' => $email,
        'password' => $password,
        'phone' => $phone,
        'picture' => $filePath,
    );

    $payload = json_encode($data);

    curl_setopt($ch, CURLOPT_POSTFIELDS, $payload);
    curl_setopt($ch, CURLOPT_HTTPHEADER, array('Content-Type:application/json'));
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);

    $result = curl_exec($ch);
    $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
    curl_close($ch);

    //Redirection de l'utilisateur vers la page de connexion
    
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
