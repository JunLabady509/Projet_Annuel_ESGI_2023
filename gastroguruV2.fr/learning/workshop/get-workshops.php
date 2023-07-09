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
    <link rel="stylesheet" href="../../css/animate.css">
    <!-- Icomoon Icon Fonts-->
    <link rel="stylesheet" href="../../css/icomoon.css">
    <!-- Bootstrap  -->
    <link rel="stylesheet" href="../../css/bootstrap.css">

    <!-- Magnific Popup -->
    <link rel="stylesheet" href="../../css/magnific-popup.css">

    <!-- Owl Carousel  -->
    <link rel="stylesheet" href="../../css/owl.carousel.min.css">
    <link rel="stylesheet" href="../../css/owl.theme.default.min.css">

    <!-- Flexslider  -->
    <link rel="stylesheet" href="../../css/flexslider.css">

    <!-- Pricing -->
    <link rel="stylesheet" href="../../css/pricing.css">

    <!-- Theme style  -->
    <link rel="stylesheet" href="../../css/style.css">

    <!-- Modernizr JS -->
    <script src="../../js/modernizr-2.6.2.min.js"></script>
    <!-- FOR IE9 below -->
    <!--[if lt IE 9]>
    <script src="js/respond.min.js"></script>
    <![endif]-->

    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/OwlCarousel2/2.3.4/assets/owl.carousel.min.css">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.6.0/jquery.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/OwlCarousel2/2.3.4/owl.carousel.min.js"></script>

</head>

<body>
    <!-- ............................................................................................................................-->
    <!-- ...................................................APERÇU DES ATELIERS............................................................-->
    <div id="fh5co-course">
        <div class="container" style="border-bottom: 1px solid #ccc;">
            <div class="row animate-box">
                <div class="col-md-6 col-md-offset-3 text-center fh5co-heading">
                    <h2>Nos Ateliers</h2>
                    <p>Apprenez à cuisiner de délicieux plats et découvrez de nouvelles recettes salées ou de pâtisserie
                        lors
                        de nos ateliers sur site. Nos professionnels passionnés vous guideront tout au long de
                        l'expérience.
                    </p>
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
                                <a href="#" class="course-img" style="background-image: url(' . "../" . $workshop['insight'] . ');"></a>
                                <div class="desc">
                                    <h3><a href="#">' . $workshop['id'] . '</a></h3>
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
                                // Si le nombre de cours affichés atteint 2, fermer la div "row" et en ouvrir une nouvelle
                                if ($count % 2 == 0) {
                                    echo '</div><div class="row">';
                                }

                            }
                            // Fermer la dernière div "row" si le nombre d'ateliers affichés n'est pas pair
                            if ($count % 2 == 0) {
                                echo '</div>';
                            }
                        } else {
                            echo 'Erreur lors de la récupération des ateliers.';
                        }
                        ?>
                    </div>
                </div>
            </div>




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


            <div class="gototop js-top">
                <a href="#" class="../../js-gotop"><i class="icon-arrow-up"></i></a>
            </div>

            <!-- jQuery -->
            <script src="../../js/jquery.min.js"></script>
            <!-- jQuery Easing -->
            <script src="../../js/jquery.easing.1.3.js"></script>
            <!-- Bootstrap -->
            <script src="../../js/bootstrap.min.js"></script>
            <!-- Waypoints -->
            <script src="../../js/jquery.waypoints.min.js"></script>
            <!-- Stellar Parallax -->
            <script src="../../js/jquery.stellar.min.js"></script>
            <!-- Carousel -->
            <script src="../../js/owl.carousel.min.js"></script>
            <!-- Flexslider -->
            <script src="../../js/jquery.flexslider-min.js"></script>
            <!-- countTo -->
            <script src="../../js/jquery.countTo.js"></script>
            <!-- Magnific Popup -->
            <script src="../../js/jquery.magnific-popup.min.js"></script>
            <script src="../../js/magnific-popup-options.js"></script>
            <!-- Count Down -->
            <script src="../../js/simplyCountdown.js"></script>
            <!-- Main -->
            <script src="../../js/main.js"></script>
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