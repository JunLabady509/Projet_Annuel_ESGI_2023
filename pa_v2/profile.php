<?php session_start();
if(empty($_SESSION['email'])) {
    header('location: index.php');
    exit;
}
?>

<?php
$title = 'Profile';

?>

<body>
<?php
    include('includes/head.php');
    ?> 
    <?php
    include('includes/header.php');
    ?>
   <?php
		// Connexion à la Base de données
		try{
      $bdd = new PDO('mysql:host=localhost;dbname=projet2023' ,'root', '',[PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION]);
  }catch(Exception $e){
      die('Erreur : ' . $e->getMessage());
  }

    $q = 'SELECT * FROM clients WHERE email = :email';
    $req = $bdd->prepare($q);
    $req->execute([
    'email' => $_SESSION['email']
      ]);
    $result = $req->fetch(PDO::FETCH_ASSOC);
 

  
?>
    
    <main>


    <div class="container">
      <div class="card">   
        <div class="card-header">
          <div class="avatar">
            <div class="user-online-indicator"></div>

          
             <img src="uploads/<?=$result['image']?>" alt="image">
          </div>
          
          <div class="email"><h3>Prenom: <?= $result['prenom'] ?> </h3></div>
          <div class="email"><h3>Nom: <?= $result['nom'] ?> </h3></div>
          <div class="email"><h3>Email: <?= $result['email'] ?> </h3></div>
          <!-- <div class="follower"><h3>Nb follower: <?= $result['fo'] ?> </h3></div> -->
          <div class="bouton"><button type="submit"><a class="aa"href="edit_profil.php"> Modifier votre profile</a></div>
        
      
        </div>
      </div>
    </div>

     
</main>
</body>

</html>


