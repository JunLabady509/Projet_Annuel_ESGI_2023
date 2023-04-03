#include <stdlib.h>
#include <stdio.h>
#include <curl/curl.h>

int CURL_Connectivity_Test(char *url)
{
  CURL *curl;
  CURLcode res;
  curl = curl_easy_init(); // Initialize curl
  if (curl)
  {
    curl_easy_setopt(curl, CURLOPT_URL, url); // Set URL
    res = curl_easy_perform(curl);            // Perform curl
    if (res != CURLE_OK)
    {
      fprintf(stderr, "curl_easy_perform() failed: %s\n", curl_easy_strerror(res));
      return 1;
    }
    else
    {
      printf("Connection to DB successful\n");
      curl_easy_cleanup(curl);
      return 0;
    }
  }
}
