<?php
session_start();
try{
  $bdd = new PDO('mysql:host=localhost;dbname=projet2023','root', '',[PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION]);
}catch(Exception $e){
  die('Erreur : ' . $e->getMessage());
}
 
	if(isset($_POST['email']) && !empty($_POST['email'])){
    	setcookie('email', $_POST['email'], time() + 24 * 3600);
    }

    if (
    	!isset($_POST['email'])
    	|| empty($_POST['email'])
    	|| !isset($_POST['mdp'])
    	|| empty($_POST['mdp'])
    ) {
    	header('location:log_sign.php?message=Vous devez remplir les 2 champs.');
      header('location:log_sign.php?email=Votre email est invalide.');
      exit;
    }

   if (!filter_var($_POST['email'], FILTER_VALIDATE_EMAIL)){
      header('location:log_sign.php?messageEmail invalide.');
      exit;
   }
   
$q = 'SELECT id_clients FROM CLIENTS WHERE email = :email AND mdp = :mdp';
$req = $bdd->prepare($q);
$req->execute([
      'email' => $_POST['email'],
      'mdp' => hash('sha512', $_POST['mdp'])
      ]);

$results = $req->fetchAll();
if(empty($results)){

  // function writeLine(false, $_POST['email']);

  header('location:log_sign.php?message=Identifiants incorrects.&type=danger');
  exit;
}

$q = 'SELECT id_clients FROM CLIENTS WHERE email = :email';
$req = $bdd->prepare($q);
$req->execute([
      'email' => $_POST['email']
      ]);
$results = $req->fetchAll(PDO::FETCH_ASSOC);

session_start();

$_SESSION['email'] = $_POST['email'];
$_SESSION['id_clients'] = $results['id_clients'];
// function writeLine(true, $_POST['email']);

header('location: index.php?message=Content de vous revoir !&type=succes');
exit;


?>