#include "biblio.h"
#include <gtk/gtk.h>
#include <stddef.h>
#include <curl/curl.h>
#include <json-c/json.h>

// Déclaration de la variable globale pour stocker la clé API
const char *api_key = "cceaeaefe34e4f388e3be6bff9458c78";
char *api_url_with_key = NULL;

void on_search_button_clicked(GtkButton *button, gpointer user_data)
{
  ApiWidgets *api_widgets = (ApiWidgets *)user_data;
  GtkWidget *api_spoonacular = api_widgets->api_spoonacular;
  GtkWidget *api_openfoodfacts = api_widgets->api_openfoodfacts;
  GtkWidget *entry = api_widgets->entry;

  const gchar *product_name = gtk_entry_get_text(GTK_ENTRY(entry));

  // URL de l'API en fonction des boutons radios sélectionnés
  const char *api_url = NULL;
  // Vérifier quel bouton radio est sélectionné
  if (gtk_toggle_button_get_active(GTK_TOGGLE_BUTTON(api_spoonacular)))
  {
    api_url = "https://api.spoonacular.com/food/products/search?query=";

    // Ajout de la clé API à l'URL de l'API
    api_url_with_key = malloc(strlen(api_url) + strlen(api_key) + strlen(product_name) + 1);
    sprintf(api_url_with_key, "%s%s&apiKey=%s", api_url, product_name, api_key);
  }
  else if (gtk_toggle_button_get_active(GTK_TOGGLE_BUTTON(api_openfoodfacts)))
  {
    api_url = "https://world.openfoodfacts.org/cgi/search.pl?search_terms=";
  }
  else
  {
    fprintf(stderr, "Aucune API n'a été sélectionnée.\n");
    return;
  }

  // Initialisation de Curl
  CURL *curl = curl_easy_init();
  if (!curl)
  {
    fprintf(stderr, "Erreur lors de l'initialisation de libcurl.\n");
    return;
  }

  // Initialisation de la structure pour stocker les données de réponse
  struct ResponseData response_data;
  response_data.response = malloc(1);
  response_data.size = 0;

  // Définition des options de la requête curl
  curl_easy_setopt(curl, CURLOPT_URL, api_url_with_key);
  curl_easy_setopt(curl, CURLOPT_FOLLOWLOCATION, 1L);
  curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_response);
  curl_easy_setopt(curl, CURLOPT_WRITEDATA, &response_data);

  // Lancement de la requête
  CURLcode res = curl_easy_perform(curl);
  if (res != CURLE_OK)
  {
    fprintf(stderr, "Erreur lors de la requête : %s\n", curl_easy_strerror(res));
    curl_easy_cleanup(curl);
    free(response_data.response);
    return;
  }

  // Analyse de la réponse JSON et affichage
  json_object *json_response = json_tokener_parse(response_data.response);
  if (json_response == NULL)
  {
    fprintf(stderr, "Erreur lors de l'analyse de la réponse JSON.\n");
  }
  else
  {
    const char *response_string = json_object_to_json_string_ext(json_response, JSON_C_TO_STRING_PRETTY);
    printf("Réponse de l'API :\n%s\n", response_string);

    // Affichage de la réponse dans une boîte de dialogue GTK+
    GtkWidget *dialog = gtk_message_dialog_new(NULL, GTK_DIALOG_MODAL, GTK_MESSAGE_INFO, GTK_BUTTONS_OK, "Réponse de l'API :\n%s", response_string);
    gtk_dialog_run(GTK_DIALOG(dialog));
    gtk_widget_destroy(dialog);

    // Test de Gestion de l'affichage
    // Affichage de la réponse dans une boîte de dialogue GTK+
    // GtkWidget *dialog = gtk_dialog_new(NULL, GTK_DIALOG_MODAL);
    // gtk_window_set_title(GTK_WINDOW(dialog), "Réponse de l'API");
    // gtk_window_set_transient_for(GTK_WINDOW(dialog), GTK_WINDOW(window));
    // gtk_window_set_position(GTK_WINDOW(dialog), GTK_WIN_POS_CENTER);
    // gtk_dialog_add_button(GTK_DIALOG(dialog), "OK", GTK_RESPONSE_OK);
    // gtk_widget_set_size_request(dialog, 500, 400);

    // // Création d'un GtkTextView pour afficher la réponse JSON
    // GtkWidget *text_view = gtk_text_view_new();
    // gtk_text_view_set_editable(GTK_TEXT_VIEW(text_view), FALSE);
    // gtk_text_view_set_cursor_visible(GTK_TEXT_VIEW(text_view), FALSE);
    // gtk_text_view_set_wrap_mode(GTK_TEXT_VIEW(text_view), GTK_WRAP_WORD_CHAR);

    // // Obtention du buffer du GtkTextView
    // GtkTextBuffer *buffer = gtk_text_view_get_buffer(GTK_TEXT_VIEW(text_view));

    // // Ajout de la réponse JSON au buffer
    // gtk_text_buffer_set_text(buffer, response_string, -1);

    // // Création d'un GtkScrolledWindow pour ajouter une barre de défilement
    // GtkWidget *scrolled_window = gtk_scrolled_window_new(NULL, NULL);
    // gtk_scrolled_window_set_policy(GTK_SCROLLED_WINDOW(scrolled_window), GTK_POLICY_AUTOMATIC, GTK_POLICY_AUTOMATIC);
    // gtk_container_add(GTK_CONTAINER(scrolled_window), text_view);

    // // Ajout du GtkScrolledWindow à la boîte de dialogue
    // GtkWidget *content_area = gtk_dialog_get_content_area(GTK_DIALOG(dialog));
    // gtk_box_pack_start(GTK_BOX(content_area), scrolled_window, TRUE, TRUE, 0);

    // // Affichage de la boîte de dialogue et attente de la réponse de l'utilisateur
    // gtk_widget_show_all(dialog);
    // gtk_dialog_run(GTK_DIALOG(dialog));
    // gtk_widget_destroy(dialog);

    // Libération de la mémoire utilisée par la réponse JSON
    json_object_put(json_response);
  }

  // Nettoyage et libération de la mémoire
  curl_easy_cleanup(curl);
  free(response_data.response);
}

// Fonction de rappel pour gérer les données de réponse
size_t write_response(void *ptr, size_t size, size_t nmemb, struct ResponseData *data)
{
  size_t total_size = size * nmemb;
  data->response = realloc(data->response, data->size + total_size + 1);
  if (data->response == NULL)
  {
    fprintf(stderr, "Échec de l'allocation mémoire.\n");
    return 0;
  }

  memcpy(data->response + data->size, ptr, total_size);
  data->size += total_size;
  data->response[data->size] = '\0';

  return total_size;
}

void on_api_openfoodfacts_toggled(GtkToggleButton *togglebutton, gpointer user_data)
{
}

void on_api_spoonacular_toggled(GtkToggleButton *togglebutton, gpointer user_data)
{
}
