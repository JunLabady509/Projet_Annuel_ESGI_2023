#include <stdlib.h>
#include <stdio.h>
#include <curl/curl.h>
#include <json-c/json.h>


size_t write_callback(void *ptr, size_t size, size_t nmemb, void *userdata) {
    FILE *file = (FILE *)userdata;
    return fwrite(ptr, size, nmemb, file);
}

int CURL_MakeAGET(char *url)
{
  // Initialisation de cURL
  curl_global_init(CURL_GLOBAL_DEFAULT);
  CURL *curl = curl_easy_init();
  if (curl == NULL)
  {
    printf("Erreur : impossible d'initialiser cURL\n");
    return 1;
  }

  // Configuration de la requête
  curl_easy_setopt(curl, CURLOPT_URL, url); // URL de l'API
  curl_easy_setopt(curl, CURLOPT_HTTPGET, 1L);                                                                                      // Méthode HTTP GET
  curl_easy_setopt(curl, CURLOPT_FOLLOWLOCATION, 1L);                                                                               // Suivre les redirections
  curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_callback);                                                                    // Callback d'écriture
  FILE *file = fopen("response.json", "wb");                                                                                        // Ouverture du fichier en mode écriture binaire
  if (file == NULL)
  {
    printf("Erreur : impossible d'ouvrir le fichier\n");
    curl_easy_cleanup(curl);
    curl_global_cleanup();
    return 1;
  }

  curl_easy_setopt(curl, CURLOPT_WRITEDATA, file); // Passage du fichier comme userdata pour le callback

  // Exécution de la requête
  CURLcode res = curl_easy_perform(curl);
  if (res != CURLE_OK)
  {
    printf("Erreur cURL : %s\n", curl_easy_strerror(res));
    fclose(file);
    curl_easy_cleanup(curl);
    curl_global_cleanup();
    return 1;
  }

  // Nettoyage
  fclose(file);
  curl_easy_cleanup(curl);
  curl_global_cleanup();

  printf("Réponse enregistrée dans le fichier response.json\n");
  return 0;
}

int parseAndPrint(char *filename) {

  // Charger le fichier JSON
    FILE * file = fopen(filename, "r");
    if (file == NULL) {
        printf("Erreur : impossible d'ouvrir le fichier\n");
        return 1;
    }
    fseek(file, 0, SEEK_END);
    long file_size = ftell(file);
    fseek(file, 0, SEEK_SET);
    char *json_buffer = (char *)malloc(file_size + 1);
    fread(json_buffer, file_size, 1, file);
    fclose(file);
    json_buffer[file_size] = '\0';

    // Analyser le fichier JSON
    struct json_object *root = json_tokener_parse(json_buffer);
    if (root == NULL) {
        printf("Erreur : impossible de parser le fichier JSON\n");
        free(json_buffer);
        return 1;
    }

    // Accéder aux valeurs et objets JSON
    struct json_object *results;
    
    json_object_object_get_ex(root, "results", &results);
    if(json_object_get_type(results) == json_type_array){
        int array_len = json_object_array_length(results);
        printf("Le tableau contient %d éléments :\n", array_len);
        for (int i = 0; i < array_len; i++) {
            struct json_object *array_elem = json_object_array_get_idx(results, i);
            if (array_elem != NULL) {
                printf("Élément %d : %s\n", i, json_object_get_string(array_elem));
            }
        }
    } else {
        printf("Erreur : le fichier JSON ne contient pas un tableau\n");
    };

    // Libérer la mémoire
    json_object_put(root);
    free(json_buffer);

    return 0;
}
