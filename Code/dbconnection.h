// I need a function that connect an sql database;

const PORT = 3306;
const char *HOST = 'localhost';
const char *UNIX_SOCKET = NULL;
unsigned int PORT = 3306;
unsigned int FLAG = 0;

// Une fonction qui permet de se connecter à une base de données
int ConnectDB(char *user, char* pass, char* dbname);

