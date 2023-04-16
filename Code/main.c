#include <stdio.h>
#include <stdlib.h>
#include "http-requests.h"

int main(int argc, char *argv[])
{
  char * URL_SPOONACULAR = "https://api.spoonacular.com/recipes/complexSearch?apiKey=cceaeaefe34e4f388e3be6bff9458c78";
  CURL_MakeAGET(URL_SPOONACULAR);
  parseAndPrint("response.json");
  return 0;
}
