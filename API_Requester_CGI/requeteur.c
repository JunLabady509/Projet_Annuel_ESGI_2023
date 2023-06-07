#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <curl/curl.h>
#include <cjson/cJSON.h>

struct MemoryStruct
{
    char *memory;
    size_t size;
};

size_t write_callback(void *contents, size_t size, size_t nmemb, void *userp)
{
    size_t realsize = size * nmemb;
    struct MemoryStruct *mem = (struct MemoryStruct *)userp;

    char *ptr = realloc(mem->memory, mem->size + realsize + 1);
    if (ptr == NULL)
    {
        /* out of memory! */
        printf("not enough memory (realloc returned NULL)\n");
        return 0;
    }

    mem->memory = ptr;
    memcpy(&(mem->memory[mem->size]), contents, realsize);
    mem->size += realsize;
    mem->memory[mem->size] = 0;

    return realsize;
}

int main(int argc, char **argv)
{
    // Callback pour écrire les données de réponse
    struct MemoryStruct chunk;

    chunk.memory = malloc(1); /* will be grown as needed by the realloc above */
    chunk.size = 0;           /* no data at this point */

    if (argc < 2)
    {
        fprintf(stderr, "Usage: %s <search term>\n", argv[0]);
        return 1;
    }

    char *api_key = getenv("API_KEY");
    if (api_key == NULL)
    {
        fprintf(stderr, "API_KEY non définie dans l'environnement\n");
        return 1;
    }

    char url[256];
    snprintf(url, sizeof(url), "https://api.spoonacular.com/recipes/complexSearch?query=%s&apiKey=%s", argv[1], api_key);

    CURL *curl = curl_easy_init();

    if (curl)
    {
        curl_easy_setopt(curl, CURLOPT_URL, url);
        curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_callback);
        curl_easy_setopt(curl, CURLOPT_WRITEDATA, (void *)&chunk);
        CURLcode res = curl_easy_perform(curl);

        if (res != CURLE_OK)
        {
            fprintf(stderr, "curl_easy_perform() failed: %s\n", curl_easy_strerror(res));
        }
        curl_easy_cleanup(curl);
    }

    if (chunk.memory)
    {
        cJSON *json = cJSON_Parse(chunk.memory);
        if (json == NULL)
        {
            fprintf(stderr, "Failed to parse JSON\n");
            return 1;
        }

        cJSON *results = cJSON_GetObjectItemCaseSensitive(json, "results");
        fprintf(stderr, "Number of results: %d\n", cJSON_GetArraySize(results));

        cJSON *item = NULL;

        printf("Content-type: text/html\n\n");
        printf("<html><body>\n");
        printf("<h1>Results for '%s'</h1>\n", argv[1]);
        printf("<ul>\n");
        cJSON_ArrayForEach(item, results)
        {
            cJSON *title = cJSON_GetObjectItemCaseSensitive(item, "title");
            cJSON *id = cJSON_GetObjectItemCaseSensitive(item, "id");
            if (cJSON_IsString(title) && (title->valuestring != NULL) && cJSON_IsNumber(id))
            {
                printf("<li><a href=\"https://api.spoonacular.com/recipes/%d/information\">%s</a></li>\n", id->valueint, title->valuestring);
            }
        }
        printf("</ul>\n");
        printf("</body></html>\n");

        cJSON_Delete(json);
        free(chunk.memory);
    }

    return 0;
}
