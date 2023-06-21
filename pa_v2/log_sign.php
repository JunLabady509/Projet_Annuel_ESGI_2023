<?php session_start();  ?>
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <link rel="stylesheet" type="text/css" href="CSS/log.css">
        <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-gH2yIJqKdNHPEq0n4Mqa/HGKIhSkIHeL5AyhkYV8i59U5AR6csBvApHHNl/vI1Bx" crossorigin="anonymous">
        <title>Connexion</title>
    </head>
    <body>

        <main>
         
          
<div class="background">

</div>
 <div class="rec_vert">
          <img class="fleuris">
            <img class="fleuris1">
            <h2 class="identifier">Connexion</h2>
                
          
             <?php
                if (isset($_GET['message']) && !empty($_GET['message'])) {
                    echo '<h3>' . htmlspecialchars($_GET['message']) . '</h3>';
                }

            ?>
                <form method="POST" action="verification.php">
                    <div class="form-group w-35 me-5 ms-5 pt-3">
                        <label for="exampleInputEmail1">Email</label>
                        <input type="email" class="form-control" name="email" aria-describedby="emailHelp" value="<?= isset($_COOKIE['email']) ? $_COOKIE['email'] : '' ?>">
                    </div>
                    
                    <div class="form-group w-35 me-5 ms-5">
                        <label for="exampleInputPassword1">Mot de passe</label>
                        <input type="password" class="form-control" name="mdp">
                      </div>

                      <button type="submit" style="background-color: #14FFFA; margin-left:50px; margin-top:50px ;" class="btn mt-3 ml-2">Se connecter</button>

                    </form>
                    
         

            <div class="sign_in"><p>Appuyez</p>
            <a href="sign.php">ici</a>
            <p>pour vous s'inscrire</p>
            </div>
             <div class="log_admin "><p>Connectez vous en tant qu'Administrateur</p>
            <a href="admin\connexion.php">ici</a>
                       <br>
            </div>
  </div> 
               
        </main>
    
        </body>


</html>

</html>
