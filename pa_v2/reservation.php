<?php session_start(); ?>
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  
  <link rel="stylesheet" type="text/css" href="CSS/reservation.css">
  <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-gH2yIJqKdNHPEq0n4Mqa/HGKIhSkIHeL5AyhkYV8i59U5AR6csBvApHHNl/vI1Bx" crossorigin="anonymous">
  <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-1BmE4kWBq78iYhFldvKuhfTAU6auU8tT94WrHftjDbrCEXSU1oBoqyl2QvZ6jIW3" crossorigin="anonymous">
  
  <script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
  <script src="https://code.jquery.com/ui/1.12.1/jquery-ui.js"></script>
  
  <title>Réservation</title>
  
  <link rel="stylesheet" href="https://code.jquery.com/ui/1.12.1/themes/base/jquery-ui.css">
  
  <script>
    $(function() {
      $("#date").datepicker({
        dateFormat: 'yy-mm-dd',
        minDate: 0 // Limite la date de réservation à partir d'aujourd'hui
      });
    });
  </script>
</head>
<body>
<?php include('includes/header.php');
$title = 'Accueil'?>

<div class="ee">
  <h1>GastroGuru</h1>
  <p>Une cuisine inspirée de la nature, pour vous aucune limite car l’art de la cuisine est à votre portée</p>
</div>

<div class="text">
  <h1>Reservation</h1>
  <div class="ligne"></div>
  <p>Reservez-Maintenant</p>
</div>

<form method="POST" action="traitement_reservation.php">
  <label for="name">Nom :</label>
  <input type="text" id="name" name="name" required>

  <label for="email">Email :</label>
  <input type="email" id="email" name="email" required>

  <label for="date">Date :</label>
  <input type="text" id="date" name="date" required>

  <label for="time">Heure :</label>
  <input type="time" id="time" name="time" required>

  <label for="guests">Nombre de personnes :</label>
  <select id="guests" name="guests" required>
    <option value="">Sélectionnez</option>
    <option value="1">1 personne</option>
    <option value="2">2 personnes</option>
    <option value="3">3 personnes</option>
    <option value="4">4 personnes</option>
  </select>

  <input type="submit" value="Réserver">
</form>

<footer>
  <p>&copy; 2023 Mon Site Web. Tous droits réservés.</p>
</footer>

</body>
</html>
