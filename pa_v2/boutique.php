<?php session_start();?>
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <link rel="stylesheet" type="text/css" href="CSS/boutique.css">
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-gH2yIJqKdNHPEq0n4Mqa/HGKIhSkIHeL5AyhkYV8i59U5AR6csBvApHHNl/vI1Bx" crossorigin="anonymous">
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-1BmE4kWBq78iYhFldvKuhfTAU6auU8tT94WrHftjDbrCEXSU1oBoqyl2QvZ6jIW3" crossorigin="anonymous">
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-ka7Sk0Gln4gmtz2MlQnikT1wXgYsOg+OMhuP+IlRH9sENBO0LRn5q+8nbTov4+1p" crossorigin="anonymous"></script>
<script src="javascript.js"></script>
  <title>Boutique</title>
  </head>
<?php $title = 'Boutique';

?>

<body>
<?php
    include('includes/head.php');
    ?> 
    <?php
    include('includes/header.php');
    ?>

<h1>Achat d'articles</h1>

<div id="catalogue">
    <h2>Catalogue d'articles</h2>
    <ul id="liste-articles">
        <!-- Les articles seront générés dynamiquement avec JavaScript -->
   
    </ul>
</div>

<div id="panier">
    <h2>Panier</h2>
    <ul id="liste-panier">
        <!-- Les articles ajoutés au panier seront générés dynamiquement avec JavaScript -->
    </ul>
    <button id="btn-acheter">Acheter</button>
</div>


    </body>
    </html>