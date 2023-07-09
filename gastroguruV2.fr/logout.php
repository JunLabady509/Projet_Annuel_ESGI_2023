<?php
// Début de session
session_start();

// Détruire toutes les variables de session
$_SESSION = array();

// Supprimer le cookie de session
if (isset($_COOKIE[session_name()])) {
    setcookie(session_name(), '', time() - 42000, '/');
}

// Détruire la session
session_destroy();

// Rediriger l'utilisateur vers une page de déconnexion réussie ou une autre page de votre choix
header("Location: index.php");
exit();
?>
