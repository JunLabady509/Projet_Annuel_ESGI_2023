package subscribe

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

// Structure pour représenter les données d'abonnement choisi par l'utilisateur
type SubscriptionChoice struct {
	UserID           int    `json:"user_id"`
	SubscriptionType string `json:"subscription_type"`
	PaymentMethod    string `json:"payment_method"`
}

// Fonction pour permettre à un utilisateur de choisir un abonnement et de démarrer le processus de paiement
func ChooseSubscription(c echo.Context) error {

	// Récupérer les données d'abonnement depuis la requête
	subscriptionChoice := new(SubscriptionChoice)
	if err := c.Bind(subscriptionChoice); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Créer une session de paiement avec Stripe
	stripe.Key = "YOUR_STRIPE_SECRET_KEY"
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				Price:    stripe.String(subscriptionChoice.SubscriptionType),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		SuccessURL: stripe.String("https://your-website.com/success"),
		CancelURL:  stripe.String("https://your-website.com/cancel"),
	}
	session, err := session.New(params)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Générer un ID d'abonnement aléatoire
	subscriptionID := generateRandomSubscriptionID()

	// Enregistrer l'ID d'abonnement dans la base de données
	if err := saveSubscriptionID(subscriptionChoice.UserID, subscriptionID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Mettre à jour le champ Subscribed_ID de l'utilisateur dans la base de données
	if err := updateSubscriptionID(subscriptionChoice.UserID, subscriptionID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Redirection vers la page de paiement
	return c.Redirect(http.StatusFound, session.URL)
}

// Fonction pour générer un ID d'abonnement aléatoire
func generateRandomSubscriptionID() string {
	// Logique pour générer un ID d'abonnement aléatoire
	// Peut-être utiliser une bibliothèque pour générer des ID uniques, comme uuid
}

// Fonction pour enregistrer l'ID d'abonnement dans la base de données
func saveSubscriptionID(userID int, subscriptionID string) error {
	// Logique pour enregistrer l'ID d'abonnement dans la base de données
}

// Fonction pour mettre à jour le champ Subscribed_ID de l'utilisateur dans la base de données
func updateSubscriptionID(userID int, subscriptionID string) error {
	// Logique pour mettre à jour le champ Subscribed_ID de l'utilisateur dans la base de données
}
