<?php session_start(); ?>
<!DOCTYPE html>
<html>
<head>
<head>    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <link rel="stylesheet" type="text/css" href="CSS/reservation.css">
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-gH2yIJqKdNHPEq0n4Mqa/HGKIhSkIHeL5AyhkYV8i59U5AR6csBvApHHNl/vI1Bx" crossorigin="anonymous">
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-1BmE4kWBq78iYhFldvKuhfTAU6auU8tT94WrHftjDbrCEXSU1oBoqyl2QvZ6jIW3" crossorigin="anonymous">
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-ka7Sk0Gln4gmtz2MlQnikT1wXgYsOg+OMhuP+IlRH9sENBO0LRn5q+8nbTov4+1p" crossorigin="anonymous"></script>
<title>Réservation</title>
<link rel="stylesheet" href="https://code.jquery.com/ui/1.12.1/themes/base/jquery-ui.css">
  <script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
  <script src="https://code.jquery.com/ui/1.12.1/jquery-ui.js"></script>

  <title>Catalogue</title>
  <link rel="stylesheet" type="text/css" href="CSS/catalogue.css">
  </head>
<body>
<?php include('includes/header.php');

$title= 'Accueil'?>
 <div class="ee">
  <h1>GastroGuru</h1>
  <p>Une cuisine inspirée de la nature, pour vous aucune limte car l’art de la cuisine est à votre porter</p>
 </div>

<div class="text">
<h1>Catalogue</h1>
<div class="ligne"></div>
<p>Nous disposons des espaces evènementiels equipés pour tous vos besoins: Mariage, Anniversaires, dîner. </p></div>

  <div class="product">
    <img src="images\img_pa2023\image1.webp" alt="Produit 1">
    <h2>NOS ESPACES </h2>
    <p>Description du produit 1.</p>
    <button>Acheter</button>
  </div>

  <div class="product">
    <img src="images\img_pa2023\TOO_restaurant.webp" alt="Produit 2">
    <h2>NOS PRESTATIONS</h2>
    <p>Description du produit 2.</p>
    <button>Acheter</button>
  </div>

  <div class="product">
    <img src="images\img_pa2023\show_detail.jpg" alt="Produit 3">
    <h2>NOS EVENEMENT</h2>
    <p>Description du produit 3.</p>
    <button>Acheter</button>
  </div>

  <!-- Ajoutez d'autres éléments de catalogue selon vos besoins -->

</body>
</html>
