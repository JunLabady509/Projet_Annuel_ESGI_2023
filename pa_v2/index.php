
<?php session_start(); ?>
<!DOCTYPE html>
<html>
<head>
  <head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Gastro Guru</title>
    <link rel="stylesheet" type="text/css" href="CSS/index.css">
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-gH2yIJqKdNHPEq0n4Mqa/HGKIhSkIHeL5AyhkYV8i59U5AR6csBvApHHNl/vI1Bx" crossorigin="anonymous">
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-1BmE4kWBq78iYhFldvKuhfTAU6auU8tT94WrHftjDbrCEXSU1oBoqyl2QvZ6jIW3" crossorigin="anonymous">
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-ka7Sk0Gln4gmtz2MlQnikT1wXgYsOg+OMhuP+IlRH9sENBO0LRn5q+8nbTov4+1p" crossorigin="anonymous"></script>
</head>

<body>
<?php include('includes/header.php');

    $title= 'Accueil'?>

  
           
 <div class="ee">
  <h1>GastroGuru</h1>
  <p>Une cuisine inspirée de la nature, pour vous aucune limte car l’art de la cuisine est à votre porter</p>
 </div>
 <div class="text">
 <h1>GastroGuru</h1>
 <div class="ligne"></div>
  <p>Une chaîne d’espaces d’évènementiels situés à paris, dediés à la gastronomie.</p>
 </div>
  <div id="carouselExampleIndicators" class="carousel slide" data-bs-ride="carousel">
  <div class="carousel-indicators">
    <button type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide-to="0" class="active" aria-current="true" aria-label="Slide 1"></button>
    <button type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide-to="1" aria-label="Slide 2"></button>
    <button type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide-to="2" aria-label="Slide 3"></button>
  </div>

  
  <div class="carousel-inner">
    
    <div class="carousel-item active">
      <img src="images\img_pa2023\image12.jpg" class="d-block w-100 h-50" alt="image">
 
    </div>
    <div class="carousel-item">
      <img src="images\img_pa2023\img7.jpg" class="d-block w-100" alt="...">
    </div>
    <div class="carousel-item">
      <img src="images\img_pa2023\img3.jpg" class="d-block w-100" alt="...">
    
    </div>
    
  </div>
  <button class="carousel-control-prev" type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide="prev">
    <span class="carousel-control-prev-icon" aria-hidden="true"></span>
    <span class="visually-hidden">Previous</span>
  </button>
  <button class="carousel-control-next" type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide="next">
    <span class="carousel-control-next-icon" aria-hidden="true"></span>
    <span class="visually-hidden">Next</span>
  </button>
</div>
</div>
  <footer>
    <p>&copy; 2023 Mon Site Web. Tous droits réservés.</p>
  </footer>

  <script>
    // Fonction pour afficher ou masquer le menu burger
    function toggleNavbar() {
      var navbar = document.getElementById("myNavbar");
      if (navbar.className === "navbar") {
        navbar.className += " responsive";
      } else {
        navbar.className = "navbar";
      }
    }
  </script>
</body>
</html>
