<style>
  .profile-picture {
    width: 40px;
    /* Ajustez la taille de l'image selon vos besoins */
    height: 40px;
    border-radius: 50%;
  }
</style>


<nav class="fh5co-nav" role="navigation">
  <div class="top">
    <div class="container">
      <div class="row">
        <div class="col-xs-12 text-right">
          <p class="site">gastroguru.fr</p>
          <p class="num">Appelez au: +33 7 49 84 49 90</p>
          <ul class="fh5co-social">
            <li><a href="#"><i class="icon-facebook2"></i></a></li>
            <li><a href="#"><i class="icon-twitter2"></i></a></li>
            <li><a href="#"><i class="icon-dribbble2"></i></a></li>
            <li><a href="#"><i class="icon-github"></i></a></li>
          </ul>
        </div>
      </div>
    </div>
  </div>

  <div class="top-menu">
    <div class="container">
      <div class="row">
        <div class="col-xs-2">
          <div id="fh5co-logo">
            <a href="index.php">
              <i class="icon-food"></i>Gastroguru<span>.</span></a>
          </div>
        </div>
        <div class="col-xs-10 text-right menu-1">
          <ul>
            <li class="active"><a href="index.php">Accueil</a></li>
            <li class="has-dropdown">
              <a href="courses.php">Formations</a>
              <ul class="dropdown">
                <li><a href="#">Cours en Ligne</a></li>
                <li><a href="#">Cours à domicile</a></li>
                <li><a href="#">Ateliers sur site</a></li>
                <li><a href="#">Se reconvertir dans la restauration</a></li>
              </ul>
            </li>
            <li><a href="teacher.html">Enseignants</a></li>
            <li><a href="about.html">À propos</a></li>
            <li><a href="pricing.html">Abonnement</a></li>
            <li class="has-dropdown">
              <a href="blog.html">Blog</a>
              <ul class="dropdown">
                <li><a href="#">Web Design</a></li>
                <li><a href="#">eCommerce</a></li>
                <li><a href="#">Branding</a></li>
                <li><a href="#">API</a></li>
              </ul>
            </li>
            <li><a href="contact.html">Contact</a></li>

            <?php if (isset($_SESSION['access_token'])) {
              $token = $_SESSION['access_token'];

              // Décoder le token d'accès JWT
              $jwtParts = explode('.', $token);
              $jwtClaims = base64_decode($jwtParts[1]);
              $claims = json_decode($jwtClaims, true);

              // Récupérer le nom de l'utilisateur
              $userFirstName = $claims['first_name'];
              $userLastName = $claims['last_name'];
              $userPicture = $claims['picture']
                ?>
              <!-- Utilisateur connecté -->
              <li class="dropdown">
                <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true"
                  aria-expanded="false">
                  <span class="user-name">
                    <?php echo $userLastName; ?>
                  </span>
                  <img src=" <?php echo $userPicture ?> " class="profile-picture" alt="Profile Picture" />
                </a>

                <ul class="dropdown-menu">
                  <li><a href="my_profile.php">Mon Profil</a></li>
                  <li><a href="#">Paramètres</a></li>
                  <li><a href="logout.php">Déconnexion</a></li>
                </ul>
              </li>
            <?php } else { ?>
              <!-- Utilisateur non connecté -->
              <li class="btn-cta"><a href="login.html"><span>Connexion</span></a></li>
              <li class="btn-cta"><a href="events/new-event.html"><span>Créer une Formation</span></a></li>
            <?php } ?>


          </ul>
        </div>
      </div>
    </div>
  </div>
</nav>