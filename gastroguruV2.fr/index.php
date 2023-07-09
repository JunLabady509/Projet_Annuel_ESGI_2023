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

	<div class="fh5co-loader"></div>
	<div id="page">
		<?php include 'includes/navbar.php'; ?>

		<aside id="fh5co-hero">
			<div class="flexslider">
				<ul class="slides">

					<?php
					$slides = [
						[
							'image' => 'images/slider-image3.jpg',
							'title' => 'Révélez votre chef intérieur, libérez votre génie culinaire',
							'subtitle' => 'Présenté par <a href="http://gastroguru.fr/" target="_blank">Gastroguru</a>',
						],
						[
							'image' => 'images/slider-image2.jpg',
							'title' => 'L\'excellence culinaire à son apogée',
							'subtitle' => 'Présenté par <a href="http://gastroguru.fr/" target="_blank">Gastroguru</a>',
						],
						[
							'image' => 'images/slide-accueil.jpg',
							'title' => 'Les racines de l\'éducation sont amères, mais les fruits en sont doux.',
							'subtitle' => 'Présenté par <a href="http://gastroguru.fr/" target="_blank">Gastroguru</a>',
						],
					];

					foreach ($slides as $slide) {
						?>
						<li style="background-image: url(<?php echo $slide['image']; ?>);">
							<div class="overlay-gradient"></div>
							<div class="container">
								<div class="row">
									<div class="col-md-8 col-md-offset-2 text-center slider-text">
										<div class="slider-text-inner">
											<h1>
												<?php echo $slide['title']; ?>
											</h1>
											<h2>
												<?php echo $slide['subtitle']; ?>
											</h2>
											<p><a class="btn btn-primary btn-lg" href="register.php">Commencez votre apprentissage dès Maintenant!</a></p>
										</div>
									</div>
								</div>
							</div>
						</li>
						<?php
					}
					?>
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
								<i class="icon-group" ></i>
							</span>
							<div class="desc">
								<h3><a href="learning/workshop/get-workshops.php">Ateliers variés sur site</a></h3>
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

					<?php include('events/get-events.php'); ?> 

				<div class="row">
					<?php
					// Exemples supplémentaires d'éléments de blog
					$blogItems = [
						[
							'image' => 'project-1.jpg',
							'title' => 'Healty Lifestyle &amp; Living',
							'date' => 'March. 15th',
							'comments' => 21,
							'description' => 'Far far away, behind the word mountains, far from the countries Vokalia and Consonantia, there live the blind texts.'
						],
						[
							'image' => 'project-2.jpg',
							'title' => 'Healty Lifestyle &amp; Living',
							'date' => 'March. 15th',
							'comments' => 21,
							'description' => 'Far far away, behind the word mountains, far from the countries Vokalia and Consonantia, there live the blind texts.'
						],
						[
							'image' => 'project-3.jpg',
							'title' => 'Healty Lifestyle &amp; Living',
							'date' => 'March. 15th',
							'comments' => 21,
							'description' => 'Far far away, behind the word mountains, far from the countries Vokalia and Consonantia, there live the blind texts.'
						],
					];

					foreach ($blogItems as $item) {
						echo '
                    <div class="col-lg-4 col-md-4">
                        <div class="fh5co-blog animate-box">
                            <a href="#" class="blog-img-holder" style="background-image: url(images/' . $item['image'] . ');"></a>
                            <div class="blog-text">
                                <h3><a href="#">' . $item['title'] . '</a></h3>
                                <span class="posted_on">' . $item['date'] . '</span>
                                <span class="comment"><a href="">' . $item['comments'] . '<i class="icon-speech-bubble"></i></a></span>
                                <p>' . $item['description'] . '</p>
                            </div>
                        </div>
                    </div>';
					}
					?>
				</div>
			</div>
		</div>
 
		
	<!-- Partie Abonnement ou Pricing -->
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
						<?php
						$plans = [
							[
								'title' => 'Trial',
								'price' => '0',
								'period' => 'per year',
								'features' => [
									'One Day Trial',
									'Limited Courses',
									'Free 3 Lessons',
									'No Supporter',
									'No Tutorial',
									'No Ebook',
									'Limited Registered User'
								],
							],
							[
								'title' => 'Silver',
								'price' => '79',
								'period' => 'per year',
								'features' => [
									'One Year Standard Access',
									'Limited Courses',
									'300+ Lessons',
									'Random Supporter',
									'View Only Ebook',
									'Standard Tutorials',
									'Unlimited Registered User'
								],
							],
							[
								'title' => 'Gold',
								'price' => '499',
								'period' => 'per year',
								'features' => [
									'Life Time Access',
									'Unlimited All Courses',
									'7000+ Lessons & Growing',
									'Free Supporter',
									'Free Ebook Downloads',
									'Premium Tutorials',
									'Unlimited Registered User'
								],
							]
						];

						foreach ($plans as $plan) {
							?>
							<div class="col-md-4 animate-box">
								<div class="pricing__item">
									<div class="wrap-price">
										<h3 class="pricing__title">
											<?php echo $plan['title']; ?>
										</h3>
									</div>
									<div class="pricing__price">
										<span class="pricing__anim pricing__anim--1">
											<span class="pricing__currency">$</span>
											<?php echo $plan['price']; ?>
										</span>
										<span class="pricing__anim pricing__anim--2">
											<span class="pricing__period">
												<?php echo $plan['period']; ?>
											</span>
										</span>
									</div>
									<div class="wrap-price">
										<ul class="pricing__feature-list">
											<?php foreach ($plan['features'] as $feature) { ?>
												<li class="pricing__feature">
													<?php echo $feature; ?>
												</li>
											<?php } ?>
										</ul>
										<button class="pricing__action">Choose plan</button>
									</div>
								</div>
							</div>
						<?php } ?>
					</div>
				</div>
			</div>
		</div>

<!-- Partie Promo disponible -->
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