#include <stdlib.h>
#include <stdio.h>
#include <gtk/gtk.h>

// Structure pour stocker les widgets des API
typedef struct {
    GtkWidget *api_spoonacular;
    GtkWidget *api_openfoodfacts;
    GtkWidget *entry;
} ApiWidgets;

// Structure pour stocker les données de réponse
struct ResponseData {
    char* response;
    size_t size;
};

// Signature des fonctions pour les possibles actions
void on_search_button_clicked(GtkButton *button, gpointer user_data);
void on_api_openfoodfacts_toggled(GtkToggleButton *togglebutton, gpointer user_data);
void on_api_spoonacular_toggled(GtkToggleButton *togglebutton, gpointer user_data);


size_t write_response(void *ptr, size_t size, size_t nmemb, struct ResponseData *data);



