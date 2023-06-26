<!DOCTYPE html>
<html>
<head>
    <title>Mon Profil</title>
    <style>
        /* Styles CSS */
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 0;
            display: flex;
            justify-content: center;
            align-items: center;
            min-height: 100vh;
        }

        .profile-card {
            max-width: 400px;
            background: linear-gradient(to bottom right, #4c68d7, #8395f0);
            padding: 20px;
            border-radius: 5px;
            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
            text-align: center;
        }

        .profile-card img {
            width: 100px;
            height: 100px;
            border-radius: 50%;
            margin-bottom: 10px;
        }

        .profile-card h1 {
            color: #fff;
            margin-bottom: 10px;
        }

        .section-heading {
            font-weight: bold;
            margin-top: 20px;
            margin-bottom: 10px;
            color: #fff;
        }

        .section-content {
            margin-bottom: 20px;
            color: #fff;
        }

        .logout-link {
            color: #fff;
            text-decoration: none;
        }
    </style>
</head>
<body>
    <div class="profile-card" >
        <img src="profile-picture.jpg" alt="Photo de profil">
        <h1>Nom Complet</h1>

        <div class="section-heading">Informations personnelles</div>
        <div class="section-content">
            <p><strong>Adresse e-mail:</strong> user@example.com</p>
            <p><strong>Date d'inscription:</strong> 01/01/2023</p>
            <p><strong>Numéro de téléphone:</strong> +33 7 12345678</p>
            <!-- Ajoutez d'autres informations de contact si nécessaire -->
        </div>

        <div class="section-heading">Historique des cours</div>
        <div class="section-content">
            <ul>
                <li>Nom du cours 1 - Date</li>
                <li>Nom du cours 2 - Date</li>
                <li>Nom du cours 3 - Date</li>
                <!-- Ajoutez d'autres cours si nécessaire -->
            </ul>
        </div>

        <div class="section-heading">Paramètres</div>
        <div class="section-content">
            <p>Préférences de notification, langue, etc.</p>
            <!-- Ajoutez d'autres paramètres personnalisés si nécessaire -->
        </div>

        <div class="section-heading">Options de connexion</div>
        <div class="section-content">
            <p>Connecté avec: Google, Facebook, etc.</p>
            <!-- Ajoutez d'autres options de connexion si nécessaire -->
        </div>

        <div class="section-heading">Statut de l'abonnement</div>
        <div class="section-content">
            <p>Abonnement actif jusqu'au 01/01/2024</p>
            <!-- Ajoutez d'autres informations d'abonnement si nécessaire -->
        </div>

        <a class="logout-link" href="logout.php">Se déconnecter</a>
    </div>
</body>
</html>
