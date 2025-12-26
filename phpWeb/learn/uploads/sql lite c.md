
## opening database
```c 
sqlite* db; //a pointer for errors and all
int rc = sqlite3_open("databasename",&db);

if (rc){
	printf("cant open database %s", sqlite3_errmsg(db));
}
else {
	printf("opened database successfully");
}

//sqlite3_open(filename, outdbhandle); creates or opens
//sqlite3_errmsg(); retrieves error message
//db handles everything rc is just a interger
sqlite3_close(db);
```
## executing sql commands
```c
char *sql = "CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY, name TEXT)";
char *errmsg = 0;

if (sqlite3_exec(db,sql,0,0,&errmsg) != SQLITE_OK){
	cerr<<"error: "<<errmsg<<endl;
	sqlite3_free(errmsg)
}
/*int sqlite3_exec(
  sqlite3*,                                  /* An open database 
  const char *sql,                           /* SQL to be evaluated 
  int (*callback)(void*,int,char**,char**),  /* Callback function 
  void *,                                    /* 1st argument to callback 
  char **errmsg                              /* Error msg written here which needs to be free later
);

SQLITE_OK: checks if there was any error
sqlite3_free: frees the error message used 
*/

```
## preparing statement
```c
sqlite3_stmt *stmt; //structure for stmt code
const char *sql = "select * from users"
rc = sqlite3_prepare_v2(db,sql,-1,&stmt,NULL); //compiles sqlcode to prepared statement executed multiple times
if (rc != SQLITE_OK){
	printf("failed to prepare statement: %s\n",sqlite3_errmsg(db));
}

int sqlite3_prepare_v2(
  sqlite3 *db,            /* Database handle */
  const char *zSql,       /* SQL statement, UTF-8 encoded */
  int nByte,              /* Maximum length of zSql in bytes. if negative number scan still null character*/
  sqlite3_stmt **ppStmt,  /* OUT: Statement handle */
  const char **pzTail     /* OUT: Pointer to unused portion of zSql */
);


```
## binding values to prepared statements
```c
sqlite3_bind_int(stmt,1,20);//bind 20 to first placeholder
sqlite3_bind_text(stmt,2,"shf");

```
## executing prepared statement
```c
while ((rc = sqlite3_step(stmt))== SQLITE_ROW){
	int id = sqlite3_column_int(stmt,0); //extract integer
	char *name = sqlite3_column_text(stmt,1);
	int age = sqlite3_column_int(stmt,2);
}
// sqlite3_step() executes the statment and fetch next if exist

sqlite3_finalize(stmt); // destroys and free space

```
# crud operation
## insert
```c
const char *sql = "INSERT INTO users (name, age) VALUES (?, ?);";
sqlite3_stmt *stmt;
sqlite3_prepare_v2(db, sql, -1, &stmt, NULL);
sqlite3_bind_text(stmt, 1, "Alice", -1, SQLITE_STATIC);
sqlite3_bind_int(stmt, 2, 30);

if (sqlite3_step(stmt) != SQLITE_DONE) {
    printf("Insert failed: %s\n", sqlite3_errmsg(db));
}
sqlite3_finalize(stmt);

```
## read / select
```c
const char *sql = "SELECT * FROM users WHERE age > ?";
sqlite3_stmt *stmt;
sqlite3_prepare_v2(db, sql, -1, &stmt, NULL);
sqlite3_bind_int(stmt, 1, 20);

while (sqlite3_step(stmt) == SQLITE_ROW) {
    int id = sqlite3_column_int(stmt, 0);
    const unsigned char *name = sqlite3_column_text(stmt, 1);
    int age = sqlite3_column_int(stmt, 2);
    printf("ID: %d, Name: %s, Age: %d\n", id, name, age);
}
sqlite3_finalize(stmt);

```
## update
```c
const char *sql = "UPDATE users SET age = ? WHERE name = ?;";
sqlite3_stmt *stmt;
sqlite3_prepare_v2(db, sql, -1, &stmt, NULL);
sqlite3_bind_int(stmt, 1, 35);
sqlite3_bind_text(stmt, 2, "Alice", -1, SQLITE_STATIC);

if (sqlite3_step(stmt) != SQLITE_DONE) {
    printf("Update failed: %s\n", sqlite3_errmsg(db));
}
sqlite3_finalize(stmt);

```
## delete
```c
const char *sql = "DELETE FROM users WHERE name = ?;";
sqlite3_stmt *stmt;
sqlite3_prepare_v2(db, sql, -1, &stmt, NULL);
sqlite3_bind_text(stmt, 1, "Alice", -1, SQLITE_STATIC);

if (sqlite3_step(stmt) != SQLITE_DONE) {
    printf("Delete failed: %s\n", sqlite3_errmsg(db));
}
sqlite3_finalize(stmt);

```
# transactions
```c
sqlite3_exec(db, "BEGIN TRANSACTION;", NULL, NULL, NULL);
// Perform multiple inserts/updates
sqlite3_exec(db, "COMMIT;", NULL, NULL, NULL);

```
