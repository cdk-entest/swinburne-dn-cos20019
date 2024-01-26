---
title: build a photo album with php and mariadb
author: haimtran
date: 26/01/2024
---

## Setup Mariadb

## Photo Table

Create the photot table

```sql
CREATE TABLE IF NOT EXISTS photo(
    id int auto_increment primary key,
    title text,
    description text,
    date DATE DEFAULT CURRENT_DATE,
    image text,
    keyword text);
```

Insert a record in the photo table

```sql
INSERT INTO photo(title, description, date, image, keyword)
VALUES ('Vim Book', 'Photo of the vim book', '2024-01-26', 'hello.jpg', 'Vim Book');
```

Select records from the photo table

```sql
select * from photo;
```

Update a record in the photo table

```sql
UPDATE photo
SET image = "golang-idiomatic.jpg"
WHERE id = 2;
```

Delete a record in the photo table

```sql
delete from photo where id=1;
```
