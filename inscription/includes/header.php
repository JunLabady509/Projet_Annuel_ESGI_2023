<header>
<link rel="stylesheet" type="text/css" href="CSS/headers.css">
    <nav class="navbar">
       

        <div class="nav-links">
        <img class="logo" src="images\img_pa2023\logoG.png" alt="icon">
           
            <ul>
            <li class="active"><a href="index.php">Accueil</a></li>
            <?php 
        if(isset($_SESSION['email']) || !empty($_SESSION['email'])){
            
            echo ' <li class="camp"><a href="camping.php">Le Camping</a></li>';
            echo ' <li class="heb"><a href="heb.php">Hébergement</a></li>';
            echo '<li class="profile"><a href="profile.php">Profile</a></li>';
            echo '<li class="log_out"><a href="deconnexion.php">Déconnexion</a>';
        }else{
         
            echo ' <li  class="log"><a href="log_sign.php"> Se connecter</a></li>';
            echo ' <li  class="sign"><a href="sign.php"> S\'inscrire</a></li>';
             
        }
        ?>
         
           
         
        </div>

        <img src="images/menu-btn.jpg" alt="menu hamburger" class="menu-hamburger">
        
      
        
       

    </ul>
    </nav>

</header>

        
          
</header>