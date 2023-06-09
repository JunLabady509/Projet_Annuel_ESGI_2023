<!DOCTYPE HTML>
<html>

<head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<title>Gastroguru &mdash; Formations</title>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<meta name="description" content="Free HTML5 Website Template by freehtml5.co" />
	<meta name="keywords"
		content="free website templates, free html5, free template, free bootstrap, free website template, html5, css3, mobile first, responsive" />
	<meta name="author" content="freehtml5.co" />

	<!-- Facebook and Twitter integration -->
	<meta property="og:title" content="" />
	<meta property="og:image" content="" />
	<meta property="og:url" content="" />
	<meta property="og:site_name" content="" />
	<meta property="og:description" content="" />
	<meta name="twitter:title" content="" />
	<meta name="twitter:image" content="" />
	<meta name="twitter:url" content="" />
	<meta name="twitter:card" content="" />

	<link href="https://fonts.googleapis.com/css?family=Source+Sans+Pro:300,400,600,700" rel="stylesheet">
	<link href="https://fonts.googleapis.com/css?family=Roboto+Slab:300,400" rel="stylesheet">

	<!-- Animate.css -->
	<link rel="stylesheet" href="css/animate.css">
	<!-- Icomoon Icon Fonts-->
	<link rel="stylesheet" href="css/icomoon.css">
	<!-- Bootstrap  -->
	<link rel="stylesheet" href="css/bootstrap.css">

	<!-- Magnific Popup -->
	<link rel="stylesheet" href="css/magnific-popup.css">

	<!-- Owl Carousel  -->
	<link rel="stylesheet" href="css/owl.carousel.min.css">
	<link rel="stylesheet" href="css/owl.theme.default.min.css">

	<!-- Flexslider  -->
	<link rel="stylesheet" href="css/flexslider.css">

	<!-- Pricing -->
	<link rel="stylesheet" href="css/pricing.css">

	<!-- Theme style  -->
	<link rel="stylesheet" href="css/style.css">

	<!-- Modernizr JS -->
	<script src="js/modernizr-2.6.2.min.js"></script>
	<!-- FOR IE9 below -->
	<!--[if lt IE 9]>
	<script src="js/respond.min.js"></script>
	<![endif]-->

	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/OwlCarousel2/2.3.4/assets/owl.carousel.min.css">
	<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.6.0/jquery.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/OwlCarousel2/2.3.4/owl.carousel.min.js"></script>


</head>

<body>

	<div class="fh5co-loader"></div>
	<div id="page">
		<div id="page">
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
								<div id="fh5co-logo"><a href="index.php"><i class="icon-food"></i>Gastroguru<span>.</span></a></div>
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
										?>
										<!-- Utilisateur connecté -->
										<li class="dropdown">
											<a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true"
												aria-expanded="false">
												<span class="user-name">
													<?php echo $userLastName; ?>
												</span>
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
										<li class="btn-cta"><a href="learning/home_courses/create-home-course.php"><span>Créer une
													Formation</span></a></li>
									<?php } ?>

								</ul>
							</div>
						</div>
					</div>
				</div>
			</nav>

			<aside id="fh5co-hero">
				<div class="flexslider">
					<ul class="slides">
						<li style="background-image: url(images/unsplash_1.jpg);">
							<div class="overlay-gradient"></div>
							<div class="container">
								<div class="row">
									<div class="col-md-8 col-md-offset-2 text-center slider-text">
										<div class="slider-text-inner">
											<h1 class="heading-section">Nos Formations</h1>
										</div>
									</div>
								</div>
							</div>
						</li>
					</ul>
				</div>
			</aside>

					

			<!-- ............................................................................................................................-->
			<!-- ...................................................APERÇU DES ATELIERS............................................................-->
			<div id="fh5co-course">
				<div class="container" style="border-bottom: 1px solid #ccc;">
					<div class="row animate-box">
						<div class="col-md-6 col-md-offset-3 text-center fh5co-heading">
							<h2>Nos Ateliers</h2>
							<p>Apprenez à cuisiner de délicieux plats et découvrez de nouvelles recettes salées ou de pâtisserie lors
								de nos ateliers sur site. Nos professionnels passionnés vous guideront tout au long de l'expérience.</p>
						</div>
					</div>
					<div class="row">
						<div class="col-md-12">
							<div class="owl-carousel owl-theme">
								<?php
								// Effectuer la requête GET vers l'API des ateliers
								$workshops = file_get_contents('http://localhost:44446/workshops');

								// Vérifier si la requête a réussi
								if ($workshops !== false) {
									// Convertir les données JSON en tableau associatif
									$workshopsData = json_decode($workshops, true);
									// Parcourir les données des ateliers
								
									$count = 0;
									foreach ($workshopsData['workshops'] as $workshop) {
										// Construire le HTML pour un atelier
										$workshopHTML = '
										<div class="item">
								<div class="course">
									<a href="#" class="course-img" style="background-image: url(' . "profile_photos/logoG.png" . ');"></a>
									<div class="desc">
																<h3><a href="#">' . $workshop['title'] . '</a></h3>
																<p>' . $workshop['start_time'] . '</p>
																<p>' . $workshop['end_time'] . '</p>
																<p>' . $workshop['instructor_id'] . '</p>
																<span><a href="#" class="btn btn-primary btn-sm btn-course">Participer à cet atelier</a></span>
																</div>
																</div>
															</div>
														';

										// Afficher l'atelier
										echo $workshopHTML;
										$count++;
										// Si le nombre de cours affichés atteint 2, arrêter la boucle
										if ($count == 2) {
											break;
										}
									}
								} else {
									echo 'Erreur lors de la récupération des ateliers.';
								}
								?>
							</div>
						</div>
					</div>
				</div>
			
			<!-- ............................................................................................................................-->

			<!-- ...................................................APERÇU DES COURS EN LIGNE............................................................-->
			<div id="fh5co-course">
				<div class="container" style="border-bottom: 1px solid #ccc;">
					<div class="row animate-box">
						<div class="col-md-6 col-md-offset-3 text-center fh5co-heading">
							<h2>Nos formations sur Web</h2>
							<p>Participez à nos leçons de cuisine en ligne. Apprenez à votre propre rythme et découvrez une multitude
								de recettes, de conseils et de démonstrations culinaires.</p>
						</div>
					</div>
					<div class="row">
						<div class="col-md-12">
							<div class="owl-carousel owl-theme">
								<?php
								// Effectuer la requête GET vers l'API des cours en ligne
								$onlineLessons = file_get_contents('http://localhost:44446/onlinelessons');
								// Vérifier si la requête a réussi
								if ($onlineLessons !== false) {
									// Convertir les données JSON en tableau associatif
									$onlineLessonsData = json_decode($onlineLessons, true);
									// Parcourir les données des cours en ligne
									$count = 0; // Compteur pour suivre le nombre de cours affichés
									foreach ($onlineLessonsData['online_lessons'] as $onlineLesson) {
										// Construire le HTML pour un cours en ligne
										$onlineLessonHTML = '
							<div class="item">
								<div class="course">
									<a href="#" class="course-img" style="background-image: url(' . "profile_photos/logoG.png" . ');"></a>
									<div class="desc">
										<h3><a href="#">' . $onlineLesson['title'] . '</a></h3>
										<p>' . $onlineLesson['start_time'] . '</p>
										<p>' . $onlineLesson['end_time'] . '</p>
										<p>' . $onlineLesson['instructor_id'] . '</p>
										<span><a href="#" class="btn btn-primary btn-sm btn-course">Rejoindre ce cours</a></span>
									</div>
								</div>
							</div>
						';
										// Afficher le cours en ligne
										echo $onlineLessonHTML;

										$count++;
										// Si le nombre de cours affichés atteint 2, arrêter la boucle
										if ($count == 2) {
											break;
										}
									}
								} else {
									echo 'Erreur lors de la récupération des cours en ligne.';
								}
								?>
							</div>
						</div>
					</div>
					<div class="row">
					</div>
				</div>
			</div>
			<!-- ............................................................................................................................-->

			<!-- ...................................................APERÇU DES COURS À DOMICILE............................................................-->
			<div class="container" style="border-bottom: 1px solid #ccc;">
				<div class="row animate-box">
					<div class="col-md-6 col-md-offset-3 text-center fh5co-heading">
						<h2>Nos cours à domicile</h2>
						<p>Participez à nos cours de cuisine à domicile. Apprenez de nouvelles recettes et techniques culinaires
							chez vous, encadrés par nos chefs expérimentés.</p>
					</div>
				</div>
				<div class="row">
					<div class="col-md-12">
						<div class="owl-carousel owl-theme">
							<?php
							// Effectuer la requête GET vers l'API des cours à domicile prévus
							$homeCourses = file_get_contents('http://localhost:44446/homecourses');

							// Vérifier si la requête a réussi
							if ($homeCourses !== false) {
								// Convertir les données JSON en tableau associatif
								$homeCoursesData = json_decode($homeCourses, true);
								// Parcourir les données des cours en ligne
								$count = 0; // Compteur pour suivre le nombre de cours affichés
								foreach ($homeCoursesData['homecourses'] as $homeCourse) {
									// Construire le HTML pour un cours à domicile prévu
									$homeCourseHTML = '
										<div class="item">
										<div class="course">
											<a href="#" class="course-img" style="background-image: url(' . "profile_photos/logoG.png" . ');"></a>
											<div class="desc">
            <h3><a href="#">' . $homeCourse['title'] . '</a></h3>
            <p>' . $homeCourse['start_time'] . '</p>
            <p>' . $homeCourse['end_time'] . '</p>
            <p>' . $homeCourse['instructor_id'] . '</p>
						<span><a href="#" class="btn btn-primary btn-sm btn-course">Rejoindre ce cours</a></span>
						</div>
					</div>
				</div>
			';

									// Afficher le cours à domicile prévu
									echo $homeCourseHTML;
									$count++;
									// Si le nombre de cours affichés atteint 2, arrêter la boucle
									if ($count == 2) {
										break;
									}
								}
							} else {
								echo 'Erreur lors de la récupération des cours à domicile prévus.';
							}
							?>
						</div>
					</div>
				</div>
			</div>

			<!-- ............................................................................................................................-->


			</div>
			<div id="fh5co-register" style="background-image: url(images/img_bg_2.jpg);">
				<div class="overlay"></div>
				<div class="row">
					<div class="col-md-8 col-md-offset-2 animate-box">
						<div class="date-counter text-center">
							<h2>Get 400 of Online Courses for Free</h2>
							<h3>By Mike Smith</h3>
							<div class="simply-countdown simply-countdown-one"></div>
							<p><strong>Limited Offer, Hurry Up!</strong></p>
							<p><a href="#" class="btn btn-primary btn-lg btn-reg">Register Now!</a></p>
						</div>
					</div>
				</div>
			</div>

			<footer id="fh5co-footer" role="contentinfo" style="background-image: url(images/img_bg_4.jpg);">
				<div class="overlay"></div>
				<div class="container">
					<div class="row row-pb-md">
						<div class="col-md-3 fh5co-widget">
							<h3>About Education</h3>
							<p>Facilis ipsum reprehenderit nemo molestias. Aut cum mollitia reprehenderit. Eos cumque dicta adipisci
								architecto culpa amet.</p>
						</div>
						<div class="col-md-2 col-sm-4 col-xs-6 col-md-push-1 fh5co-widget">
							<h3>Learning</h3>
							<ul class="fh5co-footer-links">
								<li><a href="#">Course</a></li>
								<li><a href="#">Blog</a></li>
								<li><a href="#">Contact</a></li>
								<li><a href="#">Terms</a></li>
								<li><a href="#">Meetups</a></li>
							</ul>
						</div>

						<div class="col-md-2 col-sm-4 col-xs-6 col-md-push-1 fh5co-widget">
							<h3>Learn &amp; Grow</h3>
							<ul class="fh5co-footer-links">
								<li><a href="#">Blog</a></li>
								<li><a href="#">Privacy</a></li>
								<li><a href="#">Testimonials</a></li>
								<li><a href="#">Handbook</a></li>
								<li><a href="#">Held Desk</a></li>
							</ul>
						</div>

						<div class="col-md-2 col-sm-4 col-xs-6 col-md-push-1 fh5co-widget">
							<h3>Engage us</h3>
							<ul class="fh5co-footer-links">
								<li><a href="#">Marketing</a></li>
								<li><a href="#">Visual Assistant</a></li>
								<li><a href="#">System Analysis</a></li>
								<li><a href="#">Advertise</a></li>
							</ul>
						</div>

						<div class="col-md-2 col-sm-4 col-xs-6 col-md-push-1 fh5co-widget">
							<h3>Legal</h3>
							<ul class="fh5co-footer-links">
								<li><a href="#">Find Designers</a></li>
								<li><a href="#">Find Developers</a></li>
								<li><a href="#">Teams</a></li>
								<li><a href="#">Advertise</a></li>
								<li><a href="#">API</a></li>
							</ul>
						</div>
					</div>

					<div class="row copyright">
						<div class="col-md-12 text-center">
							<p>
								<small class="block">&copy; 2016 Free HTML5. All Rights Reserved.</small>
								<small class="block">Designed by <a href="http://freehtml5.co/" target="_blank">FreeHTML5.co</a> Demo
									Images: <a href="http://unsplash.co/" target="_blank">Unsplash</a> &amp; <a
										href="https://www.pexels.com/" target="_blank">Pexels</a></small>
							</p>
						</div>
					</div>

				</div>
			</footer>
		</div>

		<div class="gototop js-top">
			<a href="#" class="js-gotop"><i class="icon-arrow-up"></i></a>
		</div>

		<!-- jQuery -->
		<script src="js/jquery.min.js"></script>
		<!-- jQuery Easing -->
		<script src="js/jquery.easing.1.3.js"></script>
		<!-- Bootstrap -->
		<script src="js/bootstrap.min.js"></script>
		<!-- Waypoints -->
		<script src="js/jquery.waypoints.min.js"></script>
		<!-- Stellar Parallax -->
		<script src="js/jquery.stellar.min.js"></script>
		<!-- Carousel -->
		<script src="js/owl.carousel.min.js"></script>
		<!-- Flexslider -->
		<script src="js/jquery.flexslider-min.js"></script>
		<!-- countTo -->
		<script src="js/jquery.countTo.js"></script>
		<!-- Magnific Popup -->
		<script src="js/jquery.magnific-popup.min.js"></script>
		<script src="js/magnific-popup-options.js"></script>
		<!-- Count Down -->
		<script src="js/simplyCountdown.js"></script>
		<!-- Main -->
		<script src="js/main.js"></script>
		<script>
			var d = new Date(new Date().getTime() + 1000 * 120 * 120 * 2000);

			// default example
			simplyCountdown('.simply-countdown-one', {
				year: d.getFullYear(),
				month: d.getMonth() + 1,
				day: d.getDate()
			});

			//jQuery example
			$('#simply-countdown-losange').simplyCountdown({
				year: d.getFullYear(),
				month: d.getMonth() + 1,
				day: d.getDate(),
				enableUtc: false
			});
		</script>

		<script>
			$(document).ready(function () {
				$(".owl-carousel").owlCarousel({
					loop: true,
					nav: true,
					dots: false,
					responsive: {
						0: {
							items: 1
						},
						768: {
							items: 2
						}
					}
				});
			});
		</script>


</body>

</html>