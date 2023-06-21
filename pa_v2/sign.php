<?php session_start();  ?>
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <link rel="stylesheet" type="text/css" href="CSS/Sign.css">
        <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-gH2yIJqKdNHPEq0n4Mqa/HGKIhSkIHeL5AyhkYV8i59U5AR6csBvApHHNl/vI1Bx" crossorigin="anonymous">
        <title>inscription</title>
    </head>
    <body>


        <main>
         
        <?php include('includes/message.php');?>
<div class="background">

</div>
 <div class="rec_vert">
          <img class="fleuris">
            <img class="fleuris1">
            <h2 class="identifier">S'inscrire</h2>
                
          
               <?php
                    if (isset($_GET['message']) && !empty($_GET['message'])) {
                    echo '<h3 style="font-size: 15px; color: red;" class="text-center">' . htmlspecialchars($_GET['message']) . '</h3>';
                    }

                ?>

                    <form method="POST" action="verification_inscription.php" enctype="multipart/form-data">

                        <div class="form-group w-75 me-5 ms-5">
                            <label for="exampleInputEmail1">Prénom</label>
                            <input type="text" class="form-control" name="prenom" aria-describedby="emailHelp" required>
                        </div>
                            <div class="form-group w-75 me-5 ms-5">
                            <label for="exampleInputEmail1">Nom</label>
                            <input type="text" class="form-control " name="nom" aria-describedby="emailHelp" required>
                        </div>
                        <div class="form-group w-75 me-5 ms-5">
                            <label for="exampleInputEmail1">Email</label>
                            <input type="email" class="form-control" name="email" aria-describedby="emailHelp" required>
                        </div>
                                         
                          <div class="form-group w-75 me-5 ms-5">
                            <label for="exampleInputPassword1">Mot de passe</label>
                            <input type="password" class="form-control" name="mdp" required>
                          </div>
                          <div class="form-group w-75 me-5 ms-5">
                            <label for="exampleInputPassword1">Confirmez votre Mot de passe</label>
                            <input type="password" class="form-control" name="password_confirmation" required>
                          </div>

                          <div class="form-group w-75 me-5 ms-5">
                            <label class="image-input-label"> Ajoutez une photo de profile: </label>
                            <input class="image-input-button" type="file" name="image" accept="image/png, image/jpeg">
                        </div>
                          <button type="submit" name="submit" style="background-color: #14FFFA"; class="btn mt-3"> S'inscrire</button>

                        </form>
                        <p> Si vous avez un compte, appuyez sur</p>
                    <div class="sign_in">
                             <a href="log_sign.php"> <br> connectez-vous</a>
                    </div>
                        
           
                </div>
              

                

        </main>
    
        </body>


</html>
