// A function that test the reading and the writing functionality of curl library with a given database

int CURL_Connectivity_Test(char *url);

int CURL_Read_Test(char *url, char *method, char *data, char *response);

int CURL_Write_Test(char *url, char *method, char *data, char *response);

int CURL_Update_Test(char *url, char *method, char *data, char *response);

int CURL_Delete_Test(char *url, char *method, char *data, char *response);

int CURL_Performance_Test(char *url, char *method, char *data, char *response, int num_requests, double *avg_time, double *avg_request_size, double *avg_response_size);

void CURL_FILE_UPLOAD_Test(char *url, char *filename);

void CURL_FILE_DOWNLOAD_Test(char *url, char *filename);
