# SQLITE3

## Data Types
- `NULL` The value is a NULL value.
- `INTEGER` The value is a signed integer, stored in 1, 2, 3, 4, 6, or 8 bytes depending on the magnitude of the value.
- `REAL` The value is a floating point value, stored as an 8-byte IEEE floating point number.
- `TEXT` The value is a text string, stored using the database encoding (UTF-8, UTF-16BE or UTF-16LE).
- `BLOB` The value is a blob of data, stored exactly as it was input.
| `Example` Typenames From The CREATE TABLE Statement or CAST Expression | Resulting Affinity | Rule Used To Determine Affinity |
|----------------------------------------------------------------------|--------------------|---------------------------------|
| INT                                                                  | INTEGER            | 1                               |
| INTEGER                                                              | INTEGER            | 1                               |
| TINYINT                                                              | INTEGER            | 1                               |
| SMALLINT                                                             | INTEGER            | 1                               |
| MEDIUMINT                                                            | INTEGER            | 1                               |
| BIGINT                                                               | INTEGER            | 1                               |
| UNSIGNED BIG INT                                                     | INTEGER            | 1                               |
| INT2                                                                 | INTEGER            | 1                               |
| INT8                                                                 | INTEGER            | 1                               |
| CHARACTER(20)                                                        | TEXT               | 2                               |
| VARCHAR(255)                                                         | TEXT               | 2                               |
| VARYING CHARACTER(255)                                               | TEXT               | 2                               |
| NCHAR(55)                                                            | TEXT               | 2                               |
| NATIVE CHARACTER(70)                                                 | TEXT               | 2                               |
| NVARCHAR(100)                                                        | TEXT               | 2                               |
| TEXT                                                                 | TEXT               | 2                               |
| CLOB                                                                 | TEXT               | 2                               |
| BLOB                                                                 | BLOB               | 3                               |
| no datatype specified                                                | BLOB               | 3                               |
| REAL                                                                 | REAL               | 4                               |
| DOUBLE                                                               | REAL               | 4                               |
| DOUBLE PRECISION                                                     | REAL               | 4                               |
| FLOAT                                                                | REAL               | 4                               |
| NUMERIC                                                              | NUMERIC            | 5                               |
| DECIMAL(10,5)                                                        | NUMERIC            | 5                               |
| BOOLEAN                                                              | NUMERIC            | 5                               |
| DATE                                                                 | NUMERIC            | 5                               |
| DATETIME                                                             | NUMERIC            | 5                               |

## Comparison Expressions
SQLite version 3 has the usual set of SQL comparison operators including:
- `=`
- `==`
- `<`
- `<=`
- `>`
- `>=`
- `!=`
- ` `
- `IN`
- `NOT IN`
- `BETWEEN`
- `IS`
- `IS NOT`

## References & Resources
- [sqlite3 cli commands](https://www.sqlite.org/cli.html#special_commands_to_sqlite3_dot_commands_)
- https://www.ffmpeg.org/
