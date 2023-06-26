<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>Inscription</title>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link href="https://fonts.googleapis.com/css?family=Source+Sans+Pro:300,400,600,700" rel="stylesheet">
  <link href="https://fonts.googleapis.com/css?family=Roboto+Slab:300,400" rel="stylesheet">
  <link rel="stylesheet" href="css/animate.css">
  <link rel="stylesheet" href="css/icomoon.css">
  <link rel="stylesheet" href="css/bootstrap.css">
  <link rel="stylesheet" href="css/style.css">

  <style>

    body {
      display: flex;
      align-items: center;
      justify-content: center;
      height: 100vh;
      margin: 0;
      background-color: #f1f1f1;
    }

    .login-container {
      background-color: #fff;
      padding: 20px;
      border-radius: 5px;
      box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
      display: flex;
      flex-direction: column;
      align-items: center;
    }

    .logo {
      max-width: 150px; /* Ajustez la taille maximale selon vos besoins */
      margin-bottom: 20px; /* Espacement en bas du logo */
    }

    .login-form {
      text-align: center;
      width: 300px; /* Ajustez la largeur du formulaire */
    }

    .login-form h1 {
      font-family: 'Roboto Slab', serif;
      font-size: 24px;
      margin-bottom: 20px;
    }

    .login-form input[type="text"],
    .login-form input[type="password"] {
      width: 100%;
      padding: 10px;
      margin-bottom: 15px;
      border: 1px solid #ccc;
      border-radius: 4px;
    }

    .login-form button[type="submit"] {
      background-color: #333;
      color: #fff;
      padding: 10px 20px;
      border: none;
      border-radius: 4px;
      cursor: pointer;
    }

    .signup-link {
      margin-top: 15px;
    }

  </style>

</head>

<body>
  <div class="login-container">
    <a href="index.php">
      <img src="images/logoG.png" alt="Logo" class="logo">
    </a>
    <div class="login-form">
      <h1>Inscription</h1>
      <form method="POST" action="register-check.php" id="myForm" enctype="multipart/form-data">
        <input type="text" class="form-control" name="last_name" id="last_name" placeholder="Nom" required>
        <input type="text" class="form-control" name="first_name" id="first_name" placeholder="Prénom" required>
        <input type="text" class="form-control" name="address" id="address" placeholder="Adresse" required>
        <input type="email" class="form-control" name="email" id="email" placeholder="Email" required>
        <small id="emailError" style="display:none; color:red;">Email déjà utilisé</small>
        <input type="password" class="form-control" name="password" id="password" placeholder="Mot de passe" required>
        <input type="password" class="form-control" name="confirm_password" id="confirm_password" placeholder="Confirmez votre mot de passe" required>
        <small id="passwordError" style="display:none; color:red;">Les mots de passe ne correspondent pas</small>
        <input type="tel" class="form-control" name="phone" id="phone" placeholder="Téléphone" required>
        <label for="profile_photo">Ajouter une photo de profil</label>      
        <input type="file" class="form-control" name="profile_photo" id="profile_photo" accept="image/*">
        <button type="submit" name="submit" id="submitButton" class="btn btn-primary" disabled>S'inscrire</button>
      </form>
      <div class="signup-link">
        Vous avez déjà un compte ? <a href="login.html">Se connecter</a>
      </div>
    </div>
  </div>

  <!-- Vérification Existence de l'adresse mail côté utilisateur -->
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
  <script>
    $(document).ready(function() {
      let timeoutId;
      $('#email').on('input', function(e) {
        clearTimeout(timeoutId);
        timeoutId = setTimeout(function() {
          const email = $('#email').val();
          if (email !== '') {
            $.ajax({
              type: 'POST',
              url: 'http://localhost:44446/users/checkEmail',
              data: {
                'email': email
              },
              dataType: 'json'
            }).done(function(data) {
              $('#email').removeClass('is-invalid');
              $('#submitButton').removeAttr('disabled');
              $('#emailError').text('').hide();
            }).fail(function(jqXHR, textStatus, errorThrown) {
              if (jqXHR.status === 409) {
                $('#email').addClass('is-invalid');
                $('#submitButton').attr('disabled', 'disabled');
                $('#emailError').text('Email déjà utilisé').show();
              } else {
                console.log(textStatus, errorThrown);
              }
            }).always(function() {
              validatePassword();
            });
          } else {
            $('#email').removeClass('is-invalid');
            $('#submitButton').removeAttr('disabled');
            $('#emailError').text('').hide();
            validatePassword();
          }
        }, 4500);
      });

      $('#myForm').submit(function(e) {
        if ($('#submitButton').prop('disabled')) {
          e.preventDefault(); // Empêcher l'envoi du formulaire si le bouton de soumission est désactivé
          alert("Les mots de passe ne correspondent pas ou l'e-mail est déjà utilisé.");
        }
      });
    });

    // Confirmation du mot de passe côté utilisateur
    var password = document.getElementById("password");
    var confirm_password = document.getElementById("confirm_password");
    var submitButton = document.getElementById("submitButton");
    var passwordError = document.getElementById("passwordError");

    function validatePassword() {
      if (password.value !== confirm_password.value) {
        confirm_password.classList.add("error");
        passwordError.style.display = "block";
        submitButton.disabled = true;
      } else {
        confirm_password.classList.remove("error");
        passwordError.style.display = "none";
        submitButton.disabled = false;
      }
    }

    password.addEventListener('input', validatePassword);
    confirm_password.addEventListener('input', validatePassword);
  </script>

  <script src="js/modernizr-2.6.2.min.js"></script>
  <script src="js/respond.min.js"></script>
</body>

</html>
