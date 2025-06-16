# How to run

From the root directory

```bash
go run cmd/main.go
```

remember to replace in **db/conn.go** the "go_db" for "localhost" for development

## Docker

1. to build the image

```bash
sudo docker build -t go-products-management .
```

2. to run the database isolated (for development)

```bash
sudo docker compose up -d go_db
```

3. to run both images after build

```bash
sudo docker compose up -d
```

## Tables Creation

1. Create product table

```sql
CREATE TABLE product (
  id SERIAL PRIMARY KEY,
  product_name VARCHAR(255) NOT NULL,
  price NUMERIC(10, 2)
)
```

2. Insert product into product table

```sql
INSERT INTO product (
	product_name,
	price
)
VALUES (
	'Weed',
	4.20
)
```
