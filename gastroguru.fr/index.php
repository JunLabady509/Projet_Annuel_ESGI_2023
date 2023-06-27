<?php
session_start();
?>

<!DOCTYPE HTML>
<html>

<head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<title>Gastroguru &mdash; Accueil</title>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<meta name="description" content="Free HTML5 Website Template by freehtml5.co" />
	<meta name="keywords"
		content="free website templates, free html5, free template, free bootstrap, free website template, html5, css3, mobile first, responsive" />
	<meta name="author" content="freehtml5.co" />

	<!--
	//////////////////////////////////////////////////////

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

</head>

<body>
	<!-- <style>
		nav .user-photo img {
			width: 30px;
			height: 30px;
			border-radius: 50%;
			margin-right: 5px;
		}
	</style> -->

	<div class="fh5co-loader"></div>

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
									<li class="btn-cta"><a href="events/create-event.php"><span>Créer une
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
					<li style="background-image: url(images/slider-image3.jpg);">
						<div class="overlay-gradient"></div>
						<div class="container">
							<div class="row">
								<div class="col-md-8 col-md-offset-2 text-center slider-text">
									<div class="slider-text-inner">
										<h1>Révélez votre chef intérieur, libérez votre génie culinaire</h1>
										<h2>Présenté par <a href="http://gastroguru.fr/" target="_blank">Gastroguru</a></h2>
										<p><a class="btn btn-primary btn-lg" href="#">Commencez votre apprentissage dès Maintenant!</a></p>
									</div>
								</div>
							</div>
						</div>
					</li>
					<li style="background-image: url(images/slider-image2.jpg );">
						<div class="overlay-gradient"></div>
						<div class="container">
							<div class="row">
								<div class="col-md-8 col-md-offset-2 text-center slider-text">
									<div class="slider-text-inner">
										<h1>L'excellence culinaire à son apogée</h1>
										<h2>Présenté par <a href="http://gastroguru.fr/" target="_blank">Gastroguru</a></h2>
										<p><a class="btn btn-primary btn-lg btn-learn" href="#">Commencez votre apprentissage dès
												Maintenant!</a></p>
									</div>
								</div>
							</div>
						</div>
					</li>
					<li style="background-image: url(images/slide-accueil.jpg);">
						<div class="overlay-gradient"></div>
						<div class="container">
							<div class="row">
								<div class="col-md-8 col-md-offset-2 text-center slider-text">
									<div class="slider-text-inner">
										<h1>Les racines de l'éducation sont amères, mais les fruits en sont doux.</h1>
										<h2>Présenté par <a href="http://gastroguru.fr/" target="_blank">Gastroguru</a></h2>
										<p><a class="btn btn-primary btn-lg btn-learn" href="#">Commencez votre apprentissage dès
												Maintenant!</a></p>
									</div>
								</div>
							</div>
						</div>
					</li>
				</ul>
			</div>
		</aside>

		<div id="fh5co-course-categories">
			<div class="container">
				<div class="row animate-box">
					<div class="col-md-6 col-md-offset-3 text-center fh5co-heading">
						<h2>Nos types de formations</h2>
						<p>Gastroguru propose une gamme diversifiée de formations culinaires, comprenant des ateliers variés sur
							site pour apprendre à cuisiner de délicieux plats et découvrir de nouvelles recettes salées ou de
							pâtisserie, dispensés par des professionnels passionnés. Nous proposons également des cours animés à
							domicile pour une expérience personnalisée, ainsi que des leçons de cuisine en ligne accessibles via notre
							site Web. De plus, Cook Master propose des formations professionnelles pour ceux qui souhaitent se
							reconvertir dans le domaine de la restauration.</p>
					</div>
				</div>
				<div class="row">
					<div class="col-md-3 col-sm-6 text-center animate-box">
						<div class="services">
							<span class="icon">
								<i class="icon-group"></i>
							</span>
							<div class="desc">
								<h3><a href="#">Ateliers variés sur site</a></h3>
								<p>Apprenez à cuisiner de délicieux plats et découvrez de nouvelles recettes salées ou de pâtisserie
									lors de nos ateliers sur site. Nos professionnels passionnés vous guideront tout au long de
									l'expérience.</p>
							</div>
						</div>
					</div>
					<div class="col-md-3 col-sm-6 text-center animate-box">
						<div class="services">
							<span class="icon">
								<i class="icon-home"></i>
							</span>
							<div class="desc">
								<h3><a href="#">Cours animés à domicile</a></h3>
								<p>Profitez d'une expérience personnalisée avec nos cours animés à domicile. Nos instructeurs qualifiés
									se rendront chez vous pour vous enseigner les techniques de cuisine et vous aider à améliorer vos
									compétences culinaires.</p>
							</div>
						</div>
					</div>
					<div class="col-md-3 col-sm-6 text-center animate-box">
						<div class="services">
							<span class="icon">
								<i class="icon-media-play"></i>
							</span>
							<div class="desc">
								<h3><a href="#">Leçons de cuisine en ligne</a></h3>
								<p>Accédez à notre site Web pour participer à nos leçons de cuisine en ligne. Apprenez à votre propre
									rythme et découvrez une multitude de recettes, de conseils et de démonstrations culinaires.</p>
							</div>
						</div>
					</div>
					<div class="col-md-3 col-sm-6 text-center animate-box">
						<div class="services">
							<span class="icon">
								<i class="icon-study"></i>
							</span>
							<div class="desc">
								<h3><a href="#">Formations professionnelles</a></h3>
								<p>Si vous envisagez une reconversion dans le domaine de la restauration, nous proposons des formations
									professionnelles pour vous aider à acquérir les compétences nécessaires pour réussir dans ce secteur.
								</p>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>

		<div id="fh5co-counter" class="fh5co-counters" style="background-image: url(images/img_bg_4.jpg);"
			data-stellar-background-ratio="0.5">
			<div class="overlay"></div>
			<div class="container">
				<div class="row">
					<div class="col-md-10 col-md-offset-1">
						<div class="row">
							<div class="col-md-3 col-sm-6 text-center animate-box">
								<span class="icon"><i class="icon-study"></i></span>
								<span class="fh5co-counter js-counter" data-from="0" data-to="3700" data-speed="5000"
									data-refresh-interval="50"></span>
								<span class="fh5co-counter-label">Membres</span>
							</div>
							<div class="col-md-3 col-sm-6 text-center animate-box">
								<span class="icon"><i class="icon-bulb"></i></span>
								<span class="fh5co-counter js-counter" data-from="0" data-to="5034" data-speed="5000"
									data-refresh-interval="50"></span>
								<span class="fh5co-counter-label">Formations Disponible</span>
							</div>
							<div class="col-md-3 col-sm-6 text-center animate-box">
								<span class="icon"><i class="icon-head"></i></span>
								<span class="fh5co-counter js-counter" data-from="0" data-to="1080" data-speed="5000"
									data-refresh-interval="50"></span>
								<span class="fh5co-counter-label">Professeurs Certifiés</span>
							</div>
							<div class="col-md-3 col-sm-6 text-center animate-box">
								<span class="icon"><i class="icon-world"></i></span>
								<span class="fh5co-counter js-counter" data-from="0" data-to="3297" data-speed="5000"
									data-refresh-interval="50"></span>
								<span class="fh5co-counter-label">Recettes Partagées</span>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>

		<div id="fh5co-testimonial" style="background-image: url(images/school.jpg);">
			<div class="overlay"></div>
			<div class="container">
				<div class="row animate-box">
					<div class="col-md-6 col-md-offset-3 text-center fh5co-heading">
						<h2><span>Testimonials</span></h2>
					</div>
				</div>
				<div class="row">
					<div class="col-md-10 col-md-offset-1">
						<div class="row animate-box">
							<div class="owl-carousel owl-carousel-fullwidth">
								<div class="item">
									<div class="testimony-slide active text-center">
										<div class="user" style="background-image: url(images/person1.jpg);"></div>
										<span>Mary Walker<br><small>Students</small></span>
										<blockquote>
											<p>&ldquo;Far far away, behind the word mountains, far from the countries Vokalia and Consonantia,
												there live the blind texts. Separated they live in Bookmarksgrove right at the coast of the
												Semantics, a large language ocean.&rdquo;</p>
										</blockquote>
									</div>
								</div>
								<div class="item">
									<div class="testimony-slide active text-center">
										<div class="user" style="background-image: url(images/person2.jpg);"></div>
										<span>Mike Smith<br><small>Students</small></span>
										<blockquote>
											<p>&ldquo;Separated they live in Bookmarksgrove right at the coast of the Semantics, a large
												language ocean.&rdquo;</p>
										</blockquote>
									</div>
								</div>
								<div class="item">
									<div class="testimony-slide active text-center">
										<div class="user" style="background-image: url(images/person3.jpg);"></div>
										<span>Rita Jones<br><small>Teacher</small></span>
										<blockquote>
											<p>&ldquo;Far from the countries Vokalia and Consonantia, there live the blind texts. Separated
												they live in Bookmarksgrove right at the coast of the Semantics, a large language ocean.&rdquo;
											</p>
										</blockquote>
									</div>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>

		<div id="fh5co-blog">
			<div class="container">
				<div class="row animate-box">
					<div class="col-md-8 col-md-offset-2 text-center fh5co-heading">
						<h2>Blog &amp; Events</h2>
						<p>Dignissimos asperiores vitae velit veniam totam fuga molestias accusamus alias autem provident. Odit ab
							aliquam dolor eius.</p>
					</div>
				</div>
				<div class="row row-padded-mb">
					<?php
					$response = file_get_contents('http://localhost:44446/events');
					if ($response === false) {
						echo "Erreur lors de la récupération des événements";
						return;
					}

					$data = json_decode($response, true);
					if ($data === null || !isset($data['events'])) {
						echo "Erreur lors du décodage de la réponse JSON ou données manquantes";
						return;
					}

					$events = $data['events'];

					$eventHTML = ''; // Initialise la variable
					
					foreach ($events as $event) {
						$dateString = $event['date'];

						// Convertir la chaîne de date en objet DateTime
						$date = new DateTime($dateString);
						// Récupérer le jour
						$day = $date->format('d');
						// Récupérer le mois (en texte complet, en anglais)
						$monthName = $date->format('F');

						// Accédez aux données de chaque événement en utilisant $event['property_name']
						// Par exemple, $event['title'], $event['description'], $event['date'], etc.
						$eventHTML .= '
    <div class="col-md-4 animate-box">
        <div class="fh5co-event">
            <div class="date text-center"><span>' . $day . '<br>' . $monthName . '</span></div>
            <h3><a href="#">' . $event['title'] . '</a></h3>
            <p>'.$event['description'].'</p>
            <p><a href="#">Read More</a></p>
        </div>
    </div>';
					}

					echo $eventHTML;
					// Assurez-vous d'adapter le code en fonction de votre structure de données JSON spécifique
					?>

					<div class="col-md-4 animate-box">
						<div class="fh5co-event">
							<div class="date text-center"><span>15<br>Mar.</span></div>
							<h3><a href="#">USA, International Triathlon Event</a></h3>
							<p>Far far away, behind the word mountains, far from the countries Vokalia and Consonantia, there live the
								blind texts.</p>
							<p><a href="#">Read More</a></p>
						</div>
					</div>
					<div class="col-md-4 animate-box">
						<div class="fh5co-event">
							<div class="date text-center"><span>15<br>Mar.</span></div>
							<h3><a href="#">New Device Develope by Microsoft</a></h3>
							<p>Far far away, behind the word mountains, far from the countries Vokalia and Consonantia, there live the
								blind texts.</p>
							<p><a href="#">Read More</a></p>
						</div>
					</div>
				</div>
				<div class="row">
					<div class="col-lg-4 col-md-4">
						<div class="fh5co-blog animate-box">
							<a href="#" class="blog-img-holder" style="background-image: url(images/project-1.jpg);"></a>
							<div class="blog-text">
								<h3><a href="#">Healty Lifestyle &amp; Living</a></h3>
								<span class="posted_on">March. 15th</span>
								<span class="comment"><a href="">21<i class="icon-speech-bubble"></i></a></span>
								<p>Far far away, behind the word mountains, far from the countries Vokalia and Consonantia, there live
									the blind texts.</p>
							</div>
						</div>
					</div>
					<div class="col-lg-4 col-md-4">
						<div class="fh5co-blog animate-box">
							<a href="#" class="blog-img-holder" style="background-image: url(images/project-2.jpg);"></a>
							<div class="blog-text">
								<h3><a href="#">Healty Lifestyle &amp; Living</a></h3>
								<span class="posted_on">March. 15th</span>
								<span class="comment"><a href="">21<i class="icon-speech-bubble"></i></a></span>
								<p>Far far away, behind the word mountains, far from the countries Vokalia and Consonantia, there live
									the blind texts.</p>
							</div>
						</div>
					</div>
					<div class="col-lg-4 col-md-4">
						<div class="fh5co-blog animate-box">
							<a href="#" class="blog-img-holder" style="background-image: url(images/project-3.jpg);"></a>
							<div class="blog-text">
								<h3><a href="#">Healty Lifestyle &amp; Living</a></h3>
								<span class="posted_on">March. 15th</span>
								<span class="comment"><a href="">21<i class="icon-speech-bubble"></i></a></span>
								<p>Far far away, behind the word mountains, far from the countries Vokalia and Consonantia, there live
									the blind texts.</p>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>

		<div id="fh5co-pricing" class="fh5co-bg-section">
			<div class="container">
				<div class="row animate-box">
					<div class="col-md-6 col-md-offset-3 text-center fh5co-heading">
						<h2>Plan &amp; Pricing</h2>
						<p>Dignissimos asperiores vitae velit veniam totam fuga molestias accusamus alias autem provident. Odit ab
							aliquam dolor eius.</p>
					</div>
				</div>
				<div class="row">
					<div class="pricing pricing--rabten">
						<div class="col-md-4 animate-box">
							<div class="pricing__item">
								<div class="wrap-price">
									<!-- <div class="icon icon-user2"></div> -->
									<h3 class="pricing__title">Trial</h3>
									<!-- <p class="pricing__sentence">Single user license</p> -->
								</div>
								<div class="pricing__price">
									<span class="pricing__anim pricing__anim--1">
										<span class="pricing__currency">$</span>0
									</span>
									<span class="pricing__anim pricing__anim--2">
										<span class="pricing__period">per year</span>
									</span>
								</div>
								<div class="wrap-price">
									<ul class="pricing__feature-list">
										<li class="pricing__feature">One Day Trial</li>
										<li class="pricing__feature">Limited Courses</li>
										<li class="pricing__feature">Free 3 Lessons</li>
										<li class="pricing__feature">No Supporter</li>
										<li class="pricing__feature">No Tutorial</li>
										<li class="pricing__feature">No Ebook</li>
										<li class="pricing__feature">Limited Registered User</li>
									</ul>
									<button class="pricing__action">Choose plan</button>
								</div>
							</div>
						</div>
						<div class="col-md-4 animate-box">
							<div class="pricing__item">
								<div class="wrap-price">
									<!-- <div class="icon icon-store"></div> -->
									<h3 class="pricing__title">Silver</h3>
									<!-- <p class="pricing__sentence">Up to 5 users</p> -->
								</div>
								<div class="pricing__price">
									<span class="pricing__anim pricing__anim--1">
										<span class="pricing__currency">$</span>79
									</span>
									<span class="pricing__anim pricing__anim--2">
										<span class="pricing__period">per year</span>
									</span>
								</div>
								<div class="wrap-price">
									<ul class="pricing__feature-list">
										<li class="pricing__feature">One Year Standard Access</li>
										<li class="pricing__feature">Limited Courses</li>
										<li class="pricing__feature">300+ Lessons</li>
										<li class="pricing__feature">Random Supporter</li>
										<li class="pricing__feature">View Only Ebook</li>
										<li class="pricing__feature">Standard Tutorials</li>
										<li class="pricing__feature">Unlimited Registered User</li>
									</ul>
									<button class="pricing__action">Choose plan</button>
								</div>
							</div>
						</div>
						<div class="col-md-4 animate-box">
							<div class="pricing__item">
								<div class="wrap-price">
									<!-- <div class="icon icon-home2"></div> -->
									<h3 class="pricing__title">Gold</h3>
									<!-- <p class="pricing__sentence">Unlimited users</p> -->
								</div>
								<div class="pricing__price">
									<span class="pricing__anim pricing__anim--1">
										<span class="pricing__currency">$</span>499
									</span>
									<span class="pricing__anim pricing__anim--2">
										<span class="pricing__period">per year</span>
									</span>
								</div>
								<div class="wrap-price">
									<ul class="pricing__feature-list">
										<li class="pricing__feature">Life Time Access</li>
										<li class="pricing__feature">Unlimited All Courses</li>
										<li class="pricing__feature">7000+ Lessons &amp; Growing</li>
										<li class="pricing__feature">Free Supporter</li>
										<li class="pricing__feature">Free Ebook Downloads</li>
										<li class="pricing__feature">Premium Tutorials</li>
										<li class="pricing__feature">Unlimited Registered User</li>
									</ul>
									<button class="pricing__action">Choose plan</button>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
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

		<div id="fh5co-gallery" class="fh5co-bg-section">
			<div class="row text-center">
				<h2><span>Instagram Gallery</span></h2>
			</div>
			<div class="row">
				<div class="col-md-3 col-padded">
					<a href="#" class="gallery" style="background-image: url(images/project-5.jpg);"></a>
				</div>
				<div class="col-md-3 col-padded">
					<a href="#" class="gallery" style="background-image: url(images/project-2.jpg);"></a>
				</div>
				<div class="col-md-3 col-padded">
					<a href="#" class="gallery" style="background-image: url(images/project-3.jpg);"></a>
				</div>
				<div class="col-md-3 col-padded">
					<a href="#" class="gallery" style="background-image: url(images/project-4.jpg);"></a>
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
</body>

</html>