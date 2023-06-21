#include "biblio.h"
#include <gtk/gtk.h>
#include <glib-2.0/glib.h>

int main(int argc, char *argv[])
{
  gtk_init(&argc, &argv);

  GtkWidget *api_spoonacular = NULL;
  GtkWidget *api_openfoodfacts = NULL;
  GtkWidget *entry = NULL;

  // Créer une fenêtre principale
  GtkWidget *window = gtk_window_new(GTK_WINDOW_TOPLEVEL);
  gtk_window_set_title(GTK_WINDOW(window), "Requeteur d'API");
  gtk_window_set_default_size(GTK_WINDOW(window), 300, 200);
  gtk_container_set_border_width(GTK_CONTAINER(window), 10);

  // Connection de la fenêtre à un signal de destruction
  g_signal_connect(window, "destroy", G_CALLBACK(gtk_main_quit), NULL);

  // Création d'un conteneur vertical pour organiser les widgets
  GtkWidget *vbox = gtk_box_new(GTK_ORIENTATION_VERTICAL, 5);
  gtk_container_add(GTK_CONTAINER(window), vbox);

  // Ajout d'un champ de saisie pour le nom du plat à rechercher
  entry = gtk_entry_new();
  gtk_box_pack_start(GTK_BOX(vbox), entry, FALSE, FALSE, 0);

  // Ajout de 2 boutons radios pour choisir l'API à utiliser
  api_spoonacular = gtk_radio_button_new_with_label(NULL, "API Spoonacular");
  api_openfoodfacts = gtk_radio_button_new_with_label_from_widget(GTK_RADIO_BUTTON(api_spoonacular), "API Open Food Facts");
  gtk_box_pack_start(GTK_BOX(vbox), api_spoonacular, FALSE, FALSE, 0);
  gtk_box_pack_start(GTK_BOX(vbox), api_openfoodfacts, FALSE, FALSE, 0);

  // Ajout du bouton de recherche
  GtkWidget *search_button = gtk_button_new_with_label("Rechercher");
  gtk_box_pack_start(GTK_BOX(vbox), search_button, FALSE, FALSE, 0);

  ApiWidgets api_widgets;
  api_widgets.api_spoonacular = api_spoonacular;
  api_widgets.api_openfoodfacts = api_openfoodfacts;
  api_widgets.entry = entry;

  // Connection des signaux d'évenements aux gestionnaires d'événements
  g_signal_connect(search_button, "clicked", G_CALLBACK(on_search_button_clicked), &api_widgets);
  g_signal_connect(api_openfoodfacts, "toggled", G_CALLBACK(on_api_openfoodfacts_toggled), NULL);
  g_signal_connect(api_spoonacular, "toggled", G_CALLBACK(on_api_spoonacular_toggled), NULL);

  // Afficher tous les éléments!
  gtk_widget_show_all(window);

  // Lancement de la boucle principale de GTK+
  gtk_main();

  return 0;
}
