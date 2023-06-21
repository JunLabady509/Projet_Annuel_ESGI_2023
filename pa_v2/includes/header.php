<header>
  <link rel="stylesheet" type="text/css" href="CSS/headers.css">

  <img class="logo" src="images\img_pa2023\logoG.png" alt="icon">
    <nav class="navbar" id="myNavbar">
      <ul>
     
        <li class="active"><a href="index.php">Accueil</a></li>
            <?php 
        if(isset($_SESSION['email']) || !empty($_SESSION['email'])){
            
            echo ' <li class="reservation"><a href="reservation.php">Reservation</a></li>';
            echo ' <li class="catalogue"><a href="catalogue.php">Catalogue</a></li>';
            echo '<li class="profile"><a href="profile.php">Profile</a></li>';
            echo '<li class="boutique"><a href="boutique.php">Boutique</a></li>';
            echo '<li class="log_out"><a href="deconnexion.php">Déconnexion</a>';
        }else{
         
            echo ' <li  class="log"><a href="log_sign.php">Connexion</a></li>';
            echo ' <li  class="sign"><a href="sign.php">Inscription</a></li>';
             
        }
        ?>
      </ul>
    </nav>
    <div class="burger-menu" onclick="toggleNavbar()">&#9776;</div>
  </header>