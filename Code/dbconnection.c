#include <stdio.h>
#include <stdlib.h>
#include <mysql.h>


int ConnectDB(char *user, char* pass, char* dbname)
{
  MYSQL *conn;
  conn = mysql_init(NULL);
  if (conn == NULL)
  {
    fprintf(stderr, "Error: Failed to initialize MySQL connection: %s\n", mysql_error(conn));
    exit(1);
  }

  if (!mysql_real_connect(conn, HOST, user, pass, dbname, PORT, UNIX_SOCKET, FLAG))
  {
    fprintf(stderr, "\nError: %s [%d]\n", mysql_error(conn), mysql_errno(conn));
    exit(1);
  }
  printf("Connection Successful to db");

  mysql_close(conn);
  return EXIT_SUCCESS;
}


