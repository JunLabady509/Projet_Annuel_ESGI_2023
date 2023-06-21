<?php 
session_start();
try{
    $bdd = new PDO('mysql:host=localhost;dbname=kano' ,'root', 'esgi',[PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION]);
}catch(Exception $e){
    die('Erreur : ' . $e->getMessage());
}
if(isset($_POST['email']) && !empty($_POST['email'])){
	setcookie('email', $_POST['email'],time() + 24 * 3600);
}
if(
	!isset($_POST['email'])
	|| empty($_POST['email'])
	|| !isset($_POST['mdp'])
	|| empty($_POST['mdp'])
	|| !isset($_POST['nom'])
	|| empty($_POST['nom'])
	|| !isset($_POST['prenom'])
	|| empty($_POST['prenom'])

){
	header('location:sing.php?message=vous devez remplir les 5 champs.&type=danger');
	exit; 
}

if(!filter_var($_POST['email'], FILTER_VALIDATE_EMAIL)) 
{
	header('location:sign.php?message=Email invalide.&type=danger');
	exit;
}

$q = 'SELECT id_clients FROM CLIENTS WHERE email = :email';
$req = $bdd->prepare($q);
$req->execute(['email' => $_POST['email']] );


$results = $req->fetchall();

if(!empty($results)){
	header('location:sign.php?message=Adresse email déjà utilisée.&type=danger');
	exit;
}


$test = $_POST['mdp'];
$pattern = "/^[^0-9][A-Z][a-z]|[0-9]/";

$res = preg_match($pattern, $test);


if(strlen($_POST['mdp']) < 8 || $res == 0){
	header('location:sign.php?message=Le mot de passe doit comporter entre 8 caracteres, 1 Maj, 1 Min et 1 Chiffre.&type=danger');
	exit;
}

if ($_FILES['image']['error'] == 0) {

    $acceptable = ['image/jpeg', 'image/png'];

    if (!in_array($_FILES['image']['type'], $acceptable)) {
        header('location:sign.php?message=Format d\'image invalide (jpeg, png).&type=alert');
        exit;
    }

    $max_size = 2 * 1024 * 1024;
    if ($_FILES['image']['size'] > $max_size) {
        header('location:sign.php?message=Image trop volumineux (2Mo max).&type=alert');
        exit;
    }

    $chemin = 'uploads';
    if (!file_exists($chemin)) {
        mkdir($chemin);
    }

    $array = explode('.', $_FILES['image']['name']);
    $extension = end($array);
    $file_name = 'image-' . time() . '.' . $extension;
    $destination = 'uploads/' . $file_name;
    move_uploaded_file($_FILES['image']['tmp_name'], $destination);
}



$sq = 'SELECT MAX(id_clients) FROM CLIENTS';
$query = $bdd->prepare($sq);
$query->execute();
$maxId = $query->fetch();
$count = $maxId['MAX(id_clients)'] + 1;
$q = 'INSERT INTO CLIENTS (id_clients, email, nom, prenom, image, mdp) VALUES (:id_clients, :email, :nom, :prenom, :image, :mdp )';
$req = $bdd->prepare($q);
$succes = $req->execute([
						'id_clients'=>$count,
						'email' => $_POST['email'],
						'mdp' => hash('sha512',$_POST['mdp']),
						'prenom' => $_POST['prenom'],
						'nom' => $_POST['nom'],
						'image' => isset($file_name) ? $file_name : 'default.png'
						]);


if (!$succes){
	header('location:sign.php?message=Erreur lors de l enregistrement.&type=danger');
}

else{
	header('location:index.php?message=Votre compte a été creer avec succes.&type=success');
}
?>


