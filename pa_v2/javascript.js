// Définition des articles disponibles dans le catalogue
var articles = [
    { id: 1, nom: "Article 1", prix: 10 },
    { id: 2, nom: "Article 2", prix: 15 },
    { id: 3, nom: "Article 3", prix: 20 }
];

// Fonction pour générer dynamiquement les articles du catalogue
function genererCatalogue() {
    var listeArticles = document.getElementById("liste-articles");

    articles.forEach(function(article) {
        var item = document.createElement("li");
        item.innerHTML = article.nom + " - " + article.prix + " €";
        item.setAttribute("data-id", article.id);
        item.addEventListener("click", ajouterAuPanier);
        listeArticles.appendChild(item);
    });
}

// Fonction pour ajouter un article au panier
function ajouterAuPanier(event) {
    var articleId = event.target.getAttribute("data-id");

    // Rechercher l'article correspondant dans le catalogue
    var article = articles.find(function(a) {
        return a.id.toString() === articleId;
    });

    if (article) {
        var listePanier = document.getElementById("liste-panier");

        // Vérifier si l'article est déjà présent dans le panier
        var articleExistant = listePanier.querySelector("[data-id='" + article.id + "']");
        if (articleExistant) {
            // Incrémenter la quantité
            var quantite = parseInt(articleExistant.getAttribute("data-quantite")) + 1;
            articleExistant.setAttribute("data-quantite", quantite);
            articleExistant.innerHTML = article.nom + " - " + article.prix + " € (x" + quantite + ")";
        } else {
            // Ajouter l'article au panier avec une quantité de 1
            var item = document.createElement("li");
            item.innerHTML = article.nom + " - " + article.prix + " € (x1)";
            item.setAttribute("data-id", article.id);
            item.setAttribute("data-quantite", 1);
            listePanier.appendChild(item);
        }
    }
}

// Fonction pour effectuer l'achat
function effectuerAchat() {
    var listePanier = document.getElementById("liste-panier");
    var articlesAchetes = [];

    // Récupérer les articles du panier
    var articlesPanier = listePanier.querySelectorAll("li");
    articlesPanier.forEach(function(item) {
        var articleId = item.getAttribute("data-id");
        var article = articles.find(function(a) {
            return a.id.toString() === articleId;
        });

        var quantite = parseInt(item.getAttribute("data-quantite"));
        var prixTotal = article.prix * quantite;

        articlesAchetes.push({
            nom: article.nom,
            quantite: quantite,
            prixTotal: prixTotal
        });
    });

    // Afficher le récapitulatif des articles achetés
    console.log(articlesAchetes);

    // Réinitialiser le panier
    listePanier.innerHTML = "";
}

// Générer le catalogue au chargement de la page
window.onload = function() {
    genererCatalogue();

    var btnAcheter = document.getElementById("btn-acheter");
    btnAcheter.addEventListener("click", effectuerAchat);
};
